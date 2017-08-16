/*
Package floats provides a set of helper routines for dealing with slices
of float64. The functions avoid allocations to allow for use within tight
loops without garbage collection overhead.

The convention used is that when a slice is being modified in place, it has
the name dst.
*/
package floats // import "gonum.org/v1/gonum/floats"
