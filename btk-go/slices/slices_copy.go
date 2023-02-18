package slices

import "fmt"

func SlicesCopy() {
	programmingLanguages := []string{"GO", "C#", "JAVA"}
	fmt.Println(programmingLanguages, "Len :", len(programmingLanguages), "Cap :", cap(programmingLanguages))
	programmingLanguagesCopy := make([]string, len(programmingLanguages))
	fmt.Println(programmingLanguagesCopy)
	copy(programmingLanguagesCopy, programmingLanguages)
	fmt.Println(programmingLanguagesCopy)
	programmingLanguages = append(programmingLanguages, "SCALA")
	//This print only three element
	//Because go memory stack hold to first and second slices different address.
	fmt.Println(programmingLanguagesCopy)
	//Sub slice from main slice -> Use for pagination in Data Api
	fmt.Println(programmingLanguages[2:4])
}
