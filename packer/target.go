package packer

// Returns the next amount you can make from 'requested' when counting in 'unitSize'.
// For example, given 251 and 250 it's 500.
// For example, given 251 and 10 it's 260.
// For example, given 251 and 5 it's 255.
// For example, given 251 and 2 it's 252.
// For example, given 251 and 1 it's 251.
func MakeTarget(requested, unitSize int) int {
	f := requested / unitSize
	if requested %unitSize != 0 { // round up
		f += 1
	}
	return f * unitSize
}
