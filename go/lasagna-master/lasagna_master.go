package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg int) (estimate int) {
	if avg <= 0 {
		avg = 2
	}

	estimate = len(layers) * avg

	return estimate
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		} else if v == "sauce" {
			sauce += float64(0.2)
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(otherIngredient []string, myIngredient []string) {
	myIngredient[len(myIngredient)-1] = otherIngredient[len(otherIngredient)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) (newQuantities []float64) {
	for _, v := range quantities {
		perServing := v / float64(2)
		newQuantities = append(newQuantities, perServing*float64(portion))
	}

	return newQuantities
}
