package main

import (
    "fmt"
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

    ActionMap["hare"] = HareAct
    DefaultMap["hare"] = Entity{
        Rune: 'h',
        Mass: 1.5,
        Dead: false,
        Passable: true,
    }

}

func RatAct(e *Entity) {

    e.RandomWalk()

    if rand.Float64() < 0.001 {
        e.Destroy()
        LogChan <- "rat randomly died\n"
    }
}

func HareAct(e *Entity) {

    e.RandomWalk()

    b := e.GetBlock()
    t := b.Tile

    if t.Class == "grass" {
        if t.Mass > 0 && t.Dead == false {
            t.Mass -= 0.01
            if t.Mass < 0 {
                t.Mass = 0
            }
            if t.Mass == 0 {
                e.World.SetTileByClass(e.X, e.Y, "dirt")
                LogChan <- fmt.Sprintf("hare ate all the grass at (%d,%d)\n", e.X, e.Y)
            }
        }
    }
}
