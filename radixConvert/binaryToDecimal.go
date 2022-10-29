package radixConvert

import (
	"fmt"
	"strconv"
)

/*
   ParseInt を使う
*/
func BinaryToDecimal1(target string) int {
	i, err := strconv.ParseInt(target, 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return int(i)
}

/*
   自作
   ※ こっちの方が速度は上
*/
func BinaryToDecimal2(target string) int {

	binary, err := strconv.Atoi(target)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	var weight = 1
	var decimal int

	for binary != 0 {
		n := binary % 10
		decimal += weight * n
		weight <<= 1
		binary = binary / 10
	}

	return decimal

}
