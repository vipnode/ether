package ether

import "errors"

// ErrInvalidUnit is returned when we attempt to parse an unknown denomination unit.
var ErrInvalidUnit = errors.New("invalid ether unit")

// ErrInvalidValue is returned when we fail to parse the numeric component of an ether value.
var ErrInvalidValue = errors.New("invalid ether value")
