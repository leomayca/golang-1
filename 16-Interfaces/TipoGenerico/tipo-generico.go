package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func generica2(interf any) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Tipo genérico")

	generica("string")
	generica(1)
	generica(true)

	generica2("string")
	generica2(1)
	generica2(true)

	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(mapa)
}
