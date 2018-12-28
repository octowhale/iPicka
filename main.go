package main

func main() {

	// logrus.Debugln(config.Storage)
	// Logger := logger.LogConfig{ENV: "debug"}
	// Logger.SetEnv(Logger.ENV)
	// Logger.SetLevel(Logger.Level)

	var HTTPSSchema string
	if config.Storage.HTTPS {
		HTTPSSchema = "https://"
	} else {
		HTTPSSchema = "http://"
	}

	file := "/data/tmp/naruto.jpg"

	Once(file, HTTPSSchema)

}
