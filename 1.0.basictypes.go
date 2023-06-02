package main

import "fmt"

var intEight int8      // 8 bits -128 to 127
var intSixteen int16   // 16 bits -2^15 to 2^15-1
var intThirtyTwo int32 // 32 bits -2^31 to 2^31-1
var sixtyFour int64    // 64 bits -2^63 to 2^63-1
var integer int        // system dependent, 32 or 64 bits

var unsignedInt8 uint8   // 0 to 255
var unsignedInt16 uint16 // 0 to 2^16-1

var floatThirtyTwo float32
var floatSixtyFour float64

func main() {
	f := 123.45 // allocated to float 64

	fmt.Println(f)
}
