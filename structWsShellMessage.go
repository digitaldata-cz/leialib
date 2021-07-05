package leialib

type (
	WsShellMsg struct {
		Status   *string            `json:"status,omitempty"`   // Status message (response to login)
		Connect  *WsShellMsgConnect `json:"connect,omitempty"`  // Open connection to remote
		Login    *string            `json:"login,omitempty"`    //
		Data     *[]byte            `json:"d,omitempty"`        // Terminal data (keyboard / display)
		TermSize *string            `json:"termSize,omitempty"` // Set new terminal size
	}
	WsShellMsgConnect struct {
		Remote     string   `json:"remote"`        // Remote ifentification (name)
		RemoteUser string   `json:"remoteUser"`    // Remote user (recieved from admin API)
		OTP        string   `json:"otp"`           // OTP auth from user (ltc)
		Env        []string `json:"env,omitempty"` // Enviroment variables
	}
)
