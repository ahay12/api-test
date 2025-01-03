# How to start
# Step 1
Clone a repository

```bash 
https://github.com/ahay12/api-app.git
```
Buat lah ```.env``` dengan isi :
sebagai contoh :
```bash
JWT_SECRET=VERY_SECRET_KEY
```	

# Step 2
buka project dan jalankan perintah dibawah untuk menginstall dependency
```bash
go get
```

# Step 3
jalankan program dengan perintah :
```bash
go run main.go
```

# Penggunaan API
disini saya menggunakan postman untuk menguji coba API :
endpoint API adalah :

#### Cara mengambil data project :
#### Method GET
```bash 
localhost:4000/api/v1/project
```

response
```
{
    "success": true,
    "message": "Projects fetched successfully",
    "data": {
        "data": [
            {
                "ID": 1,
                "CreatedAt": "2024-12-20T18:38:14.893+07:00",
                "UpdatedAt": "2024-12-20T18:38:14.893+07:00",
                "DeletedAt": null,
                "title": "Project 1",
                "description": "ini desc 2",
                "goals": 1000000,
                "fund": 200000,
                "category": "Type A",
                "tag": "Best",
                "expired": "2025-01-02T22:04:05+07:00"
            },
        ],
        "meta": {
            "limit": 10,
            "page": 1,
            "total": 3,
            "totalPages": 1
        }
    },
    "error": null
}
```
sebelum menambah kan project harus mempunyai akun dan ```Role``` harus ```admin``` bisa menggubah manual langsung dari database, **karena tidak ada akun default admin**

cara untuk Sign Up/Register
#### Method POST

```bash 
localhost:4000/api/v1/signup
```

dan masukan payload JSON, sebagai contoh :
```bash
{

"username":  "amir",

"password"  :  "password123",

"name":  "amir hakim",

"email":  "user@mail.com",

"address":  "Jaktim"

}
```
#### Login
#### Method POST
```bash 
localhost:4000/api/v1/signin
```

response :
```
{
    "success": true,
    "message": "token",
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAbWFpbC5jb20iLCJleHAiOjE3MzQ3MDczMjEsInJvbGUiOiJhZG1pbiIsInVzZXJJRCI6MX0.0zKnySZzrbq-NlSERlcyI-4IT0E9nyIZbLTepbPB09c",
    "error": null
}
```


### POST
untuk menambah project harus mempunyai ```token```dan ```Role admin```
```bash
localhost:4000/api/v1/project
```

dan masukan token di ```Authorization``` dengan Auth Type ```Bearer Token```

payload JSON untuk menambahkan project :

```bash
{

"title"  :  "Tempat Makan",

"description":  "membuat tempat makan dengan dana yang dibutuhkan Rp.10.000.000",

"goals":  10000000,

"fund":  1000000,

"category":  "Food",

"tag":  "UMKM",

"expired":  "2025-01-02T15:04:05Z"

}
```

response :
```
{
    "success": true,
    "message": "Successfully created Project",
    "data": {
        "ID": 3,
        "CreatedAt": "2024-12-20T21:43:34.971+07:00",
        "UpdatedAt": "2024-12-20T21:43:34.971+07:00",
        "DeletedAt": null,
        "title": "Pendanaan Tempat Makan",
        "description": "membuat tempat makan dengan dana yang dibutuhkan Rp.10.000.000",
        "goals": 10000000,
        "fund": 1000000,
        "category": "Food",
        "tag": "UMKM",
        "expired": "2025-01-02T15:04:05Z"
    },
    "error": null
}
```

### Filter by title
untuk mencari by title
```bash
localhost:4000/api/v1/project?title={name}
```
contoh

```
localhost:4000/api/v1/project?title=Project 1
```

response

```
{
    "success": true,
    "message": "Projects fetched successfully",
    "data": {
        "data": [
            {
                "ID": 1,
                "CreatedAt": "2024-12-20T18:38:14.893+07:00",
                "UpdatedAt": "2024-12-20T18:38:14.893+07:00",
                "DeletedAt": null,
                "title": "Project 1",
                "description": "ini desc 2",
                "goals": 1000000,
                "fund": 200000,
                "category": "Type A",
                "tag": "Best",
                "expired": "2025-01-02T22:04:05+07:00"
            }
        ],
        "meta": {
            "limit": 10,
            "page": 1,
            "total": 1,
            "totalPages": 1
        }
    },
    "error": null
}
```

### Sort by fund
```bash
localhost:4000/api/v1/project?page=1&limit=5&sort=fund&order=desc
```

response
```
{
    "success": true,
    "message": "Projects fetched successfully",
    "data": {
        "data": [
            {
                "ID": 3,
                "CreatedAt": "2024-12-20T21:43:34.971+07:00",
                "UpdatedAt": "2024-12-20T21:43:34.971+07:00",
                "DeletedAt": null,
                "title": "Pendanaan Tempat Makan",
                "description": "membuat tempat makan dengan dana yang dibutuhkan Rp.10.000.000",
                "goals": 10000000,
                "fund": 1000000,
                "category": "Food",
                "tag": "UMKM",
                "expired": "2025-01-02T22:04:05+07:00"
            },
            {
                "ID": 1,
                "CreatedAt": "2024-12-20T18:38:14.893+07:00",
                "UpdatedAt": "2024-12-20T18:38:14.893+07:00",
                "DeletedAt": null,
                "title": "Project 1",
                "description": "ini desc 2",
                "goals": 1000000,
                "fund": 200000,
                "category": "Type A",
                "tag": "Best",
                "expired": "2025-01-02T22:04:05+07:00"
            },
            {
                "ID": 2,
                "CreatedAt": "2024-12-20T18:57:30.294+07:00",
                "UpdatedAt": "2024-12-20T18:57:30.294+07:00",
                "DeletedAt": null,
                "title": "Project 2",
                "description": "ini desc 2",
                "goals": 10000000,
                "fund": 200000,
                "category": "Type B",
                "tag": "Best",
                "expired": "2025-01-02T22:04:05+07:00"
            }
        ],
        "meta": {
            "limit": 5,
            "page": 1,
            "total": 3,
            "totalPages": 1
        }
    },
    "error": null
}
```

### Update Project
#### Method PUT
contoh
```bash
localhost:4000/api/v1/project/1
```

payload
```bash
{

"title"  :  "Laundry",

"description":  "membuat Laundry dengan dana yang dibutuhkan Rp.30.000.000",

"goals":  30000000,

"fund":  1000000,

"category":  "Retail",

"tag":  "UMKM",

"expired":  "2025-01-02T15:04:05Z"

}
```

response
```
{
    "success": true,
    "message": "Successfully update Project",
    "data": {
        "ID": 1,
        "CreatedAt": "2024-12-20T18:38:14.893+07:00",
        "UpdatedAt": "2024-12-20T22:37:11.484+07:00",
        "DeletedAt": null,
        "title": "Laundry",
        "description": "membuat Laundry dengan dana yang dibutuhkan Rp.30.000.000",
        "goals": 30000000,
        "fund": 1000000,
        "category": "Retail",
        "tag": "UMKM",
        "expired": "2025-01-02T15:04:05Z"
    },
    "error": null
}
```
setelah update
```
{
    "success": true,
    "message": "Projects fetched successfully",
    "data": {
        "data": [
            {
                "ID": 1,
                "CreatedAt": "2024-12-20T18:38:14.893+07:00",
                "UpdatedAt": "2024-12-20T22:37:11.484+07:00",
                "DeletedAt": null,
                "title": "Laundry",
                "description": "membuat Laundry dengan dana yang dibutuhkan Rp.30.000.000",
                "goals": 30000000,
                "fund": 1000000,
                "category": "Retail",
                "tag": "UMKM",
                "expired": "2025-01-02T22:04:05+07:00"
            }
        ],
        "meta": {
            "limit": 10,
            "page": 1,
            "total": 1,
            "totalPages": 1
        }
    },
    "error": null
}
```

### Delete Project
#### Method DELETE
```bash
localhost:4000/api/v1/project/3
```

response
```
{
    "success": true,
    "message": "Project deleted successfully",
    "data": null,
    "error": null
}
```

setelah delete
```
{
    "success": true,
    "message": "Projects fetched successfully",
    "data": {
        "data": [
            {
                "ID": 1,
                "CreatedAt": "2024-12-20T18:38:14.893+07:00",
                "UpdatedAt": "2024-12-20T22:37:11.484+07:00",
                "DeletedAt": null,
                "title": "Laundry",
                "description": "membuat Laundry dengan dana yang dibutuhkan Rp.30.000.000",
                "goals": 30000000,
                "fund": 1000000,
                "category": "Retail",
                "tag": "UMKM",
                "expired": "2025-01-02T22:04:05+07:00"
            },
            {
                "ID": 2,
                "CreatedAt": "2024-12-20T18:57:30.294+07:00",
                "UpdatedAt": "2024-12-20T18:57:30.294+07:00",
                "DeletedAt": null,
                "title": "Project 2",
                "description": "ini desc 2",
                "goals": 10000000,
                "fund": 200000,
                "category": "Type B",
                "tag": "Best",
                "expired": "2025-01-02T22:04:05+07:00"
            }
        ],
        "meta": {
            "limit": 10,
            "page": 1,
            "total": 2,
            "totalPages": 1
        }
    },
    "error": null
}
```