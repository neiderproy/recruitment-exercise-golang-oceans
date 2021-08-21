package main

import (
	"fmt"
	 fact "github.com/neiderproy/recruitment-exercise-golang-oceans/factory"
	"log"
)

const carsAmount = 100

func main() {
	factory := fact.New()

	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id
	logListener := make(chan *fact.LogAssembled)
	err := make(chan error)
	isDone := make(chan bool)

	go factory.StartAssemblingProcess(carsAmount,logListener,err,isDone)

	for{
		select {
			case logs :=<-logListener:
				log.Println(fmt.Sprintf("vehicle: %v \n testing: %v \n assemble: %v \n Vehicle Done", logs.VehicleID,logs.LogTest,logs.LogAss))
			case factoryError := <- err:
				panic(factoryError)
			case <-isDone:
				return
		}
	}
}
