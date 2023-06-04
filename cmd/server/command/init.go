package command

import (
	"fmt"
	"github.com/kenlabs/permastar/pkg/system"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func InitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize server config file.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := checkPermastarRoot(); err != nil {
				return err
			}

			configFile := filepath.Join(Opt.PermastarRoot, Opt.ConfigFile)
			fmt.Println("Init permastar-server configs at ", configFile)
			if err := checkConfigExists(configFile); err != nil {
				return err
			}
			if err := saveConfig(configFile); err != nil {
				return err
			}

			fmt.Println("init complete.")

			return nil
		},
	}
}

func checkPermastarRoot() error {
	const failedError = "check permastar root failed:\n\t%v\n"
	rootExists, err := system.IsDirExists(Opt.PermastarRoot)
	if err != nil {
		return fmt.Errorf(failedError, err)
	}
	if !rootExists {
		fmt.Printf("permastar root %v does not exist, trying to create...\n", Opt.PermastarRoot)
		err := os.MkdirAll(Opt.PermastarRoot, 0755)
		if err != nil {
			return fmt.Errorf("create permastar root %s failed: %v", Opt.PermastarRoot, err)
		}
	}

	rootWritable, err := system.IsDirWritable(Opt.PermastarRoot)
	if err != nil {
		return fmt.Errorf(failedError, err)
	}
	if !rootWritable {
		return fmt.Errorf("permastar root %s is not writable\n", Opt.PermastarRoot)
	}

	return nil
}

func checkConfigExists(configFile string) error {
	configExists, err := system.IsFileExists(configFile)
	if err != nil {
		return fmt.Errorf("init config failed: %v", err)
	}
	if configExists {
		return fmt.Errorf("config file exists: %s", configFile)
	}
	return nil
}

func saveConfig(configFile string) error {
	buff, err := yaml.Marshal(Opt)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0755)
	defer func(file *os.File) { _ = file.Close() }(file)
	if err != nil {
		return err
	}
	_, err = file.WriteString(string(buff))
	if err != nil {
		return err
	}

	return nil
}
