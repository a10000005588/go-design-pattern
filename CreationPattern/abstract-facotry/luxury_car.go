package abstract_factory

type LuxuryCar struct{}

// Doors are implemented in car.go. Because different types of car have variety number of doors.
func (*LuxuryCar) NumDoors() int {
	return 4
}

func (*LuxuryCar) NumWheels() int {
	return 4
}

func (*LuxuryCar) NumSeats() int {
	return 5
}
