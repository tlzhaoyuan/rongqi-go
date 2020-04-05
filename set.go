package rongqi

type intSet struct {
	data map[int]struct{}
}

// NewIntSet returns a new int set with data initialized
func NewIntSet() *intSet {
	return &intSet{data: map[int]struct{}{}}
}

// Add adds a new element to the existing set, or return false if it already exists
func (is *intSet) Add(el int) bool {
	if _, ok := is.data[el]; ok {
		return false
	}
	is.data[el] = struct{}{}
	return true
}

// Contains check if a set contains an element
func (is *intSet) Contains(el int) bool {
	_, ok := is.data[el]
	return ok
}

// Remove removes an element from a set
func (is *intSet) Remove(el int) bool {
	if _, ok := is.data[el]; ok {
		delete(is.data, el)
		return true
	}
	return false
}

// Len returns the cardinality of a set
func (is *intSet) Len() int {
	return len(is.data)
}
