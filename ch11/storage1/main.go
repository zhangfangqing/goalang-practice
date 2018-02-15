package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

func bytesInUse(username string) int64 { return 0 }

const sender = "xxx@xxx.com"
const password = "xxxx"
const hostname = "smtp.xxx.com"

const template = `Warning: you are using %d bytes of storage, %d%% of your quota`

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 100000000
	percent := 100 * used / quota
	if percent < 90 {
		return
	}
	msg := fmt.Sprintf(template, used, percent)
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}
