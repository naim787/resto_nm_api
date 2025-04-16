package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"resto_nm_api/internal/models"
	"testing"
)

func Test_PostUsers(t *testing.T) {
	data := []models.Users{
		{
			Name : "Naim",
			Id : "123",
			Email : "naimmmmab@gmail.com",
			Password : "123",
			Bis_Loc : "paniki",
			Date_Loc : "2025",
			Year : "2004",
			Role : "admin",
		},
	}

	jsonData, _ := json.Marshal(data)

	resp, err := http.Post("http://127.0.0.1:3000/create-users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)

		return
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&data)

	fmt.Println("Response:",data)
}



func Test_GetUsers(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:3000/users")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}

func Test_DeleteUsers(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:3000/delete-users")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}




func Test_logika(t *testing.T) {
	var d []models.Users
	 users := models.Users{
		Name : "naim",
		Id : "12355",
		Email : "naimm@gmail.com",
		Password : "123",
	}

	d = append(d, users)
	fmt.Println(d)
}