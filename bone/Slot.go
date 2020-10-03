package bone

import (
	wrapper "dragonBones/dragonBones"
	"log"

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
	textureData := wrapper.SwigcptrSlot(om.slot.Swigcptr()).GetX_textureData()
	if textureData.Swigcptr() == 0 || om.renderDisplay == nil {
		return
	}
	om.textureScale = textureData.GetParent().GetScale() * om.slot.GetX_armature().GetX_armatureData().GetScale()
	textureDataImpl := boneObjectLookup(textureData.Swigcptr()).(*TextureDataImpl)
	switch frameDisplay := om.renderDisplay.(type) {
	case *Sprite:
		frameDisplay.setSpriteFrame(textureDataImpl.TextureResource)
	}
}

func (om *overwrittenMethodsOnSlot) X_updateMesh() {

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
	case *ArmatureDisplay:
		transformMatrix.Val[6] = tx
		transformMatrix.Val[7] = ty
		frameDisplay.SetTransform(transformMatrix)
	}
}
