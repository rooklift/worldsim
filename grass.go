package main

func init() {
    ActionMap["grass"] = nil

    DefaultMap["grass"] = Entity{
        Rune: ',',
        Mass: 0.2,
        Passable: true,
    }
}
