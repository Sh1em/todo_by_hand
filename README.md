
```
make build && make run
```

docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
and
migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up    

http://localhost:8000/swagger/index.html

```
make migrate
```
MinGW32-make  если используете Windows