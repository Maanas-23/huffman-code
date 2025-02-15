package huffman

import "fmt"

type Decoder struct {
	encoded    string
	symbolData []Data

	decoded   string
	symbolMap map[string]rune
}

func NewDecoder(encoded string, symbolData []Data) *Decoder {
	decoder := Decoder{
		encoded:    encoded,
		symbolData: symbolData,
	}

	decoder.symbolMap = make(map[string]rune)
	for _, symbol := range symbolData {
		decoder.symbolMap[symbol.Code] = symbol.Symbol
	}

	return &decoder
}

func (d *Decoder) Decode() (string, error) {
	var current string
	for _, char := range d.encoded {
		current += string(char)
		symbol, ok := d.symbolMap[current]
		if ok {
			d.decoded += string(symbol)
			current = ""
		}
	}

	if current != "" {
		return "", fmt.Errorf("invalid encoded message")
	}
	return d.decoded, nil
}
