package bone

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type DragonBoneEntity struct {
	*ecs.BasicEntity
	*common.RenderComponent
	*common.SpaceComponent
}

type DragonBoneSystem struct {
	renderer *common.RenderSystem
}

func (d *DragonBoneSystem) New(world *ecs.World) {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			d.renderer = sys
		}
	}
}

func (d *DragonBoneSystem) Add(basic *ecs.BasicEntity, render *common.RenderComponent, space *common.SpaceComponent) {
	_, ok := render.Drawable.(*ArmatureDisplay)
	if !ok {
		panic("Add Wrong Render To DragonBoneSystem")
	}
	render.SetShader(DragonBoneDefaultShader)
	d.renderer.Add(basic, render, space)
}

func (d *DragonBoneSystem) Remove(basic ecs.BasicEntity) {
	d.renderer.Remove(basic)
}

func (ui *DragonBoneSystem) Update(dt float32) {
	Factory.Update(dt)
}
