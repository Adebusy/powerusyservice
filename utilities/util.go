package utilities

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	//"google.golang.org/appengine"
	//pb "google.golang.org/appengine/v2/mail"

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

func CreateLog() {
	dt := time.Now()
	filename := "log-" + dt.Format("01-02-2006") + ".txt"
	_, Ferr := os.Stat(filename)
	if os.IsNotExist(Ferr) {
		ret, err := os.Create(filename)
		if err != nil {
			fmt.Println("Unable to create logfile for today" + filename)
		}
		fmt.Println(ret)
		defer ret.Close()
		log.SetOutput(ret)
	}
}

func LogError(err error) {
	errLog := &log.Logger{}
	errLog.Println(err)
}

func SendEmail(c *gin.Context, toAddress []string, authcode string, r *http.Request) bool {
	// ctx := appengine.NewContext(r)
	// msg := &pb.Message{
	// 	Sender:  "noreply@powerusy.com",
	// 	To:      toAddress,
	// 	Subject: "Powerusy authentication code",
	// 	Body:    "Kindly complete your registration with code " + authcode,
	// }
	// if err := pb.Send(ctx, msg); err != nil {
	// 	fmt.Println(err)
	// 	return false
	// 	//log.Errorf(c, "Alas, my user, the email failed to sendeth: %v", err)
	// }
	return true
}
