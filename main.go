package main

import (
	"github.com/BoruTamena/infra/config"
)

func main() {

	// setting the env path
	cf_viper := config.EnvConfig()
	vp := cf_viper.GetViper()
	vp.SetConfigFile(".env")

	//

}
