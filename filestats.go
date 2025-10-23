package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Filestats struct {
	filename string
	lines    int
	words    int
	chars    int
	wordFreq map[string]float64
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run filestats.go <file1.txt> [file2.txt ...]")
		os.Exit(1)
	}

	flag.Parse()
	filenames := flag.Args()

	for _, fn := range filenames {

		f, err := os.Open(fn)
		if err != nil {
			fmt.Printf("Error: File %s not found.\n", fn)
			continue
		}
		stats := &Filestats{
			filename: fn,
			wordFreq: make(map[string]float64),
		}
		processFile(f, stats)
		defer f.Close()

		printStats(*stats)
	}

}

func processFile(file *os.File, stats *Filestats) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stats.lines += 1
		words := strings.Fields(line)
		stats.words += len(words)

		for _, word := range words {
			word = strings.ToLower(word)
			stats.chars += len(word)
			if _, ok := stats.wordFreq[word]; !ok {
				stats.wordFreq[word] = 0
			}
			stats.wordFreq[word] += 1
		}
	}
}

func printStats(stats Filestats) {
	fmt.Printf("File: %s\n", stats.filename)
	fmt.Printf("Lines: %d\n", stats.lines)
	fmt.Printf("words: %d\n", stats.words)
	fmt.Printf("Characters: %d\n", stats.chars)
	fmt.Println()

	// sort and print word frequencies
	sortWordMap(stats.wordFreq, stats.words)
	fmt.Println()
}

func sortWordMap(wm map[string]float64, totWords int) {
	type keyVal struct {
		key string
		val float64
	}

	for k := range wm {
		wm[k] /= float64(totWords)
		wm[k] *= 100
	}

	var ss []keyVal
	for k, v := range wm {
		ss = append(ss, keyVal{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	fmt.Println("Word histogram (top-5)")
	for i, kv := range ss {
		fmt.Printf("%s: %.2f%%\n", kv.key, kv.val)
		if i >= 4 {
			break
		}
	}
}
