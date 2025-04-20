package testing

import (
	"fmt"
	"testing"
)

type st struct {
	name string
}

type d interface {
	get()
}

func (v *st) get() {
	hasil := st{
		name:"naimm",
	}
	fmt.Println(hasil)
}

func Test_Interface(t *testing.T) {
	instens := &st{}
	instens.get()
}