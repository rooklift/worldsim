package worldsim

import (
    "fmt"
    "encoding/json"
    "os"
)

var ActionMap map[string]func(e *Entity) = make(map[string]func(*Entity))
var DefaultMap map[string]Entity = make(map[string]Entity)

type Entity struct {

    // It's important to bear in mind that, when an entity moves, not only does its own
    // .x and .y need to change, but the old and new owning blocks need updating. Therefore,
    // never set .x and .y directly, but instead call methods like TryMove().

    Class string        `json:"class"`
    World *World        `json:"-"`          // ignored in json
    x int               `json:"x"`
    y int               `json:"y"`
    Acted bool          `json:"acted"`
    Rune rune           `json:"rune"`
    Mass float64        `json:"mass"`
    Hunger int          `json:"hunger"`
    Dead bool           `json:"dead"`
    Passable bool       `json:"passable"`
}

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
    e.x = x
    e.y = y

    return e, nil
}

func (e *Entity) X() int {
    return e.x
}

func (e *Entity) Y() int {
    return e.y
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
    return e.World.Blocks[e.x][e.y]
}

func (e *Entity) BecomeTile() error {

    w := e.World

    if w.InBounds(e.x, e.y) == false {
        return fmt.Errorf("BecomeTile() called with out of bounds entity; x, y == (%d,%d)", e.x, e.y)
    }
    w.Blocks[e.x][e.y].Tile = e
    return nil
}

func (e *Entity) Destroy() {

    // This is achieved simply by removing the entity from the slice of critters in the block

    err := e.World.DelinkCritter(e.x, e.y, e)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Destroy(): %v\n", err)
    }
}
