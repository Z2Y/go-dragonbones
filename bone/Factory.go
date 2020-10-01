package bone

import "C"

import (
	wrapper "dragonBones/dragonBones"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"runtime"
	"unsafe"
)

type DragonBoneFactoryFace interface {
	wrapper.BaseFactory

	deleteFactory()
	IsFactory()
}

type DragonBoneFactory struct {
	wrapper.BaseFactory

	basePath string
}

func (s *DragonBoneFactory) deleteFactory() {
	wrapper.DeleteDirectorBaseFactory(s.BaseFactory)
}

func (s *DragonBoneFactory) IsFactory() {}

func NewFactory() *DragonBoneFactory {
	om := &overwrittenMethodsOnFactory{dragonBones: wrapper.NewDragonBones(NewArmatureDisplay())}

	factoryFace := wrapper.NewDirectorBaseFactory(om)
	// wrapper.DirectorBaseFactoryX_onClear(factoryFace)
	om.base = factoryFace

	factory := &DragonBoneFactory{BaseFactory: factoryFace}
	return factory
}

func (factory *DragonBoneFactory) LoadDragonBonesData(reader io.Reader, name string, scale float32) (wrapper.DragonBonesData, error) {
	bytes, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	runtime.SetFinalizer(&bytes, func(*[]uint8) { fmt.Printf("bytes final.\n") })
	data := factory.ParseDragonBonesData(*(*string)(unsafe.Pointer(&bytes)), name, scale)
	return data, nil
}

func (factory *DragonBoneFactory) LoadTextureAtlasData(reader io.Reader, name string, scale float32) {
	bytes, err := ioutil.ReadAll(reader)

	if err != nil {
		return
	}

	factory.ParseTextureAtlasData(string(bytes), uintptr(0), name, scale)
}

func (factory *DragonBoneFactory) BuildArmatureDisplay(args ...interface{}) wrapper.Armature {
	return factory.BuildArmature(args...)
}

type overwrittenMethodsOnFactory struct {
	base wrapper.BaseFactory

	dragonBones wrapper.DragonBones
}

func (om *overwrittenMethodsOnFactory) X_buildTextureAtlasData(data wrapper.TextureAtlasData, textureAtlas uintptr) wrapper.TextureAtlasData {
	log.Println("build texture", data.Swigcptr(), textureAtlas)
	if data.Swigcptr() == 0 {
		return NewTextureAtlasData()
	} else {
		log.Println(data.GetImagePath())
	}
	return data
}

func (om *overwrittenMethodsOnFactory) X_buildArmature(dataPackage wrapper.BuildArmaturePackage) wrapper.Armature {
	log.Println("BuildArmature")
	a := wrapper.BaseObjectBorrowArmatureObject()
	armatureDisplay := NewArmatureDisplay()
	boneObjectAdd(armatureDisplay.Swigcptr(), armatureDisplay)
	a.Init(dataPackage.GetArmature(), armatureDisplay, armatureDisplay.Swigcptr(), om.dragonBones)
	return a
}

func (om *overwrittenMethodsOnFactory) X_buildSlot(dataPackage wrapper.BuildArmaturePackage, slotData wrapper.DragonBones_SlotData, armature wrapper.Armature) wrapper.Slot {
	slot := NewSlot()
	slot.Init(slotData, armature, uintptr(0), uintptr(0))
	log.Println("BuildSlot", slot.Swigcptr())
	return slot
}
