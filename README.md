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