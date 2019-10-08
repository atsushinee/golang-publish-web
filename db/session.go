package db

import (
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/atsushinee/golang-publish-web/utils"
	"github.com/jeanphorn/log4go"
	"time"
)

func InsertSession(uid int64) *models.Session {
	session := &models.Session{
		SessionId: utils.NewUUID(),
		Uid:       uid,
		TTL:       utils.T2S(time.Now().Add(24 * time.Hour * 3)),
	}
	_, err := x.InsertOne(session)
	if err != nil {
		log4go.Error("InsertSession error:", err)
		panic(err)
	}
	return session
}

func GetAllSession() []*models.Session {
	var sessions []*models.Session
	session := new(models.Session)
	rows, err := x.Rows(session)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		s := new(models.Session)
		err = rows.Scan(s)
		if err != nil {
			panic(err)
		}
		sessions = append(sessions, s)
	}
	return sessions
}

func DeleteSession(sessionId string) {
	session := new(models.Session)
	session.SessionId = sessionId
	_, err := x.Delete(session)
	if err != nil {
		log4go.Error("DeleteSession error:", err)
		panic(err)
	}
}
