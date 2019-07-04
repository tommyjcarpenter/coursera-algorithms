package main

import (
	"fmt"
	"math/big"
)

func Karatsuba(num1 *big.Int, num2 *big.Int) *big.Int {

	if num1.Cmp(big.NewInt(10)) == -1 || num2.Cmp(big.NewInt(10)) == -1 {
		var product = new(big.Int)
		return product.Mul(num1, num2)
	}

	num1dig := num1.String()
	num2dig := num2.String()

	n := len(num1dig)
	n2 := len(num2dig)

	// first we pad until same length
	if n < n2 {
		for i := 0; i < n2-n; i++ {
			num1dig = "0" + num1dig
		}
	} else {
		for i := 0; i < n-n2; i++ {
			num2dig = "0" + num2dig
		}
	}
	// now we need to pad until even
	// NOT doing this seems to preserve a,b,c,d, however the split is wrong when we don't do this!
	// see https://cs.stackexchange.com/questions/75099/karatsuba-multiplication-on-numbers-with-odd-length
	if len(num1dig)%2 == 1 {
		num1dig = "0" + num1dig
		num2dig = "0" + num2dig
	}

	split := len(num1dig) / 2

	a := new(big.Int)
	a.SetString(num1dig[:split], 10)
	b := new(big.Int)
	b.SetString(num1dig[split:], 10)
	c := new(big.Int)
	c.SetString(num2dig[:split], 10)
	d := new(big.Int)
	d.SetString(num2dig[split:], 10)

	var ac = new(big.Int)
	ac = Karatsuba(a, c)
	var bd = new(big.Int)
	bd = Karatsuba(b, d)
	//(a+b)(c+d) == ac + ad + bc + bd
	//ad + bc = (a+b)(c+d) - ac - bd
	var inter = new(big.Int)
	inter = Karatsuba(new(big.Int).Add(a, b), new(big.Int).Add(c, d))
	inter.Sub(inter, ac)
	inter.Sub(inter, bd)

	//mid := Karatsuba(a+b, c+d) - ac - bd // AD+BC

	var ten1 = new(big.Int)
	ten1.Exp(big.NewInt(10), big.NewInt(int64(split)*int64(2)), nil)
	ten1.Mul(ten1, ac)

	var ten2 = new(big.Int)
	ten2.Exp(big.NewInt(10), big.NewInt(int64(split)), nil)
	ten2.Mul(ten2, inter)

	var result = new(big.Int)
	result.Add(ten1, ten2)
	result.Add(result, bd)

	//fmt.Println(num1, num2, num1dig, num2dig, split, a, b, c, d, result)

	return result
}

func main() {
	//fmt.Println(Karatsuba(big.NewInt(23345344), big.NewInt(857345345)))
	s1 := "3141592653589793238462643383279502884197169399375105820974944592"
	s2 := "2718281828459045235360287471352662497757247093699959574966967627"
	a := new(big.Int)
	a.SetString(s1, 10)
	b := new(big.Int)
	b.SetString(s2, 10)
	fmt.Println(Karatsuba(a, b))
}
