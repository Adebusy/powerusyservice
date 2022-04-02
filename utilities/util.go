package utilities

import (
	"crypto/rand"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {
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

func LogError(err error) {
	dt := time.Now()
	filename := "log-" + dt.Format("01-02-2006") + ".log"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "Today ", log.LstdFlags)
	logger.Println(err)
}
