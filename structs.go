package discorduser

type Session struct {
	Authorization string
}

type User struct {
	Id           string `json:"id"`
	Username     string `json:"username"`
	Tags         string `json:"discriminator"`
	Avatar       string `json:"avatar"`
	Bot          bool   `json:"bot"`
	System       bool   `json:"system"`
	Mfa_enabled  bool   `json:"mfa_enabled"`
	Banner       string `json:"banner"`
	BannerColor  int    `json:"accent_color"`
	Locale       string `json:"locale"`
	Verified     bool   `json:"verified"`
	Email        string `json:"email"`
	Flags        int    `json:"flags"`
	NitroType    int    `json:"premium_type"`
	Public_Flags int    `json:"public_flags"`
}
