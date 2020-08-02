package packer

// Exclude all kinds of box which are larger than we require.
// There's also an opportunity to quit early if we find an exact match.
func ExcludeUnhelpfulKinds(target, baseKind int, kinds []int) []int {
	o := make([]int, 0)
	for _, i := range kinds {

		if i > target {
			continue
		}

		if i == target {
			return []int{i} // quit early
		}

		// Exclude kinds which are not multiples of baseKind.
		if i != baseKind && i%baseKind != 0 {
			continue
		}

		o = append(o, i)
	}
	return o
}
