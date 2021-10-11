package gomod

// GetModeration represents the response object that returned from Get Moderation API.
type GetModeration struct {
	Data Moderation `json:"data"`
	Meta Meta       `json:"meta"`
}

// GetModerations represents the response object that returned from Get List of Moderation API.
type GetModerations struct {
	Data Moderations `json:"data"`
	Meta Meta        `json:"meta"`
}

// PostModeration respresents the created object that returned from Create Moderation API.
type PostModeration struct {
	Data Moderation `json:"data"`
	Meta Meta       `json:"meta"`
}

// Moderation represents Moderation object.
type Moderation struct {
	ID            string  `json:"id"`
	Answer        string  `json:"answer"`
	CreditCharged float32 `json:"credit_charged"`
	CustomID      string  `json:"custom_id"`
	Source        string  `json:"data"`
	PostbackURL   string  `json:"postback_url"`
	ProcessedAt   string  `json:"processed_at"`
	ProjectID     int     `json:"project_id"`
	Status        string  `json:"status"`
}

// Moderations represents list of Moderation object.
type Moderations []Moderation
