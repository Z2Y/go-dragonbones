package bone

import (
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/EngoEngine/engo/math"
	"github.com/EngoEngine/gl"
)

type IDisplay interface {
	GetParent() IDisplay
	SetParent(IDisplay)
	AddChild(IDisplay)
	RemoveChild(IDisplay)
	GetChildIndex(IDisplay) int
	ReplaceChildAt(IDisplay, int)
	GetChildren() []IDisplay
	RemoveFromParent()
	UpdateTransform(force bool)
	SetTransform(*engo.Matrix)
	GetGlobalTransform() *engo.Matrix
	SetVisible(bool)

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

func (d *Display) SetVisible(v bool) {
	d.Hidden = !v
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

func (d *Display) GetChildIndex(child IDisplay) int {
	for id, c := range d.Children {
		if c == child {
			return id
		}
	}
	return -1
}

func (d *Display) ReplaceChildAt(child IDisplay, index int) {
	prev := d.Children[index]
	child.SetParent(d)
	prev.SetParent(nil)
	d.Children[index] = child
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

func (d *Display) restoreTransform() {
	ta, tb, tc, td := d.globalTransform.Val[0], d.globalTransform.Val[1], d.globalTransform.Val[3], d.globalTransform.Val[4]
	tx, ty := d.globalTransform.Val[6], d.globalTransform.Val[7]

	determ := (ta * td) - (tb * tc)

	if ta != 0 || tb != 0 {
		r := math.Sqrt((ta * ta) + (tb * tb))

		if tb > 0 {
			d.SpaceComponent.Rotation = math.Acos(ta/r) * engo.RadToDeg
		} else {
			d.SpaceComponent.Rotation = -math.Acos(ta/r) * engo.RadToDeg
		}

		d.RenderComponent.Scale.X = r
		d.RenderComponent.Scale.Y = determ / r
	} else if tc != 0 || td != 0 {
		s := math.Sqrt((tc * tc) + (td * td))
		if td > 0 {
			d.SpaceComponent.Rotation = (math.Pi/2 - math.Acos(-tc/s)) * engo.RadToDeg
		} else {
			d.SpaceComponent.Rotation = (math.Pi/2 + math.Acos(tc/s)) * engo.RadToDeg
		}
		d.RenderComponent.Scale.X = determ / s
		d.RenderComponent.Scale.Y = s
	} else {
		d.RenderComponent.Scale.X = 0
		d.RenderComponent.Scale.Y = 0
	}

	d.SpaceComponent.Position.X, d.SpaceComponent.Position.Y = tx, ty
}

func (d *Display) UpdateTransform(force bool) {
	if d.Parent == nil {
		d.globalTransform = engo.IdentityMatrix().Translate(d.Position.X, d.Position.Y).Rotate(d.Rotation)
	} else if d.transformMatrix != nil && (d.transformDirty || force) {
		d.globalTransform = engo.IdentityMatrix().Set(d.Parent.GetGlobalTransform().Val[:]).Multiply(d.transformMatrix)

		d.restoreTransform()
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
