package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "oktfolio",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m, m2, m3)

	fmt.Println("Travsersing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)

	// retrieve a doesn't exist key, will return a empty value
	causeName := m["cause"]
	fmt.Println(causeName)

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")

	name, ok = m["name"]
	fmt.Println(name, ok)
}
