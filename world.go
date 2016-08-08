package main

import (
    "fmt"
    "os"
)

type World struct {
    Width int
    Height int
    Blocks [][]*Block
}

type Block struct {
    Tile *Entity
    Critters []*Entity
}

func NewWorld(width, height int) *World {

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

func (w *World) Iterate() {

    var err error

    // Call each critter's action function. The critter's function is responsible for changing its X and Y coordinates.
    // But we then adjust what block owns it in this function here.

    for x := 0; x < w.Width; x++ {
        for y := 0; y < w.Height; y++ {
            for _, critter := range w.Blocks[x][y].Critters {
                critter.Acted = false
            }
        }
    }

    for x := 0; x < w.Width; x++ {
        for y := 0; y < w.Height; y++ {
            for _, critter := range w.Blocks[x][y].Critters {
                if critter.Acted == false {

                    if x != critter.X || y != critter.Y {
                        fmt.Fprintf(os.Stderr, "Iterate(): In block (%d,%d), found %s with .X == %d, .Y == %d\n", x, y, critter.Class, critter.X, critter.Y)
                    }

                    err = critter.Act()
                    if err != nil {
                        fmt.Fprintf(os.Stderr, "Iterate(): %v\n", err)
                    }
                    critter.Acted = true

                    if critter.X != x || critter.Y != y {
                        err = w.RemoveCritter(x, y, critter)
                        if err != nil {
                            fmt.Fprintf(os.Stderr, "Iterate(): %v\n", err)
                        }
                        err = w.PlaceCritter(critter)
                        if err != nil {
                            fmt.Fprintf(os.Stderr, "Iterate(): %v\n", err)
                        }
                    }
                }
            }
        }
    }
}

func (w *World) GetTile(x, y int) *Entity {

    // Assumes uniform length of columns and rows, i.e. no raggedy 2D arrays

    if w.InBounds(x, y) == false {
        return nil
    }
    return w.Blocks[x][y].Tile
}

func (w *World) SetTile(x, y int, e *Entity) error {

    // Assumes uniform length of columns and rows, i.e. no raggedy 2D arrays

    if w.InBounds(x, y) == false {
        return fmt.Errorf("SetTile() called with out of bounds x, y == (%d,%d)", x, y)
    }
    w.Blocks[x][y].Tile = e
    return nil
}

func (w *World) RemoveCritter(x, y int, e *Entity) error {

    if w.InBounds(x, y) == false {
        return fmt.Errorf("RemoveCritter() called with out of bounds x, y == (%d,%d)", x, y)
    }

    for i, c := range w.Blocks[x][y].Critters {
        if c == e {
            w.Blocks[x][y].Critters = append(w.Blocks[x][y].Critters[:i], w.Blocks[x][y].Critters[i + 1:]...)
            return nil      // One can only do the above trick once inside a range loop, which is all we need here
        }
    }

    return fmt.Errorf("RemoveCritter() couldn't find the critter in block (%d,%d)", x, y)
}

func (w *World) PlaceCritter(e *Entity) error {

    // Assumes the entity has its .X and .Y already validly set.

    x, y := e.X, e.Y

    if w.InBounds(x, y) == false {
        return fmt.Errorf("PlaceCritter() called with out of bounds x, y == (%d,%d)", x, y)
    }

    // FIXME: check not already present

    w.Blocks[x][y].Critters = append(w.Blocks[x][y].Critters, e)

    return nil
}

func (w *World) String() string {

    // First make a 2D slice of the runes to print

    var r [][]rune
    var glyph rune
    var err error

    for x := 0; x < w.Width; x++ {

        var column []rune

        for y := 0; y < w.Height; y++ {

            if len(w.Blocks[x][y].Critters) > 0 {
                glyph, err = w.Blocks[x][y].Critters[0].Glyph()
                if err != nil {
                    fmt.Fprintf(os.Stderr, "While printing world critter: %v\n", err)
                }
            } else {
                glyph, err = w.Blocks[x][y].Tile.Glyph()
                if err != nil {
                    fmt.Fprintf(os.Stderr, "While printing world tile: %v\n", err)
                }
            }
            column = append(column, glyph)
        }

        r = append(r, column)
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

func (w *World) TryMove(e *Entity, desired_x int, desired_y int) bool {

    // Adjust the entity's .X and .Y if to the requested values if possible. Do nothing else.
    // In particular, note that this function should not fix block ownership of the entity.

    if w.InBounds(desired_x, desired_y) == false {
        return false
    }

    block := e.World.Blocks[desired_x][desired_y]

    if block.Tile.Passable == false {
        return false
    }

    for _, other_critter := range block.Critters {
        if other_critter.Passable == false {
            return false
        }
    }

    e.X = desired_x
    e.Y = desired_y

    return true
}

func (w *World) CrittersInRect(centre_x int, centre_y int, dist int) []*Entity {

    var result []*Entity

    if dist < 0 {
        return result
    }

    start_x := centre_x - dist
    start_y := centre_y - dist

    end_x := centre_x + dist
    end_y := centre_y + dist

    if start_x < 0 {
        start_x = 0
    }
    if start_x >= w.Width {
        start_x = w.Width - 1
    }
    if start_y < 0 {
        start_y = 0
    }
    if start_y >= w.Height {
        start_y = w.Height - 1
    }

    for x := start_x; x <= end_x; x++ {
        for y := start_y; y <= end_y; y++ {
            result = append(result, w.Blocks[x][y].Critters...)
        }
    }

    return result
}

func (w *World) CrittersNearCritter(e *Entity, dist int) []*Entity {

    result_with_self := w.CrittersInRect(e.X, e.Y, dist)

    var result []*Entity

    for _, ent := range result_with_self {
        if ent != e {
            result = append(result, ent)
        }
    }

    return result
}
