package tools

var Arch = map[string]map[string]map[string]string{
	Asdf: {
		"darwin": {
			"amd64": "amd64",
			"arm64": "arm64",
		},
		"linux": {
			"amd64": "amd64",
			"arm64": "arm64",
		},
	},
	Bat: {
		"darwin": {
			"amd64": "x86_64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm64": "aarch64",
		},
	},
	Bottom: {
		"darwin": {
			"amd64": "x86_64",
			"arm64": "aarch64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm64": "aarch64",
		},
	},
	Cloudflared: {
		"darwin": {
			"amd64": "amd64",
		},
		"linux": {
			"amd64": "amd64",
			"arm64": "arm64",
		},
	},
	Flyctl: {
		"darwin": {
			"amd64": "x86_64",
			"arm64": "arm64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm64": "arm64",
		},
	},
	Hugo: {
		"darwin": {
			"amd64": "universal",
			"arm64": "universal",
		},
		"linux": {
			"amd64": "amd64",
			"arm64": "arm64",
		},
	},
	K9s: {
		"darwin": {
			"amd64": "amd64",
			"arm64": "arm64",
		},
		"linux": {
			"amd64": "amd64",
			"arm64": "arm64",
		},
	},
	Kubectx: {
		"darwin": {
			"amd64": "x86_64",
			"arm64": "arm64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm64": "arm64",
		},
	},
	Ripgrep: {
		"darwin": {
			"amd64": "x86_64",
			"arm64": "aarch64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm64": "aarch64",
		},
	},
	Shellcheck: {
		"darwin": {
			"amd64": "x86_64",
			"arm64": "aarch64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm64": "aarch64",
		},
	},
	UPX: {
		"linux": {
			"amd64": "amd64",
			"arm64": "arm64",
		},
	},
	Uv: {
		"darwin": {
			"amd64": "x86_64",
			"arm64": "aarch64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm64": "aarch64",
		},
	},
	Xh: {
		"darwin": {
			"amd64": "x86_64",
			"arm64": "aarch64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm64": "aarch64",
		},
	},
	Yj: {
		"darwin": {
			"amd64": "amd64",
			"arm64": "arm64",
		},
		"linux": {
			"amd64": "amd64",
			"arm64": "arm64",
		},
	},
}

var OS = map[string]map[string]map[string]string{
	Asdf: {
		"darwin": {
			"amd64": "darwin",
			"arm64": "darwin",
		},
		"linux": {
			"amd64": "linux",
			"arm64": "linux",
		},
	},
	Bat: {
		"darwin": {
			"amd64": "apple-darwin",
		},
		"linux": {
			"amd64": "unknown-linux-gnu",
			"arm64": "unknown-linux-gnu",
		},
	},
	Bottom: {
		"darwin": {
			"amd64": "apple-darwin",
			"arm64": "apple-darwin",
		},
		"linux": {
			"amd64": "unknown-linux-gnu",
			"arm64": "unknown-linux-gnu",
		},
	},
	Cloudflared: {
		"darwin": {
			"amd64": "darwin",
		},
		"linux": {
			"amd64": "linux",
			"arm64": "linux",
		},
	},
	Flyctl: {
		"darwin": {
			"amd64": "macOS",
			"arm64": "macOS",
		},
		"linux": {
			"amd64": "Linux",
			"arm64": "Linux",
		},
	},
	Hugo: {
		"darwin": {
			"amd64": "darwin",
			"arm64": "darwin",
		},
		"linux": {
			"amd64": "linux",
			"arm64": "linux",
		},
	},
	K9s: {
		"darwin": {
			"amd64": "Darwin",
			"arm64": "Darwin",
		},
		"linux": {
			"amd64": "Linux",
			"arm64": "Linux",
		},
	},
	Kubectx: {
		"darwin": {
			"amd64": "darwin",
			"arm64": "darwin",
		},
		"linux": {
			"amd64": "linux",
			"arm64": "linux",
		},
	},
	Ripgrep: {
		"darwin": {
			"amd64": "apple-darwin",
			"arm64": "apple-darwin",
		},
		"linux": {
			"amd64": "unknown-linux-musl",
			"arm64": "unknown-linux-gnu",
		},
	},
	Shellcheck: {
		"darwin": {
			"amd64": "darwin",
			"arm64": "darwin",
		},
		"linux": {
			"amd64": "linux",
			"arm64": "linux",
		},
	},
	UPX: {
		"linux": {
			"amd64": "linux",
			"arm64": "linux",
		},
	},
	Uv: {
		"darwin": {
			"amd64": "apple-darwin",
			"arm64": "apple-darwin",
		},
		"linux": {
			"amd64": "unknown-linux-gnu",
			"arm64": "unknown-linux-gnu",
		},
	},
	Xh: {
		"darwin": {
			"amd64": "apple-darwin",
			"arm64": "apple-darwin",
		},
		"linux": {
			"amd64": "unknown-linux-musl",
			"arm64": "unknown-linux-musl",
		},
	},
	Yj: {
		"darwin": {
			"amd64": "macos",
			"arm64": "macos",
		},
		"linux": {
			"amd64": "linux",
			"arm64": "linux",
		},
	},
}

var URL = map[string]string{
	Asdf:        "https://github.com/asdf-vm/asdf",
	Bat:         "https://github.com/sharkdp/bat",
	Bottom:      "https://github.com/ClementTsang/bottom",
	Cloudflared: "https://github.com/cloudflare/cloudflared",
	Flyctl:      "https://github.com/superfly/flyctl",
	Hugo:        "https://github.com/gohugoio/hugo",
	K9s:         "https://github.com/derailed/k9s",
	Kubectx:     "https://github.com/ahmetb/kubectx",
	Ripgrep:     "https://github.com/BurntSushi/ripgrep",
	Shellcheck:  "https://github.com/koalaman/shellcheck",
	UPX:         "https://github.com/upx/upx",
	Uv:          "https://github.com/astral-sh/uv",
	Xh:          "https://github.com/ducaale/xh",
	Yj:          "https://github.com/sclevine/yj",
}
