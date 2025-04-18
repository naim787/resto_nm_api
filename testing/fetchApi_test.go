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
    data := models.Users{
        Name:     "Naim",
        Id:       "123",
        Email:    "naimmmmab@gmail.com",
        Password: "123",
        Bis_Loc:  "paniki",
        Date_Loc: "2025",
        Year:     "2004",
        Role:     "admin",
    }

    jsonData, _ := json.Marshal(data)

    resp, err := http.Post("http://127.0.0.1:3000/create-users", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        t.Fatalf("Error: %v", err)
    }
    defer resp.Body.Close()

    var response []models.Users
    json.NewDecoder(resp.Body).Decode(&response)

    fmt.Println("Response:", response)

    // Validasi apakah data berhasil disimpan
    if len(response) == 0 || response[0].Name != "Naim" {
        t.Fatalf("Data not saved correctly")
    }
}



func Test_GetUsersAfterPost(t *testing.T) {
    // Panggil endpoint /users
    resp, err := http.Get("http://127.0.0.1:3000/users")
    if err != nil {
        t.Fatalf("Error: %v", err)
    }
    defer resp.Body.Close()

    var response struct {
        Message string          `json:"message"`
        Data    []models.Users  `json:"data"`
    }
    json.NewDecoder(resp.Body).Decode(&response)

    fmt.Println("Response:", response)

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