package db

import (
	"fmt"
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/atsushinee/golang-publish-web/utils"
	"github.com/jeanphorn/log4go"
)

func GetAllDocsByPid(projectId int64) ([]*models.Doc, error) {
	var docs []*models.Doc
	rows, err := x.Where("pid=?", projectId).Rows(new(models.Doc))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		i := new(models.Doc)
		err = rows.Scan(i)
		if err != nil {
			return nil, err
		}
		docs = append(docs, i)
	}
	return docs, nil
}

func GetDocById(id int64) *models.Doc {
	doc := new(models.Doc)
	b, err := x.ID(id).Get(doc)
	if err != nil {
		return nil
	}
	if !b {
		return nil
	}
	return doc
}

func RecordDocViewLog(uid, docId int64) error {
	log := &models.DocViewLog{
		DocId:      docId,
		ViewUserId: uid,
		ViewTime:   utils.NowTimeString(),
	}
	_, err := x.InsertOne(log)
	if err != nil {
		log4go.Error(fmt.Sprintf("RecordDocViewLog error: %s", err))
		return err
	}
	return err
}

func CountDocViewLog(docId int64) (int64, error) {
	i, err := x.Count(&models.DocViewLog{DocId: docId})
	return i, err
}

func GetDocViewLogByDocId(docId int64) ([]*models.DocViewLog, error) {
	var logs []*models.DocViewLog
	log := &models.DocViewLog{
		DocId: docId,
	}
	rows, err := x.Desc("view_time").Rows(log)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		i := new(models.DocViewLog)
		err = rows.Scan(i)
		if err != nil {
			return nil, err
		}
		user := GetUserById(i.ViewUserId)
		if user != nil {
			i.ViewUserName = user.Name
		}
		logs = append(logs, i)
	}
	return logs, nil
}

func CreateDoc(productId, uid int64, name, ext, url, createTime string) error {

	doc := &models.Doc{
		Pid:        productId,
		Uid:        uid,
		Name:       name,
		Type:       ext,
		Url:        url,
		CreateTime: createTime,
	}
	_, err := x.InsertOne(doc)
	if err != nil {
		log4go.Error(fmt.Sprintf("CreateDoc error: %s", err))
	}
	return err
}
