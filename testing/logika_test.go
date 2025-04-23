package testing

import (
	"fmt"
	"resto_nm_api/internal/service"
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




func Test_FindEmail(t *testing.T) {

	email := "naim@gmail.com"
    found, err := service.FindEmail[map[string]interface{}]("users", email)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if found {
        fmt.Println("Email nnya ada cuy")
    } else {
        fmt.Println("Email nnya tidak ada cuy")
    }
}