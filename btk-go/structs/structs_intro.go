package structs

type User struct {
	name string
	age  int
}

func StructsIntro() User {
	user := User{name: "Çağrı", age: 23}
	return user
}
