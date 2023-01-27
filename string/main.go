package main

func main() {
	str := "aiueoあいうえお"

	for i, rune := range str {
		println(i, string(rune))
	}

	// print str with bytes array
	for i, b := range []byte(str) {
		println(i, b)
	}
}
