package db

import (
	"fmt"
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/atsushinee/golang-publish-web/utils"
	"github.com/jeanphorn/log4go"
	"os"
)

func CreateApplication(productId, uid int64, name, url, versionName string, versionCode int64, desc, createTime string) bool {

	app := &models.Application{
		Name:        name,
		Uid:         uid,
		ProductId:   productId,
		VersionName: versionName,
		VersionCode: versionCode,
		Desc:        desc,
		Url:         url,
		CreateTime:  createTime,
	}
	_, err := x.InsertOne(app)
	if err != nil {
		log4go.Error(fmt.Sprintf("CreateItem error: %s", err))
		return false
	}

	return true
}

func GetApplicationsByProductId(productId int64) []*models.Application {
	var applications []*models.Application
	application := &models.Application{
		ProductId: productId,
	}
	rows, err := x.Desc("version_code").Rows(application)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		i := new(models.Application)
		err = rows.Scan(i)
		if err != nil {
			panic(err)
		}
		user := GetUserById(i.Uid)
		if user == nil {
			log4go.Error("user is not found,id: ", i.Uid)
		} else {
			i.Author = user.Name
		}

		info, _ := os.Stat(models.UploadDir + i.Url)
		if info != nil {
			i.FileName = info.Name()
			i.FileSize = utils.FileSize2MB(info.Size())
		}

		applications = append(applications, i)
	}
	return applications
}

func GetApplicationById(cid int64) *models.Application {
	application := &models.Application{}
	b, err := x.ID(cid).Get(application)
	if err != nil {
		log4go.Error(fmt.Sprintf("GetApplicationById error: %s", err))
		return nil
	}

	if !b {
		return nil
	}
	return application
}

func RecordApplicationDownloadLog(uid int64, cid int64) bool {

	log := &models.ApplicationDownloadLog{
		ApplicationId:  cid,
		DownloadUserId: uid,
		DownloadTime:   utils.NowTimeString(),
	}
	_, err := x.InsertOne(log)
	if err != nil {
		log4go.Error(fmt.Sprintf("RecordApplicationDownloadLog error: %s", err))
		return false
	}

	return true
}

func CountApplicationDownloadLog(applicationId int64) (int64, error) {
	i, err := x.Count(&models.ApplicationDownloadLog{ApplicationId: applicationId})
	return i, err
}

func GetDownloadLogByApplicationId(applicationId int64) ([]*models.ApplicationDownloadLog, error) {
	var logs []*models.ApplicationDownloadLog
	log := &models.ApplicationDownloadLog{
		ApplicationId: applicationId,
	}
	rows, err := x.Desc("download_time").Rows(log)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		i := new(models.ApplicationDownloadLog)
		err = rows.Scan(i)
		if err != nil {
			return nil, err
		}
		user := GetUserById(i.DownloadUserId)
		if user != nil {
			i.DownloadUserName = user.Name
		}
		logs = append(logs, i)
	}
	return logs, nil
}

//func RateItemById(cid int64) error {
//	_, err := x.Exec("update item set rate_count=rate_count+1 where id=?", cid)
//	if err != nil {
//		log4go.Error(fmt.Sprintf("RateItemById error: %s", err))
//	}
//	return err
//}
