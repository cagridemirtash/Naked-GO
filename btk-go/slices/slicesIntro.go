package slices

import "fmt"

func SlicesIntro() {
	programmingLanguages := make([]string, 3)
	programmingLanguages[0] = "Emin MTas"
	programmingLanguages[1] = "Çağrı"
	programmingLanguages[2] = "Samet"
	programmingLanguages = append(programmingLanguages, "ÇAğrı")
	programmingLanguages = append(programmingLanguages, "ÇAğrı")
	programmingLanguages = append(programmingLanguages, "ÇAğrı")
	programmingLanguages = append(programmingLanguages, "ÇAğrı")
	fmt.Println(programmingLanguages)
	fmt.Println(len(programmingLanguages), cap(programmingLanguages))
}
