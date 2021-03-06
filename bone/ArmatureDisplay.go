package bone

import (
	wrapper "dragonBones/dragonBones"
	"log"
)

type ArmatureDisplayFace interface {
	wrapper.IArmatureProxy
	deleteIArmatureProxy()
	IsIArmatureProxy()
}

type ArmatureDisplay struct {
	Display
	wrapper.IArmatureProxy
	om *overwrittenMethodsOnArmatureDisplay
}

func (a *ArmatureDisplay) IsIArmatureProxy() {}

func (a *ArmatureDisplay) deleteIArmatureProxy() {
	wrapper.DeleteDirectorIArmatureProxy(a.IArmatureProxy)
}

func NewArmatureDisplay() *ArmatureDisplay {
	om := &overwrittenMethodsOnArmatureDisplay{}

	face := wrapper.NewDirectorIArmatureProxy(om)
	om.base = face

	display := &ArmatureDisplay{IArmatureProxy: face, om: om}
	display.RenderComponent.Drawable = display
	boneObjectAdd(display.Swigcptr(), display)
	return display
}

type overwrittenMethodsOnArmatureDisplay struct {
	base wrapper.IArmatureProxy

	armature wrapper.Armature
}

func (om *overwrittenMethodsOnArmatureDisplay) DbInit(armature wrapper.Armature) {
	log.Println("DbInit")
	om.armature = armature
}

func (om *overwrittenMethodsOnArmatureDisplay) DbClear() {
	om.armature = nil
}

func (om *overwrittenMethodsOnArmatureDisplay) DbUpdate() {
	// log.Println("DbUpdate")
}

func (om *overwrittenMethodsOnArmatureDisplay) Dispose(bool) {
	if om.armature != nil {
		om.armature.Dispose()
		om.armature = nil
	}
}

func (om *overwrittenMethodsOnArmatureDisplay) GetArmature() wrapper.Armature {
	return om.armature
}

func (om *overwrittenMethodsOnArmatureDisplay) HasDBEventListener(name string) bool {
	return false
}

func (om *overwrittenMethodsOnArmatureDisplay) GetAnimation() wrapper.Animation {
	return om.armature.GetAnimation()
}
