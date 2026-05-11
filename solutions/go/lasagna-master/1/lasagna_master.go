package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
    if avgPrepTime == 0 {
        avgPrepTime = 2 
    }
    return len(layers) * avgPrepTime
    
}

func Quantities(layers []string) (int, float64) {
    noodles := 0
    sauce := 0.0
    
    for _, layer := range layers {
        if layer == "noodles" {
            noodles += 50
        } else if layer == "sauce" {
            sauce += 0.2 
        }
    }
    
    return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) {
    
    secret := friendsList[len(friendsList)-1]
    
    
    myList[len(myList)-1] = secret
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
    scaledQuantities := make([]float64, len(quantities))
    
    for i := 0; i < len(quantities); i++ {
        scaledQuantities[i] = quantities[i] * float64(portions) / 2
    }
    
    return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
