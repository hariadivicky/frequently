package appearance

import (
	"testing"
)

func TestWordsGetAppearance(t *testing.T) {
	data, err := readSplittedTestSeed()
	if err != nil {
		t.Fatalf("can not read test data: %v", err)
	}

	wTree := NewDictionary()
	wTree.InsertMany(data)

	t.Run("must has correct appearance count when it is valid word", func(st *testing.T) {
		appearance, found := wTree.GetAppearance([]byte("et"))
		if !found {
			st.Fatalf("expected word to be valid result; got %v", found)
		}

		if appearance != 310 {
			st.Fatalf("expected word appearance to be 310; got %v", appearance)
		}
	})

	t.Run("must be invalid result when searching for unknown word", func(st *testing.T) {
		_, found := wTree.GetAppearance([]byte("ha"))
		if found {
			st.Fatalf("expected word to be not found; got %v", found)
		}

		_, found = wTree.GetAppearance([]byte{})
		if found {
			st.Fatalf("expected word to be not found; got %v", found)
		}
	})
}

func TestWordsVisitAll(t *testing.T) {
	wTree := NewDictionary()

	wTree.Insert([]byte("hello"))
	wTree.Insert([]byte("hallo"))
	wTree.Insert([]byte("numero"))
	wTree.Insert([]byte("halo"))
	wTree.Insert([]byte("halo"))
	wTree.Insert([]byte("halo"))

	words := wTree.VisitAll()
	if len(words) != 4 {
		t.Fatalf("expected words len to be 3; got %v", len(words))
	}
}

func BenchmarkDictionaryInsert(b *testing.B) {
	data, err := readSplittedTestSeed()
	if err != nil {
		b.Fatalf("can not read test data: %v", err)
	}

	for i := 0; i < b.N; i++ {
		dict := NewDictionary()
		dict.InsertMany(data)
	}
}

func BenchmarkDictionaryVisitAll(b *testing.B) {
	data, err := readSplittedTestSeed()
	if err != nil {
		b.Fatalf("can not read test data: %v", err)
	}

	dict := NewDictionary()
	dict.InsertMany(data)

	for i := 0; i < b.N; i++ {
		dict.VisitAll()
	}
}
