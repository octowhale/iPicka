package storage

type Config struct {
	Storage      string
	AccKey       string
	AccSec       string
	CustomDomain string
	HTTPSSchema  bool
	Endpoint     string
	Region       string
	Bucket       string
	Internal     bool
	Prefix       string
}
