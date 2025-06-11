package generator

import "strconv"

func GenerateFlatStringJson() {
	dirname := "flat"
	filename := "random_flat.json"

	flatMap := make(map[string]interface{})
	for i := 0; i < 10000; i++ {
		key := "key_" + strconv.Itoa(i)
		flatMap[key] = randomValue()
	}

	writeFile(dirname, filename, flatMap)
}
