package radixConvert_test

import (
	"algorithmBook/radixConvert"
	"fmt"
	"testing"
)

func TestBinaryToDecimal(t *testing.T) {
	for _, testCase := range []struct {
		req string
		exp int
	}{
		{
			req: "0000",
			exp: 0,
		},
		{
			req: "0001",
			exp: 1,
		},
		{
			req: "1111",
			exp: 15,
		},
		{
			req: "1010",
			exp: 10,
		},
		{
			req: "00000000",
			exp: 0,
		},
		{
			req: "00000001",
			exp: 1,
		},
		{
			req: "11111111",
			exp: 255,
		},
		{
			req: "10101010",
			exp: 170,
		},
	} {
		i := radixConvert.BinaryToDecimal2(testCase.req)
		if err := checkBinaryToDecimal(i, testCase.exp); err != nil {
			t.Fatal(err)
		}
	}
	t.Log("Test binary to decimal OK")
}

func checkBinaryToDecimal(result, expect int) error {
	if result != expect {
		return fmt.Errorf("binary to decimal error, result: %d / expect: %d,", result, expect)
	}
	return nil
}

func BenchmarkBinaryToDecimal1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		radixConvert.BinaryToDecimal1("00000001")
	}
	b.Log("Benchmark finished.")
}

func BenchmarkBinaryToDecimal2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		radixConvert.BinaryToDecimal2("00000001")
	}
	b.Log("Benchmark finished.")
}
