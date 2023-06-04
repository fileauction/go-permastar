package estuary

type CreateUserAPIKeyResp struct {
	Expiry    string `json:"expiry,omitempty"`
	Label     string `json:"label,omitempty"`
	Token     string `json:"token,omitempty"`
	TokenHash string `json:"tokenHash,omitempty"`
}

type AddContentResp struct {
	Cid       string
	EstuaryId uint64
	Providers []string
}

type APIErrorResp struct {
	Code    int    `json:"code"`
	Details string `json:"details"`
	Reason  string `json:"reason"`
}

type Collection struct {
	Cid         string `json:"cid,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	UserId      int32  `json:"userId,omitempty"`
	Uuid        string `json:"uuid,omitempty"`
}
