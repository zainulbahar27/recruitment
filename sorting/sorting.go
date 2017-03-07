package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cetak(a []int) int {

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
	//fmt.Println(data)
	for _, i := range data {
		if tinggi < i {
			tinggi = i
		}
	}
	fmt.Println(tinggi)
	for i := tinggi; i > 0; i-- {
		for _, k := range data {
			if i <= k {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
	for _, i := range data {
		fmt.Print(i)
		fmt.Print(" ")
	}
}
