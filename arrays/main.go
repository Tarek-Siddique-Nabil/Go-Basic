package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		? array of 6 integers
	*/

	// var a [6]int32
	// a = [6]int32{7, 5, 3, 1, 9, 8}
	// fmt.Println(a[3])
	// fmt.Println(a[1:4])
	// fmt.Println(&a[0])
	// fmt.Println(&a[4])

	/*
	 * array slice
	 */
	// slice := []int32{4, 5, 6}

	// fmt.Printf("the length %v with capacity %v", len(slice), cap(slice))

	// slice = append(slice, 2)
	// fmt.Printf("\nthe length %v with capacity %v", len(slice), cap(slice))

	// slice = make([]int32, 0, 4)
	// slice = append(slice, 1, 2, 3, 4)
	// fmt.Printf("\nthe length %v with capacity %v", len(slice), cap(slice))
	// fmt.Println(slice)

	// m := make(map[string]int32)
	// m["alice"] = 30
	// m["bob"] = 54
	// age, exists := m["alic"]
	// if exists {
	// 	fmt.Println("Charlie's age is", age)
	// } else {
	// 	fmt.Println("Charlie is not in the map.")
	// }
	// fmt.Println(m)
	// delete(m, "bOb")
	// fmt.Println(m)

	// for key, value := range m {
	// 	fmt.Printf("\n%v is %v years old", key, value)
	// }

	/*
	 * difference slice between with preallocation and without preallocation
	 */

	n := 10000000
	slice1 := []int{}
	slice2 := make([]int, 0, n)
	fmt.Printf("total time take without preallocation %v\n", timeLoop(slice1, n))
	fmt.Printf("total time take with preallocation %v\n", timeLoop(slice2, n))

	/*
	* real life example of golang map

	 */

	// text := "Hey , whatsapp sokina. are you feeling good ? I am feeling good "

	// words := strings.Fields(text)
	// fmt.Println(words)

	// frequency := make(map[string]int32)

	// for _, word := range words {
	// 	frequency[word]++
	// }

	// for word, count := range frequency {
	// 	fmt.Printf("total %v times use %v in this sentance \n", count, word)
	// }

	/*
	 *loop
	 */

	// start := time.Now()
	// x := big.NewInt(1)
	// for i := 1; i < 10000; i++ {
	// 	x.Mul(x, big.NewInt(int64(i)))
	// }

	// elapsed := time.Since(start)
	// fmt.Println(x)
	// fmt.Printf("it take %v time", elapsed)

}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
