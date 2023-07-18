package model

type OutwardsVisible[T OutUser] interface {
	// Out outputs a sanitized version of this object
	// with the purpose of exposing its value outside of the program.
	Out() T
}
