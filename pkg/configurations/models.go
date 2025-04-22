package configurations

type DynamicConfig struct {
	Sbi Sbi `json:"sbi,omitempty"`
}

type Sbi struct {
	OAuth2 OAuth2 `json:"oauth2,omitempty"`
}

type OAuth2 struct {
	// Enable indicates whether the OAuth2 authentication is enabled in SBI
	Enable bool `json:"enable,omitempty"`

	// TokenExpiration is the expiration time of the token in milliseconds
	Period int32 `json:"period,omitempty"`
}
