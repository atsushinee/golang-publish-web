package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/satori/go.uuid"
	"os"
	"strings"
	"time"
)

func NewUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}

func B64UrlEncode(m string) string {
	return base64.URLEncoding.EncodeToString([]byte(m))
}

func B64UrlDecode(m string) string {
	s, err := base64.URLEncoding.DecodeString(m)
	if err != nil {
		return ""
	}
	return string(s)
}

const defaultTimeLayout = "2006-01-02 15:04:05"

func NowTimeString() string {
	return time.Now().Format(defaultTimeLayout)
}

func T2S(t time.Time) string {
	return t.Format(defaultTimeLayout)
}

func S2T(s string) time.Time {
	t, _ := time.ParseInLocation(defaultTimeLayout, s, time.Local)
	return t
}

func FileSize2MB(size int64) string {
	return fmt.Sprintf("%.2fMB", float64(size)/1024000)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func GetBasePath(path string) string {
	sep := "/"
	if !strings.Contains(path, sep) {
		return path
	}
	return sep + strings.Split(path, sep)[1]
}

func FileNameSplit(fileName string) (string, string) {
	split := strings.Split(fileName, ".")
	return strings.Join(split[:len(split)-1], "."), strings.Join(split[len(split)-1:], "")
}

func MD5Bytes(str string) []byte {
	m := md5.New()
	m.Write([]byte(str))
	return m.Sum(nil)
}

const salt = "cb8bcc6a-56f0-49af-8af1-35330b5d3d65"

func MD5(str string) string {
	return hex.EncodeToString(MD5Bytes(str + salt))
}
