package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func cetak(isi []int, atas int) {
	for i := atas; i > 0; i-- {
		for _, k := range isi {
			if i <= k {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
	for _, i := range isi {
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println("")
}

func main() {
	var data []int
	tinggi := 0

	reader := bufio.NewReader(os.Stdin)
	read, _ := reader.ReadString('\r')
	read = strings.Replace(read, "\r", "", -1)
	temp := strings.Split(read, " ")
	for _, i := range temp {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		data = append(data, j)
	}
	// fmt.Println(data)
	for _, i := range data {
		if tinggi < i {
			tinggi = i
		}
	}
	// fmt.Println(tinggi)
	cetak(data, tinggi)
	// for i, j := range data {
	// 	for index := 0; index < i+1; index++ {
	// 		if j < data[index] {
	// 			j = data[index]
	// 			data[index] = data[i]
	// 			data[i] = j
	// 		}
	// 	}
	// }
	sort.Sort(sort.IntSlice(data))
	cetak(data, tinggi)
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	cetak(data, tinggi)
}
