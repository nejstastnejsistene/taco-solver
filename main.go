package main

import "fmt"

const (
	RedTile     Tile = 'X'
	SolidTile        = '='
	CrackedTile      = '-'
	NoTile           = ' '

	Taco       Object = 'P'
	BlueBlock         = '#'
	RedBlock          = '^'
	TrelloCoin        = 'J'
	NoObject          = ' '
)

var (
	Up    = Direction{0, -1}
	Down  = Direction{0, 1}
	Left  = Direction{-1, 0}
	Right = Direction{1, 0}

	Directions = Path{Up, Down, Left, Right}
)

func Solve(level Level) (minPath Path) {
	seen := map[Level]bool{level: true}
	var solve func(state State, path Path)
	solve = func(state State, path Path) {
		for _, dir := range Directions {
			newState, ok := state.MaybeMove(dir)
			if ok && !seen[newState.Level] {
				seen[newState.Level] = true
				newPath := make(Path, len(path))
				copy(newPath, path)
				newPath = append(newPath, dir)
				if newState.CoinsRemaining == 0 &&
					(len(minPath) == 0 || len(newPath) < len(minPath)) {
					minPath = newPath
				} else {
					solve(newState, newPath)
				}
			}
		}
	}
	solve(level.NewState(), nil)
	return
}

type State struct {
	Row, Col       int
	CoinsRemaining int
	Level
}

func (v State) MaybeMove(dir Direction) (state State, ok bool) {
	state = v
	toCell, ok := state.Level.MaybeCell(state.Row+dir.Dy, state.Col+dir.Dx)
	// Destination is out of bounds.
	if !ok {
		return
	}
	// Destination is a bottomles pit.
	if toCell.Tile == NoTile {
		ok = false
		return
	}
	switch toCell.Object {
	case BlueBlock:
		ok = false
		return
	case RedBlock:
		var nextCell *Cell
		nextCell, ok = state.Level.MaybeCell(state.Row+2*dir.Dy, state.Col+2*dir.Dx)
		// You can only push red blocks if their other side is in-bounds and vacant.
		if !ok || nextCell.Object != NoObject {
			return
		}
		// If there is no tile, they fall to make a red tile.
		if nextCell.Tile == NoTile {
			nextCell.Tile = RedTile
			// Otherwise they are pushed to the next tile.
		} else {
			nextCell.Object = RedBlock
		}
	case TrelloCoin:
		state.CoinsRemaining -= 1
	}
	fromCell := state.Level.Cell(state.Row, state.Col)
	// Move Taco.
	fromCell.Object = NoObject
	toCell.Object = Taco
	state.Row += dir.Dy
	state.Col += dir.Dx
	// Damange the tile Taco just moved off of.
	switch fromCell.Tile {
	case SolidTile:
		fromCell.Tile = CrackedTile
	case CrackedTile:
		fromCell.Tile = NoTile
	}
	return
}

type Level struct {
	Rows, Cols int
	Cells      [25]Cell
}

func MustParseLevel(num int, ss []string) Level {
	if len(ss) == 0 {
		panic(fmt.Sprintf("level %d: level should have at least one row", num))
	}
	if len(ss[0]) == 0 || len(ss[0])%2 != 0 {
		panic(fmt.Sprintf("level %d: rows should have an even, non-zero number of characters", num))
	}
	rows := len(ss)
	cols := len(ss[0]) / 2
	level := Level{Rows: rows, Cols: cols}
	for r, s := range ss {
		if len(s) != len(ss[0]) {
			panic(fmt.Sprintf("level %d: each row should be the same length", num))
		}
		for c := 0; c < cols; c++ {
			cell := Cell{Tile(s[2*c]), Object(s[2*c+1])}
			switch cell.Tile {
			case RedTile, SolidTile, CrackedTile, NoTile:
			default:
				panic(fmt.Sprintf("level %d: invalid tile: \"%s\"", num, cell.Tile))
			}
			switch cell.Object {
			case Taco, BlueBlock, RedBlock, TrelloCoin, NoObject:
			default:
				panic(fmt.Sprintf("level %d: invalid object: \"%s\"", num, cell.Object))
			}
			level.Cells[r*cols+c] = Cell{Tile(s[2*c]), Object(s[2*c+1])}
		}
	}
	return level
}

func (v Level) NewState() State {
	r, c := v.FindTaco()
	coins := 0
	for r := 0; r < v.Rows; r++ {
		for c := 0; c < v.Cols; c++ {
			if v.Cell(r, c).Object == TrelloCoin {
				coins++
			}
		}
	}
	return State{r, c, coins, v}
}

func (v Level) FindTaco() (r, c int) {
	for r = 0; r < v.Rows; r++ {
		for c = 0; c < v.Cols; c++ {
			if v.Cell(r, c).Object == Taco {
				return
			}
		}
	}
	panic("taco is nowhere to be found!")
}

func (v *Level) Cell(r, c int) *Cell {
	return &v.Cells[r*v.Cols+c]
}

func (v *Level) MaybeCell(r, c int) (cell *Cell, ok bool) {
	if !(0 <= r && r < v.Rows && 0 <= c && c < v.Cols) {
		return
	}
	return v.Cell(r, c), true
}

func (v Level) String() (s string) {
	for i, cell := range v.Cells {
		if i%v.Cols == 0 {
			s += "\n"
		}
		s += cell.String()
	}
	return
}

type Cell struct {
	Tile
	Object
}

func (v Cell) String() string {
	return string(v.Tile) + string(v.Object)
}

type Tile rune
type Object rune

type Direction struct {
	Dx, Dy int
}

func (v Direction) String() string {
	switch v {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		panic("invalid direction")
	}
}

type Path []Direction

func main() {
	for i := 0; i < len(Levels); i++ {
		fmt.Printf("Level %d: ", i+1)
		fmt.Println(Solve(Levels[i]))
	}
}
