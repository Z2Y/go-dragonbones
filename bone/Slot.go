package bone

import (
	wrapper "dragonBones/dragonBones"
	"log"
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
	textureDataImpl := boneObjectLookup(textureData.Swigcptr()).(*TextureDataImpl)
	switch frameDisplay := om.renderDisplay.(type) {
	case *Sprite:
		frameDisplay.setSpriteFrame(textureDataImpl.TextureResource)
	}
	log.Println("UpdateFrame:", textureData)
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
	transform := om.slot.GetGlobal()
	pivotX := om.slot.GetX_pivotX()
	pivotY := om.slot.GetX_pivotY()
	globalTransformMatrix := om.slot.GetGlobalTransformMatrix()

	x := transform.GetX() - (globalTransformMatrix.GetA()*pivotX + globalTransformMatrix.GetC()*pivotY)
	y := transform.GetY() - (globalTransformMatrix.GetB()*pivotX + globalTransformMatrix.GetD()*pivotY)

	switch frameDisplay := om.renderDisplay.(type) {
	case *Sprite:
		frameDisplay.SpaceComponent.Position.X = x
		frameDisplay.SpaceComponent.Position.Y = y
		log.Println("Update Sprite Transform", frameDisplay.SpaceComponent.Position)
	case *ArmatureDisplay:
		log.Println("Update ArmatureDisplay Transform")
	}
}
