package main

func init() {
    RuneMap["grass"] = ','
    SpawnMap["grass"] = nil
    ActionMap["grass"] = nil

    StatsMap["grass"] = Stats{
        Mass: 0.2,
    }
}
