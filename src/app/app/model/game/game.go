package game

import (
    "fmt"
    "github.com/alexandreBeaulieu94/awesomeProject2/board"
    "github.com/alexandreBeaulieu94/awesomeProject2/game/player"
    "github.com/alexandreBeaulieu94/awesomeProject2/grid"
    . "github.com/alexandreBeaulieu94/awesomeProject2/grid/tile"
    . "github.com/alexandreBeaulieu94/awesomeProject2/position"
    "math/rand"
)

type Game interface {
    GetPlayerO() player.Player
    GetPlayerX() player.Player
    NextTurn() Game
    GetBoard() board.Board
    GetWinnerFromGrid(tiles grid.TileGrid) player.Player
    GetCurrentPlayer() player.Player
    Reset() Game
}

func NewGame(playerO player.Player, playerX player.Player, board board.Board) Game {
    if !playerO.IsCurrent() && !playerX.IsCurrent() ||
        playerO.IsCurrent() && playerX.IsCurrent() {
        rdm := rand.Float64() >= 0.5
        playerO.SetCurrent(rdm)
        playerX.SetCurrent(!rdm)
    }
    //listener := events.NewClickListener()
      g := &game{playerO, playerX, board}
   // listener.Listen(g.onClick)
   // listener.Resume()
   //
      return g
}

type game struct {
    playerO       player.Player
    playerX       player.Player
    board         board.Board
}

// Private
func (g *game) onClick() {
  /*  t, pos := g.board.GetTileUnderCursor()
    if t != nil && t.Value == EMPTY {
        if g.playerO.IsCurrent() {
            t.Value = O
        } else {
            t.Value = X
        }
        g.GetBoard().DrawTile(t, pos)
        g.GetBoard().SetCurrentPos(t.Position)
        g.NextTurn()
    }*/
}

func (g *game) GetPlayerO() player.Player {
    return g.playerO
}

func (g *game) GetPlayerX() player.Player {
    return g.playerX
}

func (g *game) GetBoard() board.Board {
    return g.board
}

func (g *game) GetWinnerFromGrid(tiles grid.TileGrid) player.Player {
    cells := tiles.GetWinningTiles()
    if cells[0] != nil {
        for _, cell := range cells {
            cell.Winning = true
        }

        if cells[0].Value == X {
            return g.playerX
        } else {
            return g.playerO
        }
    }
    return nil
}

func (g *game) GetWinningPos(tiles grid.TileGrid) Position {
    var winningTilePosition Position
    for col, columns := range tiles {
        for row := range columns {
            gridTempo := tiles.Clone()
            if gridTempo[row][col].Value == EMPTY {
                gridTempo[row][col].Value = X
                if g.GetWinnerFromGrid(gridTempo) == g.playerX {
                    winningTilePosition = gridTempo[row][col].Position
                    fmt.Println(winningTilePosition)
                }
            }
            gridTempo = nil
        }
    }

    return winningTilePosition
}

func (g *game) NextTurn() Game {
    g.GetWinnerFromGrid(g.board.CurrentGrid())
    g.playerX.SetCurrent(!g.playerX.IsCurrent())
    g.playerO.SetCurrent(!g.playerO.IsCurrent())

    if g.playerX.IsCurrent() {
        g.PlayAINextMove()
    }
    return g
}

func (g *game) Reset() Game {
    g.GetBoard().ResetAll()
    return g
}

func (g *game) GetCurrentPlayer() player.Player {
    cur := g.playerO
    if g.playerX.IsCurrent() {
        cur = g.playerX
    }
    return cur
}


/////////////////////////////////////////////////////////////////// AI


func (g *game) PlayAINextMove() {
    var posibility = g.GetPosibility()
    var choice = g.GetNextMove(posibility)
    g.GetBoard().CurrentGrid().At(choice).Value = X
}

func (g *game) GetPosibility() []Position {
    var posibility []Position
    x, y := g.GetBoard().GetCurrentPos().GetXY()
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if g.GetBoard().CurrentGrid()[i+x][j+y].Value == EMPTY {
                posibility = append(posibility, PositionAt(i+x, j+y))
            }
        }
    }

    return posibility
}

func (g *game) GetNextMove(choices []Position) Position {
    choice := g.GetWinningPos(g.GetBoard().CurrentGrid())

    // s'il peut gagner, il effectue directement le choix de gagner
    if choice == INVALID {
        choicesSorted := g.SortChoices(choices, 0)
        fmt.Println(choicesSorted)
        choice = choicesSorted[0]
    }

    return choice
}

func (g *game) SortChoices(choices []Position, number int) []Position {
    if len(choices) > number {
        if g.isGridWithO(choices[number]) {
            choices = remove(choices, number)
        }
        return g.SortChoices(choices, number+1)
    } else {
        return choices
    }
}

func remove(slice []Position, i int) []Position {
    copy(slice[i:], slice[i+1:])
    return slice[:len(slice)-1]
}

func (g *game) isGridWithO(pos Position) bool {
    tempCurrentPos := g.board.GetCurrentPos()
    g.board.SetCurrentPos(pos)

    for _, col := range g.board.CurrentGrid() {
        for _, tile := range col {
            if tile.Value == O {
                g.board.SetCurrentPos(tempCurrentPos)
                return true
            }
        }
    }
    g.board.SetCurrentPos(tempCurrentPos)
    return false
}
