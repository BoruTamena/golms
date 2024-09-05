package config

import "github.com/spf13/viper"

type viperLoader struct {
	vp *viper.Viper
}

func EnvConfig() *viperLoader {

	return &viperLoader{
		vp: viper.New(),
	}

}

func (vpLoader viperLoader) GetViper() *viper.Viper {

	return vpLoader.vp
}
