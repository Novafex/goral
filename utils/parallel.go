package utils

import "sync"

// Parallelize takes variadic functions and runs each one in a separate Go
// routine. This function blocks until all functions return.
func Parallelize(funcs ...func()) {
	var group sync.WaitGroup
	group.Add(len(funcs))
	defer group.Wait()

	for _, f := range funcs {
		go func(copy func()) {
			defer group.Done()
			copy()
		}(f)
	}
}

// ParallelizeArgs takes a single function and runs it multiple times in go
// routines using the provided args slice as the determining parameter and
// number of routines.
//
// Use this for doing the same process with multiple inputs simultaneously.
func ParallelizeArgs[T any](fn func(in T), args []T) {
	var group sync.WaitGroup
	group.Add(len(args))
	defer group.Wait()

	for _, a := range args {
		go func(in T) {
			defer group.Done()
			fn(in)
		}(a)
	}
}

// ParallelizeArgsWithReturn takes a single function and runs it for each
// argument in args as a go routine (async). It collects the returned values and
// finally returns those when completed.
func ParallelizeArgsWithReturn[In any, Ret any](fn func(in In) Ret, args []In) []Ret {
	var group sync.WaitGroup
	group.Add(len(args))

	results := make([]Ret, len(args))
	lock := sync.Mutex{}
	for i, a := range args {
		go func(in In, ind int) {
			defer group.Done()
			ret := fn(in)

			lock.Lock()
			results[ind] = ret
			lock.Unlock()
		}(a, i)
	}

	group.Wait()
	return results
}

// ParallelizeArgsWithError takes a function in and runs it for every entry in
// args as a go routine. It collects the return values (error) and returns the
// total slice.
//
// Sugar on-top of [ParallelizeArgsWithReturn].
func ParallelizeArgsWithError[In any](fn func(in In) error, args []In) []error {
	return ParallelizeArgsWithReturn[In, error](fn, args)
}
