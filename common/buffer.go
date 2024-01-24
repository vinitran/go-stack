package common

func Equals(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Alloc(length int, value byte) []byte {
	a := make([]byte, length)
	for i := 0; i < length; i++ {
		a[i] = value
	}
	return a
}

func ReadUInt16BE(source []byte, offset int) uint16 {
	return uint16(source[offset])<<8 | uint16(source[offset+1])
}

func WriteUInt16BE(source *[]byte, value uint16, offset int) {
	(*source)[offset] = byte(value >> 8)
	(*source)[offset+1] = byte(value)
}

func ReadUInt8(source []byte, offset int) uint8 {
	return source[offset]
}

func WriteUInt8(destination *[]byte, value uint8, offset int) {
	(*destination)[offset] = value
}

func ReadUInt16LE(source []byte, offset int) uint16 {
	return uint16(source[offset]) | uint16(source[offset+1])<<8
}

func WriteUInt16LE(destination *[]byte, value uint16, offset int) {
	(*destination)[offset] = byte(value)
	value >>= 8
	(*destination)[offset+1] = byte(value)
}

func ReadUInt32BE(source []byte, offset int) uint32 {
	return uint32(source[offset])<<24 | uint32(source[offset+1])<<16 | uint32(source[offset+2])<<8 | uint32(source[offset+3])
}

func WriteUInt32BE(destination *[]byte, value uint32, offset int) {
	(*destination)[offset] = byte(value >> 24)
	(*destination)[offset+1] = byte(value >> 16)
	(*destination)[offset+2] = byte(value >> 8)
	(*destination)[offset+3] = byte(value)
}

func ReadUInt32LE(source []byte, offset int) uint32 {
	return uint32(source[offset]) | uint32(source[offset+1])<<8 | uint32(source[offset+2])<<16 | uint32(source[offset+3])<<24
}

func WriteUInt32LE(destination *[]byte, value uint32, offset int) {
	(*destination)[offset] = byte(value)
	value >>= 8
	(*destination)[offset+1] = byte(value)
	value >>= 8
	(*destination)[offset+2] = byte(value)
	value >>= 8
	(*destination)[offset+3] = byte(value)
}
