//go:build !cgo || ecies_test_race
// +build !cgo ecies_test_race

package eciesgo

import (
	"crypto/elliptic"
)

func getCurve() elliptic.Curve {
	return elliptic.P256()
}
