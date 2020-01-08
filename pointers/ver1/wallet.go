package main

import (
	"errors"
	"fmt"
)

/* type Wallet struct {
	//如果一个符号（例如变量、类型、函数等）是以小写符号开头，
	//那么它在 定义它的包之外 就是私有的。
	//我们只想让自己的方法修改这个值，而其他的不可以。
	balance int
} */

/*

//当我们在代码中更改 balance 的值时，我们处理的是来自测试的副本。
//因此，balance 在测试中没有被改变。
//我们可以用 指针 来解决这个问题。指针让我们 指向 某个值，然后修改它。
//所以，我们不是拿钱包的副本，而是拿一个指向钱包的指针，这样我们就可以改变它。
func (w Wallet) Deposit(amount int) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	//我们可以使用「receiver」变量访问结构体内部的 balance 字段
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}

*/

//因為直接修改的話,修改的值是傳入的測試檔的Wallet的副本
//測試檔的Wallet的值還是一樣的,所以如果要直接修改Wallet的值
//而不是修改副本的話,就要用Pointer,指向Wallet的地址
//接收者类型是 *Wallet 而不是 Wallet
//你可以将其解读为「指向 wallet 的指针」。

/* func (w *Wallet) Deposit(amount int) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}
func (w *Wallet) Balance() int {
	return w.balance
} */

//重構
//使用int不具有描述性,因此修改為對现有的类型创建新的类型
//语法是 type MyName OriginalType

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

//类型别名有一个有趣的特性，你还可以对它们声明 方法。
//当你希望在现有类型之上添加一些领域内特定的功能时，这将非常有用。
//我们需要更新测试中的格式化字符串，以便它们将使用 String() 方法。
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//在 Go 中，错误是值，
//因此我们可以将其重构为一个变量，并为其提供一个单一的事实来源。

//var 关键字允许我们定义包的全局值。
var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return InsufficientFundsError
	}

	w.balance -= amount

	return nil
}
