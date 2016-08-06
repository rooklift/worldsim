package main

import (
    "fmt"
    "encoding/json"
)

type Entity struct {
    Class string        `json:"class"`
    X int               `json:"x"`
    Y int               `json:"y"`
    Rune rune           `json:"rune"`
    Mass float64        `json:"mass"`
    Hunger int          `json:"hunger"`
    Dead bool           `json:"dead"`
}

var ActionMap map[string]func(e *Entity, spawn bool) error = make(map[string]func(*Entity, bool) error) // function called at act time or spawn time
var DefaultMap map[string]Entity = make(map[string]Entity) // default stats

func NewEntity(x, y int, class string) (*Entity, error) {

    e := new(Entity)

    default_stats, ok := DefaultMap[class]
    if ok == false {
        return nil, fmt.Errorf("NewEntity(): class '%s' is not in DefaultMap\n", class)
    }
    *e = default_stats

    // Must do the following after defaults are set...

    e.Class = class
    e.X = x
    e.Y = y

    // Now, call the entity's action function with spawn = true so it can do anything it needs to do at spawn time...

    act_function, ok := ActionMap[class]
    if ok == false {
        return nil, fmt.Errorf("NewEntity(): class '%s' is not in ActionMap\n", class)
    }

    if act_function != nil {
        act_function(e, true)
    }

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

    err := fn(e, false)
    if err != nil {
        return fmt.Errorf("Act(): %v\n", err)
    }

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
