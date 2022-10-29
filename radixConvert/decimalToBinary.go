package radixConvert

import (
	"fmt"
	"strconv"
	"strings"
)

/*
   Sprintf を使う
   ※ こっちの方が速い
*/
func DecimalToBinary1(target int, digit int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(digit)+"b", target)
}

/*
   自作
*/
func DecimalToBinary2(target int, digit int) string {

	const base = 2
	var binarySlice []string

	for target != 0 {
		binarySlice = append([]string{strconv.Itoa(target % base)}, binarySlice...)
		target /= base
	}

	leng := len(binarySlice)
	if leng < digit {
		for i := 0; i < digit-leng; i++ {
			binarySlice = append([]string{"0"}, binarySlice...)
		}
	}

	return strings.Join(binarySlice, "")

}
