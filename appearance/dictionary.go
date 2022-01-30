package appearance

type Dictionary struct {
	root *Term
}

type TermMap map[byte]*Term

type Term struct {
	char         byte
	isComplete   bool
	children     TermMap
	Appearance   int
	CompleteWord []byte
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		root: &Term{
			children: make(TermMap),
		},
	}
}

func (d *Dictionary) InsertMany(words [][]byte) {
	for _, word := range words {
		d.Insert(word)
	}
}

func (d *Dictionary) Insert(word []byte) {
	currentChar := d.root

	for _, char := range word {
		if child, ok := currentChar.children[char]; ok {
			currentChar = child
		} else {
			currentChar.children[char] = &Term{
				char:     char,
				children: make(TermMap),
			}

			currentChar = currentChar.children[char]
		}
	}

	if !currentChar.isComplete {
		currentChar.isComplete = true
		currentChar.CompleteWord = word
	}

	currentChar.Appearance++
}

func (d *Dictionary) GetAppearance(word []byte) (int, bool) {
	currentTerm := d.root
	for _, char := range word {
		child, ok := currentTerm.children[char]
		if !ok {
			break
		}

		currentTerm = child
	}

	return currentTerm.Appearance, currentTerm.isComplete
}

func (d *Dictionary) VisitAll() []*Term {
	queue := []*Term{d.root}
	terms := []*Term{}

	var currentTerm *Term
	for len(queue) > 0 {
		currentTerm = queue[0]
		if currentTerm.isComplete {
			terms = append(terms, currentTerm)
		}

		queue = queue[1:]
		for _, child := range currentTerm.children {
			queue = append(queue, child)
		}
	}

	return terms
}
