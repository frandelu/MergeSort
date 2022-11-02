package mergesort

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	medio := len(arr) / 2
	return merge(MergeSort(arr[:medio]), MergeSort(arr[medio:]))
}

func merge(izq, der []int) []int {
	largo := len(izq) + len(der)
	arr := make([]int, largo)
	i, j := 0, 0
	for k := 0; k < largo; k++ {
		if i > len(izq)-1 && j <= len(der)-1 {
			arr[k] = der[j]
			j++
		} else if j > len(der)-1 && i <= len(izq)-1 {
			arr[k] = izq[i]
			i++
		} else if izq[i] < der[j] {
			arr[k] = izq[i]
			i++
		} else {
			arr[k] = der[j]
			j++
		}
	}
	return arr
}
