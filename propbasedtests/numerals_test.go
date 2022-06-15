package numerals_test

import (
	"fmt"
	numerals "hello/propbasedtests"
	"testing"
	"testing/quick"
)

var cases = []struct {
	arabic uint16
	roman  string
}{
	{arabic: 1, roman: "I"},
	{arabic: 2, roman: "II"},
	{arabic: 3, roman: "III"},
	{arabic: 4, roman: "IV"},
	{arabic: 5, roman: "V"},
	{arabic: 6, roman: "VI"},
	{arabic: 7, roman: "VII"},
	{arabic: 8, roman: "VIII"},
	{arabic: 9, roman: "IX"},
	{arabic: 10, roman: "X"},
	{arabic: 14, roman: "XIV"},
	{arabic: 18, roman: "XVIII"},
	{arabic: 20, roman: "XX"},
	{arabic: 39, roman: "XXXIX"},
	{arabic: 40, roman: "XL"},
	{arabic: 47, roman: "XLVII"},
	{arabic: 49, roman: "XLIX"},
	{arabic: 50, roman: "L"},
	{arabic: 90, roman: "XC"},
	{arabic: 100, roman: "C"},
	{arabic: 400, roman: "CD"},
	{arabic: 500, roman: "D"},
	{arabic: 900, roman: "CM"},
	{arabic: 1000, roman: "M"},
	{arabic: 1984, roman: "MCMLXXXIV"},
	{arabic: 3999, roman: "MMMCMXCIX"},
	{arabic: 2014, roman: "MMXIV"},
	{arabic: 1006, roman: "MVI"},
	{arabic: 798, roman: "DCCXCVIII"},
}

func TestToRomanConvertion(t *testing.T) {

	for _, tt := range cases {
		input := tt.arabic
		want := tt.roman

		t.Run(fmt.Sprint("convert ", input), func(t *testing.T) {
			got := numerals.ToRoman(input)

			if want != got {
				t.Errorf("want %q, got %q", want, got)
			}
		})
	}
}

func TestToArabicConvertion(t *testing.T) {
	for _, tt := range cases {
		input := tt.roman
		want := tt.arabic

		t.Run(fmt.Sprint("convert ", input), func(t *testing.T) {
			got := numerals.ToArabic(input)

			if want != got {
				t.Errorf("want %d, got %d", want, got)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	t.Run("the conversion can be reversed to the original input", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic > 3999 {
				return true
			}
			roman := numerals.ToRoman(arabic)
			fromRoman := numerals.ToArabic(roman)
			return arabic == fromRoman
		}

		assertProperties(t, assertion)
	})

	t.Run("can't have more than 3 consecutive symbols of the same kind", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic > 3999 {
				return true
			}

			roman := numerals.ToRoman(arabic)

			var repeated int
			var previousSymbol int32
			maxConsecutiveSymbols := 3

			for _, symbol := range roman {
				if symbol == previousSymbol {
					repeated++
				} else {
					repeated = 0
				}
				previousSymbol = symbol
			}

			return repeated < maxConsecutiveSymbols
		}

		assertProperties(t, assertion)
	})
}

func assertProperties(t *testing.T, assertion func(arabic uint16) bool) {
	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
