package Controller

import (
	"ginEssential/model"
	"ginEssential/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"path"
)

func Install(context *gin.Context){
	var cmd  *exec.Cmd
	var output []byte
	var err error
	absolutepath := context.PostForm("absolutepath")
	//判断一下是否为空
	if absolutepath == ""{
		response.Response(context,http.StatusUnprocessableEntity,422,nil,"输入不能为空")
		return
	}

	//执行shell命令
	cmd = exec.Command("adb", "install", absolutepath)  //name字段是总命令，args是命令中的各种参数
	output , err = cmd.Output()
	if err != nil{
		response.Response(context,http.StatusUnprocessableEntity,422,gin.H{
			"err":err,
			"output":output,
		},"命令执行错误")
		return
	}
	response.Success(context,nil,string(output))
}

func Download(context *gin.Context){
	url := context.PostForm("url")

	//判断url是否为空
	if url == ""{
		response.Response(context,http.StatusUnprocessableEntity,422,nil,"输入不能为空")
		return
	}

	fileNameWithSuffix := path.Base(url)  //获取文件名称带后缀
	fileType := path.Ext(fileNameWithSuffix)  //获取文件的后缀
	fileContentType := model.HttpContentType[fileType]    //获取文件类型对应的http ContentType类型
	if fileContentType == ""{
		context.JSON(http.StatusNotFound,gin.H{
			"code":404,
			"msg":"file http contentType not found",
		})
		return
	}

	context.Header("Content-Type",fileContentType)
	context.File(url)
}