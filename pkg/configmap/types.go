package configmap

type Config struct {
	Database Database `validate:"required" yaml:"database"`
}

type Database struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
