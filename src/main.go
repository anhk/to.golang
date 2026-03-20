/**
 * to anywhere
 **/
package main

import (
	"os"

	"github.com/spf13/cobra"
)

const cfgFile = "/usr/local/etc/to.json"

var printPassword bool

func main() {
	var rootCmd = &cobra.Command{
		Use:   "to [target]",
		Short: "SSH 连接工具，快速连接到配置的服务器",
		Long:  `根据配置的服务器列表，通过标签名快速建立 SSH 连接。`,
		Args:  cobra.MaximumNArgs(1),
		RunE:  run,
	}

	rootCmd.Flags().BoolVarP(&printPassword, "print-password", "p", false, "打印目标服务器的密码")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	var serverList ServerList

	if err := serverList.Load(cfgFile); err != nil {
		return err
	}

	// 无参数 或 带 -p：打印 serverList.Show()，-p 时输出密码
	if len(args) == 0 || printPassword {
		serverList.Show(printPassword)
		return nil
	}

	// 否则执行 server.Run()
	target := args[0]
	server := serverList.Find(target)
	if server == nil {
		serverList.Show(false)
		return nil
	}

	server.Run()
	return nil
}
