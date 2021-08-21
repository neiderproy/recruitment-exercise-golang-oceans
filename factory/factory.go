package factory

import (
	"errors"
	"fmt"
	"github.com/neiderproy/recruitment-exercise-golang-oceans/assemblyspot"
	 vehicleCar "github.com/neiderproy/recruitment-exercise-golang-oceans/vehicle"
)

const assemblySpots int = 5

type Factory struct {
	AssemblingSpots chan *assemblyspot.AssemblySpot
}

func New() *Factory {
	factory := &Factory{
		AssemblingSpots: make(chan *assemblyspot.AssemblySpot, assemblySpots),
	}

	totalAssemblySpots := 0

	for {
		factory.AssemblingSpots <- &assemblyspot.AssemblySpot{}

		totalAssemblySpots++

		if totalAssemblySpots >= assemblySpots {
			break
		}
	}

	return factory
}
type LogAssembled struct {
	VehicleID int
	LogTest string
	LogAss string
}

//HINT: this function is currently not returning anything, make it return right away every single vehicle once assembled,
//(Do not wait for all of them to be assembled to return them all, send each one ready over to main)
func (f *Factory) StartAssemblingProcess(amountOfVehicles int ,log chan <-*LogAssembled, err chan <-error, isDone chan <- bool) {
	if amountOfVehicles <= 0{
		err <- errors.New("not valid amount")
	}
	vehicleList := f.generateVehicleLots(amountOfVehicles)
	
	for _, vehicle := range vehicleList {
		fmt.Println("Assembling vehicle...")

		idleSpot := <-f.AssemblingSpots
		idleSpot.SetVehicle(&vehicle)

		car:= make(chan *vehicleCar.Car)
		errorCh  := make(chan error)
		go idleSpot.AssembleVehicle(car,errorCh)

		select {
			case vehicle := <- car:
				vehicle.TestingLog = f.testCar(vehicle)
				vehicle.AssembleLog = idleSpot.GetAssembledLogs()
				idleSpot.SetVehicle(nil)

				log <- &LogAssembled{
					VehicleID: vehicle.Id,
					LogTest:   vehicle.TestingLog,
					LogAss:    vehicle.AssembleLog,
				}
			case assembleVehicleErr := <- errorCh:
				err <- assembleVehicleErr

		}

	}
	isDone <- true
}

func (Factory) generateVehicleLots(amountOfVehicles int) []vehicleCar.Car {
	var vehicles = []vehicleCar.Car{}
	var index = 0

	for {
		vehicles = append(vehicles, vehicleCar.Car{
			Id:            index,
			Chassis:       "NotSet",
			Tires:         "NotSet",
			Engine:        "NotSet",
			Electronics:   "NotSet",
			Dash:          "NotSet",
			Sits:          "NotSet",
			Windows:       "NotSet",
			EngineStarted: false,
		})

		index++

		if index >= amountOfVehicles {
			break
		}
	}

	return vehicles
}

func (f *Factory) testCar(car *vehicleCar.Car) string {
	logs := ""

	log, err := car.StartEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnLeft()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnRight()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.StopEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	return logs
}
