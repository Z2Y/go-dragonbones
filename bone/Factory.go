package bone

import "C"

import (
	wrapper "dragonBones/dragonBones"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"unsafe"
)

var (
	Factory = NewFactory()
)

type DragonBoneFactoryFace interface {
	wrapper.BaseFactory

	deleteFactory()
	IsFactory()
}

type DragonBoneFactory struct {
	wrapper.BaseFactory
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

func (factory *DragonBoneFactory) SetAssetPath(base string) {
	factory.DirectorInterface().(*overwrittenMethodsOnFactory).basePath = base
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

func (factory *DragonBoneFactory) BuildArmatureDisplay(args ...interface{}) *ArmatureDisplay {
	armature := factory.BuildArmature(args...)
	if armature.Swigcptr() != 0 {
		factory.dragonBonesInstance().GetClock().Add(armature)
		return boneObjectLookup(armature.GetDisplay()).(*ArmatureDisplay)
	}
	return nil
}

func (factory *DragonBoneFactory) dragonBonesInstance() wrapper.DragonBones {
	return factory.DirectorInterface().(*overwrittenMethodsOnFactory).dragonBones
}

func (factory *DragonBoneFactory) Update(dt float32) {
	factory.dragonBonesInstance().AdvanceTime(dt)
}

type overwrittenMethodsOnFactory struct {
	base wrapper.BaseFactory

	dragonBones wrapper.DragonBones
	basePath    string
}

func (om *overwrittenMethodsOnFactory) X_buildTextureAtlasData(data wrapper.TextureAtlasData, textureAtlas uintptr) wrapper.TextureAtlasData {
	log.Println("build texture", data.Swigcptr(), textureAtlas)
	if data.Swigcptr() == 0 {
		textureAtlasData := NewTextureAtlasData()
		return textureAtlasData
	} else {
		textureAtlasData := boneObjectLookup(data.Swigcptr()).(*TextureAtlasDataImpl)
		texture, err := LoadTextureAtlas(filepath.Join(om.basePath, textureAtlasData.GetImagePath()))
		if err != nil {
			log.Println("TextureLoaded Error", err)
		} else {
			textureAtlasData.setRenderTexture(texture)
		}
	}
	return data
}

func (om *overwrittenMethodsOnFactory) X_buildArmature(dataPackage wrapper.BuildArmaturePackage) wrapper.Armature {
	log.Println("BuildArmature")
	a := wrapper.BaseObjectBorrowArmatureObject()
	armatureDisplay := NewArmatureDisplay()
	a.Init(dataPackage.GetArmature(), armatureDisplay, armatureDisplay.Swigcptr(), om.dragonBones)
	return a
}

func (om *overwrittenMethodsOnFactory) X_buildSlot(dataPackage wrapper.BuildArmaturePackage, slotData wrapper.DragonBones_SlotData, armature wrapper.Armature) wrapper.Slot {
	slot := NewSlot()
	sprite := NewSprite()
	boneObjectAdd(uintptr(unsafe.Pointer(sprite)), sprite)
	slot.Init(slotData, armature, uintptr(unsafe.Pointer(sprite)), uintptr(unsafe.Pointer(sprite)))
	log.Println("BuildSlot", slot.Swigcptr())
	return slot
}
