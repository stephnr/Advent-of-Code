package main

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Disk struct {
	ID     int
	Length int
	IsFree bool
}

func run(filepath string) int {
	lines := ReadInput(filepath)
	sum := 0

	diskCount := -1
	diskmap := []Disk{}

	for line := range lines {
		l := strings.Split(line, "")

		for i := 0; i < len(l); i++ {
			// File Block
			if i%2 == 0 {
				diskCount++
				a, _ := strconv.Atoi(l[i])
				diskmap = append(diskmap, Disk{ID: diskCount, Length: a, IsFree: false})
			}

			// Free-Space Block
			if i%2 == 1 {
				a, _ := strconv.Atoi(l[i])
				diskmap = append(diskmap, Disk{Length: a, IsFree: true})
			}
		}
	}

	cont := false

	for {
		diskmap, cont = moveFileBlocks(diskmap)

		if !cont {
			break
		}
	}

	// Calculate the checksum
	diskString := getDiskMap(diskmap)
	fmt.Printf("%+v\n", diskString)

	id := 0
	for _, disk := range diskmap {
		if disk.IsFree {
			break
		}

		for i := 0; i < disk.Length; i++ {
			sum += (id * disk.ID)
			id++
		}
	}

	return sum
}

// Sends true if restart
func moveFileBlocks(diskmap []Disk) ([]Disk, bool) {
	// Squeeze from right side into left-most empty bits
	for i, disk := range diskmap {
		if !disk.IsFree {
			continue
		}

		// Find right-most disk
		rD := -1

		for j := len(diskmap) - 1; j >= 0; j-- {
			if !diskmap[j].IsFree && diskmap[j].Length > 0 {
				rD = j
				break
			}
		}

		if rD <= i {
			return diskmap, false
		}

		// Fill free space when space is equals
		if disk.Length == diskmap[rD].Length {
			diskmap[i].ID = diskmap[rD].ID
			diskmap[i].IsFree = false
			diskmap[rD].IsFree = true
			return diskmap, true
		} else if disk.Length < diskmap[rD].Length {
			// Fill when free needs less than rightmost
			diskmap[i].ID = diskmap[rD].ID
			diskmap[i].IsFree = false

			diskmap[rD].Length -= diskmap[i].Length

			diskmap = slices.Insert(diskmap, rD+1, Disk{
				ID:     -1,
				Length: diskmap[i].Length,
				IsFree: true,
			})
			return diskmap, true
		} else if disk.Length > diskmap[rD].Length {
			diskmap[i].Length -= diskmap[rD].Length
			diskmap[rD].IsFree = true

			// Fill when free space is larger than rightmost
			diskmap = slices.Insert(diskmap, i, Disk{
				ID:     diskmap[rD].ID,
				Length: diskmap[rD].Length,
				IsFree: false,
			})

			return diskmap, true
		}
	}

	return diskmap, false
}

func getDiskMap(diskmap []Disk) string {
	var s bytes.Buffer

	for _, disk := range diskmap {
		for j := 0; j < disk.Length; j++ {
			if disk.IsFree {
				s.WriteString(".")
			} else {
				s.WriteString(strconv.Itoa(disk.ID))
			}
		}
	}

	return s.String()
}

// Part Two: FAILED
// 15679713401325 = too high
// 15679713400223 = too high
// 15707113247944 = too high
// 12900781798635 = X
// 12844947102917 = X

func main() {
	println(run("./inputs/1.txt"))
}
