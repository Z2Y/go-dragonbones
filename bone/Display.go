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
	wrapper.IArmatureProxy
}

func (a *ArmatureDisplay) IsIArmatureProxy() {}

func (a *ArmatureDisplay) deleteIArmatureProxy() {
	wrapper.DeleteDirectorIArmatureProxy(a.IArmatureProxy)
}

func NewArmatureDisplay() *ArmatureDisplay {
	om := &overwrittenMethodsOnArmatureDisplay{}

	face := wrapper.NewDirectorIArmatureProxy(om)
	om.base = face

	return &ArmatureDisplay{IArmatureProxy: face}
}

type overwrittenMethodsOnArmatureDisplay struct {
	base wrapper.IArmatureProxy
}

func (om *overwrittenMethodsOnArmatureDisplay) DbInit(armature wrapper.Armature) {
	log.Println("DbInit")
}

func (om *overwrittenMethodsOnArmatureDisplay) DbUpdate() {
	log.Println("DbUpdate")
}

func (om *overwrittenMethodsOnArmatureDisplay) HasDBEventListener(name string) bool {
	return false
}
