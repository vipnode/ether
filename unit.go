package ether

import (
	"math/big"
	"strings"
)

// Unit is a denomination of Ethereum currency by exponent value from Wei.
type Unit int

const (
	Wei        Unit = 0
	Kwei            = 3
	Mwei            = 6
	Gwei            = 9
	Microether      = 12
	Milliether      = 15
	Ether           = 18

	// Aliases
	Babbage  = Kwei
	Lovelace = Mwei
	Shannon  = Gwei
	Szabo    = Microether
	Finney   = Milliether
	Eth      = Ether

	Unknown Unit = -1
)

// Num returns the denomination as a numerator.
func (u Unit) Num() *big.Int {
	r := big.NewInt(10)
	return r.Exp(r, big.NewInt(int64(u)), nil)
}

// String returns the lowercase string representation of the denomination.
func (u Unit) String() string {
	switch u {
	case Wei:
		return "wei"
	case Kwei:
		return "kwei"
	case Mwei:
		return "mwei"
	case Gwei:
		return "gwei"
	case Microether:
		return "microether"
	case Milliether:
		return "milliether"
	case Ether:
		return "ether"
	}
	return "unknown"
}

// ParseUnit takes a string and returns a unit. Unit is Unknown if parsing failed. It is case-insensitive.
func ParseUnit(s string) Unit {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "wei":
		return Wei
	case "kwei":
		return Kwei
	case "mwei":
		return Mwei
	case "gwei":
		return Gwei
	case "microether":
		return Microether
	case "milliether":
		return Milliether
	case "ether":
		return Ether
	}
	return Unknown
}
