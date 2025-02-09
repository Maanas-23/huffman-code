package main

import (
	"fmt"
	"log"
	"os"

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
	encoder.CalculateMetrics()
	fmt.Printf("Encoded message:\n%q\n", encoded)

	decoder := huffman.NewDecoder(encoded, data)
	decoded, err := decoder.Decode()
	if err != nil {
		log.Fatalf("Error decoding message: %v", err)
	}
	fmt.Printf("Decoded message:\n%q\n", decoded)
}
