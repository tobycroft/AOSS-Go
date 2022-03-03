package AttachmentModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "ao_attachment"

func Api_find(md5 interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"md5": md5,
	}
	db.Where(where)
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_insert(token, name, path, mime, size, ext, md5 interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"token": token,
		"name":  name,
		"path":  path,
		"mime":  mime,
		"size":  size,
		"ext":   ext,
		"md5":   md5,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
