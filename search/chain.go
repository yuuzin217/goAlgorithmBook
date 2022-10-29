package search

import "fmt"

// チェイン法

type Chain struct {
	Data    int
	Pointer int
}

func generateHash(seed int) int {
	return seed%5 + 1
}

func SearchChain(target int, source map[int]*Chain) (int, int) {
	if source == nil {
		fmt.Println("source is nil")
		return 0, 0
	}
	hash := generateHash(target)
	for hash != 0 {
		if source[hash] == nil {
			fmt.Println("データがありません")
			return 0, 0
		}
		if source[hash].Data == target {
			break
		}
		hash = source[hash].Pointer
	}
	if hash <= 0 {
		fmt.Println("該当データなし")
		return 0, 0
	}
	fmt.Println("探索成功")
	return hash, source[hash].Data
}

func SetChain(insert int, source map[int]*Chain) map[int]*Chain {
	if source == nil {
		source = make(map[int]*Chain)
	}
	hash, index := generateHash(insert), 1
	for ; ; index++ {
		if source[hash] == nil {
			source[hash] = &Chain{
				Data: insert,
			}
			break
		}
		source[hash] = &Chain{
			Pointer: index,
		}
		hash = index
	}
	fmt.Println("セット成功")
	return source
}
