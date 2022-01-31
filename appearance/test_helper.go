package appearance

import (
	"bytes"
	"io/ioutil"
	"os"
)

func readTestSeed() ([]byte, error) {
	file, err := os.OpenFile("../data/GoLang_Test.txt", os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func readSplittedTestSeed() ([][]byte, error) {
	c := NewCounter()

	content, err := readTestSeed()
	if err != nil {
		return nil, err
	}

	return c.split(bytes.ToLower(content)), nil
}
