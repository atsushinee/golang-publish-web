package service

import (
	"github.com/atsushinee/golang-publish-web/db"
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/atsushinee/golang-publish-web/utils"
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, session *models.Session) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		MaxAge: -1,
	})
	RemoveSession(session.SessionId)
}

func AuthSession(w http.ResponseWriter, r *http.Request) (*models.Session, string, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil, "请先登录", false
	}
	value, ok := models.SessionMap.Load(cookie.Value)
	if !ok {
		return nil, "请先登录", false
	}
	s := value.(*models.Session)
	if time.Now().Sub(utils.S2T(s.TTL)) > 0 {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
		RemoveSession(s.SessionId)
		return nil, "您的登录已超时,请重新登录", false
	}
	return s, "", true
}

func RemoveSession(sessionId string) {
	db.DeleteSession(sessionId)
	models.SessionMap.Delete(sessionId)
}

func LoadSessions() {
	sessions := db.GetAllSession()
	for _, session := range sessions {
		models.SessionMap.Store(session.SessionId, session)
	}
}
