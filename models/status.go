package models

// Direction - Custom type to hold value for week day ranging from 1-4
type Status int

// Declare related constants for each direction starting with index 1
const (
	New        Status = iota + 1 // EnumIndex = 1
	Paid                         // EnumIndex = 2
	Processing                   // EnumIndex = 3
	Done                         // EnumIndex = 4
)

// String - Creating common behavior - give the type a String function
func (s Status) String() string {
	return [...]string{"New", "Paid", "Processing", "Done"}[s-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex functio
func (s Status) EnumIndex() int {
	return int(s)
}
