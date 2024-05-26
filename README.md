# Lures

A work-in-progress collection of strange attractor functions.

Attractors... lures... get it?

These are just the functions to do a single iteration of a given attractor. No rendering logic, no other manipulation.

Each attractor includes default/suggested parameters.

Examples are rendered using:

- https://github.com/bit101/blcairo
- https://github.com/bit101/bitlib
- https://github.com/bit101/wire

## 3d attractors

The 3d attractors are generally defined like:

```
dx/dt = ... some formula using current x, y, z + params
dy/dt = ... same
dz/dt = ... same
```

So the code for each ends up like:

```
dx = ...
dy = ...
dz = ...
return x + dx*dt, y + dy*dt, z + dz*dt
```

The `dt` value usually affects the density of the final form. Lower `dt` values can require rendering more points to complete the shape of the attractor.

1. Create a new attractor, e.g.:

```
attr := lures.NewAizawa()
```

2. Get initial point values:

```
x, y, z := attr.InitVals()
```

3. Possibly, create some kind of point object, scale it using `attr.Scale`, project it to 2d, render it.

4. Get the next point using current `x`, `y`, `z`

```
x, y, z = attr.Iterate(x, y, z)
```

5. Repeat steps 3 and 4 as many times as you want, usually in the tens of thousands.

Alternately, get all the points and add to a list, then scale, project and render each point in the list. This is how [wire](https://github.com/bit101/wire) does it.

Or, you could just generate a list of point values in whatever format you want, and import that into another program for rendering.


### Using custom parameters in 3d attractors

You can ignore `InitX`, `InitY`, `InitZ` and `Scale`. These are not used internally and are only there as suggestions to create visually pleasing examples. 

For the actual attractor parameters (`A` through `Z`) and `Dt` these _are_ used internally. You can change them after you create your attractor object:

```
attr := l3d.NewAizawa()
attr.Params.A = 0.83
attr.Dt = 0.02
```

Then render the attractor however you like. It will use the updated parameter value.

It's suggested to change parameters by small amounts at first. A small change might make a large difference in the output.

Parameter values chosen completely at random will usually not produce interesting attractors. But you might get lucky now and then if you generate many random attractors.

## 2d attractors

The 2d attractors are usually in the format:

```
x(n+1) = ... formula using some combinition of x(n), y(n) and params
y(n+1) = ... same
```

So the code is like:

```
x1 = ...
y1 = ...
return x1, y1
```

The `dt` value is not used in these types of attractors.

1. Create a new attractor, e.g.:

```
attr := lures.NewPickover()
```

2. Get initial point values:

```
x, y := attr.InitVals()
```

3. Possibly, create some kind of point object, scale it using `attr.Scale` and render it.

4. Get the next point using current `x`, `y`

```
x, y = attr.Iterate(x, y)
```

5. Repeat steps 3 and 4 as many times as you want, usually in the tens of thousands at least.

Rather than rendering directly, you may want to save the pixel hits in a grid and then do coloring based on density (how many times the pixel was visited).

### Using custom parameters in 2d attractors

Exactly the same as in 3d attractors, but no `dt` value.

2d attractors often seem to be less dependent on the initial point values and tend to converge to the general shape of the attractor no matter what point you start with. Experimenting with random parameters is more useful.
