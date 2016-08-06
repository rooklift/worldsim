package main

func init() {
    RuneMap["dirt"] = '.'
    SpawnMap["dirt"] = nil
    ActionMap["dirt"] = nil

    StatsMap["dirt"] = Stats{
        Mass: 10,
        Dead: true,
    }
}
