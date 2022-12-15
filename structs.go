package discorduser

type Session struct {
	Authorization string
}

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Tags        string `json:"discriminator"`
	Avatar      string `json:"avatar"`
	Bot         bool   `json:"bot"`
	System      bool   `json:"system"`
	MfaEnabled  bool   `json:"mfa_enabled"`
	Banner      string `json:"banner"`
	BannerColor int    `json:"accent_color"`
	Locale      string `json:"locale"`
	Verified    bool   `json:"verified"`
	Email       string `json:"email"`
	Flags       int    `json:"flags"`
	NitroType   int    `json:"premium_type"`
	PublicFlags int    `json:"public_flags"`
}

type Guild struct {
	Id                          string          `json:""`
	Name                        string          `json:""`
	Icon                        string          `json:""`
	Icon_hash                   string          `json:""`
	Splash                      string          `json:""`
	DiscoverySplash             string          `json:""`
	Owner                       bool            `json:""`
	OwnerId                     string          `json:""`
	Permissions                 int             `json:""`
	Region                      string          `json:""`
	AfkChannelId                string          `json:""`
	AfkTimeout                  int             `json:""`
	WidgetEnabled               bool            `json:""`
	WidgetChannelId             string          `json:""`
	VerificationLevel           int             `json:""`
	DefaultMessageNotifications string          `json:""`
	ExplicitContentFilter       string          `json:""`
	Roles                       []Role          `json:""`
	Emojis                      []Emoji         `json:""`
	Features                    []Feature       `json:""`
	MfaLevel                    int             `json:""`
	ApplicationId               string          `json:""`
	SystemChannelId             string          `json:""`
	SystemChannelFlags          int             `json:""`
	RulesChannelId              string          `json:""`
	MaxPresences                int             `json:""`
	MaxMembers                  int             `json:""`
	VanityUrlCode               string          `json:""`
	Description                 string          `json:""`
	Banner                      string          `json:""`
	PremiumTier                 int             `json:""`
	PremiumSubscriptionCount    int             `json:""`
	PreferredLocale             string          `json:""`
	PublicUpdatesChannelId      string          `json:""`
	MaxVideoChannelUsers        int             `json:""`
	ApproximateMemberCount      int             `json:""`
	ApproximatePresenceCount    int             `json:""`
	WelcomeScreen               []WelcomeScreen `json:""`
	NsfwLevel                   int             `json:""`
	Stickers                    []Sticker       `json:""`
	PremiumProgressBarEnabled   bool            `json:""`
}

type Role struct {
}

type Emoji struct {
}

type Feature interface {
}

type WelcomeScreen struct {
}

type Sticker struct {
}
