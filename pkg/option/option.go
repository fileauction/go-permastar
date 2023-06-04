package option

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type DaemonOptions struct {
	flags   *pflag.FlagSet `yaml:"-"`
	viper   *viper.Viper   `yaml:"-"`
	yamlStr string         `yaml:"-"`

	ConfigFile string `yaml:"-"`

	PermastarRoot string        `yaml:"PermastarRoot"`
	LogLevel      string        `yaml:"LogLevel"`
	ServerAddress ServerAddress `yaml:"ServerAddress"`
	IPFSNode      IPFSNode      `yaml:"IPFSNode"`
}

// New creates a default DaemonOptions.
func New(root *cobra.Command) *DaemonOptions {
	var opt *DaemonOptions
	if root == nil {
		opt = &DaemonOptions{
			flags: pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError),
			viper: viper.New(),
		}
	} else {
		opt = &DaemonOptions{
			flags: root.PersistentFlags(),
			viper: viper.New(),
		}
	}

	opt.flags.StringVarP(&opt.ConfigFile, "config-file", "f", "config.yaml",
		"Load server configurations from a yaml file.")
	homeRoot, err := homedir.Expand("~/.permastar")
	if err != nil {
		panic(fmt.Errorf("expand home dir failed when create new options: %v", err))
	}
	opt.flags.StringVarP(&opt.PermastarRoot, "permastar-root", "r", homeRoot,
		"Permastar root directory.")

	opt.flags.StringVar(&opt.ServerAddress.HttpAPIListenAddress, "http-api-addr", defaultHttpAPIListenAddress,
		fmt.Sprintf("Http api server listen address in multiaddress manner."))

	opt.flags.StringVar(&opt.ServerAddress.ExternalIP, "external-ip", "unknown",
		fmt.Sprintf("Permastar public IP address for api service."))

	opt.flags.StringVar(&opt.IPFSNode.GatewayIP, "ipfsnode-gatewayip", defaultIPFSNodeGatewayIP,
		"IPFS Node gateway IP address.")

	opt.flags.StringVar(&opt.IPFSNode.GatewayPort, "ipfsnode-gatewayport", defaultIPFSNodeGatewayPort,
		"IPFS Node gateway port.")

	_ = opt.viper.BindPFlags(opt.flags)

	return opt
}

func (opt *DaemonOptions) YAML() string {
	return opt.yamlStr
}

func (opt *DaemonOptions) Parse() (string, error) {
	err := opt.flags.Parse(os.Args[1:])
	if err != nil {
		return "", err
	}

	opt.viper.AutomaticEnv()
	opt.viper.SetEnvPrefix("ps")
	opt.viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if opt.ConfigFile != "" {
		opt.viper.SetConfigFile(filepath.Join(opt.PermastarRoot, opt.ConfigFile))
		opt.viper.SetConfigType("yaml")
		err := opt.viper.ReadInConfig()
		if err != nil && !os.IsNotExist(err) {
			return "", fmt.Errorf("read config file %v failed: %v", opt.ConfigFile, err)
		}
	}

	// NOTE: Workaround because viper does not treat env vars the same as other config.
	// Reference: https://github.com/spf13/viper/issues/188#issuecomment-399518663
	for _, key := range opt.viper.AllKeys() {
		val := opt.viper.Get(key)
		opt.viper.Set(key, val)
	}

	_ = opt.viper.Unmarshal(opt, func(c *mapstructure.DecoderConfig) {
		c.TagName = "yaml"
	})

	err = opt.validate()
	if err != nil {
		return "", err
	}

	buff, err := yaml.Marshal(opt)
	if err != nil {
		return "", fmt.Errorf("marshal config to yaml failed: %v", err)
	}
	opt.yamlStr = string(buff)

	return "", nil
}

func (opt *DaemonOptions) validate() error {
	// ToDo: validate flags
	return nil
}
