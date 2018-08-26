package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	ManufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	ManufacturingComplex.SetBuilder(carBuilder)
	ManufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d \n", car.Seats)
	}

	motorBikeBuilder := &BikeBuilder{}
	ManufacturingComplex.SetBuilder(motorBikeBuilder)
	ManufacturingComplex.Construct()

	bike := motorBikeBuilder.GetVehicle()

	if bike.Wheels != 2 {
		t.Errorf("Wheels on a bike must be 2 and they were %d\n", car.Wheels)
	}

	if bike.Structure != "Bike" {
		t.Errorf("Structure on a car must be 'Car and was %s\n", car.Structure)
	}

	if bike.Seats != 1 {
		t.Errorf("Seats on a car must be 1 and they were %d \n", car.Seats)
	}
}
