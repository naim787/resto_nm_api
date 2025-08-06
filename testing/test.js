const baseURL = "http://127.0.0.1:3000";

const data = {
    name: "Tono",
    id: "8888", // kalau pakai generateID di Go, field ini akan di-replace
    email: "naimmmmab@gmail.com",
    password: "155",
    bis_loc: "paniki",
    date_loc: "2025",
    year: "2004",
    role: "admin"
};

function logTitle(title) {
    console.log(`\n=== ${title.toUpperCase()} ===`);
}

async function testPostUsers() {
    logTitle('POST /create-users');
    try {
        const response = await fetch(`${baseURL}/create-users`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        });

        const result = await response.json();
        console.log('✅ Response:', result);

        if (!result.data || result.data.name !== data.name) {
            console.error('❌ Test failed: Data tidak sesuai');
        } else {
            console.log('✅ Test passed: User berhasil dibuat');
        }

        return result.data; // return untuk dapat ID-nya
    } catch (error) {
        console.error('❌ Error saat POST /create-users:', error.message);
    }
}

async function testGetUsers() {
    logTitle('GET /users');
    try {
        const response = await fetch(`${baseURL}/users`);
        const result = await response.json();

        console.log('✅ Response:', result);

        if (!Array.isArray(result.data)) {
            console.error('❌ Test failed: Response `data` bukan array');
        } else if (result.data.length === 0) {
            console.warn('⚠️ Test passed tapi kosong: Tidak ada data user');
        } else {
            console.log('✅ Test passed: Berhasil mengambil semua user');
        }

        return result.data;
    } catch (error) {
        console.error('❌ Error saat GET /users:', error.message);
    }
}

async function testDeleteUserById(userId) {
    logTitle(`DELETE /users/${userId}`);
    try {
        const response = await fetch(`${baseURL}/users/${userId}`, {
            method: 'DELETE'
        });

        const result = await response.json();
        console.log('✅ Response:', result);

        if (result.message !== 'User deleted successfully') {
            console.error('❌ Test failed: User gagal dihapus');
        } else {
            console.log('✅ Test passed: User berhasil dihapus');
        }
    } catch (error) {
        console.error(`❌ Error saat DELETE /users/${userId}:`, error.message);
    }
}

(async() => {
    const createdUser = await testPostUsers();

    // Optional: Fetch to confirm
    await testGetUsers();

    if (createdUser && createdUser.id) {
        await testDeleteUserById(createdUser.id);
    } else {
        console.warn("❗ Tidak bisa hapus user karena ID tidak ditemukan dari hasil POST.");
    }
})();