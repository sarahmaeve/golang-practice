// explore strings by rune, not by bytes
package main

import (
	"fmt"
	"unicode/utf8"
)

// is rune appropriate to return?
func letterFrequency(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, y := range s {
		//		if y != rune(" ") {
		freq[y] += 1
		//		}
	}
	return freq
}

func main() {
	arabic := "مِصْرَ أو رسمياً جُمهورِيةُ مِصرَ العَرَبيةِ هي دولة عربية تقع في الركن الشمالي الشرقي من قارة أفريقيا، ولديها امتداد"
	fmt.Println(len(arabic))
	fmt.Println(utf8.RuneCountInString(arabic))

	// "Fortunately, Go's range loop, when applied to a string, performs UTF-8
	// decoding implictly."
	// verifying with a manual version of RuneCountInString
	var count int
	for _, y := range arabic {
		count++
		// note when this loops this way the output is left to right.
		fmt.Printf("%q ", y)
	}
	fmt.Println("Count:", count)
	//
	// expanded := []rune(arabic)
	// fmt.Println(expanded)
	results := letterFrequency(arabic)
	for k, v := range results {
		fmt.Printf("%q : %d\n", k, v)
	}
}
