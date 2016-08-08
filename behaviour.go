package main

import (
    "math/rand"
)

func (e *Entity) RandomWalk() {

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

    e.TryMove(desired_x, desired_y)

    return
}
