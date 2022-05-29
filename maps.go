package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 100
	m["k2"] = 200
	fmt.Println("map:", m)

	fmt.Println("k1 =>", m["k1"])
	fmt.Println("k2 =>", m["k2"])
	fmt.Println("k3 =>", m["k3"]) // default zero value for value type
	_, prs := m["k3"]
	fmt.Println("prs for k3:", prs) // not present at this stage
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("k2 =>", m["k2"])
	fmt.Println("len:", len(m))

	m["k3"] = 300
	_, prs = m["k3"]
	fmt.Println("prs for k3:", prs) // now present

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)

}
