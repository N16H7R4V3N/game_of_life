package graphics

import "game_of_life/math"

type Graphics struct {
	Screen       Screen
	CellScale    float32              // scalar to shrink or enlarge
	TileGridSize math.Vector2DFloat32 // tile's size inside the grid
	TileMargin   math.Vector2DFloat32
	TileOffset   math.Vector2DFloat32
	TileSize     math.Vector2DFloat32 // tile's size after applying the cell's scale
}

func NewGraphics(screenWidth, screenHeight int32, cellSize float32) *Graphics {
	var g = Graphics{}
	g.Screen.Width = screenWidth
	g.Screen.Height = screenHeight
	g.CellScale = cellSize
	return &g
}

func (o *Graphics) CalcTileSize(gridX, gridY int32) {
	o.TileGridSize.X = float32(o.Screen.Width) / float32(gridX)
	o.TileGridSize.Y = float32(o.Screen.Height) / float32(gridY)
	o.TileMargin.X = o.TileGridSize.X * (1 - o.CellScale) / 2
	o.TileMargin.Y = o.TileGridSize.Y * (1 - o.CellScale) / 2
	o.TileOffset.X = o.TileGridSize.X + o.TileMargin.X
	o.TileOffset.Y = o.TileGridSize.Y + o.TileMargin.Y
	o.TileSize.X = o.TileGridSize.X - 2*o.TileMargin.X
	o.TileSize.Y = o.TileGridSize.Y - 2*o.TileMargin.Y
}

type Screen struct {
	Width  int32
	Height int32
}
