# CATEGORY BOOK SPECH

## Creat Category Book

Enpoint : POST /v1/private/category

Request Headler :
- access-token : token

Request Body :

```json
{
 "Name":"novel"
}
`````

Response body (success)

```json
{
    "data":{
        "id": 1,
        "name": "novel"
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

Enpoint : PATCH /v1/private/category/:catId

Request Headler :
- access-token : token

Request Body :

```json
{
 "Name":"update novel"
}
`````

Response body (success)

```json
{
    "data":{
        "id": 1,
        "name": "update novel"
    }
}
```

Response body (failed):

```json
{
    "errors":"unautorized, ....."
}
````

## Get All Category

Enpoint : GET /v1/public/category


Response body (success)

```json
{
    "data":[
        {
        "id": 1,
        "name": "novel"
        },
        {
        "id":2,
        "name": "fantasi"
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

Enpoint : GET /v1/public/category/:catId


Response body (success)

```json
{
    "data":
        {
        "id": 1,
        "name": "novel",
        "books":[
            {
                "id": 1,
                "tittle": "lorem ipsum dolor sit",
                "excerpt": "lorem ipsum dolor sit"
            },
              {
                "id": 2,
                "tittle": "lorem ipsum dolor sit",
                "excerpt": "lorem ipsum dolor sit"
            }
        ]
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

Enpoint : DELETE /v1/private/category/:catId


Request Headler :
- access-token : token

Response body (success)

```json
{
    "data":"delete success"
}
```

Response body (failed):

```json
{
    "errors":"category book not found"
}
````