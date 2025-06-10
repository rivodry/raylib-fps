package fps

import rl "github.com/gen2brain/raylib-go/raylib"

type Character struct {
	Camera       rl.Camera
	Pos          rl.Vector3
	Rotation     rl.Vector3
	GroundYLevel int32
}

func (ch *Character) UpdateCamera() {

	rl.UpdateCameraPro(&ch.Camera, rl.Vector3{float32(iskeydown(rl.KeyW)*0.1) - float32(iskeydown(rl.KeyS)*0.1), float32(iskeydown(rl.KeyD)*0.1) - float32(iskeydown(rl.KeyA)*0.1), float32(jump(ch.Camera.Position.Y))}, rl.Vector3{(rl.GetMouseDelta().X * 0.05), (rl.GetMouseDelta().Y * 0.05), 0}, 0)
	if ch.Camera.Position.Y > float32(ch.GroundYLevel) {
		rl.UpdateCameraPro(&ch.Camera, rl.Vector3{float32(iskeydown(rl.KeyW)*0.1) - float32(iskeydown(rl.KeyS)*0.1), float32(iskeydown(rl.KeyD)*0.1) - float32(iskeydown(rl.KeyA)*0.1), -0.01}, rl.Vector3{(rl.GetMouseDelta().X * 0.05), (rl.GetMouseDelta().Y * 0.05), 0}, 0)
	}

}
func NewCharacter(defaultpos rl.Vector3, groundylevel int32) Character {
	ch := Character{}
	ch.Camera = rl.Camera3D{}
	ch.Camera.Position = defaultpos
	ch.Camera.Target = rl.Vector3{0, defaultpos.Y, 0}
	ch.Camera.Up = rl.Vector3{0, 1, 0}
	ch.Camera.Fovy = 70
	ch.Camera.Projection = rl.CameraPerspective
	ch.GroundYLevel = groundylevel
	return ch
}
func iskeydown(key int32) float32 {
	if rl.IsKeyDown(key) {
		return 1
	} else {
		return 0
	}
}
func jump(y float32) float32 {

	if rl.IsKeyDown(rl.KeySpace) && !(y > 2) {

		return 3
	} else {

		return 0
	}
}
