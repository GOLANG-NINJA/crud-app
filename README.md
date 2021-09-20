# CRUD Приложение для Управления Списком Книг
## Пример решения тестового задания для практического проекта #1 курса GOLANG NINJA

### Стэк
- go 1.17
- postgres 

### Запуск
```go build -o app cmd/main.go && ./app```

Для postgres можно использовать Docker

```docker run -d --name ninja-db -e POSTGRES_PASSWORD=qwerty123 -v ${HOME}/pgdata/:/var/lib/postgresql/data -p 5432:5432 postgres```