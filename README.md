staticd
=======

Download statically linked binaries from GitHub.

## Install

To install `staticd`, go to the [releases page][releases] and follow the instructions below:

- Download the binary for your OS and architecture.
- Make the binary executable:

  ```
  chmod +x staticd-<os>-<arch>
  ```

From there, you can either opt for the home directory installation (recommended) or the system-wide installation.

### Home-based installation

- Create the `~/.local/bin` directory if it does not exist:
  ```
  mkdir -p ~/.local/bin
  ```

- Move the binary to `~/.local/bin`:
  ```
  mv staticd-<os>-<arch> ~/.local/bin/staticd
  ```

- Ensure `$HOME/.local/bin` is in your `$PATH`. Add the following line to your `~/.bashrc` or equivalent shell configuration file:
  ```
  export PATH="$HOME/.local/bin:$PATH"
  ```

  P.s.: use `$HOME` instead of `~/` as [tilde does not expand in quotes][SC2088].

### System-wide installation

- Move the binary to `/usr/local/bin`:
  ```
  sudo mv staticd-<os>-<arch> /usr/local/bin/staticd
  ```

After installation, you can run the tool by calling the `staticd` command.

## Usage

```
$ staticd <tool>
```

It's important to note that:

- Calling it directly as a non-root user puts the binaries into `~/.local/bin`.
- Calling it using `sudo` puts the binaries into `/usr/local/bin`.

## Tools

- `bat`: [sharkdp/bat][bat], a cat(1) clone with wings.
- `btm`: [ClementTsang/bottom][btm], yet another cross-platform graphical process/system monitor.
- `cloudflared`: [cloudflare/cloudflared][cloudflared], Argo Tunnel client.
- `flyctl`: [superfly/flyctl][flyctl], command line tools for fly.io services
- `k9s`: [derailed/k9s][k9s], Kubernetes CLI to manage your clusters in style.
- `kubectx`: [ahmetb/kubectx][kubectx], faster way to switch between clusters in kubectl
- `rg`: [BurntSushi/ripgrep][rg], recursively searches directories for a regex pattern.
- `shellcheck`: [koalaman/shellcheck][shellcheck], a static analysis tool for shell scripts.
- `upx`: [upx/upx][upx], the Ultimate Packer for eXecutables.
- `uv`: [astral-sh/uv][uv], an extremely fast Python package and project manager.
- `xh`: [ducaale/xh][xh], friendly and fast tool for sending HTTP requests.
- `yj`: [sclevine/yj][yj], convert between YAML, TOML, JSON, and HCL.

[bat]: https://github.com/sharkdp/bat
[btm]: https://github.com/ClementTsang/bottom
[cloudflared]: https://github.com/cloudflare/cloudflared
[flyctl]: https://github.com/superfly/flyctl
[k9s]: https://github.com/derailed/k9s
[kubectx]: https://github.com/ahmetb/kubectx
[releases]: https://github.com/myhro/staticd/releases
[rg]: https://github.com/BurntSushi/ripgrep
[SC2088]: https://www.shellcheck.net/wiki/SC2088
[shellcheck]: https://github.com/koalaman/shellcheck
[upx]: https://github.com/upx/upx
[uv]: https://github.com/astral-sh/uv
[xh]: https://github.com/ducaale/xh
[yj]: https://github.com/sclevine/yj
