package main

import (
    "fmt"
    "math/rand"

    "./worldsim"
)

func init() {

    worldsim.ActionMap["rat"] = RatAct
    worldsim.DefaultMap["rat"] = worldsim.Entity{
        Rune: 'r',
        Mass: 0.15,
        Dead: false,
        Passable: true,
    }

    worldsim.ActionMap["hare"] = HareAct
    worldsim.DefaultMap["hare"] = worldsim.Entity{
        Rune: 'h',
        Mass: 1.5,
        Dead: false,
        Passable: true,
    }

}

func RatAct(e *worldsim.Entity) {

    e.RandomWalk()

    if rand.Float64() < 0.001 {
        e.Destroy()
        LogChan <- "rat randomly died\n"
    }
}

func HareAct(e *worldsim.Entity) {

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
                e.World.SetTileByClass(e.X(), e.Y(), "dirt")
                LogChan <- fmt.Sprintf("hare ate all the grass at (%d,%d)\n", e.X(), e.Y())
            }
        }
    }
}
