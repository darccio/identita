package bits

type Bit bool

func Set(placeholder int, position int) int {
	placeholder |= (1 << position)
	return placeholder
}

func Clear(placeholder int, position int) int {
	mask := ^(1 << position)
	placeholder &= mask
	return placeholder
}

func Bits(value interface{}) (result []Bit) {
	return
}
