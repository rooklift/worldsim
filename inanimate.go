package main

func init() {
    ActionMap["dirt"] = nil
    DefaultMap["dirt"] = Entity{
        Rune: ' ',
        Mass: 10,
        Dead: true,
        Passable: true,
    }

    ActionMap["rock"] = nil
    DefaultMap["rock"] = Entity{
        Rune: '*',
        Mass: 100,
        Dead: true,
        Passable: false,
    }
    
}
