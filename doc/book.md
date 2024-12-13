# BOOK API SPECH

## Creat Category Book

Enpoint : POST /v1/private/book

Request Headler :
- access-token : token

Request Body :

```json
{
 "tittle":"lorem ipsum dolor sit amet",
 "exercpt":"lorem ipsum dolor sit amet",
 "category_id": 1,
 "content":"Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s"
}
`````

Response body (success)

```json
{
    "data":{
        "id": 1,
        "tittle":"lorem ipsum dolor sit amet",
        "exercpt":"lorem ipsum dolor sit amet",
        "category": "novel"
    }
}
```

Response body (failed):

```json
{
    "errors":"unautorized, ....."
}
````

## Update Category Book

Enpoint : PATCH /v1/private/book/:bookId

Request Headler :
- access-token : token

Request Body :

```json
{
 {
 "tittle":"lorem ipsum dolor sit amet update", // tidak wajib
 "exercpt":"lorem ipsum dolor sit amet update", // tidak wajib
 "category_id": 2, // tidak wajib
 "creator":"john doe andreas", // tidak wajib
 "content":"Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s update" // tidak wajib
}
}
`````

Response body (success)

```json
{
    "data":{
        "id": 1,
       "tittle":"lorem ipsum dolor sit amet update",
       "exercpt":"lorem ipsum dolor sit amet update",
       "category": "fantasi"
    }
}
```

Response body (failed):

```json
{
    "errors":"unautorized, ....."
}
````

## Get All Book

Enpoint : GET /v1/public/books


Response body (success)

```json
{
    "data":[
        {
        "id": 1,
        "tittle":"lorem ipsum dolor sit amet",
        "exercpt":"lorem ipsum dolor sit amet",
        "category": "novel"
        },
        {
         "id": 2,
        "tittle":"lorem ipsum dolor sit amet",
        "exercpt":"lorem ipsum dolor sit amet",
        "category": "fantasi"
        }
        ]
}
```

Response body (failed):

```json
{
    "errors":"category book not found"
}
````

## Get Category Book By ID

Enpoint : GET /v1/public/book/:bookId


Response body (success)

```json
{
    "data":
        {
        "id": 1,
        "tittle":"lorem ipsum dolor sit amet",
        "exercpt":"lorem ipsum dolor sit amet",
        "category": "novel",
        "creator":"john doe andreas",
        "content":"Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s"
        },
        
        
}
```

Response body (failed):

```json
{
    "errors":"category book not found"
}
````

## Delete Category Book

Enpoint : DELETE /v1/private/book/:bookId


Request Headler :
- access-token : token

Response body (success)

```json
{
    "data":"delete book success"
}
```

Response body (failed):

```json
{
    "errors":"category book not found"
}
````