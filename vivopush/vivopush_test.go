package vivopush_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/liyouping/pushapi/vivopush"
)

var appId = "your app id"
var appKey = "your app key"
var appSecret = "your app secret"
var regId = "your reg id"
var regIds = []string{"your regId1", "your regId2"} // 个数必须大于等于2，并且不能重复

func TestSend(t *testing.T) {
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
	t.Log(sendRes, err)
}

func TestSendBatch(t *testing.T) {
	client := vivopush.NewClient(appId, appKey, appSecret)

	sendReq := &vivopush.SendBatchReq{
		Notification: &vivopush.SendReq{
			RegId:          regId,
			NotifyType:     4,
			Title:          "test push title",
			Content:        "test push content",
			TimeToLive:     24 * 60 * 60,
			SkipType:       1,
			NetworkType:    -1,
			Classification: 1,
			RequestId:      strconv.Itoa(int(time.Now().UnixNano())),
		},
		MsgConfig: &vivopush.MsgConfig{
			RegIds:    regIds,
			RequestId: strconv.Itoa(int(time.Now().UnixNano())),
		},
	}
	sendRes, err := client.SendBatch(sendReq)
	t.Log(sendRes, err)
}
