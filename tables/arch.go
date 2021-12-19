package tables

var Arch = map[string]map[string]map[string]string{
	Bat: {
		"darwin": {
			"amd64": "x86_64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm":   "arm",
			"arm64": "aarch64",
		},
	},
	Bottom: {
		"darwin": {
			"amd64": "x86_64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm":   "armv7",
			"arm64": "aarch64",
		},
	},
	Cloudflared: {
		"darwin": {
			"amd64": "amd64",
		},
		"linux": {
			"amd64": "amd64",
			"arm":   "arm",
			"arm64": "arm64",
		},
	},
	DockerCompose: {
		"linux": {
			"amd64": "x86_64",
			"arm":   "armv7",
			"arm64": "aarch64",
		},
	},
	K9s: {
		"darwin": {
			"amd64": "x86_64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm":   "arm",
			"arm64": "arm64",
		},
	},
	UPX: {
		"linux": {
			"amd64": "amd64",
			"arm":   "arm",
			"arm64": "arm64",
		},
	},
	Xh: {
		"darwin": {
			"amd64": "x86_64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm":   "arm",
		},
	},
}
