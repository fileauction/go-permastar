package option

const (
	defaultHttpAPIListenAddress = "/ip4/0.0.0.0/tcp/9000"
)

type ServerAddress struct {
	HttpAPIListenAddress string `yaml:"HttpAPIListenAddress"`
	ExternalIP           string `yaml:"ExternalIP"`
}
