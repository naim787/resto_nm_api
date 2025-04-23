const data = {
    Name: "Tono",
    Id: "255",
    Email: "naimmmmab@gmail.com",
    Password: "255",
    Bis_Loc: "paniki",
    Date_Loc: "2025",
    Year: "2004",
    Role: "admin"
};
// Fungsi untuk menguji GET /users
async function testGetUsers() {
    try {
        const response = await fetch(`http://127.0.0.1:3000/users`, {
            method: 'GET'
        });

        const result = await response.json();
        console.log('GET /users Response:', result);

        if (result.data.length === 0) {
            console.error('Test failed: Data pengguna kosong, seharusnya ada data pengguna');
        } else {
            console.log('Test passed: Data pengguna berhasil diambil');
        }
    } catch (error) {
        console.error('Error saat mengirim request GET /users:', error.message);
    }
}



async function testDeleteUsers() {
    try {
        const response = await fetch(`http://127.0.0.1:3000/delete-users`, {
            method: 'GET'
        });

        const result = await response.json();
        console.log('GET /delete-users Response:', result);

        if (result.message !== "Users deleted successfully") {
            console.error('Test failed: Data pengguna tidak berhasil dihapus');
        } else {
            console.log('Test passed: Data pengguna berhasil dihapus');
        }
    } catch (error) {
        console.error('Error saat mengirim request GET /delete-users:', error.message);
    }
}


// Fungsi untuk mengirim data ke server
async function testPostUsers() {
    try {
        // Kirim permintaan POST ke endpoint
        const response = await fetch('http://127.0.0.1:3000/create-users', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        // Parse respons JSON
        const result = await response.json();

        // Cetak respons untuk melihat hasilnya
        console.log('Response:', result);

        // Validasi apakah data berhasil disimpan
        if (result.data.length === 0 || result.data[0].Name !== "Naim") {
            console.error('Test failed: Data tidak sesuai atau tidak disimpan dengan benar');
        } else {
            console.log('Test passed: Data berhasil disimpan');
        }
    } catch (error) {
        console.error('Error saat mengirim request:', error.message);
    }
}










// Jalankan fungsi test
// testGetUsers();
// testDeleteUsers();
testPostUsers();