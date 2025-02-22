package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
	JwtKey           string `mapstructure:"JWT_KEY"`
	DBUri            string `mapstructure:"DB_URI"`
	AwsS3Bucket      string `mapstructure:"AWS_S3_BUCKET"`
	AwsS3Hostname    string `mapstructure:"AWS_S3_HOSTNAME"`
	AwsS3AccessKeyId string `mapstructure:"AWS_S3_ACCESS_KEY_ID"`
	AwsS3SecretKey   string `mapstructure:"AWS_S3_SECRET_KEY"`
	AwsS3UseSSL      bool   `mapstructure:"AWS_S3_USE_SSL"`
	SMTPUsername	 string `mapstructure:"SMTP_USERNAME"`
	SMTPPassword	 string `mapstructure:"SMTP_PASSWORD"`
	AmqpURI			 string `mapstructure:"AMQP_URI"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file, %s", err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Unable to decode into struct, %s", err)
	}
	return
}
