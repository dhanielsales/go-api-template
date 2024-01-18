package error

import "runtime"

func getStack() []byte {
	buffer := make([]byte, 1024)
	for {
		layer := runtime.Stack(buffer, false)
		if layer < len(buffer) {
			return buffer[:layer]
		}

		buffer = make([]byte, 2*len(buffer))
	}
}
