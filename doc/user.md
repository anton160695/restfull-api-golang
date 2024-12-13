# USER API SPECH

## Register User
Endpoint : POST /v1/public/user/register

Request Body :

```json
{
 "username": "johndoe",
 "password": "12345678",
 "Name":"john doe andreas"
}
`````

Response body (success)

```json
{
    "data":{
        "username": "johndoe",
        "name": "john doe andrea"
    }
}
```

Response body (failed):

```json
{
    "errors":"username must be uniq, ....."
}
````

## Login User

Endpoint: POST /v1/public/user/login

Request Body :
```json
{
    "username":"johndoe",
    "password":"12345678"
}
````
Response Body (success)
```json
{
    "data":{
       "username": "johndoe",
        "name": "john doe andrea",
        "token":"jwt token"
    }
}
````

Response Body (failed):

```json
{
    "errors": "username or password wrong"
}
````

## Get User

Enpoint : GET /v1/private/user/me

Request Headler :
- access-token : token

Response body (success)

```json
{
    "data":{
        "username": "johndoe",
        "name": "john doe andrea"
    }
}
```

Response body (failed):

```json
{
    "errors":"unauthorized"
}
````

## Update User

Enpoint : PATCH /v1/private/user/me

Request Headler :
- access-token : token

Request body :
```json
{
        "username": "doeandrea", // tidak wajib
        "name": "doe andrea" // tidak wajib
}
```

Response body (success)

```json
{
    "data":{
        "username": "doeandrea",
        "name": "doe andrea"
    }
}
```

Response body (failed):

```json
{
    "errors":"unauthorized"
}
````

## Logout User

Enpoint : DELETE /v1/private/user/me

Request Headler :
- access-token : token

Response body (success)

```json
{
    "data": "oke"
}
```

Response body (failed):

```json
{
    "errors":"unauthorized"
}
````