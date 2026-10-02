package cards

import "slices"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// InBounds checks if the requested index exists within the slice
func InBounds(slice []int, index int) bool {
    if index >= 0 && index < len(slice) {
        return true
    } else {
        return false
    }    
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    // Check if the index is within bounds
    if InBounds(slice, index) {
        return slice[index]
    } else {
        return -1
    }
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    // Check if the index is within bounds
    if InBounds(slice, index) {
        slice[index] = value
    } else {
        slice = append(slice, value)
    }
    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	for i:=len(values)-1; i >= 0; i-- {
        slice = slices.Insert(slice, 0, values[i])   
    } 
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    // Check if the index is within bounds
    if InBounds(slice, index) {
    	return slices.Delete(slice, index, index+1) 
    } else {
        return slice
    }
}
