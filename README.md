Żeby zclonować:

```bash
git clone https://github.com/kamilwwwo90/Todo-app-API
```

i

```bash
cd Todo-app-API
```

Set up bazy danych:

.env.example musi być nazwany .env (docker-compose bierze tylko DB_PASSWORD)

```bash
  docker compose up -d
```

Po wszystkim:

```bash
go run ./cmd/server
```

go vul check może mieć problemy bez inicializacji modułu.

Można wypróbować te API w server_test.go
