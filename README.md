golang-embed-julia
=====

Official documentation: <https://docs.julialang.org/en/v1/manual/embedding>

## Determine the required compiler flags

You can get the required flags for your system using the `julia-config.jl` script.
It is included with the julia installation,
but the location depends on your Julia installation environment.
First, check the output of

```
echo $JULIA_DEPOT_PATH
```

If that environment variable is set,
the script might be in one of the indicated directories.
If it is unset, or you can't find it there, try

```
julia -e "println(DEPOT_PATH)"
```

## Use the cgo directive to load the Julia C libraries

Documentation: <https://golang.org/cmd/cgo/>

Your go header will have the following structure:

```go
package main

// #cgo CFLAGS: <output of julia-config.jl --cflags>
// #cgo LDFLAGS: <output of julia-config.jl --allflags starting from after the CFLAGS part>
// #include <julia.h>
import "C"
```

## Test the provided go file:
~~~
$ go run embed.go
1.4142135623730951
$
~~~
