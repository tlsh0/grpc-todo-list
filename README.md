# grpc-todo-list
### /v1/users/register POST
### /v1/users/login POST

### /v1/tasks POST - CREATE TASK
### /v1/tasks GET - LIST TASKS
### /v1/tasks/completed POST - COMPLETE TASK
### /v1/tasks/{id} DELETE - DELETE TASK

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
