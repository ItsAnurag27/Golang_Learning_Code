package main

import "fmt"

func main() {
    var i int         // 0
    var f float32     // 0.0
    var b bool        // false
    var s string      // ""
    var arr [3]int    // [0, 0, 0]
    var slc []int     // nil (uninitialized slice)
    var m map[string]int // nil (uninitialized map)
    var p *int        // nil (pointer to int)

    fmt.Println(i, f, b, s, arr, slc, m, p)
    // Output: 0 0 false "" [0 0 0] [] map[] <nil>
}