package main

import "fmt"

func main() {
	a := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	a["FEF"] = []string{"test", "test1", "test2"}

	delete(a, "bond_james")

	for k, v := range a {
		fmt.Println("======= ", k, " ======")
		for i, vv := range v {
			fmt.Printf("\t [%v] = %v \n", i, vv)
		}
	}

}
