package utils

import (
	"github.com/shopspring/decimal"
	"math"
	"math/big"
)

func HexadecimalStringToBitInt(str string) *big.Int {
	resString := ""
	if len(str) > 2 {
		if "0x" == str[:2] {
			resString = str[2:]
		} else {
			resString = str
		}
	} else {
		resString = str
	}
	if len(resString) == 0 {
		resString = "0"
	}
	desBigInt, _ := new(big.Int).SetString(resString, 16)
	return desBigInt
}

// 默认16位精度
func ToDecimal2(ivalue interface{}, decimals float64) float64 {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	case big.Int:
		value.SetBytes(v.Bytes())
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(decimals))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)
	f, _ := result.Float64()
	return f
}

func BigAmountFormatToFloat64OfDecimal(amount big.Int, decimal float64) float64 {
	return ToDecimal2(&amount, decimal)
}

// 默认16位精度
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

func AmountFormatToBigInt(amount, decimal float64) *big.Int {
	bigAmount := new(big.Float).SetFloat64(amount)
	dec := new(big.Float).SetFloat64(math.Pow(10, decimal))
	mul, _ := new(big.Float).Mul(bigAmount, dec).Float64()
	strAmount := Float64String(Decimal(mul, 0))
	setString, _ := new(big.Int).SetString(strAmount, 10)
	return setString
}

//String转换bigInt
func StringBigInt(str string) *big.Int {
	desBigInt, _ := new(big.Int).SetString(str, 10)
	return desBigInt
}

func FeeFormatToDecimalAmount(amount, decimal float64) float64 {
	bigAmount := new(big.Float).SetFloat64(amount)
	dec := new(big.Float).SetFloat64(math.Pow(10, -decimal))
	mul, _ := new(big.Float).Mul(bigAmount, dec).Float64()
	return mul
}
