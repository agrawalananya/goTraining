package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := map[string]Vertex{"shiv": {1.9, 2.9}, "hii": {1.4, 2.9}}
	if m == nil {
		fmt.Println(nil)
	}
	array := [8]Vertex{{2.9, 3.8}, {4.8, 5}}
	fmt.Println(array)
	pt := &array
	fmt.Println(pt[0])
	for i, _ := range m {
		fmt.Println(i)
	}
	var name string
	var alphabet_count int

	fmt.Scanln(&name)
	fmt.Scanln(&alphabet_count)

	// Printing the given texts
	fmt.Printf("%s %d",
		name, alphabet_count)

	var arr []int
	fmt.Println(arr)
	var s [5]int
	for i := 0; i < 5; i++ {
		s[i] = 2
	}
	fmt.Println(s)
	structs := []Vertex{{2.3, 3.4}, {3.4, 34}}
	fmt.Println(structs)
	var m map[string]Vertex

	a := make(map[string]Vertex)
	a["key"] = Vertex{2.3, 4.5}
	fmt.Println(m)

}
