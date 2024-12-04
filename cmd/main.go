package main

import (
	"os"

	"github.com/RSSU-Shellcode/Aurorium"
)

func main() {
	program := aurorium.New()
	err := program.Main()
	if err == nil {
		return
	}
	_ = os.WriteFile("error.log", []byte(err.Error()), 0600)
}
