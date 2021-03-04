# API Spec
---

## Account
### Create Account
Request :
- Method : POST
- Endpoint : `/account`
- Authorization : Bearer Token
- Header : 
    - Content-Type : application/json
    - Accept : application/json
- Body :
```json
{
    "username" : "string|unique|required|min_length=5|max_length=50",
    "password" : "string|required|min_length=5"
}
```

Response :
```json
{
    "status" : "integer",
    "message" : "string",
    "data" : {
        "id" : "integer",
        "username" : "string",
        "password" : "string"
    }
}
```

### Get Account
Request :
- Method : GET
- Endpoint : `/account/{account_id}`
- Authorization : Bearer Token
- Header : 
    - Accept : application/json

Response :
```json
{
    "status" : "integer",
    "message" : "string",
    "data" : {
        "id" : "integer",
        "username" : "string",
        "password" : "string"
    }
}
```

### Update Account
Request :
- Method : PUT
- Endpoint : `/account/{account_id}`
- Authorization : Bearer Token
- Header : 
    - Content-Type : application/json
    - Accept : application/json
- Body :
```json
{
    "username" : "string|unique|required|min_length=5|max_length=50",
    "password" : "string|required|min_length=5"
}
```

Response :
```json
{
    "status" : "integer",
    "message" : "string",
    "data" : {
        "id" : "integer",
        "username" : "string",
        "password" : "string"
    }
}
```

### List Account
Request :
- Method : GET
- Endpoint : `/account`
- Authorization : Bearer Token
- Header : 
    - Accept : application/json

Response :
```json
{
    "status" : "integer",
    "message" : "string",
    "data" : [
        {
            "id" : "integer",
            "username" : "string",
            "password" : "string"
        }
    ]
}
```

### Delete Account
Request :
- Method : DELETE
- Endpoint : `/account/{account_id}`
- Authorization : Bearer Token
- Header : 
    - Accept : application/json

Response :
```json
{
    "status" : "integer",
    "message" : "string"
}
```