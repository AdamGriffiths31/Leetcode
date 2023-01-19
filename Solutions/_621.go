type task struct {
	freq    int
	cooling int
}

func leastInterval(tasks []byte, n int) int {
	hashMap := make(map[byte]int)
	heap := make([]*task, 0)

	for _, char := range tasks {
		hashMap[char]++
	}

	for _, v := range hashMap {
		heap = append(heap, &task{v, 0})
	}

	for i := len(heap) / 2; i >= 0; i-- {
		heapify(&heap, i, len(heap))
	}

	time := 0
	coolingTime := n

	for len(heap) != 0 {
		if heap[0].cooling <= time {
			heap[0].freq -= 1
			if heap[0].freq == 0 {
				//pop the elem
				heap[0] = heap[len(heap)-1]
				j := len(heap) - 1
				heap = heap[:j]
				heapify(&heap, 0, j)
			} else {
				heap[0].cooling += coolingTime + 1
				heapify(&heap, 0, len(heap))
			}
		}
		time++
	}

	return time
}

func heapify(heap *[]*task, parent, size int) {
	leftChild := 2*parent + 1
	rightChild := 2*parent + 2

	largSmall := parent

	if leftChild < size && (*heap)[leftChild].cooling < (*heap)[largSmall].cooling {
		largSmall = leftChild
	} else if leftChild < size && (*heap)[leftChild].cooling == (*heap)[largSmall].cooling {
		if (*heap)[leftChild].freq > (*heap)[largSmall].freq {
			largSmall = leftChild
		}
	}

	if rightChild < size && (*heap)[rightChild].cooling < (*heap)[largSmall].cooling {
		largSmall = rightChild
	} else if rightChild < size && (*heap)[rightChild].cooling == (*heap)[largSmall].cooling {
		if (*heap)[rightChild].freq > (*heap)[largSmall].freq {
			largSmall = rightChild
		}
	}

	if parent != largSmall {
		(*heap)[largSmall], (*heap)[parent] = (*heap)[parent], (*heap)[largSmall]
		heapify(heap, largSmall, size)
	}
}