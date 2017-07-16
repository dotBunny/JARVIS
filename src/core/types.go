package core

// Enum type hack
type Enum interface {
	name() string
	ordinal() int
	valueOf() *[]string
}

type DataModifier interface {
	GetDataValue() string
	SetDataValue(string, bool)
	ShouldUpdate() bool
}
