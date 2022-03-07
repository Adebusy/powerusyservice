package utilities

import (
	"crypto/rand"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file" + err.Error())
	}

	return os.Getenv(key)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateAuthCode() string {
	max, _ := strconv.Atoi(GoDotEnvVariable("AuthCodeLength"))
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
