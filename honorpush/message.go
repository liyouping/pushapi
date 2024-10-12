package honorpush

const (
	Host     = "https://iam.developer.honor.com"
	PushHost = "https://push-api.cloud.honor.com"
	AuthURL  = "/auth/token" // 推送鉴权接口
)

type AuthReq struct {
	GrantType    string `json:"grant_type"`    //固定值为:client_credentials
	ClientID     string `json:"client_id"`     //开发者平台开通该应用PUSH服务后应用的Client ID
	ClientSecret string `json:"client_secret"` //开发者平台开通该应用PUSH服务后应用的Client Secret
}

type AuthRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
type SendReq struct {
	Data         string        `json:"data,omitempty"` //自定义消息负载，通知栏消息支持JSON格式字符串，透传消息支持普通字符串或者JSON格式字符串。
	Notification *Notification `json:"notification,omitempty"`
	Android      *Android      `json:"android,omitempty"`
	Token        []string      `json:"token,omitempty"` //按照Token向目标用户推消息
}
type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Image string `json:"image"`
}

type Android struct {
	Ttl            string               `json:"ttl"` //消息缓存时间，单位是秒。在用户设备离线时，消息在Push服务器进行缓存，在消息缓存时间内用户设备上线，消息会下发，超过缓存时间后消息会丢弃，默认值为“86400s”（1天），最大值为“1296000s”（15天）。
	BiTag          string               `json:"biTag"`
	Data           string               `json:"data"`
	Notification   *AndroidNotification `json:"notification"`
	TargetUserType int                  `json:"targetUserType"`
}

type AndroidNotification struct {
	Title       string        `json:"title"`
	Body        string        `json:"body"`
	ClickAction *ClickAction  `json:"clickAction,omitempty"`
	Image       string        `json:"image"`
	Style       int           `json:"style"`
	BigTitle    string        `json:"bigTitle"`
	BigBody     string        `json:"bigBody"`
	Importance  string        `json:"importance"`
	When        string        `json:"when"`
	Buttons     []*ButtonItem `json:"buttons,omitempty"`
	Badge       *Badge        `json:"badge,omitempty"`
	NotifyId    int           `json:"notifyId"`
	Tag         string        `json:"tag"`
	Group       string        `json:"group"`
}

type ClickAction struct {
	Type   int    `json:"type"`
	Intent string `json:"intent"`
	Url    string `json:"url"`
	Action string `json:"action"`
}

type ButtonItem struct {
	Name       string `json:"name"`
	ActionType int    `json:"actionType"`
	IntentType int    `json:"intentType"`
	Intent     string `json:"intent"`
	Data       string `json:"data"`
}

type Badge struct {
	AddNum     int    `json:"addNum"`
	BadgeClass string `json:"badgeClass"`
	SetNum     int    `json:"setNum"`
}

type SendRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		SendResult   bool     `json:"sendResult"`
		RequestId    string   `json:"requestId"`
		FailTokens   []string `json:"failTokens"`
		ExpireTokens []string `json:"expireTokens"`
	} `json:"data"`
}
