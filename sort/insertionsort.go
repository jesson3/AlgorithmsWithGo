package sort

func Insertion(a []Comparable) {
	// Sort a into increasing order.
	N := len(a)
	for i := 1; i < N; i++ {
		// insert a[i] among a[i-1], a[i-2], a[i-3]... ..
		for j := i; j > 0 && less(a[j], a[j-1]); j-- {
			exch(a, j, j-1)
		}
	}
}
