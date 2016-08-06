package main

func init() {
    RuneMap["dirt"] = '.'
    SpawnMap["dirt"] = SpawnDirt
    ActionMap["dirt"] = nil
}

func SpawnDirt(e *Entity) {
    e.Mass = 10
}
