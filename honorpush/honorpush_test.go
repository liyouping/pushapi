package honorpush_test

import (
	"github.com/liyouping/pushapi/honorpush"
	"testing"
	"time"
)

var appId = "your app id"
var clientId = "your client id"
var clientSecret = "your client secret"

func TestSend(t *testing.T) {
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
		Token: []string{"BAEAAAAAB.jlyp2BnKdXgiQZIk5QtnD5s9JC2kxhgXvo7NvyUIQBWys_6tHLncqrJEYkeCMpUfK_6SMMkqp47P-9WphRJqDk_Ry8Mf9ilK7q1KGhSt52XKKhBSkl-1IM"},
	}
	sendRes, err := client.Send(sendReq)
	t.Log(sendRes, err)
}
