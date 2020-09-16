package sort

import "fmt"

func binarySearch(arr []int, start, end, search int) int {
	//var middle int = (end - start) / 2
	middle := (end - start) / 2
	if middle == 0 && search != arr[middle]{
		panic(fmt.Sprintf("could not found search item %d", search))
	}

	middle += start
	if search == arr[middle] {
		return middle
	}

	if search > arr[middle] {
		return binarySearch(arr, middle, end, search)
	}

	return binarySearch(arr, start, middle, search)
}

func bubble(arr *[]int, len int)  {
	for i := 0; i < len; i++ {
		for j := 0; j < len - i; j++ {
			if (*arr)[j] > (*arr)[j + 1] {
				(*arr)[j + 1], (*arr)[j] = (*arr)[j], (*arr)[j + 1]
			}
		}
	}
}

func quick(arr *[]int, start, end int)  {
	if start >= end {
		return
	}

	left, right := start, end
	pointer := (*arr)[start]
	for left < right {
		if (*arr)[left] > (*arr)[right] {
			(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]
		}

		if pointer == (*arr)[left] {
			right--
		} else {
			left++
		}
	}

	if start < left {
		quick(arr, start, left)
	}

	if left < end {
		quick(arr, left + 1, end)
	}
}

func shell(arr *[]int, len int)  {
	gap := len / 2
	for gap > 0 {
		for i := 0; i < len; i++ {
			for j := 0; j < len - i - gap; j++ {
				if (*arr)[j] > (*arr)[j + gap] {
					(*arr)[j], (*arr)[j + gap] = (*arr)[j + gap], (*arr)[j]
				}
			}
		}
		gap /= 2
	}
}

func heap(arr *[]int, len int) {
	for i := 0; i < len; i++ {
		adjustHeap(arr, len - i)
		(*arr)[0], (*arr)[len - i - 1] = (*arr)[len - i - 1], (*arr)[0]
	}
}

func adjustHeap(arr *[]int, len int) {
	for i := len / 2; i >= 0; i-- {
		left, right := 2 * i + 1, 2 * i + 2
		if left < len && (*arr)[i] > (*arr)[left] {
			(*arr)[i], (*arr)[left] = (*arr)[left], (*arr)[i]
		}

		if right < len && (*arr)[i] > (*arr)[right] {
			(*arr)[i], (*arr)[right] = (*arr)[right], (*arr)[i]
		}
	}
}

func selectSort(arr *[]int, len int)  {
	for i := 0; i < len; i++ {
		max := (*arr)[0]
		max_index := 0
		for j := 0; j < len - i; j++ {
			if (*arr)[j] > max {
				max = (*arr)[j]
				max_index = j
			}
		}

		if max_index < len - i - 1 {
			(*arr)[max_index], (*arr)[len - i - 1] = (*arr)[len - i - 1], (*arr)[max_index]
		}
	}
}

func insertSort(arr *[]int, len int)  {
	for i := 0; i < len - 1; i++ {
		insert(arr, i + 1, (*arr)[i + 1])
	}
}

func insert(arr *[]int, len int, data int) {
	i := 0
	for ; i < len; i++ {
		if (*arr)[i] > data {
			break
		}
	}

	for j := len; j > i; j-- {
		(*arr)[j] = (*arr)[j - 1]
	}

	(*arr)[i] = data
}

func mergeSort(arr *[]int, start int, end int) {
	if (start < end) {
		middle := start + (end - start) / 2
		mergeSort(arr, start, middle)
		mergeSort(arr, middle + 1, end)
		merge(arr, start, middle, middle + 1, end)
	}
}

func merge(arr *[]int, left_start int, left_end int, right_start int, right_end int) {
	start, end := left_start, right_end
	ret := []int{}
	for left_start <= left_end && right_start <= right_end {
		if (*arr)[left_start] <= (*arr)[right_start] {
			ret = append(ret, (*arr)[left_start])
			left_start++
		} else {
			ret = append(ret, (*arr)[right_start])
			right_start++
		}
	}

	for left_start <= left_end {
		ret = append(ret, (*arr)[left_start])
		left_start++
	}

	for right_start <= right_end {
		ret = append(ret, (*arr)[right_start])
		right_start++
	}

	for i := start; i <= end; i++ {
		(*arr)[i] = ret[i - start]
	}
}
