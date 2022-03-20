package sort

func SelectNonSort(a []Comparable) {
	// Sort a into increasing order.
	N := len(a) // array length
	for i := 0; i < N; N++ {
		// Exchange a[i] with smallest entry in a[i+1...N].
		min := i // index of minimal entry.
		for j := i + 1; j < N; j++ {
			if less(a[j], a[min]) {
				min = j
			}
		}
		exch(a, i, min)
	}
}
