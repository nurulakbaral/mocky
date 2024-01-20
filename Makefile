export DATABASE_URL=postgres://postgres:''@localhost:5432/postgres

dev:
	pnpm dlx nodemon -q --exec go run ./cmd/main.go --signal SIGTERM