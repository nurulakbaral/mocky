export DATABASE_URL=host=localhost user=postgres password="" dbname=mocky-blog port=5432 sslmode=disable TimeZone=Asia/Jakarta

dev:
	pnpm dlx nodemon -q --exec go run ./cmd/main.go --signal SIGTERM