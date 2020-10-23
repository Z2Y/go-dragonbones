package bone

import (
	"image/color"
	"log"
	"math"

	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/EngoEngine/gl"
)

const (
	defaultVertexShader = `
	attribute vec2 aVertexPosition;
	attribute vec2 aTextureCoord;

	uniform mat3 matrixProjView;
	uniform mat3 modelMatrixView;

	varying vec2 vTextureCoord;

	void main() {
	  vTextureCoord = aTextureCoord;

	  vec3 matr = matrixProjView * modelMatrixView * vec3(aVertexPosition, 1.0);
	  gl_Position = vec4(matr.xy, 0, matr.z);
	}
`

	defaultFragmentShader = `
	#ifdef GL_ES
	#define LOWP lowp
	precision mediump float;
	#else
	#define LOWP
	#endif

	
	varying vec2 vTextureCoord;

	uniform vec4 uColor;
	uniform sampler2D uf_Texture;

	void main (void) {
	  gl_FragColor = texture2D(uf_Texture, vTextureCoord);
	}
`
)

type basicMeshShader struct {
	BatchSize int

	indices      []uint16
	vertexBuffer *gl.Buffer
	uvBuffer     *gl.Buffer
	indexBuffer  *gl.Buffer
	program      *gl.Program

	vertices                     []float32
	lastTexture                  *gl.Texture
	lastRepeating                common.TextureRepeating
	lastMagFilter, lastMinFilter common.ZoomFilter

	inPosition  int
	inTexCoords int

	inColor         *gl.UniformLocation
	matrixProjView  *gl.UniformLocation
	modelMatrixView *gl.UniformLocation

	projectionMatrix *engo.Matrix
	viewMatrix       *engo.Matrix
	projViewMatrix   *engo.Matrix
	modelMatrix      *engo.Matrix
	cullingMatrix    *engo.Matrix

	projViewChange bool

	camera        *common.CameraSystem
	cameraEnabled bool
}

type meshGLData struct {
	texture *gl.Texture
}

func (m *meshGLData) bindData() {
	engo.Gl.BindTexture(engo.Gl.TEXTURE_2D, m.texture)
}

func (m *meshGLData) upload(position int) {
	engo.Gl.VertexAttribPointer(position, 2, engo.Gl.FLOAT, false, 2*4, 0)
}

func (s *basicMeshShader) Setup() error {

	var err error
	s.program, err = common.LoadShader(defaultVertexShader, defaultFragmentShader)
	if err != nil {
		log.Println("Init Shader failed!")
		return err
	}
	s.vertexBuffer = engo.Gl.CreateBuffer()
	s.indexBuffer = engo.Gl.CreateBuffer()
	s.uvBuffer = engo.Gl.CreateBuffer()

	s.inPosition = engo.Gl.GetAttribLocation(s.program, "aVertexPosition")
	s.inTexCoords = engo.Gl.GetAttribLocation(s.program, "aTextureCoord")

	s.inColor = engo.Gl.GetUniformLocation(s.program, "uColor")
	s.matrixProjView = engo.Gl.GetUniformLocation(s.program, "matrixProjView")
	s.modelMatrixView = engo.Gl.GetUniformLocation(s.program, "modelMatrixView")

	s.projectionMatrix = engo.IdentityMatrix()
	s.viewMatrix = engo.IdentityMatrix()
	s.projViewMatrix = engo.IdentityMatrix()
	s.modelMatrix = engo.IdentityMatrix()
	s.cullingMatrix = engo.IdentityMatrix()

	return nil
}

func (s *basicMeshShader) Pre() {
	engo.Gl.Enable(engo.Gl.BLEND)
	engo.Gl.BlendFunc(engo.Gl.SRC_ALPHA, engo.Gl.ONE_MINUS_SRC_ALPHA)
	// Enable shader and buffer, enable attributes in shader
	engo.Gl.UseProgram(s.program)
	engo.Gl.EnableVertexAttribArray(s.inPosition)
	engo.Gl.EnableVertexAttribArray(s.inTexCoords)
	// engo.Gl.EnableVertexAttribArray(s.inColor)

	// The matrixProjView shader uniform is projection * view.
	// We do the multiplication on the CPU instead of sending each matrix to the shader and letting the GPU do the multiplication,
	// because it's likely faster to do the multiplication client side and send the result over the shader bus than to send two separate
	// buffers over the bus and then do the multiplication on the GPU.
	if s.projViewChange {
		s.projViewMatrix = s.projectionMatrix.Multiply(s.viewMatrix)
		s.projViewChange = false
	}
	engo.Gl.UniformMatrix3fv(s.matrixProjView, false, s.projViewMatrix.Val[:])
}

func (s *basicMeshShader) PrepareCulling() {
	s.projViewChange = true
	// (Re)initialize the projection matrix.
	s.projectionMatrix.Identity()
	if engo.ScaleOnResize() {
		s.projectionMatrix.Scale(1/(engo.GameWidth()/2), 1/(-engo.GameHeight()/2))
	} else {
		s.projectionMatrix.Scale(1/(engo.CanvasWidth()/(2*engo.CanvasScale())), 1/(-engo.CanvasHeight()/(2*engo.CanvasScale())))
	}
	// (Re)initialize the view matrix
	s.viewMatrix.Identity()
	if s.cameraEnabled {
		s.viewMatrix.Scale(1/s.camera.Z(), 1/s.camera.Z())
		s.viewMatrix.Translate(-s.camera.X(), -s.camera.Y()).Rotate(s.camera.Angle())
	} else {
		scaleX, scaleY := s.projectionMatrix.ScaleComponent()
		s.viewMatrix.Translate(-1/scaleX, 1/scaleY)
	}
	s.cullingMatrix.Identity()
	s.cullingMatrix.Multiply(s.projectionMatrix).Multiply(s.viewMatrix)
	s.cullingMatrix.Scale(engo.GetGlobalScale().X, engo.GetGlobalScale().Y)
}

func (s *basicMeshShader) Draw(mesh *Sprite, space *common.SpaceComponent) {
	// If our texture (or any of its properties) has changed or we've reached the end of our buffer, flush before moving on.
	engo.Gl.BindTexture(engo.Gl.TEXTURE_2D, mesh.Drawable.Texture())
	engo.Gl.TexParameteri(engo.Gl.TEXTURE_2D, engo.Gl.TEXTURE_WRAP_S, engo.Gl.CLAMP_TO_EDGE)
	engo.Gl.TexParameteri(engo.Gl.TEXTURE_2D, engo.Gl.TEXTURE_MAG_FILTER, engo.Gl.NEAREST)
	engo.Gl.TexParameteri(engo.Gl.TEXTURE_2D, engo.Gl.TEXTURE_MIN_FILTER, engo.Gl.NEAREST)

	engo.Gl.BindBuffer(engo.Gl.ARRAY_BUFFER, s.uvBuffer)
	engo.Gl.BufferData(engo.Gl.ARRAY_BUFFER, mesh.uvs, engo.Gl.STREAM_DRAW)
	engo.Gl.VertexAttribPointer(s.inTexCoords, 2, engo.Gl.FLOAT, false, 2*4, 0)

	engo.Gl.BindBuffer(engo.Gl.ARRAY_BUFFER, s.vertexBuffer)
	engo.Gl.BufferData(engo.Gl.ARRAY_BUFFER, mesh.vertices, engo.Gl.STREAM_DRAW)
	engo.Gl.VertexAttribPointer(s.inPosition, 2, engo.Gl.FLOAT, false, 2*4, 0)

	// tint := colorToFloat32(mesh.RenderComponent.Color)
	// engo.Gl.Uniform1f(s.inColor, tint)

	engo.Gl.BindBuffer(engo.Gl.ELEMENT_ARRAY_BUFFER, s.indexBuffer)
	engo.Gl.BufferData(engo.Gl.ELEMENT_ARRAY_BUFFER, mesh.indices, engo.Gl.STATIC_DRAW)

	engo.Gl.UniformMatrix3fv(s.modelMatrixView, false, mesh.globalTransform.Val[:])

	engo.Gl.DrawElements(engo.Gl.TRIANGLES, len(mesh.indices), engo.Gl.UNSIGNED_SHORT, 0)
}

func (s *basicMeshShader) Post() {
	// Cleanup
	engo.Gl.DisableVertexAttribArray(s.inPosition)
	engo.Gl.DisableVertexAttribArray(s.inTexCoords)
	// engo.Gl.DisableVertexAttribArray(s.inColor)

	engo.Gl.BindTexture(engo.Gl.TEXTURE_2D, nil)
	engo.Gl.BindBuffer(engo.Gl.ARRAY_BUFFER, nil)
	engo.Gl.BindBuffer(engo.Gl.ELEMENT_ARRAY_BUFFER, nil)

	engo.Gl.Disable(engo.Gl.BLEND)
}

func (s *basicMeshShader) flush() {
	// If we haven't rendered anything yet, no point in flushing.
	// engo.Gl.BufferData(engo.Gl.ARRAY_BUFFER, s.vertices, engo.Gl.STATIC_DRAW)
	// We only want to draw the indicies up to the number of sprites in the current batch.
}

func colorToFloat32(c color.Color) float32 {
	colorR, colorG, colorB, colorA := c.RGBA()
	colorR >>= 8
	colorG >>= 8
	colorB >>= 8
	colorA >>= 8

	red := colorR
	green := colorG << 8
	blue := colorB << 16
	alpha := colorA << 24

	return math.Float32frombits((alpha | blue | green | red) & 0xfeffffff)
}

func (s *basicMeshShader) SetCamera(c *common.CameraSystem) {
	s.projViewChange = true
	if s.cameraEnabled {
		s.camera = c
		s.viewMatrix.Identity().Translate(-s.camera.X(), -s.camera.Y()).Rotate(s.camera.Angle())
	} else {
		scaleX, scaleY := s.projectionMatrix.ScaleComponent()
		s.viewMatrix.Translate(-1/scaleX, 1/scaleY)
	}
}
