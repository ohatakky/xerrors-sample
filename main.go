package main

import (
	"fmt"
	"os"

	"golang.org/x/xerrors"
)

func fileOpen(fname string) error {
	file, err := os.Open(fname)
	if err != nil {
		// これだとos.Openのエラーを捨ててしまっている
		// return fmt.Errorf("Error in fileOpen(): %v", err)
		// なので、Errorf()でerrをラップする
		return xerrors.Errorf("Error in fileOpen(): %w", err)
	}
	defer file.Close()
	return nil
}
func main() {
	err := fileOpen("sample.txt")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Println("OK")

}
