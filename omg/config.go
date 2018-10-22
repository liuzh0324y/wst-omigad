package omg

type httpconfig struct {
	Server struct {
		HTTPPort    string `yaml:"httpport"`
		Endpoint    string `yaml:"endpoint"`
		AccessKey   string `yaml:"accesskey"`
		SecretKey   string `yaml:"secretkey"`
		Bucket      string `yaml:"bucket"`
		FilePath    string `yaml:"filepath"`
		ManagerURL  string `yaml:"managerurl"`
		CallbackURL string `yaml:"callbackurl"`
		IsRemove    bool   `yaml:"isRemove"`
	} `yaml:"server"`
}
