package main

import "fmt"

type RailRoadWidthChecker interface {
	checkRailsWidth() int
}

type RailRoad struct {
	Width int
}

// 只有宣告為type RailRoad才能呼叫 IsCorrectSizeTrain
func (railroad *RailRoad) IsCorrectSizeTrain(trainInstance RailRoadWidthChecker) bool {
	return trainInstance.checkRailsWidth() == railroad.Width
}

// 疑問：為何不是在input中type為Train ???
// Answer: 因為由struct type為 Train實作了 struct type: RailRoadWidthCheck的CheckRailWidth()的方法

type Train struct {
	TrainWidth int
}

// 宣告為Train的struct instance 實作 checkRailsWidth()函式 並回傳 int型態
func (p Train) checkRailsWidth() int {
	return p.TrainWidth
}

func main() {
	railroad := RailRoad{Width: 10}

	passengerTrain := Train{TrainWidth: 10}
	cargoTrain := Train{TrainWidth: 15}

	canPassengerTrainPass := railroad.IsCorrectSizeTrain(passengerTrain)
	canCargoTrainPass := railroad.IsCorrectSizeTrain(cargoTrain)

	fmt.Printf("Can passengerTrain pass? :%t \n", canPassengerTrainPass)
	fmt.Printf("Can cargoTrain pass? :%t \n", canCargoTrainPass)
}
