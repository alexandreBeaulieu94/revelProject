package board

import (
    . "github.com/alexandreBeaulieu94/awesomeProject2/grid"
    "github.com/alexandreBeaulieu94/awesomeProject2/grid/tile"
    . "github.com/alexandreBeaulieu94/awesomeProject2/position"
)

var DefaultPosition = MIDDLE_CENTER

type Board interface {
    Grids() map[Position]TileGrid
    CurrentGrid() TileGrid
    GridAt(pos Position) TileGrid
    GetCurrentPos() Position
    SetCurrentPos(pos Position)
    GetPreviousPos() Position
    ResetAll()
}

type board struct {
    grids   map[Position]TileGrid
    pos     Position
    prevPos Position
}

func NewBoard() Board {
    grids := make(map[Position]TileGrid)
    for i := 1; i <= 9; i++ {
        grids[Position(i)] = NewGrid(3, 3, &tile.Tile{Value: tile.EMPTY})
    }
    return &board{
        grids:   grids,
        pos:     DefaultPosition,
        prevPos: INVALID,

    }
}
