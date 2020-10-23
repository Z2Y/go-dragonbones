package bone

import "C"

import (
	wrapper "dragonBones/dragonBones"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
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

func (factory *DragonBoneFactory) LoadDragonBonesData(reader io.Reader, name string, scale float32) (*DragonBonesData, error) {
	bytes, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	binary := *(*string)(unsafe.Pointer(&bytes))
	data := factory.ParseDragonBonesData(binary, name, scale)
	dragonBonesData := NewDragonBonesData(data, &binary)
	return dragonBonesData, nil
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
		factory.dragonBonesInstance().GetClock().AddArmature(armature)
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
	a := wrapper.BaseObjectBorrowArmatureObject()
	armatureDisplay := NewArmatureDisplay()
	a.Init(dataPackage.GetArmature(), armatureDisplay, armatureDisplay.Swigcptr(), om.dragonBones)
	return a
}

func (om *overwrittenMethodsOnFactory) X_buildSlot(dataPackage wrapper.BuildArmaturePackage, slotData wrapper.SlotData, armature wrapper.Armature) wrapper.Slot {
	slot := NewSlot()
	sprite := NewSprite()
	meshSprite := NewSprite()
	boneObjectAdd(uintptr(unsafe.Pointer(sprite)), sprite)
	boneObjectAdd(uintptr(unsafe.Pointer(meshSprite)), meshSprite)
	slot.Init(slotData, armature, uintptr(unsafe.Pointer(sprite)), uintptr(unsafe.Pointer(meshSprite)))
	return slot
}
