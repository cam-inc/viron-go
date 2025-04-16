package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/automaxprocs/maxprocs"

	"github.com/cam-inc/viron-go-example/pkg/constant"
	"github.com/cam-inc/viron-go-example/pkg/server"
	"github.com/cam-inc/viron-go-example/routes"
	pkgConstant "github.com/cam-inc/viron-go/constant"
)

type (
	options struct {
		host string
		port int
		mode string
	}
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}

	if _, err := maxprocs.Set(); err != nil {
		fmt.Print(err)
		os.Exit(2)
	}

	c := &cobra.Command{}
	o := &options{}

	c.PersistentFlags().StringVarP(&o.mode, "mode", "m", "", "datastore mode")

	// registerCommand(c)
	c.RunE = func(c *cobra.Command, args []string) error {
		if o.mode == "" {
			fmt.Print("mode is required")
			os.Exit(2)
		}
		if err := os.Setenv(pkgConstant.ENV_STORE_MODE, o.mode); err != nil {
			fmt.Print(err)
			os.Exit(2)
		}
		return run(o)
	}

	// 他の設定（host, port は環境変数から直接）
	o.host = os.Getenv(constant.SERVICE_HOST)
	o.port, _ = strconv.Atoi(os.Getenv(constant.SERVICE_PORT))

	if err := c.Execute(); err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
}

// run cobra.commandで実行される(RunE)関数
func run(o *options) error {
	s := server.New(routes.New(), o.host, o.port)
	return s.RunTLS()
}
