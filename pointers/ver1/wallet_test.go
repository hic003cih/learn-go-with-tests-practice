package main

import "testing"

/* func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()

	fmt.Println("address of balance in test is", &wallet.balance)
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
*/

//基本的測試,測試存款
/* func TestWallet(t *testing.T) {

	//建立一個struct
	wallet := Wallet{}
	////要生成 Bitcoin（比特币），你只需要用 Bitcoin(999) 的语法就可以了。
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	//要生成 Bitcoin（比特币），你只需要用 Bitcoin(999) 的语法就可以了。
	want := Bitcoin(11)

	// if got != want {
	//	t.Errorf("got %d want %d", got, want)


	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
*/

//增加提款的測試
/* func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		//建立一個struct
		wallet := Wallet{}
		////要生成 Bitcoin（比特币），你只需要用 Bitcoin(999) 的语法就可以了。
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		//要生成 Bitcoin（比特币），你只需要用 Bitcoin(999) 的语法就可以了。
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(10)

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}
*/

//重複的部分重構
func TestWallet(t *testing.T) {

	//把顯示錯誤的部分取出來重複使用
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		//t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		//建立一個struct,不給balance,表示為空錢包
		wallet := Wallet{}
		//存入10元比特幣
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		//建立struct,並把balance預存20的比特幣
		wallet := Wallet{balance: Bitcoin(20)}
		//取出20元比特幣
		wallet.Withdraw(Bitcoin(20))
		assertBalance(t, wallet, Bitcoin(10))
	})

	/* t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(10)

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}) */

	/* //帳戶超領時,顯示錯誤
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		//Withdraw必須要返回一個值
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}

	}) */

	/* //重構

	//帳戶超領時,顯示錯誤
	//並使用assertError來增加閱讀性
	//為錯誤檢查做一個方法,讓我們閱讀起來更清晰
	assertError := func(t *testing.T, err error) {
		if err == nil {
			t.Error("wanted an error but didnt get one")
		}
	}

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		//Withdraw必須要返回一個值
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err)

	}) */

	//更新一个 string 的助手方法来比较
	assertError := func(t *testing.T, got error, want string) {
		if got == nil {
			//t.Fatal的用法
			//如果它被调用，它将停止测试。这是因为我们不希望对返回的错误进行更多断言。
			//如果没有这个，测试将继续进行下一步，并且因为一个空指针而引起 panic。
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		//Withdraw必須要返回一個值
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "cannot withdraw, insufficient funds")

	})
}
