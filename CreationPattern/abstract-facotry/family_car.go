package abstract_factory

type FamilyCar struct{}

// Doors are implemented in car.go. Because different types of car have variety number of doors.
func (*FamilyCar) NumDoors() int {
	return 5
}

func (*FamilyCar) NumWheels() int {
	return 4
}

func (*FamilyCar) NumSeats() int {
	return 5
}
