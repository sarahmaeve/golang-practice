// explore strings by rune, not by bytes
package main

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

// is rune appropriate to return?
func letterFrequency(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, y := range s {
		// suppress space count
		// rune(" ") doesn't work -- no string to single rune conversion
		if y != rune(' ') {
			freq[y] += 1
		}
	}
	return freq
}

func main() {
	arabic := "مِصْرَ أو رسمياً جُمهورِيةُ مِصرَ العَرَبيةِ هي دولة عربية تقع في الركن الشمالي الشرقي من قارة أفريقيا، ولديها امتداد"
	fmt.Println("Length: ", len(arabic))
	fmt.Println("Rune count:", utf8.RuneCountInString(arabic))

	// "Fortunately, Go's range loop, when applied to a string, performs UTF-8
	// decoding implictly."
	// verifying with a manual version of RuneCountInString
	var count int
	for range arabic {
		count++
		// note when this loops this way the output is left to right.
		// %q -> escaped character literal, to get Arabic letters back
		// fmt.Printf("%q ", y)
	}
	fmt.Println("Looped count:", count)
	//
	// expanded := []rune(arabic)
	// fmt.Println(expanded)
	results := letterFrequency(arabic)

	// struct-based sort from
	// https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
	// as the map returned is unsorted (by design) and not easy to sort by value
	type kv struct {
		Key   rune
		Value int
	}

	var ss []kv
	for k, v := range results {
		ss = append(ss, kv{k, v})
	}

	sort.SliceStable(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		fmt.Printf("%q, %d\n", kv.Key, kv.Value)
	}

}
