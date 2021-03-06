package tables

var OS = map[string]map[string]map[string]string{
	Bat: {
		"darwin": {
			"amd64": "apple-darwin",
		},
		"linux": {
			"amd64": "unknown-linux-gnu",
			"arm":   "unknown-linux-gnueabihf",
			"arm64": "unknown-linux-gnu",
		},
	},
	Bottom: {
		"darwin": {
			"amd64": "apple-darwin",
		},
		"linux": {
			"amd64": "unknown-linux-gnu",
			"arm":   "unknown-linux-gnueabihf",
			"arm64": "unknown-linux-gnu",
		},
	},
	Cloudflared: {
		"darwin": {
			"amd64": "darwin",
		},
		"linux": {
			"amd64": "linux",
			"arm":   "linux",
			"arm64": "linux",
		},
	},
	DockerCompose: {
		"linux": {
			"amd64": "linux",
			"arm":   "linux",
			"arm64": "linux",
		},
	},
	K9s: {
		"darwin": {
			"amd64": "Darwin",
		},
		"linux": {
			"amd64": "Linux",
			"arm":   "Linux",
			"arm64": "Linux",
		},
	},
	UPX: {
		"linux": {
			"amd64": "linux",
			"arm":   "linux",
			"arm64": "linux",
		},
	},
	Xh: {
		"darwin": {
			"amd64": "apple-darwin",
		},
		"linux": {
			"amd64": "unknown-linux-musl",
			"arm":   "unknown-linux-gnueabihf",
			"arm64": "unknown-linux-musl",
		},
	},
	Yj: {
		"darwin": {
			"amd64": "macos",
		},
		"linux": {
			"amd64": "linux",
			"arm":   "linux",
			"arm64": "linux",
		},
	},
}
