package numerals

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)

	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

type subtractiveSymbols []uint8

func (s subtractiveSymbols) Contains(symbol uint8) bool {
	for _, subtractiveSymbol := range s {
		if symbol == subtractiveSymbol {
			return true
		}
	}

	return false
}

var keyRomanNumerals = RomanNumerals{
	{Value: 1000, Symbol: "M"},
	{Value: 900, Symbol: "CM"},
	{Value: 500, Symbol: "D"},
	{Value: 400, Symbol: "CD"},
	{Value: 100, Symbol: "C"},
	{Value: 90, Symbol: "XC"},
	{Value: 50, Symbol: "L"},
	{Value: 40, Symbol: "XL"},
	{Value: 10, Symbol: "X"},
	{Value: 9, Symbol: "IX"},
	{Value: 5, Symbol: "V"},
	{Value: 4, Symbol: "IV"},
	{Value: 1, Symbol: "I"},
}

func ToRoman(arabic uint16) (roman string) {
	var result strings.Builder

	for _, n := range keyRomanNumerals {
		for arabic >= n.Value {
			result.WriteString(n.Symbol)
			arabic -= n.Value
		}
	}

	return result.String()
}

func ToArabic(roman string) (arabic uint16) {
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]
			value := keyRomanNumerals.ValueOf(symbol, nextSymbol)

			if value != 0 {
				arabic += value
				i++
			} else {
				arabic += keyRomanNumerals.ValueOf(symbol)
			}
		} else {
			arabic += keyRomanNumerals.ValueOf(symbol)
		}
	}

	return
}

func couldBeSubtractive(index int, symbol uint8, roman string) bool {
	return isThereAnyNextSymbol(index, roman) && isSubtractive(symbol)
}

func isThereAnyNextSymbol(index int, roman string) bool {
	return index+1 < len(roman)
}

func isSubtractive(symbol uint8) bool {
	ss := subtractiveSymbols{'I', 'X', 'C'}

	return ss.Contains(symbol)
}
