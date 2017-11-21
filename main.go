package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n int

func main() {
	fmt.Println("Enter pole height: ")
	reader := bufio.NewReader(os.Stdin)
	text, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	n, err = strconv.Atoi(string(text))
	if err != nil {
		panic(err)
	}

	source := &pole{make([]int, 0)}
	temp := &pole{make([]int, 0)}
	dest := &pole{make([]int, 0)}

	pm := make(map[int]*pole) // For visual...
	pm[0] = source
	pm[1] = temp
	pm[2] = dest

	for i := 1; i <= n; i++ {
		source.disks = append([]int{i}, source.disks...)
	}
	fmt.Println(source, temp, dest)
	move(source, dest, temp, n, pm)
}

func move(source, dest, temp *pole, n int, pm map[int]*pole) {
	if n <= 0 {
		return
	}
	move(source, temp, dest, n-1, pm)
	x := source.disks[len(source.disks)-1]
	source.disks = source.disks[:len(source.disks)-1]
	dest.disks = append(dest.disks, x)
	printPoles(pm)
	move(temp, dest, source, n-1, pm)
}

type pole struct {
	disks []int
}

func printPoles(pm map[int]*pole) {
	for i := n; i >= 0; i-- {
		fmt.Printf("%s %s %s\n", pm[0].getLine(i), pm[1].getLine(i), pm[2].getLine(i))
	}
	fmt.Println(pm[0].String(), pm[1].String(), pm[2].String())
	fmt.Println(strings.Repeat("-", 20))
}

func (p *pole) getLine(i int) (str string) {
	if len(p.disks) > i {
		str = strings.Repeat("#", p.disks[i])
	}
	str += strings.Repeat(" ", n-len(str))
	return
}

func (p *pole) String() string {
	return fmt.Sprint(p.disks)
}
