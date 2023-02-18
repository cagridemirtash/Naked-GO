package slices

import "fmt"

func SlicesIntro() {
	programmingLanguages := make([]string, 3)
	programmingLanguages[0] = "GO"
	programmingLanguages[1] = "JAVA"
	programmingLanguages[2] = "C#"
	programmingLanguages = append(programmingLanguages, "RUST")
	programmingLanguages = append(programmingLanguages, "SCALA")
	programmingLanguages = append(programmingLanguages, "ERLANG")
	programmingLanguages = append(programmingLanguages, "ELİXİR")
	fmt.Println(programmingLanguages)
	fmt.Println(len(programmingLanguages), cap(programmingLanguages))
}
