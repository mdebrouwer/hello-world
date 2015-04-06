package main

import ( 
	"fmt"
	"bytes"
)

func main() {
	//basicExamples()
	//stringExamples()
	//capacityExamples()
	appendExample()
}

func appendExample() {
	var buffer [64]int

	for i := range(buffer) {
		buffer[i] = i
	}
	fmt.Println(buffer)

	slice := buffer[10:25]
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
	fmt.Println(slice)

	for i := 0; i < 1000; i++ {
		if !updateSlice(&slice, i, i*3) {
			break
		}
	}

	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println(buffer)
}

func updateSlice(slice *[]int, index int, value int) bool {
	if(index >= cap(*slice)) {
		fmt.Println("Cannot extend the slice any further!")
		return false
	}
	if index >= len(*slice) {
		// We need to increase the length of the slice
		(*slice) = (*slice)[0:len(*slice)+1]
	}
	(*slice)[index] = value
	return true
}

func capacityExamples() {
	//extendExample()
	//makeExample()
	insertExample()
}

func insertExample() {
	slice := make([]int, 10, 20)
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
	for i := 0; i < 10; i++ {
		newSlice := insert(slice, i, i)
		fmt.Printf("len: %d, cap: %d\n", len(newSlice), cap(newSlice))
	}
}

// Insert inserts the value into the slice at the specified index,
// which must be in range.
// The slice must have room for the new element.
func insert(slice []int, index, value int) []int {
    // Grow the slice by one element.
    slice = slice[0 : len(slice)+1]
    // Use copy to move the upper part of the slice out of the way and open a hole.
    copy(slice[index+1:], slice[index:])
    // Store the new value.
    slice[index] = value
    // Return the result.
    return slice
}

func makeExample() {
	slice := make([]int, 10, 15)
    fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
    newSlice := make([]int, len(slice), 2*cap(slice))
    for i := range slice {
        newSlice[i] = slice[i]
    }
    slice = newSlice
    fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))

    // The same code using the builtin copy function
	newSlice = make([]int, len(slice), 2*cap(slice))
    copy(newSlice, slice)
    fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))

    // Shorthand notation for crating a slice
    // with equal cap and len
    sliceOfInts := make([]int, 10)
    fmt.Printf("len: %d, cap: %d\n", len(sliceOfInts), cap(sliceOfInts))
}

func extendExample() {
	var iBuffer [10]int
    slice := iBuffer[0:0]
    for i := 0; i < 20; i++ {
    	// cap and len are both useful builtin functions
    	// for dealing with slices
    	if cap(slice) == len(slice) {
    		fmt.Println("slice is full!")
    		return
		}
        slice = extendSlice(slice, i)
        fmt.Println(slice)
    }
}

func extendSlice(slice []int, element int) []int {
    n := len(slice)
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}

type path []byte

func (p *path) TruncateAtFinalSlash() {
	index := bytes.LastIndex(*p, []byte("/"))
	if index >= 0 {
		*p = (*p)[0:index]
	}
}

func (p path) ToUpper() {
    for i, b := range p {
        if 'a' <= b && b <= 'z' {
            p[i] = b + 'A' - 'a'
        }
    }
}

func stringExamples() {
	//truncatePathExample()
	toUpperPathExample()
}

func toUpperPathExample() {
	pathName := path("/usr/bin/tso")
    pathName.ToUpper()
    fmt.Printf("%s\n", pathName)
}

func truncatePathExample() {
	pathName := path("/usr/bin/tso")
    pathName.TruncateAtFinalSlash()
    fmt.Printf("%s\n", pathName)
}

func basicExamples() {
	var buffer [256]byte
    slice := buffer[10:20]
    for i := 0; i < len(slice); i++ {
        slice[i] = byte(i)
    }

	passPointerToSliceExample(slice)
	//changeSliceIndexExample(slice)
	//changeSliceDataExample(buffer, slice)

	fmt.Println("slice", slice)
}

func passPointerToSliceExample(slice []byte) {
	fmt.Println("Before: len(slice) =", len(slice))
    ptrSubtractOneFromLength(&slice)
    fmt.Println("After:  len(slice) =", len(slice))
}

func changeSliceIndexExample(slice []byte) {
	fmt.Println("Before: len(slice) =", len(slice))
    newSlice := subtractOneFromLength(slice)
    fmt.Println("After:  len(slice) =", len(slice))
    fmt.Println("slice", slice)
    fmt.Println("After:  len(newSlice) =", len(newSlice))
    fmt.Println("newSlice", newSlice)
}

func changeSliceDataExample(buffer [256]byte, slice []byte) {
	fmt.Println("before", slice)
    addOneToEachElement(slice)
    fmt.Println("after", slice)
    fmt.Println("buffer", buffer)
}

func ptrSubtractOneFromLength(slicePtr *[]byte) {
    //slice := *slicePtr
    //*slicePtr = slice[0 : len(slice)-1]
    *slicePtr = (*slicePtr)[0:len(*slicePtr)-1]
}

func subtractOneFromLength(slice []byte) []byte {
    slice = slice[0 : len(slice)-1]
    return slice
}

func addOneToEachElement(slice []byte) {
    for i := range slice {
        slice[i]++
    }
}
