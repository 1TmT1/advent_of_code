package helpers

import (
	"cmp"
)

// Find partition position
func partition[T cmp.Ordered](slice []T, low, high int) ([]T, int) {
	// Choose pivot based on middle value - check first, middle and last
	pivotOptions := []T{slice[0], slice[len(slice)/2-1], slice[len(slice)-1]}
	Insertionsort(pivotOptions)
	pivot := pivotOptions[1] // After sorting median element will be in the middle [0, !1!, 2]

	// Index of greater element
	index := low - 1

	// Compare each element with pivot
	for i := low; i < high; i++ {
		if slice[i] <= pivot {
			// Current smaller than pivot -> swap it with greater element
			index = index + 1

			// Make the swap
			slice[index], slice[i] = slice[i], slice[index]
		}
	}

	// Swap pivot element with the greater element at index
	slice[index+1], slice[high] = slice[high], slice[index+1]

	// Return the position where the partition done
	return slice, index + 1
}

// Implementation of the quicksort algorithm => complexity - time avg: O(nlogn) worst: O(n^2) # Note: chances are low since the pivot more probable to be close to median because how it gets chosen in partition function, space = O(1) | Not stable sorting algorithm which means relative order of equal items is not preserved
func Quicksort[T cmp.Ordered](slice []T, low, high int) []T {
	if low < high {
		// Find pivot where element smaller than pivot goes left, if larger then right
		slice, pivot := partition(slice, low, high)

		// Recursive call on left side of the pivot
		Quicksort(slice, low, pivot-1)

		// Recursive call on right side of the pivot
		Quicksort(slice, pivot+1, high)
	}

	return slice
}

// Heapify subtree of root in index i (reshape binary heap data structure from an array of elements)
func heapify[T cmp.Ordered](slice []T, heapSize, index int) {
	largest := index // Initialize largest as root
	left := 2*index + 1
	right := 2*index + 2

	// Left child exists and greater than root
	if left < heapSize && slice[largest] < slice[left] {
		largest = left
	}

	// Right child exists and greater than root
	if right < heapSize && slice[largest] < slice[right] {
		largest = right
	}

	// Change root if i different than largest like we initialized above
	if largest != index {
		slice[index], slice[largest] = slice[largest], slice[index] // Make the swap
		heapify(slice, heapSize, largest)                           // Heapify the root until there is no difference and need to swap                      //
	}
}

// Implementation of the heap sort algorithm => complexity - time = O(nlogn), space = O(1) | Not stable sorting algorithm because of the swap which can destroy the relative ordering of elements
func Heapsort[T cmp.Ordered](slice []T) {
	heapSize := len(slice)

	// Build max heap - binary tree which value of node greater than or equal to children values
	for index := heapSize/2 - 1; index > -1; index-- {
		heapify(slice, heapSize, index)
	}

	// Put max element at last position and shift leaf to same position
	for index := heapSize - 1; index > 0; index-- {
		slice[index], slice[0] = slice[0], slice[index] // Make the swap
		heapify(slice, index, 0)
	}
}

// Implementation of the insertion sort algorithm => complexity - time = O(n^2), space = O(1) | Stable sorting algorithm - doesn't destroy order of indexes with same values
func Insertionsort[T cmp.Ordered](slice []T) {
	for index, value := range slice {
		j := index - 1
		for j >= 0 && value < slice[j] {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = value
	}
}

// Implementation of the introsort algorithm => complexity - time = O(nlogn), space = O(logn) | Not stable sorting algorithm which means relative order of equal items is not preserved
func Introsort[T cmp.Ordered](slice []T, maxDepth int) {
	n := len(slice)

	// Statistically probable that insertion sort tends to be time complexity O(n) with small collections of data # Note: here the choice is slice length of 20
	if n < 20 {
		Insertionsort(slice)
	} else if maxDepth == 0 {
		Heapsort(slice)
	} else {
		Quicksort(slice, 0, len(slice)-1)
	}
}
