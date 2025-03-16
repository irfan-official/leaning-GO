# Go datatypes

```go
 bool // %v

 string // %q

 int int8 int16 int 32 int64 // %v

 uint uint8 uint16 uint32 uint64 uintptr

 bytes // alias for unit8 //

 rune // alias for int32 // one charcter in a string
      // represents a Unicode code point

 float32 float64 // %f

 complex64 complex128
```

```go
package main

import "fmt"

func main(){
// initialize variable here

var int_type int = 10;
var float_64 float64 = 10.30;
var boolean bool = true;
var str string = "Irfan"
fmt.Printf("%v %f %v %q\n",int_type, float_64, boolean, str )

}


```
