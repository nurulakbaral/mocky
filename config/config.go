package config

import "os"


var (
	DATABASE_URL = os.Getenv("DATABASE_URL")
	PORT = "3001"
)