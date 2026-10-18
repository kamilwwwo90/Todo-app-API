Żeby sklonować:

```bash
git clone https://github.com/kamilwwwo90/Todo-app-API
```

potem należy przejść do folderu projektu

```bash
cd Todo-app-API
```

Set up bazy danych:

.env.example musi być nazwany .env, docker-compose weźnie z niego DB_PASSWORD do stworzenia kontenera

żeby stworzyć dwie bazy danych:

chociaż tylko jedna jest tak naprawde używana podczas testów.

```bash
  docker compose up -d
```

Po wszystkim:

```bash
go run ./cmd/server
```

go vul check może mieć problemy bez inicializacji modułu.

Można wypróbować te API w server_test.go 
