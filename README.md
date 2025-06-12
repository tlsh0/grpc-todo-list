# grpc-todo-list

POST /v1/users/register
POST /v1/users/login
request:
```
{
    "username":"",
    "password":""
}
```
response:
```
{
    "token":""
}
```

POST   /v1/tasks        create task
request:
```
{
    "title":"",
    "description":"",
    "token":""
}
```
response:
```
{
    "task":{
        "id":"",
        "title":"",
        "description":"",
        "completed":""
    }
}
```

GET    /v1/tasks        list tasks
request:
```
{
    "token":""
}
```
response:
```
{
    "tasks":[

    ]
}
```

POST   /v1/tasks/complete
request:
```
{
    "id":"",
    "token":""
}
```
response:
```
{
    "task":{}
}
```

DELETE /v1/tasks/{id}
request:
```
{
    "id":"",
    "token":""
}
```
response:
```
{
    "message":""
}
```
