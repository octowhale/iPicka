package backend

type Config struct {
	Dbname   string `json:"dbname,omitempty"`
	Host     string `json:"host,omitempty"`
	Password string `json:"password,omitempty"`
	Port     string `json:"port,omitempty"`
	Driver   string `json:"driver,omitempty"`
	User     string `json:"user,omitempty"`
}
