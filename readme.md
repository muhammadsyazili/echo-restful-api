# API Spec
---
## Login
Request :
- Method : POST
- Endpoint : `/login`
- Header : 
    - Content-Type : application/json
    - Accept : application/json
- Body :
```json
{
    "username" : "string|required",
    "password" : "string|required"
}
```

Response :
```json
{
    "status" : "integer",
    "message" : "string",
    "data" : {
        "token" : "string"
    }
}
```

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

## Student
### Create Student
Request :
- Method : POST
- Endpoint : `/student`
- Authorization : Bearer Token
- Header : 
    - Content-Type : application/json
    - Accept : application/json
- Body :
```json
{
    "nama" : "string|required|max_length=50",
    "nim" : "integer|required|length=14",
    "jurusan" : "string|required|max_length=50",
    "account_id" : "integer|required|max_length=11"
}
```

Response :
```json
{
    "status" : "integer",
    "message" : "string",
    "data" : {
        "id" : "integer",
        "nama" : "string",
        "nim" : "integer",
        "jurusan" : "string",
        "account_id" : "integer"
    }
}
```

### Get Student
Request :
- Method : GET
- Endpoint : `/student/{student_id}`
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
        "nama" : "string",
        "nim" : "integer",
        "jurusan" : "string",
        "account_id" : "integer"
    }
}
```

### Update Student
Request :
- Method : PUT
- Endpoint : `/student/{student_id}`
- Authorization : Bearer Token
- Header : 
    - Content-Type : application/json
    - Accept : application/json
- Body :
```json
{
    "nama" : "string|required|max_length=50",
    "nim" : "integer|required|length=14",
    "jurusan" : "string|required|max_length=50",
    "account_id" : "integer|required|max_length=11"
}
```

Response :
```json
{
    "status" : "integer",
    "message" : "string",
    "data" : {
        "id" : "integer",
        "nama" : "string",
        "nim" : "integer",
        "jurusan" : "string",
        "account_id" : "integer"
    }
}
```

### List Student
Request :
- Method : GET
- Endpoint : `/student`
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
            "nama" : "string",
            "nim" : "integer",
            "jurusan" : "string",
            "account_id" : "integer"
        }
    ]
}
```

### Delete Student
Request :
- Method : DELETE
- Endpoint : `/student/{student_id}`
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