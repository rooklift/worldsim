package main

import (
    "fmt"
    "os"
)

type World struct {
    Width int
    Height int
    Tiles [][]*Entity
    Critters []*Entity
}

func (w *World) Iterate() {
    for _, critter := range w.Critters {
        err := critter.Act()
        if err != nil {
            fmt.Fprintf(os.Stderr, "While iterating: %v\n", err)
        }
    }
}

func (w *World) GetTile(x, y int) *Entity {

    // Assumes uniform length of columns and rows, i.e. no raggedy 2D arrays

    if x < 0 || x >= w.Width || y < 0 || y >= w.Height {
        return nil
    }
    return w.Tiles[x][y]
}

func (w *World) SetTile(x, y int, e *Entity) error {

    // Assumes uniform length of columns and rows, i.e. no raggedy 2D arrays

    if x < 0 || x >= w.Width || y < 0 || y >= w.Height {
        return fmt.Errorf("SetTile() called with out of bounds x, y")
    }
    w.Tiles[x][y] = e
    return nil
}

func (w *World) String() string {

    // First make a 2D slice of the runes to print

    var r [][]rune

    for x := 0; x < w.Width; x++ {

        var column []rune

        for y := 0; y < w.Height; y++ {

            glyph, err := w.Tiles[x][y].Glyph()
            if err != nil {
                fmt.Fprintf(os.Stderr, "While printing world (tile phase): %v\n", err)
            }

            column = append(column, glyph)
        }

        r = append(r, column)
    }

    // Now modify that slice with runes of the critters

    for _, critter := range w.Critters {

        glyph, err := critter.Glyph()

        if err != nil {
            fmt.Fprintf(os.Stderr, "While printing world (critter phase): %v\n", err)
        }

        r[critter.X][critter.Y] = glyph

    }

    // Now create a 1D slice that can be converted to a string
    // (not created with string += part for speed reasons)

    var s []rune

    for y := 0; y < w.Height; y++ {
        for x := 0; x < w.Width; x++ {
            s = append(s, r[x][y])
        }
        s = append(s, '\n')
    }

    return string(s)
}

func (w *World) InBounds(x, y int) bool {
    if x >= 0 && x < w.Width && y >= 0 && y < w.Height {
        return true
    }
    return false
}
