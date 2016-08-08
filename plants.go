package main

import (
    "./worldsim"
)

func init() {

    worldsim.ActionMap["grass"] = nil
    worldsim.DefaultMap["grass"] = worldsim.Entity{
        Rune: ',',
        Mass: 0.2,
        Passable: true,
    }


    worldsim.ActionMap["tree"] = nil
    worldsim.DefaultMap["tree"] = worldsim.Entity{
        Rune: 'O',
        Mass: 14000,
        Passable: false,
    }

}
