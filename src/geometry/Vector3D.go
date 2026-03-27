package geometry

type Vector3D struct {
	x, y, z float32
}

func getMidPoint(v1 Vector3D, v2 Vector3D) Vector3D {
	return Vector3D{(v1.x + v2.x) / 2, (v1.y + v2.y) / 2, (v1.z + v2.z) / 2}
}