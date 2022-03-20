package sort

var aux []Comparable // auxiliary array for merges

func MergeSort(a []Comparable) {
	aux = make([]Comparable, len(a)) // Allocate space just once.
	sort(a, 0, len(a)-1)
}

func sort(a []Comparable, lo, hi int) {
	// Sort a[lo..hi].
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sort(a, lo, mid)      // Sort left half.
	sort(a, mid+1, hi)    // Sort right half.
	merge(a, lo, mid, hi) // Merge result.
}

func merge(a []Comparable, lo, mid, hi int) {
	// Merge a[lo..mid] with a[mid+1..hi]
	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}

	for k := lo; k < hi; i++ {
		// Merge back to a[lo..hi].
		if j > mid {
			a[k] = aux[j]
			j += 1
		} else if j > hi {
			a[k] = aux[i]
			i += 1
		} else if less(aux[j], aux[i]) {
			a[k] = aux[j]
			j += 1
		} else {
			a[k] = aux[i]
			i += 1
		}
	}
}
