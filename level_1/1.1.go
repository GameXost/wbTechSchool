package main

import "fmt"

type Human struct {
	age  uint8
	name string
}

//метод структуры Human
//работает с указателем структуры, а не с копией

func (h *Human) SayHello() {
	fmt.Printf("%s says hello\n", h.name)
}

func (h *Human) Introducing() {
	fmt.Printf("Hello, my name is %s, i'm %d old\n", h.name, h.age)
}

// встраивание реализуется в гошке через анонимное поле структуры
type Action struct {
	Human // встраиваем Human
}

func main() {
	me := Action{ // В структуре Action встраиваем структуру Human
		Human{
			age:  18,
			name: "Sero",
		},
	}

	me.SayHello() // можно неявно использовать методы Human
	me.Introducing()

	me.Human.SayHello() // ли же можно явно прописывать
	me.Human.Introducing()

}
