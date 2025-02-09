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
	fmt.Println("Encoded message: ", encoded)
	fmt.Println(data)
	//encodedMsg, codeMap, err := encoder.Process()
	//if err != nil {
	//	log.Fatalf("Processing failed: %v", err)
	//}
	//
	//fmt.Printf("Original message length: %d\n", len(encoder.Message()))
	//fmt.Printf("Original message: %q\n", encoder.Message())
	//fmt.Printf("Encoded message length: %d\n", len(encodedMsg))
	//fmt.Printf("Encoded message: %q\n", encodedMsg)
	//fmt.Printf("Number of unique symbols: %d\n", len(codeMap))
}
