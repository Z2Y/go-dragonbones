package bone

import (
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/EngoEngine/gl"
)

type IDisplay interface {
	GetParent() IDisplay
	SetParent(IDisplay)
	AddChild(IDisplay)
	RemoveChild(IDisplay)
	GetChildren() []IDisplay
	RemoveFromParent()
	UpdateTransform(force bool)
	SetTransform(*engo.Matrix)
	GetGlobalTransform() *engo.Matrix

	Texture() *gl.Texture
	Width() float32
	Height() float32
	View() (float32, float32, float32, float32)
	Close()
}

type Display struct {
	Parent   IDisplay
	Children []IDisplay

	common.RenderComponent
	common.SpaceComponent
	transformMatrix *engo.Matrix
	globalTransform *engo.Matrix
	transformDirty  bool
}

func (d *Display) GetParent() IDisplay {
	return d.Parent
}

func (d *Display) SetParent(p IDisplay) {
	d.Parent = p
}

func (d *Display) AddChild(child IDisplay) {
	child.SetParent(d)
	d.Children = append(d.Children, child)
}

func (d *Display) RemoveChild(child IDisplay) {
	delete := -1
	for idx, c := range d.Children {
		if c == child {
			delete = idx
			break
		}
	}
	d.Children = append(d.Children[:delete], d.Children[delete+1:]...)
}

func (d *Display) GetChildren() []IDisplay {
	return d.Children
}

func (d *Display) RemoveFromParent() {
	if d.Parent != nil {
		d.Parent.RemoveChild(d)
		d.Parent = nil
	}
}

func (d *Display) SetTransform(transformMatrix *engo.Matrix) {
	// log.Println("SetTransform", transformMatrix)
	d.transformMatrix = transformMatrix
	d.transformDirty = true
}

func (d *Display) GetGlobalTransform() *engo.Matrix {
	return d.globalTransform
}

func (d *Display) UpdateTransform(force bool) {
	if d.Parent == nil {
		d.globalTransform = engo.IdentityMatrix().Translate(d.Position.X, d.Position.Y).Rotate(d.Rotation)
	} else if d.transformMatrix != nil && (d.transformDirty || force) {
		d.globalTransform = engo.IdentityMatrix().Set(d.Parent.GetGlobalTransform().Val[:]).Multiply(d.transformMatrix)

		d.SpaceComponent.Position.X, d.SpaceComponent.Position.Y = d.globalTransform.TranslationComponent()
		d.SpaceComponent.Rotation = d.globalTransform.RotationComponent()
	}

	for _, child := range d.Children {
		child.UpdateTransform(d.transformDirty || force)
	}
	d.transformDirty = false
}

func (d *Display) Width() float32 {
	return 0
}

func (d *Display) Height() float32 {
	return 0
}

func (d *Display) Texture() *gl.Texture {
	return nil
}

func (d *Display) View() (float32, float32, float32, float32) {
	return 0, 0, 1, 1
}

func (d *Display) Close() {
}
