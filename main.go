#!/bin/bash
package main
import (
"github.com/01-edu/z01"
)
func main() {
for ch := 'a'; ch <= 'z'; ch++ {
z01.printRune(ch)
}
z01.printRune('\n')
}
