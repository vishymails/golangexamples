package _2_good

import (
	"fmt"
	"testing"
)

var result string

func benchToString(b *testing.B, total int) {
	people := make([]Person, total)
	for x := 0; x < total; x++ {
		people[x] = Person{ID: x, Name: fmt.Sprintf("test %d", x)}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result = toString(people)
	}
}

func BenchmarkPerson1(b *testing.B) {
	benchToString(b, 1)
}

func BenchmarkPerson10(b *testing.B) {
	benchToString(b, 10)
}

func BenchmarkPerson1000(b *testing.B) {
	benchToString(b, 1000)
}

func toString(people []Person) string {
	var out string
	for _, person := range people {
		out += fmt.Sprintf("ID: %d\nName: %s", person.ID, person.Name)
	}
	return out
}

type Person struct {
	ID   int
	Name string
}
