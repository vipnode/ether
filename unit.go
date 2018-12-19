package ether

import (
	"math/big"
	"strings"
)

// Unit is a denomination of Ethereum currency by exponent value from Wei.
type Unit int

const (
	Unknown    Unit = -1
	Wei        Unit = 0
	Kwei       Unit = 3
	Mwei       Unit = 6
	Gwei       Unit = 9
	Microether Unit = 12
	Milliether Unit = 15
	Ether      Unit = 18

	// Aliases
	Babbage  = Kwei
	Lovelace = Mwei
	Shannon  = Gwei
	Szabo    = Microether
	Finney   = Milliether
	Eth      = Ether
)

// Num returns the denomination as a numerator in wei.
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
