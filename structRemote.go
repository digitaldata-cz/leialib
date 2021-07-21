package leialib

type (
	// LEIA Remote - "register token" used to register temote server to gateway. Base64 encoded command line argument.
	RemoteRegisterToken struct {
		Gateway []string `json:"gateway"`
		Group   string   `json:"group"` // Request to group (UUID)
		CA      string   `json:"ca"`
	}

	// HTTPS request with registration sent to gateway
	RemoteRegisterRequest struct {
		Hostname string `json:"hostname"` // Remote server hostname
		Group    string `json:"group"`    // Request to group (UUID)
		Arch     string `json:"arch"`     // Server architecture (amd64, arm64, ...)
		Platform string `json:"platform"` // Server platform (linux, windows,...)
		OS       string `json:"os"`       // Operating system (CentOS 7.3, ...)
		UUID     string `json:"uuid"`     // Current UUID (register again with new token)
	}

	// HTTPS response recieved from gateway
	RemoteRegisterResponse struct {
		UUID string `json:"uuid"`
	}

	RemoteWsControlMessage struct {
		Action *string   `json:"action,omitempty"`
		User   *string   `json:"user,omitempty"`
		Env    *[]string `json:"env,omitempty"`
		UUID   *string   `json:"uuid,omitempty"`
	}
)
