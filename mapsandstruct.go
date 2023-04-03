package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {

	//maps
	statePopulations := make(map[string]int, 10)
	statePopulations = map[string]int{
		"California": 3950017,
		"Texas":      27862596,
		"Florida":    20612439,
	}
	statePopulations["Georgia"] = 10310371
	delete(statePopulations, "Texas")

	pop, ok := statePopulations["Ohio"]
	fmt.Println(statePopulations, ok, pop)
	len := len(statePopulations)
	fmt.Println(len)

	//struct
	aDoctor := Doctor{
		number:    3,
		actorName: "Prakash",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor.companions)
	bDoctor := struct{ name string }{name: "John Pertwee"}
	fmt.Println(bDoctor)

}
