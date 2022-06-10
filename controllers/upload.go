package controllers

import (
	"fmt"
	"shopadmin/services"
	"shopadmin/utils"
)

type UploadController struct {
	AdminBaseController
}

//上传单个图片
func (u *UploadController) UploadSinglePic() {
	userId := u.GetUserId()
	var attachService services.AttachmentService
	fileInfo, err := attachService.Upload(u.Ctx, "file", userId)
	if err != nil {
		u.ErrLog(err)
		fmt.Println(err)
		return
	}
	utils.HttpSuccess(fileInfo, u.Ctx)
}
