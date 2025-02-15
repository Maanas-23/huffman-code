package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/Maanas-23/huffman-code/huffman"
)

func main() {
	inputFile := "input_encode.txt"
	if len(os.Args) >= 2 {
		inputFile = os.Args[1]
	}

	encoder, err := huffman.NewEncoder(inputFile)
	if err != nil {
		log.Fatalf("Error creating encoder: %v", err)
	}
	encoded, data := encoder.Encode()
	printData(data)
	fmt.Printf("\nEncoded message:\n%q\n", encoded)

	decoder := huffman.NewDecoder(encoded, data)
	decoded, err := decoder.Decode()
	if err != nil {
		log.Fatalf("Error decoding message: %v", err)
	}
	fmt.Printf("\nDecoded message:\n%q\n", decoded)

	encoder.CalculateMetrics()
}

func printData(data []huffman.Data) {
	sort.Slice(data, func(i, j int) bool {
		if data[i].Probability == data[j].Probability {
			if len(data[i].Code) == len(data[j].Code) {
				return data[i].Code < data[j].Code
			}
			return len(data[i].Code) < len(data[j].Code)
		}
		return data[i].Probability > data[j].Probability
	})

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()
	fmt.Fprintln(w, "\nSymbol\tProbability\tCode\t")
	fmt.Fprintln(w, "------\t-----------\t----\t")
	for _, d := range data {
		fmt.Fprintf(w, "%q\t%.4f\t%s\t\n", d.Symbol, d.Probability, d.Code)
	}
}
