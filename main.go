package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
)

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

func BasePointGGet() ECPoint {
	return ECPoint{
		X: elliptic.P256().Params().Gx,
		Y: elliptic.P256().Params().Gy,
	}
}

func ECPointGen(x, y *big.Int) ECPoint {
	return ECPoint{X: x, Y: y}
}

func IsOnCurveCheck(a ECPoint) bool {
	return elliptic.P256().IsOnCurve(a.X, a.Y)
}

func AddECPoints(a, b ECPoint) ECPoint {
	x, y := elliptic.P256().Add(a.X, a.Y, b.X, b.Y)
	return ECPoint{X: x, Y: y}
}

func DoubleECPoints(a ECPoint) ECPoint {
	x, y := elliptic.P256().Double(a.X, a.Y)
	return ECPoint{X: x, Y: y}
}

func ScalarMult(k *big.Int, a ECPoint) ECPoint {
	x, y := elliptic.P256().ScalarMult(a.X, a.Y, k.Bytes())
	return ECPoint{X: x, Y: y}
}

func ECPointToString(point ECPoint) string {
	return fmt.Sprintf("(%s, %s)", point.X.String(), point.Y.String())
}

func StringToECPoint(s string) ECPoint {
	var x, y big.Int
	fmt.Sscanf(s, "(%s, %s)", &x, &y)
	return ECPoint{X: &x, Y: &y}
}

func PrintECPoint(point ECPoint) {
	fmt.Printf("(%s, %s)\n", point.X.String(), point.Y.String())
}

func main() {
	G := BasePointGGet()
	k := SetRandom(256)
	d := SetRandom(256)

	H1 := ScalarMult(d, G)
	H2 := ScalarMult(k, H1)

	H3 := ScalarMult(k, G)
	H4 := ScalarMult(d, H3)

	result := IsEqual(H2, H4)
	fmt.Println("Are k*(d*G) and d*(k*G) equal?", result)
}

func SetRandom(bits int) *big.Int {
	randomBits := make([]byte, bits/8)
	_, err := rand.Read(randomBits)
	if err != nil {
		panic(err)
	}

	randomInt := new(big.Int).SetBytes(randomBits)
	randomInt.Mod(randomInt, new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(bits)), nil))

	return randomInt
}

func IsEqual(a, b ECPoint) bool {
	return a.X.Cmp(b.X) == 0 && a.Y.Cmp(b.Y) == 0
}
