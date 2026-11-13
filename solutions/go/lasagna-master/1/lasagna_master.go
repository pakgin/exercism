package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	power := time
	if time == 0 {
		power = 2
	}

	return len(layers) * power
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodlesCounter, sauceCounter := 0, 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodlesCounter += 50
		} else if layer == "sauce" {
			sauceCounter += 0.2
		}
	}
	return noodlesCounter, sauceCounter
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friend, mine []string) []string {
	output := mine
	output[len(output)-1] = friend[len(friend)-1]
	return output
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(weights []float64, times int) []float64 {
	output := make([]float64, len(weights))
	copy(output, weights)
	for i := range output {
		output[i] *= (float64(times) / 2)
	}
	return output
}
