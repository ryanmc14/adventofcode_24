package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"

	"github.com/ryanmc14/adventofcode_24/advent_utils"
)

func main() {
    var listLeft []int
    var listRight []int
    var part1Answer  int
    var part2Answer  int

	sc := advent_utils.Input("input.txt")
	re := regexp.MustCompile("   ")
	for sc.Scan() {
		results := re.Split(sc.Text(), -1)
		intLeft, err := strconv.Atoi(results[0])
		intRight, err := strconv.Atoi(results[1])
		if err != nil {
			log.Fatal(err.Error())
		}

		listLeft = append(listLeft, intLeft)
		listRight = append(listRight, intRight)
	}

    sort.Ints(listLeft)
    sort.Ints(listRight)

    //part 1
    for i, _ := range listLeft {
        part1Answer += diff(listLeft[i], listRight[i])
    }
    fmt.Println(part1Answer)

    //part2
    for _, v := range listLeft{
        frequency := uniqueValue(v, listRight)
        part2Answer += (v*frequency)
    }

    fmt.Println(part2Answer)
}


func diff(a, b int) int {
   if a < b {
      return b - a
   }
   return a - b
}

func uniqueValue(value int, arr []int) int{
    frequency:= 0
    for _, num := range arr {
        if num != value{
            continue
        }
        frequency +=1
    }
    return frequency
}
