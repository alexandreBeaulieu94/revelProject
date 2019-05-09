package board

func (b *board) Grids() map[Position]TileGrid {
    return b.grids
}

func (b *board) GridAt(pos Position) TileGrid {
    return b.grids[pos]
}

func (b *board) CurrentGrid() TileGrid {
    return b.grids[b.pos]
}

func (b *board) GetCurrentPos() Position {
    return b.pos
}

func (b *board) SetCurrentPos(pos Position) {
    b.prevPos = b.pos
    b.pos = pos
}

func (b *board) GetPreviousPos() Position {
    return b.prevPos
}

func (b *board) ResetAll() {
    for _, grid := range b.grids {
        grid.Reset()
    }
}


