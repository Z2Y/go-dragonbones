package bone

import (
	"image/color"
	"log"
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
	common.RenderComponent
	common.SpaceComponent

	spriteFrame *common.TextureResource
}

func NewSprite() *Sprite {
	sprite := Sprite{}
	return &sprite
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
	log.Println("SpriteDrawable", sprite.RenderComponent.Drawable, textureData)
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
