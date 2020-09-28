package reverseBitsreverseBits

func reverseBitsreverseBits(num uint32) uint32 {
	var result uint32
	for i := 32; i > 0; i-- {
		result = (result << 1) + (num & 1)
		num = num >> 1
	}
	return result
}
