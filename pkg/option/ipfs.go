package option

const (
	defaultIPFSNodeGatewayIP   = "127.0.0.1"
	defaultIPFSNodeGatewayPort = "5001"
)

type IPFSNode struct {
	GatewayIP   string `yaml:"GatewayIP"`
	GatewayPort string `yaml:"GatewayPort"`
}
