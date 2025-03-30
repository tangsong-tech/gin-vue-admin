package config

type Web3 struct {
	RpcUrl        string `mapstructure:"rpc-url" json:"rpc-url" yaml:"rpc-url"`
	ApiKey        string `mapstructure:"api-key" json:"api-key" yaml:"api-key"`
	ApiSecret     string `mapstructure:"api-secret" json:"api-secret" yaml:"api-secret"`
	ApiPassphrase string `mapstructure:"api-passphrase" json:"api-passphrase" yaml:"api-passphrase"`
	Deposit       string `mapstructure:"deposit" json:"deposit" yaml:"deposit"`
	Withdraw      string `mapstructure:"withdraw" json:"withdraw" yaml:"withdraw"`
	Ucontract     string `mapstructure:"ucontract" json:"ucontract" yaml:"ucontract"`
	Tokencontract string `mapstructure:"tokencontract" json:"tokencontract" yaml:"tokencontract"`
}
