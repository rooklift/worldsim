package main

import (
    "math/rand"
)

func init() {
    ActionMap["rat"] = RatAct

    DefaultMap["rat"] = Entity{
        Rune: 'r',
        Mass: 0.15,
        Dead: false,
        Passable: true,
    }
}

func RatAct(e *Entity, spawn bool) {
    if spawn {
        return
    }

    direction := rand.Int31n(4)

    desired_x, desired_y := e.X, e.Y

    switch direction {
    case 0:
        desired_x -= 1
    case 1:
        desired_x += 1
    case 2:
        desired_y -= 1
    case 3:
        desired_y += 1
    }

    if e.World.InBounds(desired_x, desired_y) == false {
        return
    }

    target_tile := e.World.GetTile(desired_x, desired_y)

    if target_tile.Passable == false {
        return
    }

    e.X = desired_x
    e.Y = desired_y

    return
}
