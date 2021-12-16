package main

import (
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
			bitArr = LeftPad(bitArr, 4-len(bitArr))
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
	if len(*bits) < n {
		return ""
	}

	peek := (*bits)[0:n]
	*bits = (*bits)[n:]
	return peek
}

func ConvertBitsToDecimal(bits string) int64 {
	val, _ := strconv.ParseInt(bits, 2, 64)
	return val
}

func GetBitIDs(bits *string) (version, lengthTypeID int64) {
	version = ConvertBitsToDecimal(TrimBits(bits, 3))
	lengthTypeID = ConvertBitsToDecimal(TrimBits(bits, 3))

	return version, lengthTypeID
}

func ParseBits(bits *string) (versionSum int64) {
	if len(*bits) == 0 {
		return versionSum
	}

	version, lengthTypeID := GetBitIDs(bits)
	versionSum += version

	if lengthTypeID == 4 {
		// No action
		num := ""

	Scan:
		for {
			checkBits := TrimBits(bits, 5)
			num += checkBits[1:]

			if checkBits[0] == '0' {
				break Scan
			}
		}

		return versionSum
	}

	opType := ConvertBitsToDecimal(TrimBits(bits, 1))

	if opType == 0 {
		TrimBits(bits, 15)
		versionSum += ParseBits(bits)
	} else if opType == 1 {
		iter := ConvertBitsToDecimal(TrimBits(bits, 11))

		for iter > 0 {
			versionSum += ParseBits(bits)
			iter--
		}
	}

	versionSum += ParseBits(bits)
	return versionSum
}

func main() {
	bits := ConvertToBits(ReadFile("./inputs/1")[0])
	println(ParseBits(&bits))
}
