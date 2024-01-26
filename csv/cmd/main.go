package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	// Open the input file
	inputFile, err := os.Open("../output.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(inputFile)

	f, err := os.Create("aqa-german-gcse-vocab-list-[81839].html")
	check(err)
	defer f.Close()

	_, _ = f.WriteString("<!DOCTYPE html>\n")
	_, _ = f.WriteString("<html lang=\"en\">\n")
	_, _ = f.WriteString("</html>\n")
	_, _ = f.WriteString("<head>\n")
	_, _ = f.WriteString("<link rel=\"stylesheet\" href=\"style.css\">\n")
	_, _ = f.WriteString("</head>\n")
	_, _ = f.WriteString("<body>\n")
	_, _ = f.WriteString("<div class=\"rows\">\n")

	var counter = 1
	var counter2 = 1
	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line at each "|"
		parts := strings.Split(line, "|")
		if counter == 1 {
			_, _ = f.WriteString("<div class=rows" + fmt.Sprintf("%d", counter2) + ">\n")
			counter2 += 1
		}
		_, _ = f.WriteString("<div class=\"flashcard-container\">\n")
		_, _ = f.WriteString("	<div class=\"flashcard\">\n")
		_, _ = f.WriteString("		<div class=\"question\">\n")
		_, _ = f.WriteString("			<div class=\"content\">\n")
		_, _ = f.WriteString(fmt.Sprintf(parts[0]))
		_, _ = f.WriteString("			</div>\n")
		_, _ = f.WriteString("		</div>\n")
		_, _ = f.WriteString("		<div class=\"answer\">\n")
		_, _ = f.WriteString("			<div class=\"content\"\n>")
		_, _ = f.WriteString(fmt.Sprintf(parts[1]))
		_, _ = f.WriteString("			</div>\n")
		_, _ = f.WriteString("		</div>\n")
		_, _ = f.WriteString("	</div>\n")
		_, _ = f.WriteString("</div>\n")

		// Output each part on a new line
		for _, part := range parts {
			fmt.Println(strings.TrimSpace(part))
		}

		// Add an extra line break between records for readability
		fmt.Println()
		if counter == 10 {
			counter = 0
			_, _ = f.WriteString("</div>\n")
		}
		counter += 1
	}

	_, _ = f.WriteString("</div>\n")
	_, _ = f.WriteString("</body>\n")

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
}
