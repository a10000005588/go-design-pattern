package future

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

// 設置一個timeout，避免有的goroutine concurrency太久
func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Log("Timeout!")

	t.Fail()
	wg.Done()
}

// setContext為一個閉包，將msg加工後，並且回傳一個會吐msg的匿名函式
func setContext(msg string) ExecuteStringFunc {
	msg = fmt.Sprintf("%s Closure !\n", msg)

	// 只有return內的東西是concurrent，其餘閉包的內容都為concurrent !
	return func() (string, error) {
		return msg, nil
	}
}

func TestStringOrError_Execute(t *testing.T) {
	// create an instance called future
	// which contains the struct that implement the "Success" "Fail" "Execute"
	future := &MaybeString{}

	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		// Timeout !
		go timeout(t, &wg)

		// 呼叫了future.Success會回傳MaybeString的instance: s
		// 故可以再透過s 呼叫 Fail
		// 寫法就跟如下一樣
		// s = future.Success()
		// s.Fail()
		// ==========
		// Success callback一個會執行t.Log(s)以及wg.Done()的匿名函式，
		future.Success(func(s string) {
			t.Log(s)
			// substract one for the wg, (do not need to wait another goroutine)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()

			wg.Done()
		})
		// 上述動作將Success以及Fail的匿名函式callback給MaybeString組裝成future

		// 這時future的 successFunc ＆ failFunc 這兩隻function都會有content了

		// future callback return "hello world" to the Execution function.
		future.Execute(func() (string, error) {
			return "Hello World!", nil
		})
		wg.Wait()
	})

	t.Run("Closure Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		// Timeout!
		go timeout(t, &wg)

		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})
		// callback一個閉包 (閉包內除了return都是同步在執行的)
		future.Execute(setContext("Hello"))
		wg.Wait()
	})

	t.Run("Error result", func(t *testing.T) {
		var wg sync.WaitGroup

		wg.Add(1)

		future.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e.Error())
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "", errors.New("Error ocurred")
		})

		wg.Wait()
	})
}
