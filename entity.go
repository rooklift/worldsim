package main

import (
    "fmt"
    "encoding/json"
)

type Entity struct {
    Class string        `json:"class"`
    X int               `json:"x"`
    Y int               `json:"y"`
    Stats
}

type Stats struct {
    Mass float64        `json:"mass"`
    Hunger int          `json:"hunger"`
    Dead bool           `json:"dead"`
}

var RuneMap map[string]rune = make(map[string]rune) // rune to display
var SpawnMap map[string]func(*Entity) error = make(map[string]func(*Entity) error) // function called at spawn time
var ActionMap map[string]func(*Entity) error = make(map[string]func(*Entity) error) // function called at act time
var StatsMap map[string]Stats = make(map[string]Stats) // default stats

func print_map_lengths() {
    fmt.Printf("len(RuneMap): %d\n", len(RuneMap))
    fmt.Printf("len(SpawnMap): %d\n", len(SpawnMap))
    fmt.Printf("len(ActionMap): %d\n", len(ActionMap))
    fmt.Printf("len(StatsMap): %d\n", len(StatsMap))
}

func NewEntity(x, y int, class string) (*Entity, error) {

    _, ok := RuneMap[class]
    if ok == false {
        return nil, fmt.Errorf("NewEntity(): class '%s' is not in RuneMap\n", class)
    }

    e := new(Entity)

    e.Class = class
    e.X = x
    e.Y = y

    default_stats, ok := StatsMap[class]
    if ok == false {
        return nil, fmt.Errorf("NewEntity(): class '%s' is not in StatsMap\n", class)
    }
    e.Stats = default_stats

    spawn_function, ok := SpawnMap[class]
    if ok == false {
        return nil, fmt.Errorf("NewEntity(): class '%s' is not in SpawnMap\n", class)
    }

    if spawn_function != nil {
        spawn_function(e)
    }

    return e, nil
}

func (e *Entity) String() string {
    if e == nil {
        return "(nil entity)"
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

    err := fn(e)
    if err != nil {
        return fmt.Errorf("Act(): %v\n", err)
    }

    return nil
}

func (e *Entity) Glyph() (rune, error) {
    if e == nil {
        return ' ', fmt.Errorf("Glyph(): received nil pointer")
    }

    r, ok := RuneMap[e.Class]
    if ok == false {
        return ' ', fmt.Errorf("Glyph(): class '%s' was not in RuneMap", e.Class)
    }
    return r, nil
}
