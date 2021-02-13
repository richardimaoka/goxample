package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

type Employee struct {
}

func main() {
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
