# grpc-todo-list
### /v1/users/register POST
### /v1/users/login POST
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

### /v1/tasks POST - CREATE TASK
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
### /v1/tasks GET - LIST TASKS
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
### /v1/tasks/completed POST - COMPLETE TASK
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
### /v1/tasks/{id} DELETE - DELETE TASK
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