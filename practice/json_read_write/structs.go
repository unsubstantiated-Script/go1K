package json_read_write

type Permissions map[string]bool

type UserIn struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions Permissions `json:"perms,omitempty"`
}

type UserOut struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"perms,omitempty"`
}
