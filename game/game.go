package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Start() bool {
	fmt.Println("Rock-Paper-Scissors start.")
	fmt.Println("Please input number. Rock:0, Paper: 1, Scissors: 2")
	i := 0
	var s string
	for {
		if i > 5 {
			break
		}
		fmt.Println("Battle round ", i)
		if sc.Scan() {
			s = sc.Text()
			data, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("ERROR! Please input only number 0 - 2")
				continue
			}
			fmt.Println(data)
			// TODO
		}
		fmt.Println("Your ", s)
		i += 1
	}
	return true
}