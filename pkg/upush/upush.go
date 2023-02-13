/************************************************************
 * Author:        jackey
 * Date:        2022/8/30
 * Description: 友盟推送
 * Version:    V1.0.0
 **********************************************************/

package upush

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"liteIm/pkg/common"
	"liteIm/pkg/config"
	"liteIm/pkg/utils"
	"strconv"
	"time"
)

//func main() {
//	Push_IOS("测试标题", "037ae4d4791d2eb31b7f9fba2fe8d44b1228bcbd19ea4dc6cf6fd1edfb5d8a4e", 0, "")
//}

// =========友盟的接口====固定值，一般不会变动=============
var hostUmengPush = "http://msg.umeng.com"
var postPath = "/api/send"

// ===========Android的APP Key和 秘钥，不同程序会不同=============
var appKeyAndroid = config.Config.Section(common.ConfigSectionUpush).Key("appKeyAndroid").String()
var masterSecreptAndroid = config.Config.Section(common.ConfigSectionUpush).Key("masterSecreptAndroid").String()

// ===========IOS的APP Key和 秘钥，不同程序会不同===========
var appKeyIOS = config.Config.Section(common.ConfigSectionUpush).Key("appKeyIOS").String()
var masterSecretIOS = config.Config.Section(common.ConfigSectionUpush).Key("masterSecretIOS").String()

// 推送是否是生成模式
var pushProductionMode = "false"

// 设置前端收到通知后点击通知跳转的页面（这个是前端给的）
var pushAndroidActityChat = "xxx"
var pushAndroidActivityWeb = "xxx"

// UmengAndroid ========================================
type UmengAndroid struct {
	Appkey    string `json:"appkey"`    // 必填项
	Timestamp string `json:"timestamp"` // 必填项
	Type      string `json:"type"`      // 必填项

	DeviceTokens   string          `json:"device_tokens"` // 选填,用于给特定设备的推送
	ProductionMode string          `json:"production_mode"`
	Payload        *PayloadAndroid `json:"payload"` // 必填项
	Description    string          `json:"description"`
}

type PayloadAndroid struct {
	DisplayType string            `json:"display_type"` // 必填项
	Body        *BodyAndroid      `json:"body"`         // 必填项
	Extral      map[string]string `json:"extra"`
}
type BodyAndroid struct {
	Ticker    string `json:"ticker"`     // 必填项
	Title     string `json:"title"`      // 必填项
	Text      string `json:"text"`       // 必填项
	AfterOpen string `json:"after_open"` // 必填项
	Activity  string `json:"activity"`   // 必填项
}

// UPushAndroid android推送
// deviceToken:设备的编号,如果设置deviceToken，则是单播;如果未设置则是全播
func UPushAndroid(title string, deviceToken string, extrasData map[string]string) {

	body := &BodyAndroid{}
	// 必填 通知栏提示文字
	body.Ticker = title
	// 必填 通知标题
	body.Title = "xxx"
	// 必填 通知文字描述
	body.Text = title
	// 打开Android端的Activity
	body.AfterOpen = "go_activity"

	payLoad := &PayloadAndroid{}
	payLoad.DisplayType = "notification"
	payLoad.Body = body
	/*
	   额外携带的信息
	*/
	payLoad.Extral = extrasData

	messageAndroid := UmengAndroid{}
	messageAndroid.Appkey = appKeyAndroid

	// 打开聊天
	body.Activity = pushAndroidActityChat
	if deviceToken == "" {
		// 全播
		messageAndroid.Type = "broadcast"
		// 打开webview
		body.Activity = pushAndroidActivityWeb
	} else {

		// 单播
		messageAndroid.Type = "unicast"
		messageAndroid.DeviceTokens = deviceToken

		// 打开聊天
		body.Activity = pushAndroidActityChat
	}

	timeInt64 := time.Now().Unix()
	timestamp := strconv.FormatInt(timeInt64, 10)
	messageAndroid.Timestamp = timestamp
	messageAndroid.ProductionMode = pushProductionMode
	messageAndroid.Payload = payLoad
	messageAndroid.Description = title

	postBody, _ := json.Marshal(messageAndroid)
	url := hostUmengPush + postPath

	// MD5加密
	sign := utils.MD5("POST" + url + string(postBody) + masterSecreptAndroid)
	url = url + "?sign=" + sign

	_, err := resty.New().R().SetBody(messageAndroid).Post(url)
	if err != nil {
		logrus.Error("UPushAndroid", err)
	}
}

//=============================================================================================
/**
IOS推送必须项:
appkey
"timestamp":"xx",       // 必填 时间戳，10位或者13位均可，时间戳有效期为10分钟
type       //broadcast
"alert": "xx"          // 必填
MasterSecret
"production_mode":"true/false" // 可选 正式/测试模式。测试模式下，只会将消息发给测试设备。
*/

type UmengIOS struct {
	Appkey          string      `json:"appkey"`    // 必填项
	Timestamp       string      `json:"timestamp"` // 必填项
	Type            string      `json:"type"`      // 必填项
	Production_mode string      `json:"production_mode"`
	Payload         *PayloadIOS `json:"payload"`       // 必填项
	Devicetokens    string      `json:"device_tokens"` // 选填项
	Description     string      `json:"description"`
}

type PayloadIOS struct {
	Aps *ApsIOS `json:"aps"`
	//Ptype int     `json:"ptype"` // 1000:咨询,1001:政策法规;1004:法制宣传;1013:新闻
	//Purl  string  `json:"purl"`
}

type ApsIOS struct {
	Alert            string `json:"alert"` // 必填项
	ContentAvailable string `json:"content-available"`
}

// UPushIOS
// IOS推送
// deviceToken:设备的编号,如果设置deviceToken，则是单播;如果未设置则是全播
// ptype,purl:额外参数，非必选值
// ptype:区分类型
// purl:要打开的超链接的类型
func UPushIOS(title string, deviceToken string, pType int, purl string) {

	aps := &ApsIOS{}
	aps.Alert = "xxx"
	aps.ContentAvailable = title

	payLoad := &PayloadIOS{}
	payLoad.Aps = aps
	//payLoad.Ptype = ptype
	//payLoad.Purl = purl

	messageIOS := UmengIOS{}
	messageIOS.Payload = payLoad
	messageIOS.Appkey = appKeyIOS
	timeInt64 := time.Now().Unix()
	timestamp := strconv.FormatInt(timeInt64, 10)
	messageIOS.Timestamp = timestamp

	// 通过判断是否设置 deviceToken，来区分 单播 和 全播
	if deviceToken == "" {
		// 全播
		messageIOS.Type = "broadcast"
	} else {
		// 单播
		messageIOS.Type = "unicast"
		messageIOS.Devicetokens = deviceToken
	}

	messageIOS.Production_mode = pushProductionMode
	messageIOS.Description = title

	postBody, _ := json.Marshal(messageIOS)
	url := hostUmengPush + postPath

	// MD5加密
	sign := utils.MD5("POST" + url + string(postBody) + masterSecretIOS)
	url = url + "?sign=" + sign

	_, err := resty.New().R().SetBody(messageIOS).Post(url)
	if err != nil {
		logrus.Error("UPushIOS", err)
	}
}
