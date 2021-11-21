package tables

var OS = map[string]map[string]map[string]string{
	Bat: {
		"darwin": {
			"amd64": "apple-darwin",
		},
		"linux": {
			"amd64": "unknown-linux-gnu",
			"arm":   "unknown-linux-gnueabihf",
		},
	},
	Bottom: {
		"darwin": {
			"amd64": "apple-darwin",
		},
		"linux": {
			"amd64": "unknown-linux-gnu",
			"arm":   "unknown-linux-gnueabihf",
		},
	},
	Cloudflared: {
		"darwin": {
			"amd64": "darwin",
		},
		"linux": {
			"amd64": "linux",
			"arm":   "linux",
		},
	},
	DockerCompose: {
		"linux": {
			"amd64": "linux",
			"arm":   "linux",
		},
	},
	K9s: {
		"darwin": {
			"amd64": "Darwin",
		},
		"linux": {
			"amd64": "Linux",
			"arm":   "Linux",
		},
	},
	Xh: {
		"darwin": {
			"amd64": "apple-darwin",
		},
		"linux": {
			"amd64": "unknown-linux-musl",
			"arm":   "unknown-linux-gnueabihf",
		},
	},
}
