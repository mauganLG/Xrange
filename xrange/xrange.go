package xrange

import "iter"

// Xrange generate a sequence that emulate the Python range behaviour
// The sequence will be empty if you put a 0 as a step or
// if you put start inferior to stop and put a negative step
// and vice-versa
func Xrange(start, stop, step int) iter.Seq[int] {

	return func(yield func(int) bool) {

		if step == 0 || (start < stop && step < 0) || (start > stop && step > 0) {
			return
		}
		for i := start; (step > 0 && i < stop) || (step < 0 && i > stop); i += step {
			if !yield(i) {
				return
			}
		}

	}
}
