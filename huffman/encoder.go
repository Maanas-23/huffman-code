package huffman

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Maanas-23/huffman-code/priority_queue"
)

type Data struct {
	Symbol      rune
	Probability float64
	Code        []byte
}

type Encoder struct {
	inputFile  string
	message    string
	base       int
	symbolData []Data

	encoded string
	codeMap map[rune][]byte

	entropy   float64
	avgLength float64
}

func NewEncoder(inputFile string) (*Encoder, error) {
	encoder := Encoder{
		inputFile: inputFile,
		codeMap:   make(map[rune][]byte),
	}
	if err := encoder.parseInput(); err != nil {
		return nil, err
	}

	return &encoder, nil
}

func (e *Encoder) CalculateMetrics() {
	totalCodeLength := 0
	for _, code := range e.codeMap {
		totalCodeLength += len(code)
	}
	e.avgLength = float64(totalCodeLength) / float64(len(e.codeMap))

	e.entropy = 0.0
	for _, d := range e.symbolData {
		e.entropy -= d.Probability * math.Log(d.Probability)
	}
	e.entropy /= math.Log(float64(e.base))

	fmt.Println("Avg Code length: ", e.avgLength)
	fmt.Println("Entropy: ", e.entropy)
}

func (e *Encoder) Encode() (string, []Data) {
	sort.Slice(e.symbolData, func(i, j int) bool {
		return e.symbolData[i].Probability < e.symbolData[j].Probability
	})
	extra := mod(1-len(e.symbolData), e.base-1)

	// Initialize pq
	pq := priority_queue.PriorityQueue{}
	for i, item := range e.symbolData {
		pq = append(pq, &priority_queue.Element{
			Priority: item.Probability,
			Items:    []int{i},
		})
	}
	for i := 0; i < extra; i++ {
		pq = append(pq, &priority_queue.Element{
			Priority: 0,
			Items:    nil,
		})
	}
	heap.Init(&pq)

	// Huffman coding main
	for pq.Len() > 1 {
		combined := priority_queue.Element{
			Priority: 0,
			Items:    nil,
		}
		for i := e.base - 1; i >= 0; i-- {
			lowest := heap.Pop(&pq).(*priority_queue.Element)
			combined.Priority += lowest.Priority
			combined.Items = append(combined.Items, lowest.Items...)

			for _, item := range lowest.Items {
				e.symbolData[item].Code = append(e.symbolData[item].Code, byte(i))
			}
		}
		heap.Push(&pq, &combined)
	}

	for _, d := range e.symbolData {
		for i := 0; i < len(d.Code)/2; i++ {
			d.Code[i], d.Code[len(d.Code)-i-1] = d.Code[len(d.Code)-i-1], d.Code[i]
		}
	}

	// map symbols to codes
	for _, d := range e.symbolData {
		e.codeMap[d.Symbol] = d.Code
	}

	// Encode message
	var encoded []byte
	for _, char := range e.message {
		for _, code := range e.codeMap[char] {
			encoded = append(encoded, getChar(int(code)))
		}
	}
	e.encoded = string(encoded)

	return e.encoded, e.symbolData
}

func mod(a, b int) int {
	// to ensure positive mod result
	return (a%b + b) % b
}

func getChar(i int) byte {
	if i < 10 {
		return byte(i + 0x30)
	}
	return byte(i - 9 + 0x60)
}

func (e *Encoder) parseInput() error {
	content, err := os.ReadFile(e.inputFile)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	if len(lines) < 2 {
		return fmt.Errorf("file must contain at least 2 lines")
	}

	base, err := strconv.Atoi(strings.TrimSpace(lines[0]))
	if err != nil {
		return fmt.Errorf("invalid base number: %v", err)
	}
	if base <= 1 {
		return fmt.Errorf("base must be greater than 1")
	}

	e.base = base
	e.message = strings.Join(lines[1:], "\n")

	freq := make(map[rune]int)
	total := 0

	for _, char := range e.message {
		freq[char]++
		total++
	}

	e.symbolData = make([]Data, 0, len(freq))
	for char, cnt := range freq {
		probability := float64(cnt) / float64(total)
		e.symbolData = append(e.symbolData, Data{
			Symbol:      char,
			Probability: probability,
		})
	}
	return nil
}
