package aurorium

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	program := New()

	err := program.Main()
	if err != nil {
		panic(err)
	}

	m.Run()
	os.Exit(0)
}

func TestRun(t *testing.T) {

}
