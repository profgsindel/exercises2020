package main

import (
	"bufio"
	"exercises2020/fuel"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../masses.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var masses []int64
	var totalMass int64
	s := bufio.NewScanner(f)
	for s.Scan() {
		mass, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%d\n", mass)
		totalMass += mass
		masses = append(masses, mass)
	}
	log.Printf("total mass count %d", len(masses))
	log.Printf("total mass %d", totalMass)

	var totalFuel int64
	for _, m := range masses {
		totalFuel += fuel.TotalRequiredFuel(m)
	}
	log.Printf("total fuel for %d mass = %d", totalMass, totalFuel)
}
