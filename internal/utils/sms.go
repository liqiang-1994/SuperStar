package utils

import (
	"SuperStar/internal/config"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
	"math/rand"
	"strconv"
	"time"
)

func NewSms(c *sms.Client) *SmsClient {
	return &SmsClient{client: c}
}

type SmsClient struct {
	client *sms.Client
}

func NewSmsClient(cfg *config.Config) *sms.Client {
	credential := common.NewCredential(cfg.Sms.SecretId, cfg.Sms.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, err := sms.NewClient(credential, "ap-beijing", cpf)
	if err != nil {

	}
	return client
}

func SendSms(client *sms.Client, cfg config.Config, qrCode string, phoneNumber string, userInfo string) {
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppid = common.StringPtr(cfg.Sms.AppId)
	request.Sign = common.StringPtr(cfg.Sms.Sign)
	request.TemplateID = common.StringPtr(cfg.Sms.TemplateId)
	request.TemplateParamSet = common.StringPtrs([]string{qrCode})
	request.PhoneNumberSet = common.StringPtrs([]string{phoneNumber})
	request.SessionContext = common.StringPtr(userInfo)
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		fmt.Printf("An API error has returne: %s", err)
	}
	b, _ := json.Marshal(response.Response)
	// 打印返回的json字符串
	fmt.Printf("%s", b)
}

func Code() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(899999) + 100000
	res := strconv.Itoa(code) //转字符串返回
	return res
}
