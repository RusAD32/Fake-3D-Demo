package Pseudo3D

import "math"

func wp_to_sp(x, y, z float64, cam *Camera) (float64, float64, float64, float64) {
	xTrans := x - cam.X - cam.Width/2
	yTrans := y - cam.Y - cam.Height/2
	zTrans := z - cam.Z

	// rotation
	xRot := xTrans*cam.m00 + yTrans*cam.m10
	yRot := xTrans*cam.m01 + yTrans*cam.m11

	// pitch
	tmp := yRot
	yRot = -tmp*math.Sin(cam.Pitch) - zTrans*math.Cos(cam.Pitch)
	zRot := -tmp*math.Cos(cam.Pitch) + zTrans*math.Sin(cam.Pitch)

	scale := zRot / 100.0
	scale += (1 - scale) * cam.FOV
	scale = cam.Zoom / scale

	resX := xRot * scale
	resY := yRot * scale
	return resX + cam.Width/2, resY + cam.Height/2, zRot, scale
}
