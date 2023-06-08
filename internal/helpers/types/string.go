package types

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// StringValue returns the value of the string pointer passed in or
// an empty string if the pointer is nil.
func StringValue(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
