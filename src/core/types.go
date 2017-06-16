package core

// Enum type hack
type Enum interface {
	name() string
	ordinal() int
	valueOf() *[]string
}

// LineSpacer over text to align with normal start
const LineSpacer string = "\t\t\t\t"

// LetterBytes is used to generate random hashes
const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
