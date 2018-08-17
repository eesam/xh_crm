package crm

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const designColorPreUrl = "/upload/designColor/"
const designColorUploadPath = "/root/xh_crm/web/upload/designColor/"

type httpDesignColor struct {
}

func (p *httpDesignColor) onRequest(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/crm/designColor/getList") {
		p.onGetList(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/designColor/add") {
		p.onAdd(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/designColor/edit") {
		p.onEdit(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/designColor/delete") {
		p.onDelete(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/designColor/uploadPic") {
		p.onUploadPic(w, r)
	}
}

func (p *httpDesignColor) onGetList(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	offset, err := strconv.Atoi(vars.Get("offset"))
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	size, err := strconv.Atoi(vars.Get("size"))
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	designId := vars.Get("designId")

	lst, totalCount, err := instance.db.databaseDesignColor.getList(offset, size, designId)
	if err != nil {
		response(w, -1, err.Error())
		return
	}

	var designColorList DesignColorList
	designColorList.DesignColorInfos = lst
	designColorList.TotalCount = totalCount
	data, err := json.Marshal(designColorList)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, string(data))
}

func (p *httpDesignColor) onUploadPic(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		response(w, -1, "Method not post")
		return
	}
	file, head, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		response(w, -1, err.Error())
		return
	}
	defer file.Close()

	// 创建文件
	filename := uuid.Must(uuid.NewV4()).String()
	filename = strings.Replace(filename, "-", "", -1)
	filename += filepath.Ext(head.Filename)
	log.Println("onUploadPic", head.Filename, "==>", filename)
	os.MkdirAll(designColorUploadPath, os.ModePerm)
	fW, err := os.Create(designColorUploadPath + filename)
	if err != nil {
		log.Println("文件创建失败")
		response(w, -1, err.Error())
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		log.Println("文件保存失败")
		response(w, -1, err.Error())
		return
	}

	instance.designColorPicCache.add(designColorUploadPath + filename)
	response(w, 0, filename)
}

func (p *httpDesignColor) onAdd(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var info DesignColorInfo
	info.DesignId = vars.Get("designId")
	info.ColorId = vars.Get("colorId")
	info.PicUrl = designColorPreUrl + vars.Get("picUrl")
	info.Note = vars.Get("note")
	log.Println(info)

	picUrl := designColorUploadPath + vars.Get("picUrl")
	err := instance.designColorPicCache.remove(picUrl)
	if err != nil {
		response(w, -1, err.Error())
		return
	}

	err = instance.db.databaseDesignColor.insert(info)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}

func (p *httpDesignColor) onEdit(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var info DesignColorInfo
	info.DesignId = vars.Get("designId")
	info.ColorId = vars.Get("colorId")
	info.PicUrl = designColorPreUrl + vars.Get("picUrl")
	info.Note = vars.Get("note")
	log.Println(info)

	picUrl := designColorUploadPath + vars.Get("picUrl")
	err := instance.designColorPicCache.remove(picUrl)
	if err != nil {
		response(w, -1, err.Error())
		return
	}

	err = instance.db.databaseDesignColor.update(info)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}

func (p *httpDesignColor) onDelete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	designId := vars.Get("designId")
	colorId := vars.Get("colorId")
	log.Println(designId, colorId)
	err := instance.db.databaseDesignColor.delete(designId, colorId)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}
