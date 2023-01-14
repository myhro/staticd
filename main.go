package main

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/myhro/staticd/tools"
)

var version = "dev"

func exit(msg string) {
	//nolint:errcheck
	io.WriteString(os.Stderr, fmt.Sprintf("Error: %v\n", msg))
	os.Exit(1)
}

func run(name string) {
	tool := &tools.Tool{
		Name: name,
	}

	err := tool.SetURL()
	if err != nil {
		exit(err.Error())
	}

	err = tool.SetRuntime(runtime.GOARCH, runtime.GOOS)
	if err != nil {
		exit(err.Error())
	}

	fmt.Println("Checking version for", tool.Name)

	err = tool.GetVersion()
	if err != nil {
		exit(err.Error())
	}

	err = tool.SetAsset()
	if err != nil {
		exit(err.Error())
	}

	fmt.Println("Downloading version", tool.Version)

	err = tool.Download()
	if err != nil {
		exit(err.Error())
	}

	fmt.Println("Downloaded", tool.Asset.Name)

	if !tool.Asset.IsBinary {
		fmt.Println("Extracting binary")

		err := tool.Extract()
		if err != nil {
			exit(err.Error())
		}
	}

	fmt.Println("Done")
}

//nolint:funlen
func main() {
	rootCmd := &cobra.Command{
		Use:   "staticd",
		Short: "Download statically linked binaries from GitHub",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	batCmd := &cobra.Command{
		Use:   tools.Bat,
		Short: "A cat(1) clone with wings",
		Run: func(cmd *cobra.Command, args []string) {
			run(tools.Bat)
		},
	}

	bottomCmd := &cobra.Command{
		Use:   tools.Bottom,
		Short: "Yet another cross-platform graphical process/system monitor",
		Run: func(cmd *cobra.Command, args []string) {
			run(tools.Bottom)
		},
	}

	cloudflaredCmd := &cobra.Command{
		Use:   tools.Cloudflared,
		Short: "Argo Tunnel client",
		Run: func(cmd *cobra.Command, args []string) {
			run(tools.Cloudflared)
		},
	}

	dockerComposeCmd := &cobra.Command{
		Use:   tools.DockerCompose,
		Short: "Define and run multi-container applications with Docker",
		Run: func(cmd *cobra.Command, args []string) {
			run(tools.DockerCompose)
		},
	}

	flyctlCmd := &cobra.Command{
		Use:   tools.Flyctl,
		Short: "Command line tools for fly.io services",
		Run: func(cmd *cobra.Command, args []string) {
			run(tools.Flyctl)
		},
	}

	k9sCmd := &cobra.Command{
		Use:   tools.K9s,
		Short: "Kubernetes CLI To Manage Your Clusters In Style",
		Run: func(cmd *cobra.Command, args []string) {
			run(tools.K9s)
		},
	}

	upxCmd := &cobra.Command{
		Use:   tools.UPX,
		Short: "The Ultimate Packer for eXecutables",
		Run: func(cmd *cobra.Command, args []string) {
			run(tools.UPX)
		},
	}

	xhCmd := &cobra.Command{
		Use:   tools.Xh,
		Short: "Friendly and fast tool for sending HTTP requests",
		Run: func(cmd *cobra.Command, args []string) {
			run(tools.Xh)
		},
	}

	yjCmd := &cobra.Command{
		Use:   tools.Yj,
		Short: "Convert between YAML, TOML, JSON, and HCL",
		Run: func(cmd *cobra.Command, args []string) {
			run(tools.Yj)
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}

	rootCmd.AddCommand(batCmd)
	rootCmd.AddCommand(bottomCmd)
	rootCmd.AddCommand(cloudflaredCmd)
	rootCmd.AddCommand(dockerComposeCmd)
	rootCmd.AddCommand(flyctlCmd)
	rootCmd.AddCommand(k9sCmd)
	rootCmd.AddCommand(upxCmd)
	rootCmd.AddCommand(xhCmd)
	rootCmd.AddCommand(yjCmd)
	rootCmd.AddCommand(versionCmd)

	err := rootCmd.Execute()
	if err != nil {
		exit(err.Error())
	}
}
