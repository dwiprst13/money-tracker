# Money Tracker API

Aplikasi **Money Tracker** berbasis **Go + Gin + GORM (SQLite)** untuk mencatat pemasukan, pengeluaran, laporan, dan anggaran.  
Bisa dijalankan sebagai **backend API** untuk web maupun aplikasi mobile (Flutter, React Native, dsb).  

---

## Fitur Utama
- **Autentikasi User** (register & login sederhana).
- **Manajemen Transaksi** (pemasukan & pengeluaran).
- **Laporan Bulanan** (total pemasukan, pengeluaran, saldo, insight sederhana).
- **Anggaran** (atur budget bulanan per kategori).
- **Struktur Modular** (dipisah ke package: models, controllers, middleware, routes, utils).

---

## Struktur Folder
money-tracker/  
│  
├── cmd/  
│ └── main.go  
│  
├── config/  
│ └── database.go  
│  
├── models/  
│ ├── user.go  
│ ├── transaction.go  
│ └── budget.go  
│  
├── controllers/  
│ ├── auth_controller.go  
│ ├── transaction_controller.go  
│ ├── report_controller.go  
│ └── budget_controller.go  
│  
├── middleware/  
│ └── auth.go  
│  
├── routes/  
│ └── routes.go  
│  
└── utils/  
└── helper.go  


---

## Instalasi & Menjalankan

### 1. Clone Repository
```bash
git clone https://github.com/username/money-tracker.git
cd money-tracker
```

2. Install Dependencies

Pastikan sudah ada Go 1.20+
```bash
go mod tidy
```

3. Jalankan Server
```bash
go run ./cmd/main.go
```

Server akan berjalan di:

http://localhost:8080

## API Endpoints
### Auth

POST /register
```bash
{
  "email": "user@example.com",
  "password": "123456"
}
```

POST /login
```bash
{
  "email": "user@example.com",
  "password": "123456"
}
```

Response
```bash
{
  "message": "Login berhasil",
  "userID": 1
}
```
Catatan: Autentikasi masih sederhana (header X-User-ID).
Untuk produksi sebaiknya gunakan JWT atau OAuth.

### Transactions

POST /transactions

```bash
Header: X-User-ID: 1

{
  "amount": -50000,
  "category": "Makan",
  "note": "Makan siang",
  "receipt_url": "http://example.com/struk.jpg"
}
```

GET /transactions
```bash
Header: X-User-ID: 1
```
### Reports

GET /reports
```bash
Header: X-User-ID: 1
```
Response
```bash
{
  "total_income": 2000000,
  "total_expense": -1500000,
  "saldo": 500000,
  "insight": "Bulan ini pengeluaran makan naik 20% dibanding bulan lalu."
}
```
### Budgets

POST /budgets
```bash
Header: X-User-ID: 1

{
  "category": "Makan",
  "amount": 1000000
}
```

GET /budgets
```bash
Header: X-User-ID: 1
```

## Tech Stack

- Go
- Gin –> Web Framework
- GORM –> ORM
- SQLite –> Database

Rencana Pengembangan

 Ganti autentikasi ke JWT.
 Tambahkan OAuth (Google/Apple).
 Sinkronisasi real-time dengan WebSocket/Firebase.
 Dashboard grafik untuk laporan keuangan.

Lisensi

MIT License © 2025
