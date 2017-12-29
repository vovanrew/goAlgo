package searching

import "math"

//////////////////////////BINARY SEARCH///////////////////////////

func BinarySearch(array []int, find int) int {
	return BinarySearchRoutine(array, 0, len(array), find)
}

func BinarySearchRoutine(array []int, left int, right, find int) int {
	middle := 0

	for left <= right {
		middle = (left + right) / 2
		if array[middle] == find {
			return middle
		}

		if array[middle] < find {
			left = middle + 1
		} else if array[middle] > find {
			right = middle - 1
		}
	}

	return -1
}

//////////////////////////JUMP SEARCH///////////////////////////

func JumpSearch(array []int, find int) int {
	jumpAgregator := int(math.Sqrt(float64(len(array))))
	previousJump := 0

	for jumpAgregator < len(array) && array[jumpAgregator] <= find{
		previousJump = jumpAgregator
		jumpAgregator += jumpAgregator
	}

	if array[previousJump] == find {
		return previousJump
	}

	for i := previousJump; i < int(math.Min(float64(jumpAgregator), float64(len(array)))); i++ {
		if array[i] == find {
			return i
		}
	}

	return -1
}

//////////////////////////EXPONENTIAL SEARCH///////////////////////////

func ExponentialSearch(array []int, find int) int {
	i := 0

	if array[i] == find {
		return i
	}

	i = 1
	for i < int(math.Min(float64(find), float64(len(array)))) {
		i = int(math.Min(float64(i * 2), float64(len(array))))
	}

	return BinarySearchRoutine(array, i / 2, i, find)
}