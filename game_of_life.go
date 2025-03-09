package main

import (
	"game_of_life/graphics"
	"game_of_life/grid"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawGrid(grid grid.Grid, graphics graphics.Graphics) {
	for x := int32(0); x < grid.Size.X; x++ {
		for y := int32(0); y < grid.Size.Y; y++ {
			if grid.Cells[x+y*grid.Size.X] == 0 {
				continue
			}
			offsetX := int32(float32(x) * graphics.TileOffset.X)
			offsetY := int32(float32(y) * graphics.TileOffset.Y)
			rl.DrawRectangle(offsetX, offsetY, int32(graphics.TileSize.X), int32(graphics.TileSize.Y), rl.LightGray)
		}
	}
}

func main() {
	var grd = grid.NewGrid(360, 270)
	var gfx = graphics.NewGraphics(1200, 900)
	gfx.CellScale = 0.9

	grd.Randomize()
	gfx.CalcTileSize(grd.Size.X, grd.Size.Y)

	rl.InitWindow(gfx.Screen.Width, gfx.Screen.Height, "Game of Life")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		grd.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		drawGrid(*grd, *gfx)
		rl.EndDrawing()
	}
}
