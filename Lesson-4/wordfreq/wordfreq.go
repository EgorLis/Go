// приложение для подсчета повторений
// слов во входном тексте из файла
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// 1. Открываем файл
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("cannot open file: %v", err)
		return
	}

	defer f.Close()

	// 2. Получаем частоту слов
	wordfreq := WordFreq(f)

	fmt.Println("Повторений слов:")

	for k, v := range wordfreq {
		fmt.Printf("%s - %d\n", k, v)
	}
}

func WordFreq(f *os.File) map[string]int {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		counts[word]++
	}

	return counts
}
