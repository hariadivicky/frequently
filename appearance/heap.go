package appearance

type Heap struct {
	data []*Term
}

func NewHeap(maxCapacity int) *Heap {
	return &Heap{
		data: make([]*Term, 0, maxCapacity),
	}
}

func (h *Heap) Insert(c *Term) {
	if len(h.data) == 0 {
		h.data = append(h.data, c)
		return
	}

	if len(h.data) < cap(h.data) {
		h.data = append(h.data, c)
		h.rebuild(0)
		return
	}

	if c.Appearance <= h.data[0].Appearance {
		return
	}

	h.data[0] = c
	h.rebuild(0)
}

func (h *Heap) Data() []*Term {
	return h.data
}

func (h *Heap) rebuild(currentIndex int) {
	leftIndex := 2*currentIndex + 1
	rightIndex := leftIndex + 1
	smaller := currentIndex

	if h.isCurrentNodeSmaller(leftIndex, smaller) {
		smaller = leftIndex
	}

	if h.isCurrentNodeSmaller(rightIndex, smaller) {
		smaller = rightIndex
	}

	if smaller != currentIndex {
		h.data[smaller], h.data[currentIndex] = h.data[currentIndex], h.data[smaller]
		h.rebuild(smaller)
	}
}

func (h *Heap) isCurrentNodeSmaller(a, b int) bool {
	return a < len(h.data) && h.data[a].Appearance < h.data[b].Appearance
}
