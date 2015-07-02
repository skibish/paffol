package helpers

import (
	"log"
	"math/rand"
	"os"
	"regexp"
	"time"
)

var compiledRegexp *regexp.Regexp

// Random function returns random number in range (0, max]
func Random(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// CleanString strips given string from useless characters
func CleanString(s string) string {
	return compiledRegexp.ReplaceAllString(s, "")
}

// GetPort variable
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}

func init() {
	compiledRegexp, _ = regexp.Compile("[^0-9а-яА-ЯеЁ!?]")
}
