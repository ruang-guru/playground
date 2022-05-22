package sortKM

// TODO: answer here

//concrete strategy implementation
type AscendingSort struct{}

func (as *AscendingSort) Sort(array []int) {
	//choose any sort algo you want
	// TODO: answer here
	n := len(array)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}
		array[i], array[minIdx] = array[minIdx], array[i]
	}
}
