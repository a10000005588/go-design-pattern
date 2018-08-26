package main

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training!")
}

func Swim() {
	fmt.Println("Swim !")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("Eating!")
}

type Shark struct {
	Animal
	Swim func()
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

// SwimmerImpl implment the Swim function.
type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimmer Swim!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

func main() {
	// 以下的composite都用 Swim() 函式
	swimmerA := CompositeSwimmerA{
		MySwim: Swim,
	}
	// swimmerA是CompositeSwimmer type 可以呼叫MyAthlete(type 為Athlete struct又可呼叫Train)
	swimmerA.MyAthlete.Train()
	swimmerA.MySwim()

	fish := Shark{
		Swim: Swim,
	}

	fish.Swim()
	fish.Eat()

	swimmerB := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	swimmerB.Train()
	swimmerB.Swim()
}
