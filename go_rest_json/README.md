# API and static files

Combine Go-Json-Rest with other handlers.

`api.MakeHandler()` is a valid `http.Handler`, and can be combined with other handlers.
In this example the api handler is used under the `/api/` prefix, while a FileServer is instantiated under the `/static/` prefix.

curl demo:
```
curl -i http://127.0.0.1:8080/api/message
curl -i http://127.0.0.1:8080/static/main.go
```

## GORM

Demonstrate basic CRUD operation using a store based on MySQL and GORM
GORM is simple ORM library for Go. In this example the same struct is used both as the GORM model and as the JSON model.
curl demo:

```
curl -i -H 'Content-Type: application/json' \
    -d '{"Message":"this is a test"}' http://127.0.0.1:8080/reminders
curl -i http://127.0.0.1:8080/reminders/1
curl -i http://127.0.0.1:8080/reminders
curl -i -X PUT -H 'Content-Type: application/json' \
    -d '{"Message":"is updated"}' http://127.0.0.1:8080/reminders/1
curl -i -X DELETE http://127.0.0.1:8080/reminders/1
```