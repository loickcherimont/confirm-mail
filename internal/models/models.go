package models

// Data about form
// And state for submit
// Error if fields are empty/wrong
// Success if fields are correct
type Form struct {
	Email, Password string
	Error, Success  bool
}
