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
/*
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
*/
//重構
/* func TestSearch(t *testing.T) {

	//map需要兩中類型,第一個是key的類型,寫在[]內,第二個是value的類型
	//key的類型是可比較的類型,因為如果不能判斷是否相等的話,將無法知道我得到的值是否正確了
	//value的類可以是任意值,剩至是另一個map
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		//
		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
}

//測試新增文字功能
func TestAdd(t *testing.T) {

	//定義一個map類型,叫Dictionary
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"
	//調用新增
	dictionary.Add(word, definition)

	want := "this is just a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if want != got {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
} */

func TestSearch(t *testing.T) {

	//map需要兩中類型,第一個是key的類型,寫在[]內,第二個是value的類型
	//key的類型是可比較的類型,因為如果不能判斷是否相等的話,將無法知道我得到的值是否正確了
	//value的類可以是任意值,剩至是另一個map
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		//
		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
}

//測試新增文字功能
func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)

	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}
