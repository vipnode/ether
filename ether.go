// Ether package provides parsing and printing across different Ethereum unit
// denominations.
//
// It tries to retain precision by using big.Rat for internal computation, and
// big.Int for inputs and outputs.
package ether

import (
	"math/big"
	"strings"
	"unicode"
)

var ethInWei = big.NewInt(1e18)
var ethInGwei = big.NewInt(1e9)

var boundary = big.NewInt(1e4)
var gweiBoundary = new(big.Int).Div(ethInGwei, boundary)
var ethBoundary = new(big.Int).Div(ethInWei, boundary)

// Print converts some value in wei like 123 into a string like "123 wei". It
// will use gwei or ether if it is within 4 decimals of precision.
// denomation based on the value.
func Print(wei *big.Int) string {
	unit := "wei"
	denom := big.NewInt(1)

	if wei.CmpAbs(ethBoundary) >= 0 {
		unit = "ether"
		denom = ethInWei
	} else if wei.CmpAbs(gweiBoundary) >= 0 {
		unit = "gwei"
		denom = ethInGwei
	}
	s := new(big.Rat).SetFrac(wei, denom).FloatString(4)
	s = strings.TrimRight(s, "0.")
	if s == "" {
		s = "0"
	}
	return s + " " + unit
}

// Parse takes a string like "2 eth" and converts it to the wei equivalent.
func Parse(s string) (*big.Int, error) {
	// Find numbers like -123.456 and stop after the first non-number-y character.
	var splitPos int
	for pos, ch := range s {
		if !unicode.IsNumber(ch) && ch != '-' && ch != '.' {
			splitPos = pos
			break
		}
	}

	if splitPos == 0 {
		// No unit provided, assume it's wei.
		if r, ok := new(big.Int).SetString(s, 0); ok {
			return r, nil
		}
		return nil, ErrInvalidValue
	}

	number, unit := s[:splitPos], s[splitPos:]

	n, ok := new(big.Rat).SetString(number)
	if !ok {
		return nil, ErrInvalidValue
	}

	mul := new(big.Rat)
	u := ParseUnit(unit)
	if u == Unknown {
		return nil, ErrInvalidUnit
	}
	mul.SetInt(u.Num())
	n.Mul(n, mul)
	return new(big.Int).Div(n.Num(), n.Denom()), nil
}
