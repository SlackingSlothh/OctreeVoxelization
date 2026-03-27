package geometry

type Cube struct {
	v0, v1 Vector3D
}

func divideCube(parent Cube) []Cube {
	mid := getMidPoint(parent.v0, parent.v1)

	return []Cube{
		{parent.v0, mid},
		{Vector3D{mid.x, parent.v0.y, parent.v0.z}, Vector3D{parent.v1.x, mid.y, mid.z}},
		{Vector3D{parent.v0.x, mid.y, parent.v0.z}, Vector3D{mid.x, parent.v1.y, mid.z}},
		{Vector3D{mid.x, mid.y, parent.v0.z}, Vector3D{parent.v1.x, parent.v1.y, mid.z}},
		{Vector3D{parent.v0.x, parent.v0.y, mid.z}, Vector3D{mid.x, mid.y, parent.v1.z}},
		{Vector3D{mid.x, parent.v0.y, mid.z}, Vector3D{parent.v1.x, mid.y, parent.v1.z}},
		{Vector3D{parent.v0.x, mid.y, mid.z}, Vector3D{mid.x, parent.v1.y, parent.v1.z}},
		{Vector3D{mid.x, mid.y, mid.z}, parent.v1},
	}
}