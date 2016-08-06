package main

import (
    "fmt"
)

type Entity struct {
    Class string
    X int
    Y int
    Mass float64
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
    if e.Class != "" {
        return e.Class
    }
    return "(classless entity)"
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

    err := fn()
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
