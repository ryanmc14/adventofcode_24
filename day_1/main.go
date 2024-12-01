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
    var list1 []int
    var list2 []int
    var part1Answer  int
    var part2Answer  int

	sc := advent_utils.Input("input.txt")
	re := regexp.MustCompile("   ")
	for sc.Scan() {
		results := re.Split(sc.Text(), -1)
		int1, err := strconv.Atoi(results[0])
		int2, err := strconv.Atoi(results[1])
		if err != nil {
			log.Fatal(err.Error())
		}

		list1 = append(list1, int1)
		list2 = append(list2, int2)
	}

    sort.Ints(list1)
    sort.Ints(list2)

    //part 1
    for i, _ := range list1 {
        part1Answer += diff(list1[i], list2[i])
    }
    fmt.Println(part1Answer)

    //part2
    for _, v := range list1{
        frequency := uniqueValue(v, list2)
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
