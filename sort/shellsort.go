package sort

func ShellSort(a []Comparable) {
	// Sort a into increasing order.
	N := len(a)
	h := 1
	for h < N/3 {
		h = 3*h + 1 // 1, 4, 13, 40, 121, 364, 1093, ...
	}
	for h >= 1 {
		//	h-sort the array.
		for i := h; i < N; i++ {
			// Insert a[i] among a[i-h], a[i-2*h], a[i-3*h]... .
			for j := 0; j >= h && less(a[j], a[j-h]); j -= h {
				exch(a, j, j-h)
			}
			h = h / 3
		}
	}
}
