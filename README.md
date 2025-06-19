# 📘 API GraphQL Movie App (Pustaka Film)

API GraphQL Movie App adalah backend API yang dikembangkan menggunakan bahasa Go (Golang) dan framework gqlgen. Aplikasi ini dirancang untuk menyajikan informasi seputar film, termasuk film yang akan tayang (Coming Soon), film yang sedang tayang (Now Showing), serta rekomendasi film. Informasi yang ditampilkan mencakup judul, genre, rating, durasi, sinopsis, daftar aktor, dan tautan ke platform streaming resmi yang menayangkan film tersebut. Selain itu, aplikasi ini juga dilengkapi dengan fitur autentikasi berbasis JWT, dan forum untuk diskusi.
---

### 🎬Fitur Utama	
- Autentikasi menggunakan JWT (JSON Web Token)
- Registrasi dan login user
- Pencarian film berdasarkan judul
- Penayangan data film seperti judul, genre, rating, durasi, synopsis, daftar actor, dan status tayang (Coming Soon / Now Showing)
- Rekomendasi film berdasarkan rating dan jumlah views
- Forum diskusi antar user
- Integrasi tautan ekternal (platform streaming resmi) seperti, Netflix, XXI, dll

### ⚙️ Instalasi
Pastikan kamu sudah menginstal Go dan MySQL.

1. Clone repository ini:
```bash
git clone https://github.com/WahyuWardaya/movie_graphql.git
cd movie_graphql
```

2. Ubah konfigurasi database di config/config.go sesuai dengan kredensial lokalmu.
3. Install dependency:
```bash
go mod tidy
```
4. Generate file GraphQL jika belum ada:

```bash
go run github.com/99designs/gqlgen generate atau gqlgen generate (agar mengenerate data secara global)
```

###🚀 Menjalankan Server
```bash
go run server.go
```

Secara default, server akan berjalan di http://localhost:8080/query.

### 📌 Catatan Tambahan
1. JWT Token harus dikirim di header:
Authorization: Bearer <your-token>
2. GQL playground juga bisa digunakan secara lokal untuk testing

---
### Nama Kelompok

1. Putu Wahyu Wardaya (42230016)
2. I Gusti Ngurah Abi Praja Andika (42230045)
3. I Gede Agus Perdana Putra (42230058)
4. Ni Wayan Ariningsih (42230060)

