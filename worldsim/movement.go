package worldsim

import (
    "fmt"
    "math/rand"
    "os"
)

func (e *Entity) Move(desired_x int, desired_y int) {

    // If moving is possible, 3 things need to happen:
    // -- Change e.x and e.y
    // -- Delink the entity from the old block
    // -- Link the entity from the target block

    w := e.World

    if w.InBounds(desired_x, desired_y) == false {
        fmt.Fprintf(os.Stderr, "Move(): target coordinates (%d,%d) out of bounds\n", desired_x, desired_y)
        return
    }

    old_x, old_y := e.x, e.y

    err := w.DelinkCritter(old_x, old_y, e)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Move(): %v\n", err)
    }

    e.x, e.y = desired_x, desired_y

    err = w.PlaceCritter(e)     // Uses e.x and e.y to work
    if err != nil {
        fmt.Fprintf(os.Stderr, "Move(): %v\n", err)
    }
}

func (e *Entity) TryMove(desired_x int, desired_y int) bool {

    w := e.World

    if w.InBounds(desired_x, desired_y) == false {
        return false
    }

    block := e.World.blocks[desired_x][desired_y]

    if block.tile == nil {  // No walking into the void. But nil tiles shouldn't exist really.
        fmt.Fprintf(os.Stderr, "TryMove(): target coordinates had nil tile\n")
        return false
    }

    if block.tile.Passable == false {
        return false
    }

    for _, other_critter := range block.critters {
        if other_critter.Passable == false {
            return false
        }
    }

    e.Move(desired_x, desired_y)

    return true
}

func (e *Entity) RandomWalk() {

    direction := rand.Int31n(4)

    desired_x, desired_y := e.x, e.y

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
