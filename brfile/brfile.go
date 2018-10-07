package brfile

import (
	_ "bufio"
)

type offset int
type BrFile struct {
	off      offset
	original []byte
}

type index struct{
	id int
	fullName string
}

func (br *BrFile) Find(s string) (off offset) {
	for k:=range
return
}

func (br *BrFile) Read() {

}

func (br *BrFile) Write() {

}
