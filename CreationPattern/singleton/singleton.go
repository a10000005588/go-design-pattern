package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

// create a singleton instance with count variable.
var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	// return the address of the instance
	return instance
}

// only type instance
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
