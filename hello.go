package main

import (
	"fmt"
	"unicode/utf8"
	"github.com/mdebrouwer/stringutil"
)

func main() {
	sample := "!oG ,olleH"
	fmt.Println(stringutil.Reverse(sample))

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}

	const nihongo = "日本語"
	
	for index, runeValue := range nihongo {
        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }

    for i, w := 0, 0; i < len(nihongo); i += w {
        runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
        fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
        w = width
    }
}
