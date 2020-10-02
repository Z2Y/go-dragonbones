package main

import (
	"dragonBones/bone"
	"image/color"
	"os"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type DragonBoneTest struct {
}

type ArmatureEntity struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (s *DragonBoneTest) Type() string {
	return "DragonBoneTest"
}

func (s *DragonBoneTest) Preload() {
	bone.Factory.SetAssetPath("Assets")
	bFile, _ := os.Open("Assets/mecha_1002_101d_show_ske.dbbin")
	tFile, _ := os.Open("Assets/mecha_1002_101d_show_tex.json")
	bone.Factory.LoadDragonBonesData(bFile, "", 1.0)
	bone.Factory.LoadTextureAtlasData(tFile, "", 1.0)
}

func (s *DragonBoneTest) Setup(u engo.Updater) {
	w := u.(*ecs.World)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&bone.DragonBoneSystem{})

	common.SetBackground(color.White)
	armatureDisplay := bone.Factory.BuildArmatureDisplay("mecha_1002_101d", "mecha_1002_101d_show")

	armatureEntity := ArmatureEntity{BasicEntity: ecs.NewBasic()}
	armatureEntity.RenderComponent.Drawable = armatureDisplay
	armatureEntity.SpaceComponent = common.SpaceComponent{
		Width:  30,
		Height: 30,
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *bone.DragonBoneSystem:
			sys.Add(&armatureEntity.BasicEntity, &armatureEntity.RenderComponent, &armatureEntity.SpaceComponent)
		}
	}
}

func main() {
	opts := engo.RunOptions{
		Title:  "DragonBones Test",
		Width:  800,
		Height: 600,
	}

	engo.Run(opts, &DragonBoneTest{})
}
