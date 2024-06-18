package main

func sortDescInt(a [8]int) [8]int {
	var quicksortDesc func(a *[8]int, low, high int)
	var partitionDesc func(a *[8]int, low, high int) int

	quicksortDesc = func(a *[8]int, low, high int) {
		if low < high {
			pi := partitionDesc(a, low, high)
			quicksortDesc(a, low, pi-1)
			quicksortDesc(a, pi+1, high)
		}
	}

	partitionDesc = func(a *[8]int, low, high int) int {
		pivot := a[high]
		i := low - 1
		for j := low; j < high; j++ {
			if a[j] > pivot { // изменено на `>` для сортировки по убыванию
				i++
				a[i], a[j] = a[j], a[i]
			}
		}
		a[i+1], a[high] = a[high], a[i+1]
		return i + 1
	}

	quicksortDesc(&a, 0, len(a)-1)
	return a
}

func sortAscInt(a [8]int) [8]int {
	var quicksortAsc func(a *[8]int, low, high int)
	var partitionAsc func(a *[8]int, low, high int) int

	quicksortAsc = func(a *[8]int, low, high int) {
		if low < high {
			pi := partitionAsc(a, low, high)
			quicksortAsc(a, low, pi-1)
			quicksortAsc(a, pi+1, high)
		}
	}

	partitionAsc = func(a *[8]int, low, high int) int {
		pivot := a[high]
		i := low - 1
		for j := low; j < high; j++ {
			if a[j] < pivot {
				i++
				a[i], a[j] = a[j], a[i]
			}
		}
		a[i+1], a[high] = a[high], a[i+1]
		return i + 1
	}

	quicksortAsc(&a, 0, len(a)-1)
	return a
}

func sortDescFloat(a [8]float64) [8]float64 {
	var quicksortDesc func(a *[8]float64, low, high int)
	var partitionDesc func(a *[8]float64, low, high int) int

	quicksortDesc = func(a *[8]float64, low, high int) {
		if low < high {
			pi := partitionDesc(a, low, high)
			quicksortDesc(a, low, pi-1)
			quicksortDesc(a, pi+1, high)
		}
	}

	partitionDesc = func(a *[8]float64, low, high int) int {
		pivot := a[high]
		i := low - 1
		for j := low; j < high; j++ {
			if a[j] > pivot {
				i++
				a[i], a[j] = a[j], a[i]
			}
		}
		a[i+1], a[high] = a[high], a[i+1]
		return i + 1
	}

	quicksortDesc(&a, 0, len(a)-1)
	return a
}

func sortAscFloat(a [8]float64) [8]float64 {
	var quicksortAsc func(a *[8]float64, low, high int)
	var partitionAsc func(a *[8]float64, low, high int) int

	quicksortAsc = func(a *[8]float64, low, high int) {
		if low < high {
			pi := partitionAsc(a, low, high)
			quicksortAsc(a, low, pi-1)
			quicksortAsc(a, pi+1, high)
		}
	}

	partitionAsc = func(a *[8]float64, low, high int) int {
		pivot := a[high]
		i := low - 1
		for j := low; j < high; j++ {
			if a[j] < pivot {
				i++
				a[i], a[j] = a[j], a[i]
			}
		}
		a[i+1], a[high] = a[high], a[i+1]
		return i + 1
	}

	quicksortAsc(&a, 0, len(a)-1)
	return a
}
