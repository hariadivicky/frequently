package appearance

import (
	"testing"
)

func TestHeapInsert(t *testing.T) {
	terms := NewHeap(3)

	c1 := &Term{
		Appearance:   1,
		CompleteWord: []byte("c1"),
	}
	c2 := &Term{
		Appearance:   2,
		CompleteWord: []byte("c2"),
	}
	c3 := &Term{
		Appearance:   3,
		CompleteWord: []byte("c3"),
	}
	c4 := &Term{
		Appearance:   4,
		CompleteWord: []byte("c4"),
	}
	c5 := &Term{
		Appearance:   5,
		CompleteWord: []byte("c4"),
	}

	terms.Insert(c1)
	terms.Insert(c2)
	terms.Insert(c3)
	terms.Insert(c4)
	terms.Insert(c5)

	result := terms.Data()
	expected := []int{3, 4, 5}

	for k, term := range result {
		if expected[k] != term.Appearance {
			t.Errorf("expected char %s appearance to be %d; got %d", string(term.CompleteWord), expected[k], term.Appearance)
		}
	}

	terms.Insert(c1)
	result = terms.Data()
	for _, term := range result {
		if term.Appearance == c1.Appearance {
			t.Fatalf("expected c1 not to be inserted; got %v", term.Appearance)
		}
	}
}

func BenchmarkHeapInsert(b *testing.B) {
	data, err := readSplittedTestSeed()
	if err != nil {
		b.Fatalf("can not read test data: %v", err)
	}

	dict := NewDictionary()
	dict.InsertMany(data)

	heap := NewHeap(10)

	for i := 0; i < b.N; i++ {
		for _, term := range dict.VisitAll() {
			heap.Insert(term)
		}
	}
}
