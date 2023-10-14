package configmap

type Config struct {
	Database Database `validate:"required" yaml:"database"`
	Trades   Trades   `validate:"required" yaml:"trades"`
}

type Database struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Trades struct {
	CSVPath string `yaml:"csvpath"`
}
