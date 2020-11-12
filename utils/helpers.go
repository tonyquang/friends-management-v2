package utils

import (
	"crypto/md5"

	"encoding/hex"
	"io/ioutil"
	"regexp"

	"gorm.io/gorm"
)

func ValidateEmail(mail string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(mail) < 3 && len(mail) > 254 {
		return false
	}
	return emailRegex.MatchString(mail)
}

func ExtractMentionEmail(text string) []string {
	re := regexp.MustCompile(`@([a-zA-Z0-9]+)@([a-zA-Z0-9\.]+)\.([a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	rs := []string{}

	for i := 0; i < len(match); i++ {
		rs = append(rs, (match[i][0])[1:])
	}
	return rs
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// LoadFixture will load and execute SQL queries from fixture file
func LoadFixture(tx *gorm.DB, fixturePath string, rollBackName string) error {
	if fixturePath != "" {
		query, err := ioutil.ReadFile(fixturePath)
		if err != nil {
			tx.RollbackTo(rollBackName)
			panic(err)
		}
		rs := tx.Raw(string(query))
		if rs.Error != nil {
			tx.RollbackTo(rollBackName)
			panic(rs.Error)
		}
	}
	return nil
}
