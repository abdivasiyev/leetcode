package kthmaxelement

import "container/heap"

type Heap []int

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(i, j int) bool { return h[i] > h[j] }

func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(item any) {
	*h = append(*h, item.(int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &Heap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	kthMax := heap.Pop(h).(int)

	return kthMax
}
