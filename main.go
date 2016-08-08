package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"
)

const (
    WIDTH = 8
    HEIGHT = 6
)

func main() {

    rand.Seed(time.Now().UTC().UnixNano())

    go logger()

    if len(DefaultMap) != len(ActionMap) {
        fmt.Printf("len(DefaultMap): %d\n", len(DefaultMap))
        fmt.Printf("len(ActionMap): %d\n", len(ActionMap))
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

func world_gen() (*World, error) {

    w := NewWorld(WIDTH, HEIGHT)

    var err error

    err = sprinkle_world(w, "dirt", 1.0)
    if err != nil {
        return w, fmt.Errorf("world_gen(): %v", err)
    }

    err = sprinkle_world(w, "grass", 0.4)
    if err != nil {
        return w, fmt.Errorf("world_gen(): %v", err)
    }

    err = sprinkle_world(w, "tree", 0.05)
    if err != nil {
        return w, fmt.Errorf("world_gen(): %v", err)
    }

    err = add_critters(w)
    if err != nil {
        return w, fmt.Errorf("world_gen(): %v", err)
    }

    return w, nil
}

func sprinkle_world(w *World, class string, chance float64) error {

    for x := 0; x < w.Width; x++ {
        for y := 0; y < w.Height; y++ {
            if rand.Float64() < chance {
                w.SetTileByClass(x, y, class)
            }
        }
    }

    return nil
}

func create_and_place_critter(x int, y int, class string, w *World) error {

    new_ent, err := NewEntity(x, y, class, w)
    if err != nil {
        return fmt.Errorf("create_and_place_critter(): %v", err)
    }

    err = w.PlaceCritter(new_ent)
    if err != nil {
        return fmt.Errorf("create_and_place_critter(): %v", err)
    }

    return nil
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

func add_critters(w *World) error {

    var errors []error

    errors = append(errors, create_and_place_critter(2, 2, "hare", w))
    errors = append(errors, create_and_place_critter(3, 3, "rat", w))

    errors = filtered_errors(errors)

    if len(errors) == 0 {
        return nil
    }

    if len(errors) == 1 {
        return fmt.Errorf("add_critters(): %v", errors[0])
    }

    return fmt.Errorf("add_critters(): multiple errors; the first was: %v", errors[0])

}
