package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type Point struct {
	x, Y int
}

func main() {
	var dilbert Employee
	dilbert.Salary -= 5000
	fmt.Printf("%v\n", dilbert)
	fmt.Printf("%+v\n", dilbert)
	position := dilbert.Position
	position = "Senior"
	fmt.Printf("%+v\n", position)
	fmt.Printf("%+v\n", dilbert)

	var p = Point{2, 2}
	fmt.Printf("%v", p)

	cmd := exec.Command("./script.sh")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		fmt.Println(stderr.String())
		log.Fatal("exit with error")
	}
	fmt.Printf("%s", stdout.String())
}
