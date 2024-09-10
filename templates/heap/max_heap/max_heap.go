package maxheap

// MaxHeap represents a max heap data structure
type MaxHeap struct {
    heap []int
}

// NewMaxHeap creates a new empty max heap
func NewMaxHeap() *MaxHeap {
    return &MaxHeap{
        heap: []int{},
    }
}

// Insert adds a new element to the heap
func (h *MaxHeap) Insert(val int) {
    h.heap = append(h.heap, val)
    h.heapifyUp(len(h.heap) - 1)
}

// Extract removes and returns the maximum element from the heap
func (h *MaxHeap) Extract() (int, bool) {
    if len(h.heap) == 0 {
        return 0, false
    }

    max := h.heap[0]
    lastIndex := len(h.heap) - 1
    h.heap[0] = h.heap[lastIndex]
    h.heap = h.heap[:lastIndex]
    h.heapifyDown(0)

    return max, true
}

// Peek returns the maximum element without removing it
func (h *MaxHeap) Peek() (int, bool) {
    if len(h.heap) == 0 {
        return 0, false
    }
    return h.heap[0], true
}

// Size returns the number of elements in the heap
func (h *MaxHeap) Size() int {
    return len(h.heap)
}

// heapifyUp maintains the heap property by moving an element up
func (h *MaxHeap) heapifyUp(index int) {
    for index > 0 {
        parentIndex := (index - 1) / 2
        if h.heap[parentIndex] >= h.heap[index] {
            break
        }
        h.heap[parentIndex], h.heap[index] = h.heap[index], h.heap[parentIndex]
        index = parentIndex
    }
}

// heapifyDown maintains the heap property by moving an element down
func (h *MaxHeap) heapifyDown(index int) {
    lastIndex := len(h.heap) - 1
    for {
        leftChild := 2*index + 1
        rightChild := 2*index + 2
        largest := index

        if leftChild <= lastIndex && h.heap[leftChild] > h.heap[largest] {
            largest = leftChild
        }
        if rightChild <= lastIndex && h.heap[rightChild] > h.heap[largest] {
            largest = rightChild
        }

        if largest == index {
            break
        }

        h.heap[index], h.heap[largest] = h.heap[largest], h.heap[index]
        index = largest
    }
}
