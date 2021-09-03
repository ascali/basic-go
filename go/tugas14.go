package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var execTerminal2, _ = exec.Command("ifconfig").Output()
	fmt.Println(string(execTerminal2))
	var execTerminal, _ = exec.Command("ls").Output()
	fmt.Println(string(execTerminal))
	var execTerminal3, _ = exec.Command("go", "help").Output()
	fmt.Println(string(execTerminal3))
}
