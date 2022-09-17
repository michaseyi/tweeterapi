package config

import "os"

var (
	DATABASE_URI  = os.Getenv("DATABASE_URI")
	DATABASE_NAME = os.Getenv("DATABASE_NAME")
)
