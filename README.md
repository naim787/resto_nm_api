/project-root
├── /cmd              <- entry point aplikasi (main.go)
│   └── main.go
├── /internal         <- logic utama, gak boleh dipakai luar modul
│   └── /handler      <- fungsi-fungsi HTTP handler (CRUD logic)
│   └── /service      <- bisnis logic (kalau mau dipisah dari handler)
│   └── /repository   <- data access layer (ke DB)
│   └── /model        <- definisi struct / model
├── /pkg              <- kode yang reusable
├── go.mod