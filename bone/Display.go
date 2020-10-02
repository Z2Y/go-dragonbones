package bone

import "github.com/EngoEngine/gl"

type IDisplay interface {
	GetParent() IDisplay
	SetParent(IDisplay)
	AddChild(IDisplay)
	RemoveChild(IDisplay)
	GetChildren() []IDisplay
	RemoveFromParent()

	Texture() *gl.Texture
	Width() float32
	Height() float32
	View() (float32, float32, float32, float32)
	Close()
}

type Display struct {
	Parent   IDisplay
	Children []IDisplay
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
