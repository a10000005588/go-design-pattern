package singleton

type Singleton2 interface {
	AddTwo() int
}

type singleton2 struct {
	count int
}

var instance2 *singleton2

func GetInstance2() Singleton2 {
	return nil
}

func (s2 *singleton2) AddTwo() int {
	s2.count += 2
	return s2.count
}
