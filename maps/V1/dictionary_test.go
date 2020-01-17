package main

import "testing"

/*
func TestSearch(t *testing.T) {

	//map需要兩中類型,第一個是key的類型,寫在[]內,第二個是value的類型
	//key的類型是可比較的類型,因為如果不能判斷是否相等的話,將無法知道我得到的值是否正確了
	//value的類可以是任意值,剩至是另一個map
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
	}
} */

//重構
func TestSearch(t *testing.T) {

	//map需要兩中類型,第一個是key的類型,寫在[]內,第二個是value的類型
	//key的類型是可比較的類型,因為如果不能判斷是否相等的話,將無法知道我得到的值是否正確了
	//value的類可以是任意值,剩至是另一個map
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
