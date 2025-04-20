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
    // Data yang akan dikirim ke endpoint
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

    // Konversi data ke JSON
    jsonData, _ := json.Marshal(data)

    // Kirim permintaan POST ke endpoint
    resp, err := http.Post("http://127.0.0.1:3000/create-users", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        t.Fatalf("Error saat mengirim request: %v", err)
    }
    defer resp.Body.Close()

    // Struktur respons yang diharapkan
    var response map[string]any

    // Decode respons JSON
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        t.Fatalf("Error saat decoding response JSON: %v", err)
    }

    fmt.Println(response)

    // Cetak respons untuk melihat apa yang dikembalikan
    // fmt.Println("Response Message:", response.Message)
    // fmt.Println("Response Data:", response.Data)

    // // Validasi apakah data berhasil disimpan
    // if len(response.Data) == 0 || response.Data[0].Name != "Naim" {
    //     t.Fatalf("Data tidak sesuai atau tidak disimpan dengan benar")
    // }
}



func Test_GetUsersAfterPost(t *testing.T) {
    // Panggil endpoint /users
    resp, err := http.Get("http://127.0.0.1:3000/users")
    if err != nil {
        t.Fatalf("Error: %v", err)
    }
    defer resp.Body.Close()

     // Validasi status code
     if resp.StatusCode != http.StatusOK {
        t.Fatalf("Status code tidak sesuai, dapat: %d, ingin: %d", resp.StatusCode, http.StatusOK)
    }

    var response struct {
        Message string          `json:"message"`
        Data    []models.Users  `json:"data"`
    }
    
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        t.Fatalf("Error saat decoding response JSON: %v", err)
    }

    // Validasi struktur response
    if len(response.Data) == 0 {
        t.Fatalf("Data pengguna kosong, seharusnya ada data pengguna")
    }

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