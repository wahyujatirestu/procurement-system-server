# Simple Procurement System – Backend API

## Overview

**Simple Procurement System** adalah RESTful Backend API yang dibangun untuk kebutuhan _technical test_ dan simulasi sistem procurement perusahaan. Sistem ini dirancang menggunakan arsitektur **MVC + Service + Repository**, mendukung **database transaction (ACID)**, **JWT Authentication**, dan **Swagger OpenAPI Documentation**.

Project ini berfokus pada:

- Master Data Management (Auth/Users, Items, Suppliers)
- Purchasing Transaction (Header & Detail)
- Atomic database transaction
- Clean & scalable backend structure

---

## Tech Stack

- **Language**: Go (Golang)
- **Framework**: Fiber
- **ORM**: GORM
- **Database**: PostgreSQL / MySQL
- **Authentication**: JWT
- **Password Hashing**: bcrypt
- **API Documentation**: Swagger (OpenAPI 3)
- **HTTP Client**: Resty (Webhook Integration)
- **Architecture**: MVC + Service + Repository

---

## Project Structure

```
├── config/
├── controllers/
├── databases/
├── docs/
├── dto/
├── middleware/
├── models/
├── repositories/
├── routes/
├── security/
├── services/
├── utils/
├── webhook/
├── .example.env
├── .gitignore
├── go.mod
├── go.sum
├── main.go
├── README.md
└── server.go
```

---

## Authentication & Authorization

- Menggunakan **JWT Access Token**
- Token dikirim melalui header:

```
Authorization: Bearer <token>
```

- Middleware akan:
    1. Validasi token
    2. Ambil userId & role dari token
    3. Inject user ke context
    4. Validasi role (jika dibutuhkan)

---

## Password Hashing

Password disimpan menggunakan **bcrypt**.

- Hash saat register / create user
- Verify saat login

Implementasi berada di:

```
/security/password.go
```

---

## API Response Standard

Semua response API **konsisten** dengan format:

```json
{
    "code": 200,
    "message": "success",
    "data": {}
}
```

Contoh Login Success:

```json
{
    "code": 200,
    "message": "login success",
    "data": {
        "access_token": "jwt-token",
        "user": {
            "id": 1,
            "username": "admin",
            "role": "ADMIN"
        }
    }
}
```

---

## Master Data

### Items

- Create Item
- Get List Items
- Get Item By ID
- Update Item
- (Soft) Delete Item

### Suppliers

- Create Supplier
- Get List Suppliers
- Get Supplier By ID
- Update Supplier
- (Soft) Delete Supplier

    _Catatan: Soft delete menggunakan `deleted_at` (GORM soft delete)._

---

## Purchasing Module

Purchasing terdiri dari:

- **Purchasing (Header)**
- **Purchasing Detail (Items yang dibeli)**

### Business Rules

- Purchasing **TIDAK BISA di-update atau di-delete**
- Purchasing bersifat **immutable transaction**
- Data yang salah harus diperbaiki dengan transaksi baru

---

## Database Transaction (ACID)

Saat proses **Create Purchasing**, sistem melakukan **3 proses secara atomic**:

### 1. Insert Purchasing Header

```go
purchasing := models.Purchasing{
    Date: time.Now(),
    SupplierID: supplier.ID,
    UserID: user.ID,
}
```

### 2. Insert Purchasing Detail

```go
detail := models.PurchasingDetail{
    PurchasingID: purchasing.ID,
    ItemID: item.ID,
    Qty: reqItem.Qty,
    SubTotal: subTotal,
}
```

### 3. Update Stock Item

```go
item.Stock += reqItem.Qty
```

**Karena ini Purchasing (pembelian ke supplier), stock item BERTAMBAH.**

Jika salah satu proses gagal, **seluruh transaksi akan di-rollback**.

---

## Purchasing Get List (Join Lengkap)

Endpoint `GET /purchasings` akan menampilkan:

- Purchasing Header
- Supplier (ID & Name)
- User (ID & Username)
- Purchasing Details
- Item Name & Price

Contoh response:

```json
{
    "id": 1,
    "date": "2026-01-21T10:00:00Z",
    "grand_total": 500000,
    "supplier": {
        "id": 1,
        "name": "PT Supplier Jaya"
    },
    "user": {
        "id": 1,
        "username": "admin"
    },
    "details": [
        {
            "item_id": 1,
            "item_name": "Laptop",
            "qty": 5,
            "price": 100000,
            "sub_total": 500000
        }
    ]
}
```

---

## Webhook Integration

Setelah transaksi purchasing berhasil:

- Sistem mengirim payload ke webhook service
- Proses webhook dijalankan **async (goroutine)**
- Tidak mempengaruhi transaksi utama
- Implementasi HTTP client webhook menggunakan Resty (github.com/go-resty/resty/v2)
- Request dikirim dalam format JSON (HTTP POST) ke URL webhook yang dapat dikonfigurasi

    _Penggunaan Resty dipilih karena ringan, clean API, dan mendukung retry, timeout, serta logging dengan mudah untuk kebutuhan integrasi eksternal._

---

## Swagger Documentation

Swagger telah diintegrasikan untuk dokumentasi REST API.

### Install Swagger CLI (Jika Belum Terinstall)

Sebelum menjalankan atau generate Swagger, pastikan **swag CLI** sudah terinstall di mesin lokal.

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Pastikan `$GOPATH/bin` atau `$HOME/go/bin` sudah masuk ke dalam `PATH`.

Cek instalasi:

```bash
swag --version
```

### Generate Swagger Documentation

Setelah install, jalankan perintah berikut di root project:

```bash
swag init
```

Perintah ini akan:

- Generate folder `docs/`
- Generate file Swagger berdasarkan annotation di controller

### Akses Swagger UI

```
http://localhost:8888/swagger/index.html
```

> **Important Note**
>
> Aplikasi secara default berjalan di **port 8888**.
> Jika reviewer menjalankan di port lain, silakan menyesuaikan base URL Swagger.

Swagger mencakup:

- Auth API
- Item API
- Supplier API
- Purchasing API

---

## Running Application

### 1. Clone Repository

```bash
git clone https://github.com/wahyujatirestu/procurement-system-server.git
cd simple-procurement-system
```

### 2. Setup Environment

```env
DB_HOST=localhost
DB_PORT=your_dbport
DB_USERNAME=your_username
DB_PASSWORD=your_password
DB_NAME=your_dbname
API_PORT=8888
ACCESS_TOKEN=your_token_secret
JWT_APP_NAME=your_app_name
WEBHOOK_URL=https://webhook.site/{your_webhookid}
CORS_ALLOW_ORIGINS=your_client_url
```

### 3. Run Application

```bash
go run .
```

---

## Author

**Restu Adi Wahyujati**
Junior Backend / Fullstack Developer (Golang)

---
