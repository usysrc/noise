package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "noise")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	var cols [800 * 600]rl.Color
	for i := 0; i < 800; i++ {
		for j := 0; j < 600; j++ {
			cols[j*600+i] = rl.Black
			if rand.Intn(100) < 50 {
				cols[j*600+i] = rl.White
			}
		}
	}
	t := 0
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		for i := 0; i < 800; i++ {
			for j := 0; j < 600; j++ {
				rl.DrawRectangle(int32(i), int32(j), 1, 1, cols[j*600+i])
			}
		}
		if t == 0 {
			rl.TakeScreenshot("screen.png")
		}
		rl.EndDrawing()
		t++
	}
}
