package radixConvert_test

import (
	"algorithmBook/radixConvert"
	"fmt"
	"testing"
)

func TestDecimalToBinary(t *testing.T) {
	for _, testCase := range []struct {
		req   int
		digit int
		exp   string
	}{
		{
			req:   0,
			digit: 4,
			exp:   "0000",
		},
		{
			req:   1,
			digit: 4,
			exp:   "0001",
		},
		{
			req:   15,
			digit: 4,
			exp:   "1111",
		},
		{
			req:   10,
			digit: 4,
			exp:   "1010",
		},
		{
			req:   0,
			digit: 8,
			exp:   "00000000",
		},
		{
			req:   1,
			digit: 8,
			exp:   "00000001",
		},
		{
			req:   255,
			digit: 8,
			exp:   "11111111",
		},
		{
			req:   170,
			digit: 8,
			exp:   "10101010",
		},
	} {
		b := radixConvert.DecimalToBinary1(testCase.req, testCase.digit)
		if err := checkDecimalToBinary(b, testCase.exp); err != nil {
			t.Fatal(err)
		}
	}
	t.Log("Test decimal to binary OK")
}

func checkDecimalToBinary(result, expect string) error {
	if result != expect {
		return fmt.Errorf("decimal to binary error, result: %s / expect: %s,", result, expect)
	}
	return nil
}

func BenchmarkDecimalToBinary1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		radixConvert.DecimalToBinary1(11, 8)
	}
	b.Log("Benchmark finished.")
}

func BenchmarkDecimalToBinary2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		radixConvert.DecimalToBinary2(11, 8)
	}
	b.Log("Benchmark finished.")
}
