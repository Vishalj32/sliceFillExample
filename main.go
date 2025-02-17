package main

func main() {
	//var bigSlice = make([]byte, 73437)
}

func FillSliceIndex(slice []byte, value byte) {
	for j := 0; j < len(slice); j++ {
		slice[j] = value
	}
}

func FillSliceRange(slice []byte, value byte) {
	for j := range slice {
		slice[j] = value
	}
}

func FillSliceCopyTrick(slice []byte, value byte) {
	slice[0] = value
	for j := 1; j < len(slice); j *= 2 {
		copy(slice[j:], slice[:j])
	}
}

func FillSlicePatternCopyTrick(slice []byte, pattern []byte) {
	copy(slice, pattern)
	for j := len(pattern); j < len(slice); j *= 2 {
		copy(slice[j:], slice[:j])
	}
}
