package tables

var Arch = map[string]map[string]map[string]string{
	Bat: {
		"darwin": {
			"amd64": "x86_64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm":   "arm",
		},
	},
	Bottom: {
		"darwin": {
			"amd64": "x86_64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm":   "armv7",
		},
	},
	Cloudflared: {
		"darwin": {
			"amd64": "amd64",
		},
		"linux": {
			"amd64": "amd64",
			"arm":   "arm",
		},
	},
	DockerCompose: {
		"linux": {
			"amd64": "x86_64",
			"arm":   "armv7",
		},
	},
	K9s: {
		"darwin": {
			"amd64": "x86_64",
		},
		"linux": {
			"amd64": "x86_64",
			"arm":   "arm",
		},
	},
	UPX: {
		"linux": {
			"amd64": "amd64",
			"arm":   "arm",
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
