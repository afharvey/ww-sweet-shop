package packer

// GetOrderSize returns the smallest order we can send to the customer. 
// No less than requested and no more than necessary.
func GetOrderSize(requested int, kinds []int) (int, int) {
	var orderSize int
	var winner int
	for _, i := range kinds {
		target := MakeTarget(requested, i)
		if target < orderSize || orderSize == 0 {
			winner = i
			orderSize = target
		}
	}
	return orderSize, winner
}
