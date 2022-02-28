package pro190

func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		// 1. 将给定的二进制数,由低到高位逐个取出
		// 1.1 右移 i 位,
		tmp := num >> i
		// 1.2  取有效位
		tmp = tmp & 1
		// 2. 然后通过位运算将其放置到反转后的位置.
		tmp = tmp << (31 - i)
		// 3. 将上述结果再次通过运算结合到一起
		result |= tmp
	}
	return result
}
