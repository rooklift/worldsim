package main

import (
    "./worldsim"
)

func init() {
    worldsim.ActionMap["dirt"] = nil
    worldsim.DefaultMap["dirt"] = worldsim.Entity{
        Rune: ' ',
        Mass: 10,
        Dead: true,
        Passable: true,
    }

    worldsim.ActionMap["rock"] = nil
    worldsim.DefaultMap["rock"] = worldsim.Entity{
        Rune: '*',
        Mass: 100,
        Dead: true,
        Passable: false,
    }

}
