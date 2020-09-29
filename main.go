package main

import (
	"dragonBones/bone"
	"log"
	"os"
)

func main() {
	log.Println("hello")
	bFile, _ := os.Open("Assets/mecha_1002_101d_show_ske.dbbin")
	tFile, _ := os.Open("Assets/mecha_1002_101d_show_tex.json")

	factory := bone.NewFactory()
	factory.LoadDragonBonesData(bFile, "", 1.0)
	factory.LoadTextureAtlasData(tFile, "", 1.0)
	log.Println("LoadCompleted")
	factory.BuildArmatureDisplay("mecha_1002_101d", "mecha_1002_101d_show")
	log.Println("BuildCompleted")
	// display.GetAnimation().Play("idle")
	log.Println("load Finish")
}
