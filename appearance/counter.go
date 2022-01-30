package appearance

import (
	"bytes"
	"sort"
)

type Counter struct {
	dictionary *Dictionary
	splitter   []rune
}

type Result struct {
	Word       string `json:"word"`
	Appearance int    `json:"appearance"`
}

func NewCounter() *Counter {
	return &Counter{
		dictionary: NewDictionary(),
		splitter:   []rune{' ', '\n', '\t', '.', ',', '?', '!', ':'},
	}
}

func (c *Counter) Top(b []byte, max int) []Result {
	if b == nil {
		return make([]Result, 0)
	}

	c.dictionary.InsertMany(c.split(b))

	processor := NewHeap(max)
	for _, term := range c.dictionary.VisitAll() {
		processor.Insert(term)
	}

	return c.descSort(processor.Data())
}

func (c *Counter) descSort(res []*Term) []Result {
	sort.Slice(res, func(i, j int) bool {
		return res[i].Appearance > res[j].Appearance
	})

	var result []Result
	for _, term := range res {
		result = append(result, Result{
			Word:       string(term.CompleteWord),
			Appearance: term.Appearance,
		})
	}
	return result
}

func (c *Counter) split(b []byte) [][]byte {
	return bytes.FieldsFunc(b, func(r rune) bool {
		return c.mustSplit(r)
	})
}

func (c *Counter) mustSplit(r rune) bool {
	for _, splitter := range c.splitter {
		if r == splitter {
			return true
		}
	}

	return false
}
