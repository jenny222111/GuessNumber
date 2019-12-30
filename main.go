package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int
	var chance int = 10
	var usedChance int =1
	now := time.Now()
	rand.Seed(now.UnixNano())
	for {		
		fmt.Println("input the guess number, from 0-100!~ ")
		fmt.Scanln(&num)
		fmt.Printf("SO your Num is %d\n", num)
		random := rand.Intn(100) + 1
		fmt.Println("Rand Num is: ", random)
		if num == random{
			switch  {
				case usedChance==1:
					fmt.Println("guess right, you are genius!`")
				case usedChance>=2 && usedChance<=3:
					fmt.Println("guess right, you are smart！")
				case usedChance>=4 && usedChance<=9:
					fmt.Println("guess right, but it is very difficult to you！")
				case  usedChance>=chance:
					fmt.Println("guess right, too many times!~")
			}
			break
		}else {
			if usedChance >= chance{
				fmt.Printf("you have used %d times, still failed\n", usedChance)
				fmt.Println()
			}else {
				fmt.Printf("you aleardy used %d times\n", usedChance)
				fmt.Println()
			}
			usedChance++
		}
	}
}