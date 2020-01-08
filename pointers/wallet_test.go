package main

import "testing"

//重複的部分重構
func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		//建立一個struct,不給balance,表示為空錢包
		wallet := Wallet{}
		//存入10元比特幣
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		//建立struct,並把balance預存20的比特幣
		wallet := Wallet{balance: Bitcoin(20)}
		//取出20元比特幣
		wallet.Withdraw(Bitcoin(20))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {

		//建立struct,並把balance預存20的比特幣
		wallet := Wallet{Bitcoin(20)}

		//取出100元比特幣,超領,應該要傳出err
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, InsufficientFundsError)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		//Withdraw必須要返回一個值
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "cannot withdraw, insufficient funds")

	})
}

//把顯示錯誤的部分取出來重複使用
func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	//t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

//更新一个 string 的助手方法来比较
func assertError(t *testing.T, got error, want error) {
	if got == nil {
		//t.Fatal的用法
		//如果它被调用，它将停止测试。这是因为我们不希望对返回的错误进行更多断言。
		//如果没有这个，测试将继续进行下一步，并且因为一个空指针而引起 panic。
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
