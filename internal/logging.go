package logging

import (
	"fmt"
	"os"
)

func ErrorAndQuit(err error) {
	e := fmt.Errorf("Error %q", err)
	fmt.Println(e.Error())
	os.Exit(1)
}

func ErrorAndQuitF(err error, f string) {
	e := fmt.Errorf("%v %q", f, err)
	fmt.Println(e.Error())
	os.Exit(1)
}
