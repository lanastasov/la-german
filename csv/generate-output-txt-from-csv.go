package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the CSV file
	csvFile, err := os.Open("aqa-german-gcse-vocab-list-[81839].csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Create a new reader
	reader := csv.NewReader(csvFile)
	reader.Comma = ','       // Set the delimiter to comma
	reader.LazyQuotes = true // Allow lazy quotes
	reader.TrimLeadingSpace = true

	// Open the output file
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	// Read each record from csv
	for {
		record, err := reader.Read()
		if err != nil {
			if err == csv.ErrFieldCount {
				continue // Skip on bad line
			}
			break // Stop at EOF or other error
		}

		if len(record) >= 2 {
			// Write the first and second element to the output file
			rcrd := ""
			for i := 0; i < len(record); i++ {
				if strings.Contains(record[i], "::") {
					rcrd += strings.Split(record[i], "::")[1]
				}
			}

			line := fmt.Sprintf("%s | %s | %s\n", strings.TrimSpace(record[0]), strings.TrimSpace(record[1]), rcrd)
			writer.WriteString(line)
		}
	}

	// Flush remaining data
	writer.Flush()
}
