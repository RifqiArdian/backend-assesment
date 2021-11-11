#Analisis Masalah
- stock masih ada tapi kenyataannya sudah habis
Hal itu terjadi mungkin karena race condition dimana ada 2 request masuk secara bersamaan
untuk satu produk sehingga pengurangan stok hanya terjadi satu kali, untuk menyelesaikan masalah tersebut
bisa dengan menggunakan fitur mutex bawaan dari golang jadi ketika ada satu hit masuk maka stok dan balance akan 
kunci dulu sampe proses beres

- stock bisa minus
Hal itu terjadi karena tidak ada handle error jika request stock melebihi dari sisa stock produk tersebut, selain itu
bisa juga karena tidak menggunakan transactional database ketika proses transaksi, sehingga ketika ada error ketika transaksi
data tidak kembali lagi seperti semula atau di rollback
  
#Cara menggunakan aplikasi
- clone project dari github 
- buat file .env di folder paling luar untuk yang contohnya bisa di ambil dari .env.example
- buka terminal kemudian masuk ke directory project tersebut
- ketik perintah go get -u untuk mendownload library yang dibutuhkan
- ketik perintah go run main.go untuk menjalankan aplikasi
- bawaannya aplikasi berjalan di port 3000 atau bisi di ubah di .env


# API Specification

## Error status on admin auth service

- 401 : Unauthorized
- 400 : Bad Request


## Register

Request :
- Method : POST
- Endpoint : `/auth/register`
- Header :
  - Content-Type: application/json
- Body :
```json 
{
    "name" : "string",
    "email" : "string",
    "password" : "string",
    "address" : "string"
}
```
- Response :
```json 
{
    "code" : "int"
    "status" : "string"
    "data" : {
        "id": "string",
        "name": "string",
        "email": "string",
        "balance": int64,
        "token": "string",
        "address": "string",
        "created_at": int64,
        "updated_at": int64     
    }
}
```

## Login

Request :
- Method : POST
- Endpoint : `/auth/login`
- Header :
  - Content-Type: application/json
- Body :
```json 
{
    "email" : "string",
    "password" : "string",
}
```
- Response :
```json 
{
    "code" : "int"
    "status" : "string",
    "data" : {
        "id": "string",
        "name": "string",
        "email": "string",
        "balance": int64,
        "token": "string",
        "address": "string",
        "created_at": int64,
        "updated_at": int64     
    }
}
```

## Add Cart

Request :
- Method : POST
- Endpoint : `/cart`
- Header :
  - Content-Type: application/json
- Body :
```json 
{
    "product_id" : "string",
    "quantity" : int64,
}
```
- Response :
```json 
{
    "code" : "int"
    "status" : "string",
}
```

## Get Cart

Request :
- Method : GET
- Endpoint : `/cart`
- Header :
  - Content-Type: application/json
- Response :
```json 
{
    "code" : "int"
    "status" : "string",
    "data" : [
        {
            "id" : "string",
            "quantity" : int64,
            "product" : {
                "id" : "string",
                "name" : "string",
                "category" : "string",
                "image" : "string",
                "price" : int64,
                "stock" : int64,
                "created_at" : int64,
                "updated_at" : int64
            }
            "created_at" : int64,
            "updated_at" : int64 
        }    
    ]
}
```

## Edit Cart

Request :
- Method : PUT
- Endpoint : `/cart/:id`
- Header :
  - Content-Type: application/json
- Body :
```json 
{
    "quantity" : int64,
}
```
- Response :
```json 
{
    "code" : "int"
    "status" : "string",
}
```

## Delete Cart

Request :
- Method : DELETE
- Endpoint : `/cart/:id`
- Header :
  - Content-Type: application/json
- Response :
```json 
{
    "code" : "int"
    "status" : "string",
}
```

## Get Product

Request :
- Method : GET
- Endpoint : `/product`
- Header :
  - Content-Type: application/json
- Response :
```json 
{
    "code" : "int"
    "status" : "string",
    "data" : [
        {
            "id" : "string",
            "name" : "string",
            "category" : "string",
            "image" : "string",
            "price" : int64,
            "stock" : int64,
            "created_at" : int64,
            "updated_at" : int64
        },
        {
            "id" : "string",
            "name" : "string",
            "category" : "string",
            "image" : "string",
            "price" : int64,
            "stock" : int64,
            "created_at" : int64,
            "updated_at" : int64
        }
    ]
}
```

## Add Transaction

Request :
- Method : POST
- Endpoint : `/transaction`
- Header :
  - Content-Type: application/json
- Body :
```json 
{
    "product_id" : "string",
    "quantity" : int64,
    "address" : "string"
}
```
- Response :
```json 
{
    "code" : "int"
    "status" : "string",
}
```

## Get Transaction

Request :
- Method : GET
- Endpoint : `/transaction`
- Header :
  - Content-Type: application/json
- Response :
```json 
{
    "code" : "int"
    "status" : "string",
    "data" : [
        {
            "id": "string",
            "user_id": "string",
            "product_id": {
                "id": "string",
                "name": "string",
                "category": "string",
                "image": "string",
                "price": int64,
                "stock": int64,
                "created_at": int64,
                "updated_at": int64
            },
            "quantity": int64,
            "price": int64,
            "total_price": int64,
            "address": "string",
            "created_at": int64,
            "updated_at": int64
        }
    ]
}
```