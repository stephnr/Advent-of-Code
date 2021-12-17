package main

import (
	"fmt"
	"math"
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

func ParseBits(bits *string, resTable []int64) (versionSum int64, resOut []int64) {
	if len(*bits) == 0 {
		return versionSum, resTable
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

		resTable = append(resTable, ConvertBitsToDecimal(num))

		return versionSum, resTable
	}

	opType := ConvertBitsToDecimal(TrimBits(bits, 1))

	if opType == 0 {
		packetsLength := ConvertBitsToDecimal(TrimBits(bits, 15))
		pre := len(*bits)

		for int64(pre-len(*bits)) < packetsLength {
			v, r := ParseBits(bits, resTable)
			versionSum += v
			resTable = r
		}

		resTable = ComputePartTwo(lengthTypeID, resTable)
	} else if opType == 1 {
		iter := ConvertBitsToDecimal(TrimBits(bits, 11))

		for iter > 0 {
			v, r := ParseBits(bits, resTable)
			versionSum += v
			resTable = r
			iter--
		}

		resTable = ComputePartTwo(lengthTypeID, resTable)
	}

	for len(*bits) > 0 {
		v, r := ParseBits(bits, resTable)
		versionSum += v
		resTable = r
	}

	return versionSum, resTable
}

func ComputePartTwo(lengthTypeID int64, resTable []int64) []int64 {
	// Handle Length Type ID
	switch lengthTypeID {
	case 0:
		// Sum
		res := int64(0)
		for _, v := range resTable {
			res += v
		}

		resTable = []int64{res}
	case 1:
		// Product
		res := int64(1)
		for _, v := range resTable {
			res *= v
		}

		resTable = []int64{res}
	case 2:
		// Minimum
		if len(resTable) == 1 {
			break
		}

		res := resTable[1]
		for _, v := range resTable[1:] {
			res = int64(math.Min(float64(res), float64(v)))
		}

		resTable = []int64{res}
	case 3:
		// Maximum
		if len(resTable) == 1 {
			break
		}

		res := resTable[1]
		for _, v := range resTable[1:] {
			res = int64(math.Max(float64(res), float64(v)))
		}

		resTable = []int64{res}
	case 5:
		// Greater Than = 1 if (0 > 1) else 0
		if resTable[0] > resTable[1] {
			resTable = []int64{1}
		} else {
			resTable = []int64{0}
		}
	case 6:
		// Less Than = 1 if (0 < 1) else 0
		if resTable[0] < resTable[1] {
			resTable = []int64{1}
		} else {
			resTable = []int64{0}
		}
	case 7:
		// Equals = 1 if (0 == 1) else 0
		if resTable[0] == resTable[1] {
			resTable = []int64{1}
		} else {
			resTable = []int64{0}
		}
	}

	return resTable
}

func main() {
	bits := ConvertToBits(ReadFile("./inputs/1")[0])
	p1, p2 := ParseBits(&bits, make([]int64, 0))
	fmt.Printf("%d , %+v\n", p1, p2)
}
