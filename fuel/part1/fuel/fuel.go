package fuel

func TotalRequiredFuel(mass int64) int64 {
	totalFuel := baseRequiredFuel(mass)
	marginalFuel := baseRequiredFuel(totalFuel)
	var totalMarginalFuel int64
	for {
		if marginalFuel <= 0 {
			break
		}
		totalMarginalFuel += marginalFuel
		marginalFuel = baseRequiredFuel(marginalFuel)
	}
	return totalFuel + totalMarginalFuel
}

func baseRequiredFuel(mass int64) int64 {
	return (mass / 3) - 2
}
