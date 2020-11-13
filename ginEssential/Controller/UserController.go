package Controller

import (
	"ginEssential/Util"
	"ginEssential/common"
	"ginEssential/dto"
	"ginEssential/model"
	"ginEssential/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Login(context *gin.Context)  {
	DB := common.GetDB()
	//获取参数
	telephone := context.PostForm("telephone")
	password := context.PostForm("password")

	//数据验证
	if len(telephone) != 11{
		response.Response(context,http.StatusUnprocessableEntity,422,nil,"手机格式不正确")
		//context.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"msg":"手机号格式不正确",
		//})
		return
	}
	if len(password) < 6{
		response.Response(context,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
		//context.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"msg":"密码不能少于6位",
		//})
		return
	}

	//判断手机号是否存在
	var user model.User
	DB.Where("telephone=?",telephone).First(&user)
	if user.ID == 0{
		response.Response(context,http.StatusUnprocessableEntity,422,nil,"该手机号未注册")
		//context.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"msg":"该手机号未注册",
		//})
		return
	}

	//判断密码是否正确，密码需要加密，不能明文保存
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password)); err !=nil {
		response.Response(context,http.StatusBadRequest,400,nil,"密码错误")
		//context.JSON(http.StatusBadRequest,gin.H{
		//	"code":400,
		//	"msg":"密码错误",
		//})
		return
	}

	//登录成功，发放token
	token , err := common.ReleaseToken(user)
	if err != nil{
		response.Response(context,http.StatusInternalServerError,500,nil,"系统异常")
		//context.JSON(500,gin.H{
		//	"code":500,
		//	"msg":"系统异常",
		//})
		return
	}
	response.Success(context,gin.H{"token":token},"登陆成功")
	//context.JSON(200,gin.H{
	//	"code":200,
	//	"data":gin.H{"token":token},
	//	"msg":"登录成功",
	//})
}

func Register(context *gin.Context) {
	DB := common.GetDB()
	//获取参数
	name := context.PostForm("name")
	telephone := context.PostForm("telephone")
	password := context.PostForm(("password"))
	//数据验证
	if len(telephone) !=  11 {
		response.Response(context,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
		//context.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"msg":"手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		response.Response(context,http.StatusUnprocessableEntity,422,nil,"密码不能短于六位")
		//context.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"msg":"密码不能短于六位"})
		return
	}

	log.Println(name,telephone,password)
	//如果名称没有传。生成十位随机字符串
	if len(name) == 0{
		name = Util.RandomString(10)
	}

	//判断手机号是否存在
	if isTelephoneExists(DB,telephone){
		response.Response(context,http.StatusUnprocessableEntity,422,nil,"手机号已注册")
		//context.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"msg":"手机号已注册"})
		return
	}

	//创建用户,密码不能明文保存，需要加密
	hasedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)      //一个字节一个字节处理的
	if err != nil{
		response.Response(context,http.StatusInternalServerError,500,nil,"加密错误")
		//context.JSON(http.StatusInternalServerError,gin.H{
		//	"code":500,
		//	"mgs":"加密错误",
		//})
	}
	newUser :=model.User{
		Name : name,
		Telephone: telephone,
		Password: string(hasedPassword),  //字节流转化为字符串
	}
	DB.Create(&newUser)
	//返回结果
	response.Success(context,nil,"注册成功")
	context.JSON(200,gin.H{
		"msg":"注册成功",
	})
}

func isTelephoneExists(db *gorm.DB,telephone string)bool{
	var user model.User
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID != 0{
		return true
	} else {
		return false
	}
}

func Info(context *gin.Context){        //直接从上下文中获取信息
	user , _ := context.Get("user")
	response.Success(context,gin.H{"user":dto.ToUserDto(user.(model.User))},"")
	//context.JSON(http.StatusOK,gin.H{
	//	"code":200,
	//	"data":gin.H{
	//		"user":dto.ToUserDto(user.(model.User)),   //go的断言，做了一次强制类型转化，将interface{}转化为user。
	//	},
	//})
}


