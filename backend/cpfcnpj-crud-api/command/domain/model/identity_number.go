package model

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	Cpf  = "cpf"
	Cnpj = "cnpj"
)

var (
	cpfFirstDigitTable   = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfSecondDigitTable  = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjFirstDigitTable  = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjSecondDigitTable = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfRegExp, _         = regexp.Compile(`^0{11}$|^1{11}$|^2{11}$|^3{11}$|^4{11}$|^5{11}$|^6{11}$|^7{11}$|^8{11}$|^9{11}$`)
	cnpjRegExp, _        = regexp.Compile(`^0{14}$|^1{14}$|^2{14}$|^3{14}$|^4{14}$|^5{14}$|^6{14}$|^7{14}$|^8{14}$|^9{14}$`)
)

type IdentityNumber interface {
	Value() string
	IsValid() bool
	Type() string
}

type identityValue struct {
	value string
}

type IdentityCpf struct {
	identityValue
}

type IdentityCnpj struct {
	identityValue
}

func isCpfWithRepeatedValues(cpf IdentityCpf) bool {
	return cpfRegExp.MatchString(cpf.value)
}

func isCnpjWithRepeatedValues(cnpj IdentityCnpj) bool {
	return cnpjRegExp.MatchString(cnpj.value)
}

func (cpf IdentityCpf) IsValid() bool {
	if len(cpf.value) != 11 {
		return false
	}

	if isCpfWithRepeatedValues(cpf) {
		return false
	}

	firstPart := cpf.value[0:9]
	sum := sumDigit(firstPart, cpfFirstDigitTable)

	r1 := sum % 11
	d1 := 0

	if r1 >= 2 {
		d1 = 11 - r1
	}

	secondPart := firstPart + strconv.Itoa(d1)

	dsum := sumDigit(secondPart, cpfSecondDigitTable)

	r2 := dsum % 11
	d2 := 0

	if r2 >= 2 {
		d2 = 11 - r2
	}

	finalPart := fmt.Sprintf("%s%d%d", firstPart, d1, d2)
	return finalPart == cpf.value
}

func (cnpj IdentityCnpj) IsValid() bool {
	if len(cnpj.value) != 14 {
		return false
	}

	if isCnpjWithRepeatedValues(cnpj) {
		return false
	}

	firstPart := cnpj.value[:12]
	sum1 := sumDigit(firstPart, cnpjFirstDigitTable)
	rest1 := sum1 % 11
	d1 := 0

	if rest1 >= 2 {
		d1 = 11 - rest1
	}

	secondPart := fmt.Sprintf("%s%d", firstPart, d1)
	sum2 := sumDigit(secondPart, cnpjSecondDigitTable)
	rest2 := sum2 % 11
	d2 := 0

	if rest2 >= 2 {
		d2 = 11 - rest2
	}

	finalPart := fmt.Sprintf("%s%d", secondPart, d2)
	return finalPart == cnpj.value
}

func (a identityValue) Value() string {
	return a.value
}

func (a IdentityCpf) Type() string {
	return Cpf
}

func (a IdentityCnpj) Type() string {
	return Cnpj
}

func NewCpf(value string) IdentityNumber {
	return &IdentityCpf{identityValue{value: value}}
}

func NewCnpj(value string) IdentityNumber {
	return &IdentityCnpj{identityValue{value: value}}
}

func sumDigit(s string, table []int) int {

	if len(s) != len(table) {
		return 0
	}

	sum := 0

	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}

	return sum
}

func main() {

}
