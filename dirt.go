package main

func init() {
    ActionMap["dirt"] = nil

    DefaultMap["dirt"] = Entity{
        Rune: ' ',
        Mass: 10,
        Dead: true,
        Passable: true,
    }
}
