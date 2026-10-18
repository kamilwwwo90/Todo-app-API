Żeby sklonować:

```bash
git clone https://github.com/kamilwwwo90/Todo-app-API
```

Następnie przejdź do katalogu projektu.

```bash
cd Todo-app-API
```

Konfiguracja bazy danych:

.env.example musi być nazwany .env, docker-compose odczyta z niego DB_PASSWORD do stworzenia kontenera.

Żeby stworzyć dwie bazy danych:

Chociaż tylko jedna jest tak naprawdę używana podczas testów.

```bash
  docker compose up -d
```

Po wszystkim:

```bash
go run ./cmd/server
```

Mogą pojawić się ostrzeżenia od rozszerzenia go w vs code jeżeli moduł nie został zainicjializowany.

Można wypróbować te API w server_test.go .
