package models

// Inbox related types
type Inbox struct {
	BaseEntity
	AutoReplyEnabled              bool         `json:"autoReplyEnabled"`
	AutoReplyMessage              string       `json:"autoReplyMessage"`
	AutoReplySubject              string       `json:"autoReplySubject"`
	ClientOnly                    bool         `json:"clientOnly"`
	DisplayOrder                  int          `json:"displayOrder"`
	Email                         string       `json:"email"`
	EmailForwardingState          string       `json:"emailForwardingState"`
	ForwardingAddress             string       `json:"forwardingAddress"`
	HappinessRatingEnabled        bool         `json:"happinessRatingEnabled"`
	HappinessRatingMessage        string       `json:"happinessRatingMessage"`
	IconImage                     string       `json:"iconImage"`
	IsAdmin                       bool         `json:"isAdmin"`
	IsFreeDomain                  bool         `json:"isFreeDomain"`
	LanguageCode                  string       `json:"languageCode"`
	LocalPart                     string       `json:"localPart"`
	Name                          string       `json:"name"`
	NotificationsOnly             bool         `json:"notificationsOnly"`
	OnClosedLock                  string       `json:"onClosedLock"`
	OnClosedWait                  int          `json:"onClosedWait"`
	PublicIconImage               string       `json:"publicIconImage"`
	SendEmailsFrom                string       `json:"sendEmailsFrom"`
	Signature                     string       `json:"signature"`
	SMTPPassword                  string       `json:"smtpPassword"`
	SMTPProvider                  string       `json:"smtpProvider"`
	SMTPPort                      int          `json:"smtpPort"`
	SMTPSecurity                  string       `json:"smtpSecurity"`
	SMTPServer                    string       `json:"smtpServer"`
	SMTPUsername                  string       `json:"smtpUsername"`
	SpamThreshold                 int          `json:"spamThreshold"`
	Starred                       bool         `json:"starred"`
	Synced                        bool         `json:"synced"`
	SyncDays                      any          `json:"syncDays"`
	SyncAccountID                 any          `json:"syncAccountId"`
	SyncSubscriptionID            any          `json:"syncSubscriptionId"`
	TimeloggingEnabled            bool         `json:"timeloggingEnabled"`
	Type                          string       `json:"type"`
	UseTeamworkMailServer         bool         `json:"useTeamworkMailServer"`
	UsingOfficeHours              bool         `json:"usingOfficeHours"`
	Verified                      bool         `json:"verified"`
	IncludeTicketHistoryOnForward bool         `json:"includeTicketHistoryOnForward"`
	User                          any          `json:"user"`
	Users                         []InboxUser  `json:"users"`
	Inboxaliases                  []EntityRef  `json:"inboxaliases"`
	Ticketstatus                  EntityRef    `json:"ticketstatus"`
	Inboxemailrefs                any          `json:"inboxemailrefs"`
	Triggers                      []Trigger    `json:"triggers"`
	Tickettypes                   []EntityRef  `json:"tickettypes"`
	Inboxcnames                   []InboxCname `json:"inboxcnames"`
	Oauth2Token                   any          `json:"oauth2token"`
	Restricteddomains             any          `json:"restricteddomains"`
}

type InboxesResponse struct {
	Inboxes    []Inbox      `json:"inboxes"`
	Included   IncludedData `json:"included"`
	Pagination Pagination   `json:"pagination"`
	Meta       Meta         `json:"meta"`
}

type InboxResponse struct {
	Inbox    Inbox        `json:"inbox"`
	Included IncludedData `json:"included"`
}

type InboxUser struct {
	EntityRef
	Meta struct {
		Access  string `json:"access"`
		IsAdmin bool   `json:"isAdmin"`
		Starred bool   `json:"starred"`
		State   string `json:"state"`
	} `json:"meta"`
}

type Trigger struct {
	EntityRef
	Meta struct {
		DisplayOrder int `json:"displayOrder"`
	} `json:"meta"`
}

type InboxCname struct {
	EntityRef
	Meta struct {
		Domain string `json:"domain"`
	} `json:"meta"`
}
