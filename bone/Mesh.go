package bone

type MeshSprite struct {
	Sprite

	drawMode  int
	blendMode int

	uvs      []float32
	indices  []uint16
	vertices []float32
}

func NewMeshSprite() *MeshSprite {
	mesh := MeshSprite{}
	return &mesh
}
