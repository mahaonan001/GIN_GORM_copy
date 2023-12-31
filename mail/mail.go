package mail

import (
	"GIN_GORM/common"
	"GIN_GORM/model"
	"GIN_GORM/response"
	"GIN_GORM/util"
	"log"
	"net/http"
	"net/smtp"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func SendMail(Toemail string, c *gin.Context) {
	DB := common.GetDB_Email()
	smtpHost := "smtp.qq.com"                                 // SMTP服务器地址
	smtpPort := "587"                                         // SMTP服务器端口
	smtpUser := "example@qq.com"                              // SMTP用户名
	smtpPassword := viper.GetString("emailCode.smtpPassword") // SMTP密码（授权码）
	toUserEmail := Toemail                                    // 接收者邮箱地址
	code := util.RandomString(6)                              // 验证码
	infTime := time.Now().Add(1 * time.Minute)
	flag := IsEmailAccept(toUserEmail, DB)
	//验证码可用
	if flag {
		response.Response(c, http.StatusBadRequest, 402, gin.H{"msg": "err"}, "请1分钟后重试")
		return
	}
	//验证码不可用
	e := email.NewEmail()
	e.From = smtpUser                                                                                               // 发件人邮箱账号
	e.To = append(e.To, toUserEmail)                                                                                // 收件人邮箱地址                                                                               // 收件人邮箱地址
	e.Subject = "马浩楠给您的亲爱的验证码"                                                                                      // 邮件主题
	e.Text = []byte("您的验证码是：" + code)                                                                               // 邮件正文内容（纯文本）
	e.HTML = []byte("<p>欢迎使用马浩楠的验证码测试</p></br><p>您的验证码是：</p></br><strong>" + code + "</strong></br><p>有效期限5分钟</p>") // 邮件正文内容（HTML格式）
	newCode := model.EmailCode{
		Email:      toUserEmail,
		Code_email: code,
		InfTime:    infTime,
	}
	err := e.Send(smtpHost+":"+smtpPort, smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)) // 发送邮件
	if err == nil {
		DB.Create(&newCode)
		response.SuccessRe(c, "成功获取验证码", nil)
		return
	}

	// 处理发送邮件失败的情况
	// 比如打印日志或返回错误信息
	log.Println("邮件发送失败", err)
}

// 判断邮箱是否合法或在一分钟内是否申请过验证码
func IsEmailAccept(email string, DB *gorm.DB) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	// 使用MatchString()函数来判断电子邮件地址是否匹配正则表达式
	if emailRegex.MatchString(email) {
		return TimeInner(DB, email)
	}
	return false
}

// 主函数
func Code_email(c *gin.Context) {
	//获取前端传来的请求中的email
	User_email := c.PostForm("email")
	//获取验证码和验证码截至时间
	SendMail(User_email, c)
}
func TimeInner(DB *gorm.DB, email string) bool { //存在一分钟内可用的验证码
	var EmailCode model.EmailCode
	TimeNow := time.Now()
	if err := DB.Model(&model.EmailCode{}).Error; err != nil {
		log.Println("fail to get table count")
		return false
	}

	result := DB.Where("email=? and inf_time >= ?", email, TimeNow).First(&EmailCode)
	if EmailCode.ID != 0 {
		return true //第一次获取验证码
	} else if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Println(result.Error.Error())
		panic("查询出错")
	} else {
		// 查询到结果
		log.Println("存在可用验证码")
		return false
	}
}
