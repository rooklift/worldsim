package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"
)

const (
    WIDTH = 60
    HEIGHT = 30
)

func main() {

    rand.Seed(time.Now().UTC().UnixNano())

    if len(DefaultMap) != len(ActionMap) {
        fmt.Printf("len(DefaultMap): %d\n", len(DefaultMap))
        fmt.Printf("len(ActionMap): %d\n", len(ActionMap))
        return
    }

    w, err := world_gen()
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
    }

    for n := 0; n < 100; n++ {
        w.Iterate()
    }

    fmt.Println(w)
}

func world_gen() (*World, error) {

    w := make_world(WIDTH, HEIGHT)

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

func make_world(width, height int) *World {

    // Returns a 2D slice of empty blocks

    w := new(World)

    w.Width = width
    w.Height = height

    for x := 0; x < w.Width; x++ {
        var column []*Block
        for y := 0; y < w.Height; y++ {
            column = append(column, new(Block))
        }
        w.Blocks = append(w.Blocks, column)
    }

    return w
}

func sprinkle_world(w *World, class string, chance float64) error {

    for x := 0; x < w.Width; x++ {
        for y := 0; y < w.Height; y++ {
            if rand.Float64() < chance {

                new_entity, err := NewEntity(x, y, class, w)
                if err != nil {
                    return fmt.Errorf("sprinkle_world(): %v", err)
                }

                w.SetTile(x, y, new_entity)
            }
        }
    }

    return nil
}

func add_critters(w *World) error {

    new_ent, err := NewEntity(5, 5, "rat", w)
    if err != nil {
        return fmt.Errorf("add_critters(): %v", err)
    }

    err = w.PlaceCritter(new_ent)
    if err != nil {
        return fmt.Errorf("add_critters(): %v", err)
    }

    return nil
}
