package services

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"shopadmin/models"
	"shopadmin/utils"
	"strings"
	"time"
)

//文件大小 单位：字节
var validateSize int

//允许上传的文件类型
var extString string

//实际保存地址
var realTempPath string
var realSavePath string

//nginx 虚拟主机目录地址
var savePath string
var tempPath string

//图片服务器地址
var fileHost string

type AttachmentService struct {
	BaseService
}

func init() {
	validateSize, _ = beego.AppConfig.Int("attachment::validate_size")
	extString, _ = beego.AppConfig.String("attachment::validate_ext")
	tempPath, _ = beego.AppConfig.String("attachment::temp_path")
	savePath, _ = beego.AppConfig.String("attachment::save_path")

	realTempPath, _ = beego.AppConfig.String("attachment::real_temp_path")
	realSavePath, _ = beego.AppConfig.String("attachment::real_save_path")
	fileHost, _ = beego.AppConfig.String("attachment::file_host")
}

//临时目录的文件信息
type TempFileInfo struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	RealUrl string `json:"real_url"`
}

//上传单个文件
func (att *AttachmentService) Upload(ctx *context.Context, name string, userId int) (*TempFileInfo, error) {
	file, h, err := ctx.Request.FormFile(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if err := att.validateForAttachment(h); err != nil {
		return nil, err
	}
	fileExt := path.Ext(h.Filename)
	newName := utils.GenarateUUid()
	tempDir := realTempPath
	//如果文件夹不存在，则创建
	if _, err := os.Stat(tempDir); err != nil {
		os.MkdirAll(tempDir, 0777)
	}
	saveTemp := tempDir + newName + fileExt
	newFile, err := os.OpenFile(saveTemp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	defer newFile.Close()
	if _, err = io.Copy(newFile, file); err != nil {
		return nil, err
	}
	return &TempFileInfo{
		newName,
		fileHost + tempPath + newName + fileExt,
		saveTemp,
	}, nil
}

//判断文件是否合格
func (att *AttachmentService) validateForAttachment(header *multipart.FileHeader) error {
	size := int(header.Size)
	if size > validateSize {
		return errors.New("文件超过限制大小")
	}
	ext := strings.TrimLeft(path.Ext(header.Filename), ".")
	if !strings.Contains(extString, ext) {
		return errors.New("不支持该文件格式")
	}
	return nil
}

//将文件从缓存目录，存放到实际目录
//name 存放哪个模块的图片
func (att *AttachmentService) CopyFileToRealPath(name, tempFile string, userId int) (*models.Attachment, error) {
	file, err := os.Open(tempFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//文件实际保存的目录
	newSaveDir := filepath.Join(realSavePath, name, utils.GetNowDate())
	if _, err := os.Stat(newSaveDir); err != nil {
		os.MkdirAll(newSaveDir, 0777)
	}
	fileExt := path.Ext(tempFile)
	newFileName := utils.GenarateUUid()
	//实际保存的文件
	var saveFile string
	saveFile = newSaveDir + "/" + newFileName + fileExt
	newFile, err := os.OpenFile(saveFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	defer newFile.Close()
	_, err = io.Copy(newFile, file)
	if err != nil {
		return nil, err
	}
	//上传的文件信息保存到数据库
	newFileInfo, _ := newFile.Stat()
	//文件服务器地址
	url := fileHost + filepath.Join(savePath, name, utils.GetNowDate(), newFileName+fileExt)
	data, err := att.SaveAttachmentInfo(newFileInfo, url, saveFile, name, userId)
	if err != nil {
		return nil, err
	}
	return data, nil

}

//将上传的图片信息保存数据库
func (att *AttachmentService) SaveAttachmentInfo(file os.FileInfo, url, saveFile, mime string, userId int) (*models.Attachment, error) {
	var fileInfo = models.Attachment{
		AdminUserId: userId,
		SaveName:    file.Name(),
		SavePath:    saveFile,
		Url:         url,
		Extension:   path.Ext(file.Name()),
		Mime:        mime,
		Size:        file.Size(),
		Md5:         utils.GetMd5String(file.Name()),
		Sha1:        utils.GetSha1String(file.Name()),
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	o := orm.NewOrm()
	_, err := o.Insert(&fileInfo)
	if err != nil {
		return nil, err
	}
	return &fileInfo, nil
}
