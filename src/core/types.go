package core

// Enum type hack
type Enum interface {
	name() string
	ordinal() int
	valueOf() *[]string
}