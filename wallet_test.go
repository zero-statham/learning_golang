package com_learnku

import (
	"testing"
)

/**
指针
当你传值给函数或方法时，Go 会复制这些值。因此，如果你写的函数需要更改状态，你就需要用指针指向你想要更改的值

Go 取值的副本在大多数时候是有效的，但是有时候你不希望你的系统只使用副本，在这种情况下你需要传递一个引用。
例如，非常庞大的数据或者你只想有一个实例（比如数据库连接池）

nil
指针可以是 nil
当函数返回一个的指针，你需要确保检查过它是否为 nil，否则你可能会抛出一个执行异常，编译器在这里不能帮到你
nil 非常适合描述一个可能丢失的值
错误
错误是在调用函数或方法时表示失败的
通过测试我们得出结论，在错误中检查字符串会导致测试不稳定。
因此，我们用一个有意义的值重构了，这样就更容易测试代码，同时对于我们 API 的用户来说也更简单。

错误处理的故事远远还没有结束，你可以做更复杂的事情，这里只是抛砖引玉。后面的部分将介绍更多的策略。

​不要只是检查错误，要优雅地处理它们​
从现有的类型中创建新的类型。
用于为值添加更多的领域内特定的含义
可以让你实现接口
指针和错误是 Go 开发中重要的组成部分，你需要适应这些。幸运的是，如果你做错了，编译器通常会帮你解决问题，你只需要花点时间读一下错误信息。
*/

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrlnsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
