package main

import (
    "fmt"
    "os"
)

type World struct {
    Width int
    Height int
    Tiles [][]*Entity
    Creeps []*Entity
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
    var s []rune
    for y := 0; y < w.Height; y++ {
        for x := 0; x < w.Width; x++ {

            glyph, err := w.Tiles[x][y].Glyph()
            if err != nil {
                fmt.Fprintf(os.Stderr, "While printing world: %v\n", err)
            }
            s = append(s, glyph)
        }
        s = append(s, '\n')
    }
    return string(s)
}
