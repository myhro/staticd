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

func run(name string, version string) {
	tool := &tools.Tool{
		Name:    name,
		User:    &tools.UserOS{},
		Version: version,
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

func newCommand(name string, description string) *cobra.Command {
	return &cobra.Command{
		Use:   name + " [version]",
		Short: description,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			version := "latest"
			if len(args) == 1 {
				version = args[0]
			}
			run(name, version)
		},
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "staticd",
		Short: "Download statically linked binaries from GitHub",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	batCmd := newCommand(tools.Bat, "A cat(1) clone with wings")
	bottomCmd := newCommand(tools.Bottom, "Yet another cross-platform graphical process/system monitor")
	cloudflaredCmd := newCommand(tools.Cloudflared, "Argo Tunnel client")
	flyctlCmd := newCommand(tools.Flyctl, "Command line tools for fly.io services")
	k9sCmd := newCommand(tools.K9s, "Kubernetes CLI To Manage Your Clusters In Style")
	kubectxCmd := newCommand(tools.Kubectx, "Faster way to switch between clusters in kubectl")
	ripgrepCmd := newCommand(tools.Ripgrep, "Recursively searches directories for a regex pattern")
	shellcheckCmd := newCommand(tools.Shellcheck, "A static analysis tool for shell scripts")
	upxCmd := newCommand(tools.UPX, "The Ultimate Packer for eXecutables")
	uvCmd := newCommand(tools.Uv, "An extremely fast Python package and project manager")
	xhCmd := newCommand(tools.Xh, "Friendly and fast tool for sending HTTP requests")
	yjCmd := newCommand(tools.Yj, "Convert between YAML, TOML, JSON, and HCL")

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
	rootCmd.AddCommand(flyctlCmd)
	rootCmd.AddCommand(k9sCmd)
	rootCmd.AddCommand(kubectxCmd)
	rootCmd.AddCommand(ripgrepCmd)
	rootCmd.AddCommand(shellcheckCmd)
	rootCmd.AddCommand(upxCmd)
	rootCmd.AddCommand(uvCmd)
	rootCmd.AddCommand(xhCmd)
	rootCmd.AddCommand(yjCmd)
	rootCmd.AddCommand(versionCmd)

	err := rootCmd.Execute()
	if err != nil {
		exit(err.Error())
	}
}
