package configmap

type Config struct {
	Database Database `validate:"required" yaml:"database"`
	Trades   Trades   `validate:"required" yaml:"trades"`
}

type Database struct {
	Host                   string `yaml:"host"`
	Port                   int    `yaml:"port"`
	DatabaseName           string `yaml:"databaseName"`
	TransactionsCollection string `yaml:"transactionsCollection"`
	PositionsCollection    string `yaml:"positionsCollection"`
	TaxCollection          string `yaml:"taxCollection"`
}

type Trades struct {
	CSVPath string `yaml:"csvpath"`
}
