package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"

    "./worldsim"
)

const (
    WIDTH = 8
    HEIGHT = 6
)

func main() {

    rand.Seed(time.Now().UTC().UnixNano())

    go logger()

    if len(worldsim.DefaultMap) != len(worldsim.ActionMap) {
        fmt.Printf("len(DefaultMap): %d\n", len(worldsim.DefaultMap))
        fmt.Printf("len(ActionMap): %d\n", len(worldsim.ActionMap))
        return
    }

    w, err := world_gen()
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
    }

    fmt.Println(w)

    for n := 0; n < 1000; n++ {
        w.Iterate()
    }

    fmt.Println(w)
}

func world_gen() (*worldsim.World, error) {

    w := worldsim.NewWorld(WIDTH, HEIGHT)

    var err error

    err = w.SprinkleTerrain("dirt", 1.0)
    if err != nil {
        return w, fmt.Errorf("world_gen(): %v", err)
    }

    err = w.SprinkleTerrain("grass", 0.4)
    if err != nil {
        return w, fmt.Errorf("world_gen(): %v", err)
    }

    err = w.SprinkleTerrain("tree", 0.05)
    if err != nil {
        return w, fmt.Errorf("world_gen(): %v", err)
    }

    err = add_critters(w)
    if err != nil {
        return w, fmt.Errorf("world_gen(): %v", err)
    }

    return w, nil
}

func filtered_errors(errors []error) []error {

    var result []error

    for _, e := range errors {
        if e != nil {
            result = append(result, e)
        }
    }

    return result
}

func add_critters(w *worldsim.World) error {

    var errors []error

    errors = append(errors, w.CreateCritterByClass(2, 2, "hare"))
    errors = append(errors, w.CreateCritterByClass(3, 3, "rat"))

    errors = filtered_errors(errors)

    if len(errors) == 0 {
        return nil
    }

    if len(errors) == 1 {
        return fmt.Errorf("add_critters(): %v", errors[0])
    }

    return fmt.Errorf("add_critters(): multiple errors; the first was: %v", errors[0])

}
