# Lures

A work-in-progress collection of strange attractor functions.

Attractors... lures... get it?

These are just the functions to do a single iteration of a given attractor. No rendering logic, no other manipulation.

Each attractor includes default/suggested parameters.

Examples are rendered using:

- https://github.com/bit101/blcairo
- https://github.com/bit101/bitlib
- https://github.com/bit101/wire

## How to use

1. Create a new attractor, e.g.:

```
attr := lures.NewAizawa()
```

2. Get initial point values:

```
x, y, z := attr.InitVals()
```

3. Possibly, create some kind of point object, scale it using `attr.Scale`, project it to 2d, render it.

4. Get the next point using current `x`, `y`, `z` and `attr.Dt`.

```
x, y, z = attr.Iterate(x, y, z, attr.Dt)
```

5. Repeat steps 3 and 4 as many times as you want, usually in the tens of thousands.

Alternately, get all the points and add to a list, then scale, project and render each point in the list. This is how [wire](https://github.com/bit101/wire) does it.

Or, you could just generate a list of point values in whatever format you want, and import that into another program for rendering.
