package storage

type Config struct {
	Acckey     string `json:"acckey,omitempty"`
	Accsec     string `json:"accsec,omitempty"`
	Bucketname string `json:"bucketname,omitempty"`
	Internal   bool   `json:"internal,omitempty"`
	Region     string `json:"region,omitempty"`
	Driver     string `json:"driver,omitempty"`
}
