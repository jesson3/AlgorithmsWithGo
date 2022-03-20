package sort

import "fmt"

type Comparable interface {
	lessThan(Comparable) bool
}

func less(v Comparable, w Comparable) bool {
	return v.lessThan(w)
}

func exch(a []Comparable, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func show(a []Comparable) {
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

func IsSorted(a []Comparable) bool {
	for i := 1; i < len(a); i++ {
		if less(a[i], a[i-1]) {
			return false
		}
	}
	return true
}
