package bone

import (
	wrapper "dragonBones/dragonBones"
)

type TextureAtlasDataFace interface {
	wrapper.TextureAtlasData

	deleteTextureAtlas()
	IsTextureAtlas()
}

type TextureAtlasDataImpl struct {
	wrapper.TextureAtlasData
}

func (tex *TextureAtlasDataImpl) IsTextureAtlas() {}

func (tex *TextureAtlasDataImpl) deleteTextureAtlas() {
	wrapper.DeleteDirectorTextureAtlasData(tex.TextureAtlasData)
}

func NewTextureAtlasData() *TextureAtlasDataImpl {
	om := &overwrittenMethodsOnTextureAtlasData{}

	face := wrapper.NewDirectorTextureAtlasData(om)
	wrapper.DirectorTextureAtlasDataX_onClear(face)
	om.base = face

	return &TextureAtlasDataImpl{TextureAtlasData: face}
}

type overwrittenMethodsOnTextureAtlasData struct {
	base wrapper.TextureAtlasData
}

func (om *overwrittenMethodsOnTextureAtlasData) GetClassTypeIndex() wrapper.Std_size_t {
	return wrapper.GetTextureAtlasDataTypeIndex(om.base)
}

func (om *overwrittenMethodsOnTextureAtlasData) CreateTexture() wrapper.TextureData {
	return NewTextureData().TextureData
}

type TextureDataFace interface {
	wrapper.TextureData

	deleteTextureData()
	IsTextureData()
}

type TextureDataImpl struct {
	wrapper.TextureData
}

func (tex *TextureDataImpl) IsTextureData() {}

func (tex *TextureDataImpl) deleteTextureData() {
	wrapper.DeleteDirectorTextureData(tex.TextureData)
}

func NewTextureData() *TextureDataImpl {
	om := &overwrittenMethodsOnTextureData{}

	face := wrapper.NewDirectorTextureData(om)
	om.base = face

	return &TextureDataImpl{TextureData: face}
}

type overwrittenMethodsOnTextureData struct {
	base wrapper.TextureData
}
