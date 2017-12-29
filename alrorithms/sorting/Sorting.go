package sorting

////////////////////BUBLE SORT////////////////

func BubleSort(array []int) {

	if cap(array) == 0 {
		return;
	}

	temp := array[0]

	for i := 0; i < len(array) - 1; i++ {
		for j := 0; j < len(array) - 1; j++ {
			if array[j] > array[j + 1] {
				temp = array[j]
				array[j] = array[j + 1]
				array[j + 1] = temp
			}
		}
	}
}

////////////////////SELECTION SORT////////////////

func SelectionSort(array []int) {
	if cap(array) == 0 {
		return;
	}

	maxIndex := 0
	temp := array[0]

	for i := 0; i < len(array) - 1; i++ {
		maxIndex = 0
		for j := 0; j < len(array) - i; j++ {
			if array[j] > array[maxIndex] {
				maxIndex = j;
			}
		}

		temp = array[len(array) - i - 1];
		array[len(array) - i - 1] = array[maxIndex];
		array[maxIndex] = temp;
	}
}

////////////////////INSERTION SORT////////////////

func InsertionSort(array []int) {
	if cap(array) == 0 {
		return;
	}

	current := 0
	j := 0

	for i := 0; i < len(array) - 1; i++ {

		j = i + 1;
		current = array[j]

		for j > 0 && current < array[j - 1] {
			array[j] = array[j - 1]
			j--
		}

		array[j] = current
	}
}

////////////////////MERGE SORT////////////////

func copy(src []int, dst []int, from int, to int, dstFrom int) []int {
	for from <= to {
		dst[dstFrom] = src[from]
		dstFrom++
		from++;
	}

	return dst
}

func merge(array []int, left int, rihgt int) {
	middle := (rihgt + left) / 2
	l := left
	r := middle + 1

	tempArray := make([]int, rihgt + 1)
	iterTemp := 0

	for l <= middle && r <= rihgt {
		if array[l] < array[r] {
			tempArray[iterTemp] = array[l]
			l++
		} else {
			tempArray[iterTemp] = array[r]
			r++
		}
		iterTemp++;
	}

	copy(array, tempArray, l, middle, iterTemp)
	copy(array, tempArray, r, rihgt, iterTemp)

	iterTemp = 0
	for i := left; i <= rihgt; i++ {
		array[i] = tempArray[iterTemp]
		iterTemp++
	}
}

func MergeSort(array []int, left int, right int) {
	if(left < right) {
		middle := (right + left) / 2
		MergeSort(array, left, middle)
		MergeSort(array, middle + 1, right)
		merge(array, left, right)
	}
}

////////////////////QUICK SORT////////////////

func partitions(array []int, left, right int) int {
	pivot := array[right]
	i := left - 1
	temp := 0

	for j := left; j < right; j++ {
		if array[j] <= pivot {
			i++
			temp = array[i]
			array[i] = array[j]
			array[j] = temp
		}
	}

	i++
	array[right] = array[i]
	array[i] = pivot

	return i
}

func QuickSort(array []int, left, right int)  {
	if left < right {
		pivot := partitions(array, left, right)
		QuickSort(array, left, pivot - 1)
		QuickSort(array, pivot + 1, right)
	}
}

////////////////////QUICK SORT////////////////

func heapify(array []int, currentParent int, size int) {
	largest := currentParent
	left := currentParent * 2 + 1
	right := currentParent * 2 + 2

	if left < size && array[left] > array[largest] {
		largest = left
	}

	if right < size && array[right] > array[largest] {
		largest = right
	}

	if largest != currentParent {
		temp := array[largest]
		array[largest] = array[currentParent]
		array[currentParent] = temp
		heapify(array, largest, size)
	}
}

func HeapSort(array []int)  {

	for i := len(array) / 2 - 1; i >= 0; i-- {
		heapify(array, i, len(array))
	}

	temp := 0
	for i := 0; i < len(array); i++ {
		temp =  array[len(array) - i - 1]
		array[len(array) - i - 1] = array[0]
		array[0] = temp
		heapify(array, 0, len(array) - i - 1)
	}
}

////////////////////COUNTING SORT////////////////

func maxElem(array []int) int {
	max := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}

	return max
}

func CountingSort(array []int) {
	counter := make([]int, maxElem(array) + 1)

	for i := 0; i < len(array); i++ {
		counter[array[i]]++
	}

	for i := 1; i < len(counter); i++ {
		counter[i] += counter[i - 1]
	}

	temp := make([]int, len(array))

	for i := 0; i < len(array); i++ {
		temp[counter[array[i]] - 1] = array[i]
		counter[array[i]]--
	}

	copy(temp, array, 0, len(array) - 1, 0)
}

////////////////////RADIX SORT////////////////

func countingSortRoutine(array []int, eps int) {
	counter := make([]int, 10)

	for i := 0; i < len(array); i++ {
		counter[(array[i] / eps) % 10]++
	}

	for i := 1; i < len(counter); i++ {
		counter[i] += counter[i - 1]
	}

	temp := make([]int, len(array))

	for i := len(array) - 1; i >= 0; i-- {
		temp[counter[(array[i] / eps) % 10] - 1] = array[i]
		counter[(array[i] / eps) % 10]--
	}

	copy(temp, array, 0, len(array) -1, 0)
}

func RadixSort(array []int) {
	maxElem := maxElem(array)

	for i := 1; maxElem / i * 10 > 0; i *= 10 {
		countingSortRoutine(array, i)
	}
}