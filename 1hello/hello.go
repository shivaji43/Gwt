package main

// func main() {
// 	fmt.Println(Hello("hehe"))
// }

var HelloPrefix = "Hello "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == "Spanish" {
		HelloPrefix = "Hola "
	}
	return HelloPrefix + name
}
