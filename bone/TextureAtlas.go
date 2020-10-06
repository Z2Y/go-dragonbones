package bone

import (
	wrapper "dragonBones/dragonBones"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type TextureAtlasDataFace interface {
	wrapper.TextureAtlasData

	deleteTextureAtlas()
	IsTextureAtlas()
}

type DragonBonesData struct {
	wrapper.DragonBonesData
	binary *string
}

func NewDragonBonesData(data wrapper.DragonBonesData, binary *string) *DragonBonesData {
	dragonBonesData := &DragonBonesData{DragonBonesData: data}
	if len(*binary) > 4 &&
		(*binary)[0] == 'D' ||
		(*binary)[1] == 'B' ||
		(*binary)[2] == 'D' ||
		(*binary)[3] == 'T' {
		dragonBonesData.binary = binary
	}
	boneObjectAdd(data.Swigcptr(), dragonBonesData)
	return dragonBonesData
}

type TextureAtlasDataImpl struct {
	wrapper.TextureAtlasData

	TextureResource *common.TextureResource
}

func (tex *TextureAtlasDataImpl) IsTextureAtlas() {}

func (tex *TextureAtlasDataImpl) deleteTextureAtlas() {
	wrapper.DeleteDirectorTextureAtlasData(tex.TextureAtlasData)
}

func (tex *TextureAtlasDataImpl) setRenderTexture(texture *common.TextureResource) {
	if tex.TextureResource == texture {
		return
	}
	tex.TextureResource = texture

	if tex.TextureResource == nil {
		textures := tex.GetTextures().Values()
		for i := 0; i < int(textures.Size()); i++ {
			texture := textures.Get(i)
			textureImpl := boneObjectLookup(texture.Swigcptr()).(*TextureDataImpl)
			textureImpl.TextureResource = nil
		}
	} else {
		textures := tex.GetTextures().Values()
		for i := 0; i < int(textures.Size()); i++ {
			texture := textures.Get(i)
			region := texture.GetRegion()
			rotated := texture.GetRotated()
			x, y, width, height := region.GetX(), region.GetY(), region.GetWidth(), region.GetHeight()
			textureImpl := boneObjectLookup(texture.Swigcptr()).(*TextureDataImpl)
			if rotated {
				width, height = height, width
			}
			log.Println("SubTextureRegion")
			viewport := engo.AABB{
				Min: engo.Point{
					X: x / tex.TextureResource.Width,
					Y: y / tex.TextureResource.Height,
				},
				Max: engo.Point{
					X: (x + width) / tex.TextureResource.Width,
					Y: (y + height) / tex.TextureResource.Height,
				},
			}
			textureImpl.Rotated = rotated
			textureImpl.Parent = tex
			textureImpl.TextureResource = &common.TextureResource{Texture: tex.TextureResource.Texture, Width: width, Height: height, Viewport: &viewport}
		}
	}
}

func NewTextureAtlasData() *TextureAtlasDataImpl {
	om := &overwrittenMethodsOnTextureAtlasData{}

	face := wrapper.NewDirectorTextureAtlasData(om)
	wrapper.DirectorTextureAtlasDataX_onClear(face)
	om.base = face

	data := &TextureAtlasDataImpl{TextureAtlasData: face}
	boneObjectAdd(data.Swigcptr(), data)
	return data
}

type overwrittenMethodsOnTextureAtlasData struct {
	base wrapper.TextureAtlasData
}

func (om *overwrittenMethodsOnTextureAtlasData) GetClassTypeIndex() int64 {
	return wrapper.GetTextureAtlasDataTypeIndex(om.base)
}

func (om *overwrittenMethodsOnTextureAtlasData) CreateTexture() wrapper.TextureData {
	return NewTextureData()
}

type TextureDataFace interface {
	wrapper.TextureData

	deleteTextureData()
	IsTextureData()
}

type TextureDataImpl struct {
	wrapper.TextureData

	TextureResource *common.TextureResource
	Parent          *TextureAtlasDataImpl
	Rotated         bool
}

func (tex *TextureDataImpl) IsTextureData() {}

func (tex *TextureDataImpl) deleteTextureData() {
	wrapper.DeleteDirectorTextureData(tex.TextureData)
}

func NewTextureData() *TextureDataImpl {
	om := &overwrittenMethodsOnTextureData{}

	face := wrapper.NewDirectorTextureData(om)
	om.base = face

	data := &TextureDataImpl{TextureData: face}
	boneObjectAdd(data.Swigcptr(), data)

	return data
}

func LoadTextureAtlas(imagePath string) (*common.TextureResource, error) {
	// todo
	log.Println("LoadTextureAtlas", imagePath)
	f, err := os.Open(imagePath)

	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()

	imageObject := common.NewImageObject(common.ImageToNRGBA(img, bounds.Dx(), bounds.Dy()))
	texture := common.NewTextureResource(imageObject)

	return &texture, nil
}

type overwrittenMethodsOnTextureData struct {
	base wrapper.TextureData
}
