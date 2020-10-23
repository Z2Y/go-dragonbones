package bone

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

var (
	DragonBoneHUDShader     = NewDragonBoneShader(common.HUDShader)
	DragonBoneDefaultShader = NewDragonBoneShader(common.HUDShader)
)

type DragonBoneShader struct {
	wrappedShader common.Shader
	meshShader    *basicMeshShader
	culling       common.CullingShader

	isRenderMesh bool
}

func NewDragonBoneShader(shader common.Shader) *DragonBoneShader {
	if culling, ok := shader.(common.CullingShader); ok {
		return &DragonBoneShader{wrappedShader: shader, culling: culling}
	}
	return &DragonBoneShader{wrappedShader: shader}
}

func (db *DragonBoneShader) PrepareCulling() {
	if db.culling != nil {
		db.culling.PrepareCulling()
	}
	if db.meshShader != nil {
		db.meshShader.PrepareCulling()
	}
}

func (db *DragonBoneShader) ShouldDraw(rc *common.RenderComponent, sc *common.SpaceComponent) bool {
	// todo
	return true
}

func (db *DragonBoneShader) Pre() {
	if db.meshShader == nil {
		db.meshShader = &basicMeshShader{}
		db.meshShader.Setup()
	}
	db.wrappedShader.Pre()
}

func (db *DragonBoneShader) Draw(rc *common.RenderComponent, sc *common.SpaceComponent) {
	display, ok := rc.Drawable.(IDisplay)
	if !ok {
		return
	}
	display.UpdateTransform(true)
	db.DrawDisplay(rc.Drawable.(IDisplay))
}

func (db *DragonBoneShader) DrawDisplay(iDisplay IDisplay) {
	children := iDisplay.GetChildren()
	for _, subDisplay := range children {
		db.DrawDisplay(subDisplay)
	}

	switch display := iDisplay.(type) {
	case *Sprite:
		if !display.Hidden && display.Drawable != nil {
			db.meshShader.Pre()
			db.meshShader.Draw(display, &display.SpaceComponent)
			db.meshShader.Post()
			//db.wrappedShader.Draw(&display.RenderComponent, &display.SpaceComponent)
		}
	}
}

func (db *DragonBoneShader) Post() {
	db.wrappedShader.Post()
}

func (db *DragonBoneShader) Setup(w *ecs.World) error {
	return nil
}

func (db *DragonBoneShader) SetCamera(cs *common.CameraSystem) {
}
