package firstUniqChar

import "math"

func firstUniqChar(s string) int {
	length := len(s)
	if length == 0 {
		return -1
	}

	byteIndexMap := map[byte]int{s[0]: 0}

	for i := 1; i < length; i++ {
		v, ok := byteIndexMap[s[i]]
		if ok {
			if v >= 0 {
				byteIndexMap[s[i]] = -1
			}
		} else {
			byteIndexMap[s[i]] = i
		}
	}

	index := math.MaxInt32

	for _, v := range byteIndexMap {
		if v >= 0 && v < index {
			index = v
		}
	}

	if index == math.MaxInt32 {
		return -1
	}

	return index
}
