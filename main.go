package main

import (
	"github.com/cwseger/spotify-cli/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
