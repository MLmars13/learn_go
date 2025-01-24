package main

import (
	"fmt"
)

type MyStruct struct {
	value string
}

type MyInterface interface {
	plainReceiver() string
	pointerReceiver() string
}

func (m MyStruct) plainReceiver() string {
	fmt.Printf("inside m address: %p\n", &m)
	m.value = "[changed by plain receiver]"
	return "method on plain receiver " + m.value
}

func (m *MyStruct) pointerReceiver() string {
	fmt.Printf("inside: %p\n", m)
	fmt.Printf("%p, %p\n", &(m.value), &((*m).value))
	// (*m).value = "[changed by pointer receiver]"
	// return "method on pointer receiver " + (*m).value
	m.value = "[changed by pointer receiver]"
	return "method on pointer receiver " + m.value
}

func main() {
	// m := MyStruct{"MyStruct"}
	// fmt.Fprintf(os.Stdout, "%s, %s", m.plainReceiver(), m.pointerReceiver())
	// fmt.Fprintf(os.Stdout, "\n.......%s, %s", (&m).plainReceiver(), (&m).pointerReceiver())

	// fmt.Printf("\nbefore m address: %p\n", &m)
	// (&m).plainReceiver()
	// //(&m).pointerReceiver()
	// fmt.Fprintf(os.Stdout, "\n%s", m.value)

	// var s2 MyInterface
	// s2 = &m
	// fmt.Printf("before s2: %p, value = %s\n", s2, (&m).value)
	// s2.pointerReceiver()
	// fmt.Fprintf(os.Stdout, "\nafter value: %s", m.value)

	test_embedded()
}

type T struct {
	value string
}

type S struct {
	outerValue string
	*T
}

func (t T) plainReceiver() string {
	fmt.Printf("inside plainReceiver(), t address: %p, %p\n", &t, &(t.value))
	return t.value
}

func (t *T) pointerReceiver() string {
	fmt.Printf("inside pointerReceiver(), t address: %p, %p\n", t, &(t.value))
	return t.value
}

func test_embedded() {

	s := S{
		outerValue: "outer value",
		T:          &T{value: "inner value"},
	}
	// fmt.Printf("\nbefore plainReceiver(), t address: %p, %p, %p\n", s.T, &(s.value), &((*(s.T)).value))
	// (&s).plainReceiver()
	// (&s).pointerReceiver()

	// fmt.Printf("\nbefore pointerReceiver(), t address: %p, %p\n", s.T, &s.value)
	// s.pointerReceiver()

	var s3 MyInterface
	s3 = &s
	s3.plainReceiver()
	s3.pointerReceiver()
}
