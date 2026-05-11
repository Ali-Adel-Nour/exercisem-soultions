package gross
import "fmt" 

func Units() map[string]int {
    unitMeasurements := map[string]int{
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen":    6,
        "dozen":              12,
        "small_gross":        120,
        "gross":              144,
        "great_gross":        1728,
    }

    for key, value := range unitMeasurements {
        fmt.Printf("Key: %s, Value: %d\n", key, value) // Debugging (optional)
    }

    return unitMeasurements
}


func NewBill() map[string]int {
  
    emptyBill := map[string]int{}

  
    return emptyBill
}


// AddItem adds an item to customer bill.
func AddItem(bill map[string]int, units map[string]int, item string, unit string) bool {
    // Check if the unit is valid
    unitValue, ok := units[unit]
    if !ok {
        return false // Unit does not exist in the units map
    }
    
    // Add or update the item in the customer bill
    bill[item] += unitValue
    return true
}

func RemoveItem(bill map[string]int, units map[string]int, item string, unit string) bool {
    // Check if the item exists in the bill
    currentQuantity, itemExists := bill[item]
    if !itemExists {
        return false // Item does not exist in the bill
    }

    // Check if the unit is valid
    unitValue, unitExists := units[unit]
    if !unitExists {
        return false // Unit does not exist in the units map
    }

    // Calculate the new quantity
    newQuantity := currentQuantity - unitValue
    if newQuantity < 0 {
        return false // New quantity is less than 0
    }

    // Update or remove the item
    if newQuantity == 0 {
        delete(bill, item) // Completely remove the item
    } else {
        bill[item] = newQuantity // Update the quantity
    }

    return true
}


// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    // Check if the item exists in the bill
    quantity, exists := bill[item]
    if !exists {
        return 0, false // Item does not exist
    }

    return quantity, true // Return the quantity and true
}