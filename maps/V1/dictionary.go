//用maps建構一個字典

package main

//我們創建了一個Dictionary 类型,对 map 的简单封装
type Dictionary map[string]string

/* func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
} */

//在Dictionary實例上調用Search
func (d Dictionary) Search(word string) string {
	return d[word]
}
