package bone

import (
	"image/color"
	"unsafe"

	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/EngoEngine/gl"
)

type Texture struct {
	id       *gl.Texture
	width    float32
	height   float32
	viewport engo.AABB
}

type Sprite struct {
	Display

	spriteFrame *common.TextureResource

	uvs      []float32
	indices  []uint16
	vertices []float32
}

func NewSprite() *Sprite {
	sprite := Sprite{}
	return &sprite
}

func (sprite *Sprite) createTextureBuffer() {
	if sprite.spriteFrame == nil {
		panic("can't create buffer for empty sprite")
	}
	viewport := sprite.spriteFrame.Viewport
	u, v, u2, v2 := viewport.Min.X, viewport.Min.Y, viewport.Max.X, viewport.Max.Y
	sprite.vertices = make([]float32, 8)
	sprite.uvs = make([]float32, 8)
	sprite.indices = make([]uint16, 6)
	sprite.indices[0] = 0
	sprite.indices[1] = 1
	sprite.indices[2] = 2
	sprite.indices[3] = 0
	sprite.indices[4] = 2
	sprite.indices[5] = 3

	sprite.uvs[0] = u
	sprite.uvs[1] = v
	sprite.uvs[2] = u2
	sprite.uvs[3] = v
	sprite.uvs[4] = u2
	sprite.uvs[5] = v2
	sprite.uvs[6] = u
	sprite.uvs[7] = v2

	sprite.vertices[0] = 0
	sprite.vertices[1] = 0
	sprite.vertices[2] = sprite.spriteFrame.Width
	sprite.vertices[3] = 0
	sprite.vertices[4] = sprite.spriteFrame.Width
	sprite.vertices[5] = sprite.spriteFrame.Height
	sprite.vertices[6] = 0
	sprite.vertices[7] = sprite.spriteFrame.Height
}

func (sprite *Sprite) setSpriteFrame(textureData *common.TextureResource) {
	sprite.spriteFrame = textureData
	sprite.RenderComponent.Color = color.White

	viewport := engo.AABB{Max: engo.Point{X: 1.0, Y: 1.0}}
	if textureData.Viewport != nil {
		viewport = *textureData.Viewport
	}
	texture := &Texture{textureData.Texture, textureData.Width, textureData.Height, viewport}

	sprite.RenderComponent.Drawable = (*common.Texture)(unsafe.Pointer(texture))
	sprite.RenderComponent.Scale = engo.Point{X: 1, Y: 1}
}

func (sprite *Sprite) Width() float32 {
	if sprite.spriteFrame == nil {
		return 0
	}
	return sprite.spriteFrame.Width
}

func (sprite *Sprite) Height() float32 {
	if sprite.spriteFrame == nil {
		return 0
	}
	return sprite.spriteFrame.Height
}

func (sprite *Sprite) Texture() *gl.Texture {
	if sprite.spriteFrame == nil {
		return nil
	}
	return sprite.spriteFrame.Texture
}

func (sprite *Sprite) View() (float32, float32, float32, float32) {
	if sprite.spriteFrame == nil {
		return 0, 0, 0, 0
	}
	return sprite.spriteFrame.Viewport.Min.X, sprite.spriteFrame.Viewport.Min.Y, sprite.spriteFrame.Viewport.Max.X, sprite.spriteFrame.Viewport.Max.Y
}

func (sprite *Sprite) Close() {
}
