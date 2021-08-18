package main

import (
	chi "gigphil/app"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	chi.Setup()
}
