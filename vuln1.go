// vuln.go
package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Hardcoded credentials (G101)
const dbPassword = "supersecret123"
const apiKey = "AKIAIOSFODNN7EXAMPLE"

// Weak cryptography (G401)
func hashData(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

// Command injection (G204)
func runCommand(userInput string) {
	cmd := exec.Command("sh", "-c", userInput)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

// SQL injection (G201)
func getUser(db *sql.DB, username string) {
	query := "SELECT * FROM users WHERE name = '" + username + "'"
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
}

// Insecure random (G404)
func generateToken() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000000)
}

// File permission issue (G302)
func writeFile() {
	data := []byte("sensitive data")
	err := ioutil.WriteFile("secret.txt", data, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

// Path traversal risk (G304)
func readFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// Insecure HTTP (G107)
func fetchURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

// Disabled TLS verification (G402)
func insecureTLS() {
	tr := &http.Transport{
		// #nosec G402 (intentionally insecure)
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	_, _ = client.Get("https://example.com")
}

// Hardcoded temp file (G303)
func tempFile() {
	f, err := os.Create("/tmp/mytempfile")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}

func main() {
	db, _ := sql.Open("sqlite3", "test.db")
	defer db.Close()

	fmt.Println("Hash:", hashData("password"))
	runCommand("ls -la")
	getUser(db, "admin' OR '1'='1")
	fmt.Println("Token:", generateToken())
	writeFile()
	readFile("../../etc/passwd")
	fetchURL("http://example.com")
	insecureTLS()
	tempFile()
}
