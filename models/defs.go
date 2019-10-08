package models

import (
	"sync"
)

var SessionMap *sync.Map

func init() {
	SessionMap = &sync.Map{}
}

const UploadDir = "./Downloads"
const DocDir = "./Docs"
