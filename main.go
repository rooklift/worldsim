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

    fmt.Println(w)
    fmt.Println(w.Tiles[2][2])
}

func world_gen() (*World, error) {

    w := make_world(WIDTH, HEIGHT)

    var err error

    err = sprinkle_world(w, "dirt", 1.0)
    if err != nil {
        return w, fmt.Errorf("world_gen(): %v", err)
    }

    err = sprinkle_world(w, "grass", 0.2)
    if err != nil {
        return w, fmt.Errorf("world_gen(): %v", err)
    }

    return w, nil
}

func make_world(width, height int) *World {

    // Returns a 2D slice of nil pointers

    w := new(World)

    w.Width = width
    w.Height = height

    for x := 0; x < w.Width; x++ {
        var column []*Entity
        for y := 0; y < w.Height; y++ {
            column = append(column, nil)
        }
        w.Tiles = append(w.Tiles, column)
    }

    return w
}

func sprinkle_world(w *World, class string, chance float64) error {

    for x := 0; x < w.Width; x++ {
        for y := 0; y < w.Height; y++ {
            if rand.Float64() < chance {

                new_entity, err := NewEntity(x, y, class)
                if err != nil {
                    return fmt.Errorf("sprinkle_world(): %v", err)
                }

                w.SetTile(x, y, new_entity)
            }
        }
    }

    return nil
}
