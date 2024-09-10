package minheap

// MinHeap represents a min heap data structure
type MinHeap struct {
    heap []int
}

// NewMinHeap creates a new empty min heap
func NewMinHeap() *MinHeap {
    return &MinHeap{
        heap: []int{},
    }
}

// Insert adds a new element to the heap
func (h *MinHeap) Insert(val int) {
    h.heap = append(h.heap, val)
    h.heapifyUp(len(h.heap) - 1)
}

// Extract removes and returns the minimum element from the heap
func (h *MinHeap) Extract() (int, bool) {
    if len(h.heap) == 0 {
        return 0, false
    }

    min := h.heap[0]
    lastIndex := len(h.heap) - 1
    h.heap[0] = h.heap[lastIndex]
    h.heap = h.heap[:lastIndex]
    h.heapifyDown(0)

    return min, true
}

// Peek returns the minimum element without removing it
func (h *MinHeap) Peek() (int, bool) {
    if len(h.heap) == 0 {
        return 0, false
    }
    return h.heap[0], true
}

// Size returns the number of elements in the heap
func (h *MinHeap) Size() int {
    return len(h.heap)
}

// heapifyUp maintains the heap property by moving an element up
func (h *MinHeap) heapifyUp(index int) {
    for index > 0 {
        parentIndex := (index - 1) / 2
        if h.heap[parentIndex] <= h.heap[index] {
            break
        }
        h.heap[parentIndex], h.heap[index] = h.heap[index], h.heap[parentIndex]
        index = parentIndex
    }
}

// heapifyDown maintains the heap property by moving an element down
func (h *MinHeap) heapifyDown(index int) {
    lastIndex := len(h.heap) - 1
    for {
        leftChild := 2*index + 1
        rightChild := 2*index + 2
        smallest := index

        if leftChild <= lastIndex && h.heap[leftChild] < h.heap[smallest] {
            smallest = leftChild
        }
        if rightChild <= lastIndex && h.heap[rightChild] < h.heap[smallest] {
            smallest = rightChild
        }

        if smallest == index {
            break
        }

        h.heap[index], h.heap[smallest] = h.heap[smallest], h.heap[index]
        index = smallest
    }
}
