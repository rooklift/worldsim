package main

import (
    "fmt"
    "encoding/json"
)

type Entity struct {
    Class string        `json:"class"`
    World *World        `json:"-"`          // ignored in json
    X int               `json:"x"`
    Y int               `json:"y"`
    Acted bool          `json:"acted"`
    Rune rune           `json:"rune"`
    Mass float64        `json:"mass"`
    Hunger int          `json:"hunger"`
    Dead bool           `json:"dead"`
    Passable bool       `json:"passable"`
    Doom bool           `json:"doom"`       // Being destroyed this turn?
}

var ActionMap map[string]func(e *Entity) = make(map[string]func(*Entity)) // function called at act time or spawn time
var DefaultMap map[string]Entity = make(map[string]Entity) // default stats

func NewEntity(x, y int, class string, world *World) (*Entity, error) {

    e := new(Entity)

    default_stats, ok := DefaultMap[class]
    if ok == false {
        return nil, fmt.Errorf("NewEntity(): class '%s' is not in DefaultMap\n", class)
    }
    *e = default_stats

    // Must do the following after defaults are set...

    e.Class = class
    e.World = world
    e.X = x
    e.Y = y

    return e, nil
}

func (e *Entity) String() string {
    if e == nil {
        return "nil"
    }

    j, _ := json.MarshalIndent(e, "", "  ")
    return string(j)
}

func (e *Entity) Act() error {
    if e == nil {
        return fmt.Errorf("Act(): received nil pointer")
    }

    fn, ok := ActionMap[e.Class]
    if ok == false {
        return fmt.Errorf("Act(): class '%s' was not in ActionMap", e.Class)
    }

    if fn == nil {
        return nil
    }

    fn(e)

    return nil
}

func (e *Entity) Glyph() (rune, error) {
    if e == nil {
        return ' ', fmt.Errorf("Glyph(): received nil pointer")
    }

    r := e.Rune
    if r == 0 {
        return ' ', fmt.Errorf("Glyph(): entity had zero rune")
    }

    return r, nil
}

func (e *Entity) GetBlock() *Block {
    return e.World.Blocks[e.X][e.Y]
}

func (e *Entity) TryMove(desired_x int, desired_y int) bool {

    // Adjust the entity's .X and .Y if to the requested values if possible. Do nothing else.
    // In particular, note that this function should not fix block ownership of the entity.

    w := e.World

    if w.InBounds(desired_x, desired_y) == false {
        return false
    }

    block := e.World.Blocks[desired_x][desired_y]

    if block.Tile.Passable == false {
        return false
    }

    for _, other_critter := range block.Critters {
        if other_critter.Passable == false {
            return false
        }
    }

    e.X = desired_x
    e.Y = desired_y

    return true
}

func (e *Entity) BecomeTile() error {

    w := e.World

    if w.InBounds(e.X, e.Y) == false {
        return fmt.Errorf("BecomeTile() called with out of bounds entity; x, y == (%d,%d)", e.X, e.Y)
    }
    w.Blocks[e.X][e.Y].Tile = e
    return nil
}
