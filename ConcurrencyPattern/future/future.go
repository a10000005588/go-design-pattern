package future

// SuccessFuc 接受一個只吃string的匿名函式
type SuccessFunc func(string)

// FailFunc 接受一個只吃error的匿名函式
type FailFunc func(error)

// ExecuteStringFunc 接受不吃任何參數的匿名函式，並回傳string,err
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunc = f
	return s
}

func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

func (s *MaybeString) Execute(f ExecuteStringFunc) {
	// future中的Execute會以concurrency的方式執行
	// goroutine a anonymous function.
	go func(s *MaybeString) {
		// f() will return "Hello World!" from test case.
		str, err := f()
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunc(str)
		}
	}(s) // () 為呼叫匿名函式，並帶入type MaybeString 的s
}
