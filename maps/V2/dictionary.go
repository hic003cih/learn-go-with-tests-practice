//用maps建構一個字典

package main

import (
	"errors"
	"fmt"
)

//我們創建了一個Dictionary 类型,对 map 的简单封装
type Dictionary map[string]string

/* func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
} */

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

//在Dictionary實例上調用Search
func (d Dictionary) Search(word string) (string, error) {
	//map搜尋特性,返回兩個值,第二個值是bool,表示是否有找到key
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	fmt.Println(d.Search(word))

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
