package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	measurements := make(map[string]int)
    measurements["quarter_of_a_dozen"] = 3
    measurements["half_of_a_dozen"] = 6
    measurements["dozen"] = 12
    measurements["small_gross"] = 120
    measurements["gross"] = 144
    measurements["great_gross"] = 1728
    return measurements
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
    if !exists {
        return false
    }
    bill[item]+=value
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billAmt, itemExists := bill[item]
    unitAmt, unitExists := units[unit]
    if !itemExists || !unitExists || billAmt-unitAmt < 0 {
        return false
    }
    if billAmt-unitAmt == 0 {
        delete(bill, item)
    } else {
        bill[item] -= unitAmt
    }
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
	return value, exists
}
