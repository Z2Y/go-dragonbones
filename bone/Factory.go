package bone

import (
	wrapper "dragonBones/dragonBones"
	"io"
	"io/ioutil"
	"log"
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
	om := &overwrittenMethodsOnFactory{}

	factoryFace := wrapper.NewDirectorBaseFactory(om)
	om.base = factoryFace

	factory := &DragonBoneFactory{BaseFactory: factoryFace}
	return factory
}

func (factory *DragonBoneFactory) LoadDragonBonesData(reader io.Reader, name string, scale float32) (wrapper.DragonBonesData, error) {
	bytes, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}
	data := factory.ParseDragonBonesData(string(bytes), name, scale)
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
	a.Init(dataPackage.GetArmature(), armatureDisplay, armatureDisplay.Swigcptr(), wrapper.SwigcptrDragonBones(0))
	return a
}

func (om *overwrittenMethodsOnFactory) X_buildSlot(dataPackage wrapper.BuildArmaturePackage, slotData wrapper.DragonBones_SlotData, armature wrapper.Armature) wrapper.Slot {
	slot := NewSlot()
	slot.Init(slotData, armature, uintptr(0), uintptr(0))
	log.Println("BuildSlot", slot.Swigcptr())
	return slot
}
