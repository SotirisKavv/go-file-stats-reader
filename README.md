# File Stats Reader — Text analysis and word frequency in Go

A file analysis tool that reads text files and generates detailed statistics including word frequency analysis. Demonstrates file I/O, text processing, data structures, and sorting algorithms in Go.

Quick links:
- Entrypoint: `filestats.go`
- Sample files: `files/`

---

## 🚀 What is this?

A CLI tool that analyzes text files to extract statistics like line count, word count, character count, and displays a word frequency histogram. Shows practical file processing and data analysis patterns.

---

## ✨ Features

- **File Statistics:** Lines, words, characters count per file
- **Word Frequency Analysis:** Top 5 most frequent words with percentages
- **Batch Processing:** Analyze multiple files in one command
- **Error Handling:** Graceful handling of missing or unreadable files
- **Data Structures:** Maps and slices for efficient data processing

---

## 🦄 Go Skills Demonstrated

- **File I/O:** `os.Open`, `bufio.Scanner` for reading files
- **String Processing:** `strings.Fields`, `strings.ToLower`
- **Data Structures:** Maps for frequency counting, slices for sorting
- **Sorting:** Custom sorting with `sort.Slice` and comparators
- **Struct Usage:** Custom types for organizing file statistics
- **Error Handling:** File access validation and error reporting

---

## 🛠️ Usage

```powershell
# Analyze single file
go run filestats.go file1.txt

# Analyze multiple files
go run filestats.go file1.txt file2.txt

# Example output:
# File: example.txt
# Lines: 42
# Words: 350
# Characters: 1847
# 
# Word histogram (top-5)
# the: 8.57%
# and: 6.29%
# is: 4.86%
# to: 3.43%
# of: 2.86%
```

---

## Folder map

- `filestats.go`: CLI, file parsing, stats, and histogram
- `files/`: Example text files


## Next steps (ideas)

- Add stop‑word filtering and stemming
- Add per‑file vs aggregate statistics modes
- Output results as JSON/CSV

## 🎯 Learning Objectives

This project demonstrates:
- **File Processing:** Reading and parsing text files efficiently
- **Data Analysis:** Statistical analysis and frequency calculations
- **Algorithm Implementation:** Sorting and ranking algorithms
- **Memory Management:** Efficient use of maps and slices

---

**Author:** IAmSotiris

