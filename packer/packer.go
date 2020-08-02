package packer

import "sort"

func MakeOrder(ordered int, kinds []int) []int {

	// Everything in this solution depends on the kinds being sorted.
	sort.Slice(kinds, func(i, j int) bool {
		return kinds[i] > kinds[j] // descending
	})

	// 1. Find the optimum order size and the base kind.
	orderSize, baseKind := GetOrderSize(ordered, kinds)

	// 2. Exclude all the kinds which are of no use.
	kindsForThisOrder := ExcludeUnhelpfulKinds(orderSize, baseKind, kinds)

	// 3. We know the optimal size and have a list of kinds we know will make it.
	//    All we do now is add them to the order using as few as possible.
	return picker(orderSize, kindsForThisOrder)
}

// Given a target which we can hit and kinds which will add up to it then
// the final task is to add them up using the smallest number of elements.
// For example if we want 100 and we have 80 and 10 then we'd want 80+10+10.
func picker(target int, kinds []int) []int {
	output := make([]int, 0)
	remaining := target

	for _, kind := range kinds {

		// how many of this kind of box can we fit into the order?
		numberOfBoxes := remaining / kind
		remaining -= kind * numberOfBoxes

		for i := 0; i < numberOfBoxes; i++ {
			output = append(output, kind)
		}

	}

	return output
}
