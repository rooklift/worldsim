package main

func init() {
    ActionMap["rat"] = RandomWalk

    DefaultMap["rat"] = Entity{
        Rune: 'r',
        Mass: 0.15,
        Dead: false,
        Passable: true,
    }
}
