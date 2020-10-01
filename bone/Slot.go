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
	log.Println("Update Display", display)
}

func (om *overwrittenMethodsOnSlot) X_addDisplay() {
	display := boneObjectLookup(om.slot.GetX_armature().GetDisplay()).(*ArmatureDisplay)
	log.Println("Add Display", display)
}

func (om *overwrittenMethodsOnSlot) X_replaceDisplay(value uintptr, isArmatureDisplay bool) {

}

func (om *overwrittenMethodsOnSlot) X_updateZOrder() {

}

func (om *overwrittenMethodsOnSlot) X_updateFrame() {

}

func (om *overwrittenMethodsOnSlot) X_updateMesh() {

}

func (om *overwrittenMethodsOnSlot) X_updateVisible() {

}

func (om *overwrittenMethodsOnSlot) X_updateBlendMode() {

}

func (om *overwrittenMethodsOnSlot) X_updateColor() {

}

func (om *overwrittenMethodsOnSlot) X_updateTransform() {

}
