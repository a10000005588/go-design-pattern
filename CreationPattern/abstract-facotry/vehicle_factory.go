package abstract_factory

import (
	"errors"
	"fmt"
)

type VehicleFatory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType      = 1
	MotorbikeFatoryType = 2
)

func BuildFactory(f int) (VehicleFatory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFatoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recongized", f))
	}
}
