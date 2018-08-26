package mediator

import "fmt"

type One struct{}
type Two struct{}
type Three struct{}
type Four struct{}

func main() {

}

// No mediator
func (o *One) OnePlus(n interface{}) interface{} {
	switch n.(type) {
	case One:
		return &Two{}
	case Two:
		return &Three{}
	case Three:
		return &Four{}
	default:
		fmt.Println("1~4 are not found")
	}
}

// With mediator

func Sum(a, b interface{}) interface{} {
	switch a := a.(type) {
	case One:
		switch b.(type) {
		case One:
			return &Two{}
		case Two:
			return &Three{}
		case Three:
			return &Four{}
		default:
			fmt.Println("1~4 not found")
		}
	}
}
