package main

import (
	"errors"
	"fmt"
	"strings"
)

func main(){
//	var plantCapacities []float64
//
//	plantCapacities = []float64{30, 30 , 30, 60, 60, 100}
//
//	var capacity float64 =  + plantCapacities[0] + plantCapacities[1] + plantCapacities[2] +
//plantCapacities[3] + plantCapacities[4] + plantCapacities[5]
//
//var gridLoad = 75.
//
//utilization := gridLoad / capacity
//
//fmt.Printf("%-20s%.0f\n","Capacity:", capacity)
//fmt.Printf("%-20s%.0f\n","Load:", gridLoad)
//fmt.Printf("%-20s%.1f%%\n","Utilization:", utilization * 100)
//

plants := []PowerPlant{
	{hydro, 300, active},
	{wind, 30, active},
	{wind, 25, inactive},
	{wind, 35, active},
	{solar, 45, unavailable},
	{solar, 40, inactive},
}
	grid := PowerGrid{300, plants}

	if option, err := requestOption(); err ==nil {
		fmt.Println("")

		switch option{

		case"1":
			grid.generatePlantReport()
		case "2":
			grid.generateGridReport()

		default:
			fmt.Println("Unknown option, no action taken")
		}
	}else{
		fmt.Println(err.Error())
	}

}

func requestOption() (option string, err error) {
	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Print("Please choose an option: ")

	fmt.Scanln(&option)
	if option != "1" && option != "2"{
		err = errors.New("Invalid option selected")
	}
	return
}


type PlantType string

const(
	hydro PlantType = "Hydro"
	wind PlantType = "Wind"
	solar PlantType = "Solar"
)

type PlantStatus string
const (
	active PlantStatus = "Active"
	inactive PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavailable"
	)

type PowerPlant struct {
	plantType PlantType
	capacity float64
	status PlantStatus
}

type PowerGrid struct {
	load float64
	plants []PowerPlant
}

func(pg *PowerGrid) generatePlantReport(){
for idx, p := range pg.plants {
	label := fmt.Sprintf("%s%d", "Plant #", idx)
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("%-20s%s\n", "Type:", p.plantType)
	fmt.Printf("%-20s%.0f\n", "Capacity:", p.capacity)
	fmt.Printf("%-20s%.0f\n", "Status:", p.status)
}
}
func(pg *PowerGrid) generateGridReport(){
	capacity := 0.
	for _, p := range pg.plants {
		if p.status == active {
			capacity += p.capacity
		}
	}
	label := "Power Grid Report"
	fmt.Println(label)
	fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.0f\n", "Load:", pg.load)
	fmt.Printf("%-20s%.2f%%\n", "Utilization:", pg.load/capacity*100)
}