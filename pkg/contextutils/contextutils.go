package contextutils

// ContextKey is a custom type used as a key for storing and retrieving values in context.
// It is a string type, ensuring uniqueness when used as a key.
type ContextKey string

// String returns the string representation of the ContextKey.
func (c ContextKey) String() string {
	return string(c)
}
