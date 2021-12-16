package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Operator struct {
	Version    int64
	LengthType int64
	SubPackets []int64
}

type Literal struct {
	Version    int64
	LengthType int64
	Value      int64
}

func part_one(input []string) int {
	return 0
}

func ConvertToBits(line string) (bits string) {
	charTable := "ABCDEF"

	for _, r := range strings.Split(line, "") {
		val, fail := strconv.Atoi(string(r))

		if fail != nil {
			// A-F
			val = strings.Index(charTable, string(r)) + 10
		}

		bitArr := strconv.FormatInt(int64(val), 2)

		if len(bitArr) < 4 {
			bitArr = "000" + bitArr
		}

		bits += bitArr[len(bitArr)-4:]
	}

	return bits
}

func LeftPad(bits string, n int) string {
	for i := 0; i < n; i++ {
		bits = "0" + bits
	}

	return bits
}

// TrimBits takes n bits from the left and returns the selection
func TrimBits(bits *string, n int) string {
	peek := (*bits)[0:n]
	*bits = (*bits)[n:]
	return peek
}

func ConvertBitsToDecimal(bits string) int64 {
	val, _ := strconv.ParseInt(bits, 2, 64)
	return val
}

func ParseInput(input string) (out []string) {
	bits := ConvertToBits(input)
	literals := make([]Literal, 0)
	// operators := make([]Operator, 0)

	println(bits)

	for len(bits) > 0 {
		peekVersion := "0" + TrimBits(&bits, 3)

		if ConvertBitsToDecimal(peekVersion) == 0 {
			// End
			break
		}

		peekTypeID := "0" + TrimBits(&bits, 3)

		packetVersion := ConvertBitsToDecimal(peekVersion)
		packetTypeID := ConvertBitsToDecimal(peekTypeID)

		if packetTypeID == 4 {
			num := ""

		PacketLoop:
			for {
				peekVal := TrimBits(&bits, 5)
				num += peekVal[1:]

				if peekVal[0] == '0' {
					// This was the last value
					break PacketLoop
				}
			}

			val := ConvertBitsToDecimal(num)

			// Store literal
			literals = append(literals, Literal{
				Version:    packetVersion,
				LengthType: packetTypeID,
				Value:      val,
			})

			fmt.Printf("%+v\n", literals[len(literals)-1])
		}
	}

	// First 3 bits = Version
	// Next 3 bits = Type ID
	// -> 4 - LITERAL VALUE

	return out
}

func main() {
	input := ParseInput(ReadFile("./inputs/1")[0])
	println(part_one(input))
	// println(part_two())
}
