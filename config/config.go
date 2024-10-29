package config

// Config is the database configuration
type Config struct {
	Host      string `json:"host"`
	Name      string `json:"name"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Deletable bool   `json:"deletable"`
}

// New prepares the configuration for mxorm
func New(host string, port int, name, user, password string, allowDelete bool) *Config {
	return &Config{
		Host:      host,
		Port:      port,
		Name:      name,
		User:      user,
		Password:  password,
		Deletable: allowDelete,
	}
}
