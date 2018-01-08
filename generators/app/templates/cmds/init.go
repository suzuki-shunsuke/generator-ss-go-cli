package cmds

import (
	"io/ioutil"
	"os"

	"github.com/urfave/cli"

	"<%= package %>/assets"
	"<%= package %>/services"
)

func initCore(isDryRun bool, cfgFilePath string) (err error) {
	if _, err = os.Stat(cfgFilePath); err == nil {
		return nil
	}
	data, err := assets.Asset("data/config.yml")
	if err != nil {
		return err
	}
	if isDryRun {
		return nil
	}
	return ioutil.WriteFile(cfgFilePath, data, os.ModePerm)
}

func Init(c *cli.Context) error {
	cfgFilePath := c.String("config")
	isDryRun := c.Bool("dry-run")
	err := initCore(isDryRun, cfgFilePath)
	if err == nil {
		return nil
	}
	return cli.NewExitError(err, services.GetStatusCode(err))
}
