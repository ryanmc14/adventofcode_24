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
	sc := advent_utils.Input("input.txt")
	var list1 []int
	var list2 []int
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
    var sum  int
    var sum2  int
    //part 1
    for i, _ := range list1 {
        sum += diff(list1[i], list2[i])
    }
    fmt.Println(sum)

    //part2
    for _, v := range list1{
        frequency := uniqueValue(v, list2)
        sum2 += (v*frequency)
    }

    fmt.Println(sum2)
}

//two lists
// sort lists
// loop through and subtract

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
