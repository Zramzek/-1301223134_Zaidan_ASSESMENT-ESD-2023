package main

import (
	"fmt"
	"sort"
)

func main() {
	// Input data
	inputData := [][]string{
		{"Bagas", "Dimas", "Bagas", "Bagas", "Indra", "Gilang", "Gilang", "Hana", "Fajar", "Fajar"},
		{"Bagas", "Dimas", "Fajar", "Bagas", "Indra", "Gilang", "Gilang", "Bagas", "Fajar", "Fajar"},
		{"Aisyah", "Bagas", "Dewi", "Dimas", "Eka", "Fajar", "Gilang", "Hana", "Indra", "Jihan"},
	}

	counter := make(map[string]int)
	for _, data := range inputData {
		for _, name := range data {
			counter[name]++
		}
	}

	sortedNames := make([]string, 0, len(counter))
	for name := range counter {
		sortedNames = append(sortedNames, name)
	}
	sort.Slice(sortedNames, func(i, j int) bool {
		return counter[sortedNames[i]] > counter[sortedNames[j]]
	})

	statusAnak := make(map[string]string)
	for _, name := range sortedNames {
		if counter[name] > 1 {
			statusAnak[name] = "Nakal"
		} else {
			statusAnak[name] = "Baik"
		}
	}

	output := ""
	for _, name := range sortedNames {
		output += fmt.Sprintf("%s %s, ", name, statusAnak[name])
	}

	output = output[:len(output)-2]
	fmt.Println(output)
}
