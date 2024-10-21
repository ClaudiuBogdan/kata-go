package huffmancoding

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Freq == pq[j].Freq {
		if pq[i].Char != 0 && pq[j].Char != 0 {
			return pq[i].Char < pq[j].Char
		}
		// For internal nodes (Char == 0), use a stable comparison
		return i < j
	}
	return pq[i].Freq < pq[j].Freq
}
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Node)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
