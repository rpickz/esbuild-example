# `esbuild` Example

## Summary

This repository demonstrates an example build of simple Javascript front-end programs comparing webpack and `esbuild`.

`esbuild` is significantly more performant and efficient than webpack, producing similar build results in a fraction of the time,
heavily optimising the amount of memory allocations, CPU resource, etc. used to produce a build.

## Example Results

Time results for `webpack`:

```
real    0m4.639s
user    0m4.764s
sys     0m0.853s
```

Time results for `esbuild`:

```
real    0m0.069s
user    0m0.063s
sys     0m0.017s
```

As can be seen above, `esbuild` is significantly more performant than `webpack`.