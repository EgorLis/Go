package main

import (
	"bytes"
	"fmt"
)

// IntSet представляет собой множество небольших неотрицательных
// целых чисел. Нулевое значение представляет пустое множество.
type IntSet struct {
	words []uint64
}

// Has указывает, содержит ли множество неотрицательное значение х.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add добавляет неотрицательное значение x в множество,
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith делает множество s равным объединению множеств s и t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String возвращает множество как строку вида "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Вычисление количество элементов в множестве
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for i := 0; i < 64; i++ {
			if word&(1<<i) != 0 {
				count++
			}
		}
	}
	return count
}

// Удаляем неотрицательное число из множества
func (s *IntSet) Remove(x int) {
	if x < 0 {
		return
	}
	word, bit := x/64, uint64(x%64)
	s.words[word] &^= 1 << bit

	// после очистки бита:
	last := len(s.words) - 1
	for last >= 0 && s.words[last] == 0 {
		last--
	}
	s.words = s.words[:last+1]
}

// Удаляем все элементы множества
func (s *IntSet) Clear() {
	s.words = nil
}

// Копируем множество
func (s *IntSet) Copy() *IntSet {
	var newSet IntSet
	newSet.words = make([]uint64, len(s.words))
	copy(newSet.words, s.words)
	return &newSet
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(&y)

	x.UnionWith(&y)
	fmt.Println(&x)

	fmt.Printf("%t, %t\n", x.Has(9), x.Has(123)) // true false

	fmt.Printf("Длина множества x: %d\n", x.Len())

	c := *x.Copy()

	x.Remove(144)
	fmt.Println(&x)
	fmt.Println(&c)

	x.Clear()
	fmt.Println(&x)
}
