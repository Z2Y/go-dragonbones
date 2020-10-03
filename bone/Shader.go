package bone

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

var (
	DragonBoneHUDShader     = NewDragonBoneShader(common.HUDShader)
	DragonBoneDefaultShader = NewDragonBoneShader(common.DefaultShader)
)

type DragonBoneShader struct {
	wrappedShader common.Shader
	culling       common.CullingShader
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
}

func (db *DragonBoneShader) ShouldDraw(rc *common.RenderComponent, sc *common.SpaceComponent) bool {
	// todo
	return true
}

func (db *DragonBoneShader) Pre() {
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
		if display.RenderComponent.Drawable != nil {
			db.wrappedShader.Draw(&display.RenderComponent, &display.SpaceComponent)
		}
	}
}

func (db *DragonBoneShader) Post() {
	db.wrappedShader.Post()
}

func (db *DragonBoneShader) Setup(*ecs.World) error {
	return nil
}

func (db *DragonBoneShader) SetCamera(cs *common.CameraSystem) {
}
