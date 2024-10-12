# pushapi

各手机厂商推送 api 接入

[![Coverage Status](https://coveralls.io/repos/github/liyouping/pushapi/badge.svg?branch=master)](https://coveralls.io/github/liyouping/pushapi?branch=master)
[![GoDoc](https://pkg.go.dev/badge/github.com/liyouping/pushapi)](https://pkg.go.dev/github.com/liyouping/pushapi)


vivo（更新日期：2024-03-21）：

*   <https://dev.vivo.com.cn/documentCenter/doc/362>

oppo （更新日期：2024-04-25） ：

*   <https://open.oppomobile.com/new/developmentDoc/info?id=11236>

小米 （更新日期：2024-04-25）：

*   <https://dev.mi.com/distribute/doc/details?pId=1559>

华为 （更新日期：2024-03-05）：

*   <https://developer.huawei.com/consumer/cn/doc/HMSCore-References/https-send-api-0000001050986197>

荣耀 （更新日期：2024-10-11）：
*   <https://developer.honor.com/cn/docs/11002/reference/downlink-message#%E8%8E%B7%E5%8F%96%E9%89%B4%E6%9D%83%E6%8E%A5%E5%8F%A3>
## 调用示例

### vivo

```go
package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/liyouping/pushapi/vivopush"
)

var appId = "your app id"
var appKey = "your app key"
var appSecret = "your app secret"
var regId = "your reg id"

func main() {
	client := vivopush.NewClient(appId, appKey, appSecret)

	sendReq := &vivopush.SendReq{
		RegId:          regId,
		NotifyType:     4,
		Title:          "test push title",
		Content:        "test push content",
		TimeToLive:     24 * 60 * 60,
		SkipType:       1,
		NetworkType:    -1,
		Classification: 1,
		RequestId:      strconv.Itoa(int(time.Now().UnixNano())),
	}
	sendRes, err := client.Send(sendReq)
	fmt.Println(sendRes, err)
}
```

### oppo

```go
package main

import (
	"fmt"

	"github.com/liyouping/pushapi/oppopush"
)

var appKey = "your app key"
var masterSecret = "your master secret"
var regId = "your reg id"
var channelId = "your channel id"

func main() {
	client := oppopush.NewClient(appKey, masterSecret)

	sendReq := &oppopush.SendReq{
		Notification: &oppopush.Notification{
			Title:     "test push title",
			Content:   "test push content",
			ChannelID: channelId,
		},
		TargetType:  2,
		TargetValue: regId,
	}
	sendRes, err := client.Send(sendReq)
	fmt.Println(sendRes, err)
}
```

### 小米

```go
package main

import (
	"fmt"

	"github.com/liyouping/pushapi/xiaomipush"
)

var appSecret = "your app secret"
var regId = "your reg id"
var channelId = "your channel id"
var channelName = "your channel name"

func main() {
	client := xiaomipush.NewClient(appSecret)

	sendReq := &xiaomipush.SendReq{
		RegistrationId: regId,
		Title:          "test push title",
		Description:    "test push content",
		NotifyType:     2,
		Extra: &xiaomipush.Extra{
			NotifyEffect: "1",
			ChannelId:    channelId,
			ChannelName:  channelName,
		},
	}
	sendRes, err := client.Send(sendReq)
	fmt.Println(sendRes, err)
}
```

### 华为

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/liyouping/pushapi/huaweipush"
)

var appId = "your app id"
var appSecret = "your app secret"
var regId = "your reg id"
var badgeClass = "your badge class. example: com.example.hmstest.MainActivity"

func main() {
	client := huaweipush.NewClient(appId, appSecret)

	sendReq := &huaweipush.SendReq{
		Message: &huaweipush.Message{
			Android: &huaweipush.AndroidConfig{
				FastAppTarget: 2,
				Notification: &huaweipush.AndroidNotification{
					Title: "test push title",
					Body:  "test push content",
					ClickAction: &huaweipush.ClickAction{
						Type: 3,
					},
					Sound: strconv.Itoa(1),
					Badge: &huaweipush.BadgeNotification{
						AddNum: 1,
						Class:  badgeClass,
					},
				},
			},
			Tokens: []string{regId},
		},
	}
	sendRes, err := client.Send(sendReq)
	fmt.Println(sendRes, err)
}
```
### 荣耀

```go
package main

import (
	"fmt"
	"github.com/liyouping/pushapi/honorpush"
	"time"
)

var appId = "your app id"
var clientId = "your client id"
var clientSecret = "your client secret"

func main() {
	client := honorpush.NewClient(appId, clientId, clientSecret)
	now := time.Now().UTC()
	formatted := now.Format("2006-01-02T15:04:05.999999999Z")

	sendReq := &honorpush.SendReq{
		Android: &honorpush.Android{
			Ttl:            "86400s",
			BiTag:          "biTag001",
			TargetUserType: 1,

			Notification: &honorpush.AndroidNotification{
				Title: "荣耀通知测试",
				Body:  "荣耀通知测试body",
				Image: "https://res.vmallres.com/pimages//common/config/logo/SXppnESYv4K11DBxDFc2.png",
				Buttons: []*honorpush.ButtonItem{
					{
						Name:       "test",
						ActionType: 0,
						IntentType: 0,
					},
				},
				Badge: &honorpush.Badge{
					AddNum: 1,
				},
				ClickAction: &honorpush.ClickAction{
					Type: 3,
				},
				Importance: "NORMAL",
				Style:      0,
				NotifyId:   123,
				When:       formatted,
			},
		},
		Notification: &honorpush.Notification{
			Title: "荣耀通知测试",
			Body:  "荣耀通知测试body",
			Image: "https://res.vmallres.com/pimages//common/config/logo/SXppnESYv4K11DBxDFc2.png",
		},
		Token: []string{"BAEAAAAAB.jlyp2BnKlypiQZIk5QtnD5s9JC2kxhgXvo7NvyUIQBWys_4tHLncqrJEYkeCMpUfK_6SMMkqp12P-9WphRJqDk_Ry8Mf8ilK7q1KGhSt52XKKhBSkl-9IM"},
	}
	sendRes, err := client.Send(sendReq)
	fmt.Println(sendRes, err)
}
```
## License

this repo is released under the [MIT License](https://github.com/liyouping/pushapi/blob/master/LICENSE).

