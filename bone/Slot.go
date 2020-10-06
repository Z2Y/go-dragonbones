package bone

import (
	wrapper "dragonBones/dragonBones"
	"log"
	"unsafe"

	"github.com/EngoEngine/engo"
)

type SlotInterface interface {
	wrapper.Slot
	deleteSlot()
	IsSlot()
}

type Slot struct {
	wrapper.Slot
}

func (s *Slot) IsSlot() {}

func (s *Slot) deleteSlot() {
	wrapper.DeleteDirectorSlot(s.Slot)
}

func NewSlot() *Slot {
	om := &overwrittenMethodsOnSlot{}

	slotFace := wrapper.NewDirectorSlot(om)
	wrapper.DirectorSlotX_onClear(slotFace)
	om.slot = slotFace

	slot := &Slot{Slot: slotFace}

	return slot
}

func DeleteSlot(s SlotInterface) {
	s.deleteSlot()
}

type overwrittenMethodsOnSlot struct {
	slot wrapper.Slot

	textureScale  float32
	renderDisplay IDisplay
}

func (om *overwrittenMethodsOnSlot) X_initDisplay(value uintptr, isRetain bool) {
	log.Println("InitDisplay", value, isRetain)
}

func (om *overwrittenMethodsOnSlot) X_disposeDisplay(value uintptr, isRetain bool) {
	log.Println("Dispose Display", value, isRetain)
}

func (om *overwrittenMethodsOnSlot) X_onUpdateDisplay() {
	display := om.slot.GetDisplay()
	if display == 0 {
		display = om.slot.GetRawDisplay()
	}
	if display == 0 {
		return
	}

	switch _display := boneObjectLookup(display).(type) {
	case IDisplay:
		om.renderDisplay = _display
	}
	log.Println("Update Display", display)
}

func (om *overwrittenMethodsOnSlot) X_addDisplay() {
	display := boneObjectLookup(om.slot.GetX_armature().GetDisplay()).(*ArmatureDisplay)
	if om.renderDisplay != nil {
		display.AddChild(om.renderDisplay)
	}
	log.Println("Add Display", display)
}

func (om *overwrittenMethodsOnSlot) X_replaceDisplay(value uintptr, isArmatureDisplay bool) {
	display := boneObjectLookup(om.slot.GetX_armature().GetDisplay()).(*ArmatureDisplay)
	if om.renderDisplay != nil {
		display.AddChild(om.renderDisplay)
	}

	switch prevDisplay := boneObjectLookup(value).(type) {
	case IDisplay:
		display.RemoveChild(prevDisplay)
	}
	om.textureScale = 1.0
}

func (om *overwrittenMethodsOnSlot) X_removeDisplay() {
	if om.renderDisplay != nil {
		om.renderDisplay.RemoveFromParent()
	}
}

func (om *overwrittenMethodsOnSlot) X_updateZOrder() {

}

func (om *overwrittenMethodsOnSlot) X_updateFrame() {
	swigSlot := wrapper.SwigcptrSlot(om.slot.Swigcptr())
	textureData := swigSlot.GetX_textureData()

	displayIndex := swigSlot.GetX_displayIndex()
	display := swigSlot.GetX_display()

	if displayIndex >= 0 && display != uintptr(0) && textureData.Swigcptr() != 0 {
		deformVertices := om.slot.GetX_deformVertices()
		var verticesData wrapper.VerticesData
		if deformVertices.Swigcptr() != 0 {
			verticesData = deformVertices.GetVerticesData()
		}
		armatureScale := om.slot.GetX_armature().GetX_armatureData().GetScale()
		om.textureScale = textureData.GetParent().GetScale() * armatureScale
		textureDataImpl := boneObjectLookup(textureData.Swigcptr()).(*TextureDataImpl)
		switch frameDisplay := om.renderDisplay.(type) {
		case *Sprite:
			frameDisplay.setSpriteFrame(textureDataImpl.TextureResource)
		case *MeshSprite:
			if verticesData != nil {
				data := verticesData.GetData()
				offset := int(verticesData.GetOffset())
				intArray := data.GetIntArray()
				floatArray := data.GetFloatArray()
				weight := verticesData.GetWeight()

				vertexCount := int(*(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(intArray)) + uintptr(offset+0)*2)))
				triangleCount := int(*(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(intArray)) + uintptr(offset+1)*2)))
				vertexOffset := int(*(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(intArray)) + uintptr(offset+2)*2)))

				if vertexOffset < 0 {
					vertexOffset += 65536 // Fixed out of bouds bug.
				}

				log.Println(offset, vertexCount, triangleCount, *floatArray)
				uvOffset := vertexOffset + int(vertexCount)*2
				region := textureDataImpl.GetRegion()
				regionX, regionY := region.GetX(), region.GetY()
				regionWidth, regionHeight := region.GetWidth(), region.GetHeight()

				textureAtlasWidth := textureDataImpl.Parent.TextureResource.Width
				textureAtlasHeight := textureDataImpl.Parent.TextureResource.Height

				frameDisplay.vertices = make([]float32, vertexCount*2)
				frameDisplay.uvs = make([]float32, vertexCount*2)
				frameDisplay.indices = make([]uint16, triangleCount*3)

				for i := 0; i < len(frameDisplay.vertices); i++ {
					frameDisplay.vertices[i] = (*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(floatArray)) + uintptr(vertexOffset+i)*4))) * armatureScale
				}
				for i := 0; i < len(frameDisplay.indices); i++ {
					frameDisplay.indices[i] = *(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(intArray)) + uintptr(offset+4+i)*2))
				}

				for i := 0; i < len(frameDisplay.uvs); i += 2 {
					u := *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(floatArray)) + uintptr(uvOffset+i)*4))
					v := *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(floatArray)) + uintptr(uvOffset+i+1)*4))

					if textureDataImpl.Rotated {
						frameDisplay.uvs[i] = (regionX + (1.0-v)*regionWidth) / textureAtlasWidth
						frameDisplay.uvs[i+1] = (regionY + u*regionHeight) / textureAtlasHeight
					} else {
						frameDisplay.uvs[i] = (regionX + u*regionWidth) / textureAtlasWidth
						frameDisplay.uvs[i+1] = (regionY + v*regionHeight) / textureAtlasHeight
					}
				}
				frameDisplay.setSpriteFrame(textureDataImpl.TextureResource)
				if weight.Swigcptr() != 0 {
					frameDisplay.SetTransform(engo.IdentityMatrix())
				}
				log.Println(armatureScale)
				log.Println(regionX, regionY, regionWidth, regionHeight, textureAtlasWidth, textureAtlasHeight)
				log.Println("Mesh:", frameDisplay.uvs, frameDisplay.indices, frameDisplay.vertices)
				// panic("Mesh")
			}
		}
	} else {
		switch frameDisplay := om.renderDisplay.(type) {
		case *Sprite:
			frameDisplay.spriteFrame = nil
			frameDisplay.Position = engo.Point{X: 0, Y: 0}
			frameDisplay.Hidden = true
		case *MeshSprite:
			frameDisplay.spriteFrame = nil
			frameDisplay.Position = engo.Point{X: 0, Y: 0}
			frameDisplay.Hidden = true
		}
	}

}

func (om *overwrittenMethodsOnSlot) X_updateMesh() {
	dv := om.slot.GetX_deformVertices()

	deformVertices := dv.GetVertices()
	bones := dv.GetBones()
	verticesData := dv.GetVerticesData()

	weightData := verticesData.GetWeight()
	scale := om.slot.GetX_armature().GetX_armatureData().GetScale()

	hasFFD := !deformVertices.IsEmpty()

	// textureData := wrapper.SwigcptrSlot(om.slot.Swigcptr()).GetX_textureData()
	// textureDataImpl := boneObjectLookup(textureData.Swigcptr()).(*TextureDataImpl)

	meshDisplay := om.renderDisplay.(*MeshSprite)

	if weightData.Swigcptr() != 0 {
		data := verticesData.GetData()
		offset := int(verticesData.GetOffset())
		weightOffset := int(weightData.GetOffset())
		intArray := data.GetIntArray()
		floatArray := data.GetFloatArray()
		bonesSize := int(bones.Size())

		vertexCount := int(*(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(intArray)) + uintptr(offset+0)*2)))
		weightFloatOffset := int(*(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(intArray)) + uintptr(weightOffset+1)*2)))

		if weightFloatOffset < 0 {
			weightFloatOffset += 65536 // Fixed out of bouds bug.
		}

		iB := weightOffset + 2 + bonesSize
		iV := weightFloatOffset
		iF := 0
		iD := 0
		for i := 0; i < vertexCount; i++ {
			boneCount := int(*(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(intArray)) + uintptr(iB)*2)))
			iB += 1
			xG, yG := float32(0.0), float32(0.0)
			for j := 0; j < boneCount; j++ {
				boneIndex := int(*(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(intArray)) + uintptr(iB)*2)))
				iB += 1
				bone := bones.Get(boneIndex)
				if bone.Swigcptr() != 0 {
					matrix := bone.GetGlobalTransformMatrix()
					weight := *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(floatArray)) + uintptr(iV)*4))
					iV += 1

					xL := *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(floatArray)) + uintptr(iV)*4)) * scale
					iV += 1
					yL := *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(floatArray)) + uintptr(iV)*4)) * scale
					iV += 1

					if hasFFD {
						xL += deformVertices.Get(iF)
						iF += 1
						yL += deformVertices.Get(iF)
						iF += 1
					}
					a, b, c, d := matrix.GetA(), matrix.GetB(), matrix.GetC(), matrix.GetD()
					tx, ty := matrix.GetTx(), matrix.GetTy()
					xG += (a*xL + c*yL + tx) * weight
					yG += (b*xL + d*yL + ty) * weight
				}
			}
			meshDisplay.vertices[iD] = xG
			iD += 1
			meshDisplay.vertices[iD] = yG
			iD += 1
		}
	} else if hasFFD {
		data := verticesData.GetData()
		offset := int(verticesData.GetOffset())
		intArray := data.GetIntArray()
		floatArray := data.GetFloatArray()

		vertexCount := int(*(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(intArray)) + uintptr(offset+0)*2)))
		vertexOffset := int(*(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(intArray)) + uintptr(offset+2)*2)))

		if vertexOffset < 0 {
			vertexOffset += 65536 // Fixed out of bouds bug.
		}

		for i, l := 0, vertexCount*2; i < l; i += 2 {
			x := (*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(floatArray)) + uintptr(vertexOffset+i)*4))) * scale
			y := (*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(floatArray)) + uintptr(vertexOffset+i+1)*4))) * scale

			x += deformVertices.Get(i)
			y += deformVertices.Get(i + 1)

			meshDisplay.vertices[i] = x
			meshDisplay.vertices[i+1] = y
			/* if (isSurface) {
				var matrix = this._parent._getGlobalTransformMatrix(x, y);
				meshDisplay.vertices[i] = matrix.a * x + matrix.c * y + matrix.tx;
				meshDisplay.vertices[i + 1] = matrix.b * x + matrix.d * y + matrix.ty;
			}*/
		}
	}
}

func (om *overwrittenMethodsOnSlot) X_updateVisible() {

}

func (om *overwrittenMethodsOnSlot) X_updateBlendMode() {

}

func (om *overwrittenMethodsOnSlot) X_updateColor() {

}

func (om *overwrittenMethodsOnSlot) GetClassTypeIndex() int64 {
	return wrapper.GetSlotTypeIndex(om.slot)
}

func (om *overwrittenMethodsOnSlot) X_updateTransform() {
	om.slot.UpdateGlobalTransform()

	// rawDisplay := om.slot.GetRawDisplay()
	// transform := om.slot.GetGlobal()
	pivotX := om.slot.GetX_pivotX()
	pivotY := om.slot.GetX_pivotY()
	globalTransformMatrix := om.slot.GetGlobalTransformMatrix()
	a, b, c, d := globalTransformMatrix.GetA(), globalTransformMatrix.GetB(), globalTransformMatrix.GetC(), globalTransformMatrix.GetD()
	tx, ty := globalTransformMatrix.GetTx(), globalTransformMatrix.GetTy()

	transformMatrix := engo.IdentityMatrix()

	transformMatrix.Val[0] = a
	transformMatrix.Val[1] = b
	transformMatrix.Val[3] = c
	transformMatrix.Val[4] = d

	switch frameDisplay := om.renderDisplay.(type) {
	case *Sprite:
		if om.textureScale != 1.0 {
			transformMatrix.Val[0] *= om.textureScale
			transformMatrix.Val[1] *= om.textureScale
			transformMatrix.Val[4] *= om.textureScale
			transformMatrix.Val[3] *= om.textureScale
		}
		transformMatrix.Val[6] = tx - (a*pivotX + c*pivotY)
		transformMatrix.Val[7] = ty - (b*pivotX + d*pivotY)
		frameDisplay.SetTransform(transformMatrix)
	case *MeshSprite:
		if om.textureScale != 1.0 {
			transformMatrix.Val[0] *= om.textureScale
			transformMatrix.Val[1] *= om.textureScale
			transformMatrix.Val[4] *= om.textureScale
			transformMatrix.Val[3] *= om.textureScale
		}
		transformMatrix.Val[6] = tx - (a*pivotX + c*pivotY)
		transformMatrix.Val[7] = ty - (b*pivotX + d*pivotY)
		frameDisplay.SetTransform(transformMatrix)
		log.Println("Update Mesh Display Transform")
	case *ArmatureDisplay:
		transformMatrix.Val[6] = tx
		transformMatrix.Val[7] = ty
		frameDisplay.SetTransform(transformMatrix)
	}
}
