package leialib

type (
	// LEIA Gateway user connection
	UserConnect struct {
		OTP         string `json:"otp"`
		Server      string `json:"server"`
		ServerGroup string `json:"serverGroup"`
	}
)
