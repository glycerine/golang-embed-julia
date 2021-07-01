package main

// See README.md and <https://docs.julialang.org/en/v1/manual/embedding> on how to modify the paths.
// NOTE: It is important that there is NO empty line before `import "C"`, see https://golang.org/cmd/cgo.
//
// #cgo CFLAGS: -fPIC -DJULIA_INIT_DIR="/Applications/Julia-0.6.app/Contents/Resources/julia/lib" -I/Applications/Julia-0.6.app/Contents/Resources/julia/include/julia -I.
// #cgo LDFLAGS: -L/Applications/Julia-0.6.app/Contents/Resources/julia/lib/julia  -L/Applications/Julia-0.6.app/Contents/Resources/julia/lib -Wl,-rpath,/Applications/Julia-0.6.app/Contents/Resources/julia/lib -ljulia
// #include <julia.h>
import "C"

func main() {

	/* required: setup the Julia context */
	C.jl_init()

	/* run Julia commands */
	C.jl_eval_string(C.CString(`println(sqrt(2.0))`))

	/* strongly recommended: notify Julia that the
	   program is about to terminate. this allows
	   Julia time to cleanup pending write requests
	   and run all finalizers
	*/
	C.jl_atexit_hook(0)
}
