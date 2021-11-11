package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main(){
	var y = 8
	var x = 6
	var user = [1][2]int{{4,1}}
	var obstacle = [34][2]int{
		{0,0},{0,1},{0,2},{0,3},{0,4},{0,5},{0,6},{0,7},
		{5,0},{5,1},{5,2},{5,3},{5,4},{5,5},{5,6},{5,7},
		{1,0},{2,0},{3,0},{4,0},{5,0},{6,0},
		{1,7},{2,7},{3,7},{4,7},{5,7},{6,7},
		{2,2},{2,3},{2,4},{3,4},{3,6},{4,2},
	}

	firstPosition(user,x, y, obstacle)

	treasure := randomPosition()

	treasureInObstacle := false
	for {
		for _, o := range obstacle {
			if treasure[0] == o || treasure[0]==user[0] {
				treasure = randomPosition()
				treasureInObstacle = true
				break
			}
			treasureInObstacle = false
		}
		if !treasureInObstacle{
			break
		}
	}

	for  {
		info()
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		}

		positionBefore := user
		if char == 'a'{
			user[0][0] -= 1
		}else if char == 'b'{
			user[0][1] += 1
		}else if char == 'c'{
			user[0][0] += 1
		}else if char == 'd'{
			user[0][1] -= 1
		}else{
			fmt.Println("wrong input")
		}

		for _, o := range obstacle {
			if user[0] == o {
				fmt.Println("forbidden")
				user = positionBefore
			}
		}

		for i:=0; i<x; i++ {
			for j:=0; j<y; j++ {
				isObstacle := false
				//isTreasure := false
				isUser := false

				for _, u := range user {
					if i == u[0] && j == u[1] {
						isUser = true
					}
				}
				for _, o := range obstacle {
					if i == o[0] && j == o[1] {
						isObstacle = true
					}
				}
				//for _, t := range treasure {
				//	if i == t[0] && j == t[1] {
				//		isTreasure = true
				//	}
				//}

				if isObstacle {
					fmt.Print("#")
				}else if isUser{
					fmt.Print("X")
				//}else if isTreasure{
				//	fmt.Print("$")
				}else{
					fmt.Print(".")
				}

			}
			fmt.Println("")
		}

		if user == treasure{
			break
		}
	}
	fmt.Println("Horee you got the treasure")
}

func randomPosition() [1][2]int {
	rand.Seed(time.Now().UnixNano())
	var randomY = rand.Intn(7)
	var randomX = rand.Intn(5)
	return [1][2]int{{randomX,randomY}}
}

func firstPosition(user [1][2]int, x int, y int, obstacle [34][2]int){
	for i:=0; i<x; i++ {
		for j:=0; j<y; j++ {
			isObstacle := false
			isUser := false

			for _, u := range user {
				if i == u[0] && j == u[1] {
					isUser = true
				}
			}
			for _, o := range obstacle {
				if i == o[0] && j == o[1] {
					isObstacle = true
				}
			}

			if isObstacle {
				fmt.Print("#")
			}else if isUser{
				fmt.Print("X")
			}else{
				fmt.Print(".")
			}

		}
		fmt.Println("")
	}
}

func info(){
	log.Println("press a to move up")
	log.Println("press b to move right")
	log.Println("press c to move down")
	log.Println("press d to move left")
}