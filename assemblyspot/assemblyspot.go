package assemblyspot

import (
	"errors"
	"fmt"
	"github.com/neiderproy/recruitment-exercise-golang-oceans/vehicle"
	"sync"
	"time"
)

type AssemblySpot struct {
	vehicleToAssemble *vehicle.Car
	assemblyLog       string
}

func (s *AssemblySpot) SetVehicle(v *vehicle.Car) {
	s.vehicleToAssemble = v
}

func (s *AssemblySpot) GetAssembledVehicle() *vehicle.Car {
	return s.vehicleToAssemble
}

func (s *AssemblySpot) GetAssembledLogs() string {
	return s.assemblyLog
}

//hint: improve this function to execute this process concurrenlty
func (s *AssemblySpot) AssembleVehicle(vehicle chan <-*vehicle.Car, errorChanel chan<- error) {

	if s.vehicleToAssemble == nil {
		errorChanel <- errors.New("no vehicle set to start assembling")
	}

	wg := sync.WaitGroup{}
	wg.Add(7)
	s.assembleChassis(&wg)
	s.assembleTires(&wg)
	s.assembleEngine(&wg)
	s.assembleElectronics(&wg)
	s.assembleDash(&wg)
	s.assembleSeats(&wg)
	s.assembleWindows(&wg)

	wg.Wait()
	vehicle <- s.vehicleToAssemble
}

func (s *AssemblySpot) assembleChassis(wg *sync.WaitGroup) {
	defer wg.Done()
	s.vehicleToAssemble.Chassis = "Assembled"
	time.Sleep(1 * time.Second)
	s.assemblyLog += fmt.Sprintf("Chassis at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))

}

func (s *AssemblySpot) assembleTires(wg *sync.WaitGroup) {
	defer wg.Done()
	s.vehicleToAssemble.Tires = "Assembled"
	time.Sleep(1 * time.Second)
	s.assemblyLog += fmt.Sprintf("Tires at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
}

func (s *AssemblySpot) assembleEngine(wg *sync.WaitGroup) {
	defer wg.Done()
	s.vehicleToAssemble.Engine = "Assembled"
	time.Sleep(1 * time.Second)
	s.assemblyLog += fmt.Sprintf("Engine at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
}

func (s *AssemblySpot) assembleElectronics(wg *sync.WaitGroup) {
	defer wg.Done()
	s.vehicleToAssemble.Electronics = "Assembled"
	time.Sleep(1 * time.Second)
	s.assemblyLog += fmt.Sprintf("Electronics at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
}

func (s *AssemblySpot) assembleDash(wg *sync.WaitGroup) {
	defer wg.Done()
	s.vehicleToAssemble.Dash = "Assembled"
	time.Sleep(1 * time.Second)
	s.assemblyLog += fmt.Sprintf("Dash at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
}

func (s *AssemblySpot) assembleSeats(wg *sync.WaitGroup) {
	defer wg.Done()
	s.vehicleToAssemble.Sits = "Assembled"
	time.Sleep(1 * time.Second)
	s.assemblyLog += fmt.Sprintf("Sits at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
}

func (s *AssemblySpot) assembleWindows(wg *sync.WaitGroup) {
	defer wg.Done()
	s.vehicleToAssemble.Windows = "Assembled"
	time.Sleep(1 * time.Second)
	s.assemblyLog += fmt.Sprintf("Windows at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))

}
