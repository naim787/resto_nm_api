package main

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

func main() {
    fmt.Println(C.GoString(C.CString("Hello CGO")))
}
