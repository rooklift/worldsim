package main

func init() {
    RuneMap["grass"] = ','
    SpawnMap["grass"] = SpawnGrass
    ActionMap["grass"] = nil
}

func SpawnGrass(e *Entity) {
    e.Mass = 0.2
}
