package constants

import "github.com/luksrocha/house-system/util"

func EnvKeyConstants() util.Config {

	return util.Config{
		JWTSecretKey: "JWT_SECRET_KEY",
		DBDriver:     "DB_DRIVER",
		DBHost:       "DB_HOST",
		DBPort:       "DB_PORT",
		DBName:       "DB_NAME",
		DBUser:       "DB_USER",
		DBPassword:   "DB_PASSWORD",
	}

}
