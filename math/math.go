package math

type Vector2DInt32 struct {
	X, Y int32
}

type Vector2DFloat32 struct {
	X, Y float32
}

func Mod(a, b int) int {
	return (a%b + b) % b
}
