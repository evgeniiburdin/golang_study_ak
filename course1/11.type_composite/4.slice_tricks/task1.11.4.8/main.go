package main

func Shift(xs []int) (int, []int) {
	// To shift the slice, we can just append itself(except the last element) to its last element,
	// so the last element becomes the first one of the new slice, and the other elements append to it
	// in theirs original order.
	xs = append(xs[len(xs)-1:], xs[:len(xs)-1]...)
	return xs[0], xs
}
