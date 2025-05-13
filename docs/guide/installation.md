# Installing mu

You can install mu by downloading a release from GitHub or by using our installer script.

Choose your adventure below.

## Direct Download

You can download the binary from the [mu releases page](https://github.com/micro/mu/releases) on GitHub and add to your `$PATH`.

The mu_VERSION_checksums.txt file contains the SHA-256 checksum for each file.

## Installer Script

We also have an [install script](https://github.com/micro/mu/blob/main/docs/public/install.sh) which is very useful in scenarios like CI.

By default, it installs on the `./bin` directory relative to the working directory:

```bash
sh -c "$(curl --location https://micro.github.io/mu/install.sh)" -- -d
```

It is possible to override the installation directory with the -b parameter. On Linux, common choices are `~/.local/bin` and `~/bin` to install for the current user or `/usr/local/bin` to install for all users:

```bash
sh -c "$(curl --location https://micro.github.io/mu/install.sh)" -- -d -b ~/.local/bin
```

!> On macOS and Windows, ~/.local/bin and ~/bin are not added to $PATH by default.

By default, it installs the latest version available. You can also specify a tag ([available in releases](https://github.com/micro/mu/releases)) to install a specific version:

```bash
sh -c "$(curl --location https://micro.github.io/mu/install.sh)" -- -d v0.1.0
```