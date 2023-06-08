package types

// Float64 returns a pointer to the float64 value passed in.
func Float64(v float64) *float64 {
	return &v
}

// Float64Value returns the value of the float64 pointer passed in or
// an empty float64 if the pointer is nil.
func Float64Value(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

// Float32 returns a pointer to the float32 value passed in.
func Float32(v float32) *float32 {
	return &v
}

// Float32Value returns the value of the float32 pointer passed in or
// an empty float32 if the pointer is nil.
func Float32Value(v *float32) float32 {
	if v == nil {
		return 0
	}
	return *v
}
