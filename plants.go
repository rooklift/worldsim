package main

func init() {

    ActionMap["grass"] = nil
    DefaultMap["grass"] = Entity{
        Rune: ',',
        Mass: 0.2,
        Passable: true,
    }


    ActionMap["tree"] = nil
    DefaultMap["tree"] = Entity{
        Rune: 'O',
        Mass: 14000,
        Passable: false,
    }

}
