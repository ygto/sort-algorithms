package bubble

// process run count
var runCount int
/*
func main() {
	var arr []int

	//best case
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//worst case
	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Println(runCount, bubbleSort(arr))

}*/

func Sort(arr []int) []int {

	l := len(arr)
	var changed bool

	for i := 0; i < l-1; i++ {
		changed = false
		for j := 0; j < l-1; j++ {
			runCount++
			if arr[j] > arr[j+1] {
				changed = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if changed == false {
			break
		}
	}
	return arr
}
