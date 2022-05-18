package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

type Robot struct {
	Icon      string
	PaintIcon string
	x, y      int
}

const mapSize = 12

type GameMapStruct struct {
	field   [mapSize][mapSize]string
	finishX int
	finishY int
}

func (gameMap *GameMapStruct) loadMap1() {
	gameMap.field = [mapSize][mapSize]string{
		{"☒", "☒", "☒", "☒", "☒", "☒", "☒", "☒", "☒", "☒", "☒", "☒"},
		{"☒", " ", "*", " ", " ", " ", " ", " ", " ", " ", " ", "☒"},
		{"☒", " ", "*", " ", "*", "*", "*", "*", " ", "*", " ", "☒"},
		{"☒", " ", " ", " ", " ", "*", " ", "*", " ", "*", " ", "☒"},
		{"☒", " ", "*", "*", " ", "*", " ", "*", " ", " ", " ", "☒"},
		{"☒", " ", "*", " ", " ", " ", " ", "*", "*", "*", "*", "☒"},
		{"☒", " ", "*", "*", "*", "*", "*", "*", " ", " ", "🍨", "☒"},
		{"☒", " ", "*", " ", "*", " ", " ", "*", " ", "*", "*", "☒"},
		{"☒", " ", " ", " ", "*", " ", " ", "*", " ", " ", " ", "☒"},
		{"☒", " ", "*", " ", "*", "*", " ", "*", "*", "*", " ", "☒"},
		{"☒", " ", "*", " ", " ", " ", " ", " ", " ", " ", " ", "☒"},
		{"☒", "☒", "☒", "☒", "☒", "☒", "☒", "☒", "☒", "☒", "☒", "☒"}}
	gameMap.finishX = 10
	gameMap.finishY = 6
}

func clearConsole() {
	os.Stdout.Write([]byte{0x1B, 0x5B, 0x33, 0x3B, 0x4A, 0x1B, 0x5B, 0x48, 0x1B, 0x5B, 0x32, 0x4A})
}

func (gameMap *GameMapStruct) print() {
	clearConsole()
	for _, v := range gameMap.field {
		for _, v2 := range v {
			fmt.Print(v2 + "\t")
		}
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
}

var GameMap GameMapStruct

var robot Robot
var oldChar string = " "

func main() {
	robot = Robot{Icon: "🤖", PaintIcon: "🍬", x: 1, y: 1}
	GameMap.loadMap1()
	GameMap.field[robot.y][robot.x] = robot.Icon
	GameMap.print()

	if err := keyboard.Open(); err != nil {
		return
	}
	defer keyboard.Close()

	var char rune = ' '
	for char != 'q' {
		c, _, err := keyboard.GetSingleKey()
		char = c
		if err != nil {
			panic(err)
		}

		if char == 's' {
			if IsDownFree(robot.x, robot.y, true) {
				MoveDown(robot.x, robot.y)
			}
		}
		if char == 'w' {
			if IsUpFree(robot.x, robot.y, true) {
				MoveUp(robot.x, robot.y)
			}
		}
		if char == 'd' {
			if IsRightFree(robot.x, robot.y, true) {
				MoveRight(robot.x, robot.y)
			}
		}
		if char == 'a' {
			if IsLeftFree(robot.x, robot.y, true) {
				MoveLeft(robot.x, robot.y)
			}
		}
		if char == 'f' {
			Paint()
		}
		if robot.x == GameMap.finishX && robot.y == GameMap.finishY {
			fmt.Printf("Victory!")
			return
		}
		//fmt.Printf("You pressed: %q\r\n", char)
	}
}

func IsDownFree(x, y int, checkWall bool) bool {
	y++
	rangeCheck := y > mapSize-1
	return checkRange(x, y, rangeCheck, checkWall)
}

func MoveDown(x, y int) {
	GameMap.field[y][x] = oldChar
	robot.y++
	swap()
}

func IsUpFree(x, y int, checkWall bool) bool {
	y--
	rangeCheck := y < 0
	return checkRange(x, y, rangeCheck, checkWall)
}

func MoveUp(x, y int) {
	GameMap.field[y][x] = oldChar
	robot.y--
	swap()
}

func IsRightFree(x, y int, checkWall bool) bool {
	x++
	rangeCheck := x > mapSize-1
	return checkRange(x, y, rangeCheck, checkWall)
}

func MoveRight(x, y int) {
	GameMap.field[y][x] = oldChar
	robot.x++
	swap()
}

func IsLeftFree(x, y int, checkWall bool) bool {
	x--
	rangeCheck := x < 0
	return checkRange(x, y, rangeCheck, checkWall)
}

func MoveLeft(x, y int) {
	GameMap.field[y][x] = oldChar
	robot.x--
	swap()
}

func Paint() {
	oldChar = robot.PaintIcon
}

func swap() {
	oldChar = GameMap.field[robot.y][robot.x]
	GameMap.field[robot.y][robot.x] = robot.Icon
	GameMap.print()
}

func checkRange(x, y int, rangeCheck, checkWall bool) bool {
	if checkWall {
		return !(rangeCheck || checkIsWall(x, y))
	} else {
		return !(rangeCheck)
	}
}

func checkIsWall(x, y int) bool {
	return GameMap.field[y][x] == "*" || GameMap.field[y][x] == "☒"
}
