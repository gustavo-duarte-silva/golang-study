package loop

func Loops(str string, numLoops int) string {
	var concatStr string
	for i := 0; i < numLoops; i++ {
		concatStr += str
	}
	return concatStr
}
