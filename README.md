![Cover](https://capsule-render.vercel.app/api?type=waving&height=300&color=gradient&text=JetBrains%20Hacker&desc=Build%20your%20own%20license!&descAlignY=60&descAlign=70)

---

# JetBrains Hacker

JetBrains Hacker is a tool that can customize your JetBrains IDEs license.

I create it to activate my common IDEs, and hope that it can be useful to others.

## Features

- 💪 Useful - Offline activation code and online license server.
- 🔨 Customize - Any information as you want, just make your own license.

## Usage

1. Get the `ja-netfilter` from the Internet. This software is worked based on the `ja-netfilter`.

2. Add -javaagent:/path/to/ja-netfilter.jar=`${app}` to your vmoptions (manual or auto). Note that `${app}` is a parameter to `ja-netfilter` that specifies the location of the `config-${app}` and `plugins-${app}` folders, if empty the `config` and `plugins` folders will be used by default.


> for Java version 17+, you need add these 2 lines to your vmoptions file: (for manual, without any whitespace chars)

```vmoptions
--add-opens=java.base/jdk.internal.org.objectweb.asm=ALL-UNNAMED
--add-opens=java.base/jdk.internal.org.objectweb.asm.tree=ALL-UNNAMED
```

3. Get the `jetbrains hacker`. See [Installation](#installation) or [Build](#build) on bellowed.

4. There are two ways to use the `jetbrains hacker`.

### Activation Code

> This approach is used in offline scenario.

1. Run `jetbrains_hacker build-cert` to build needed certificates.
2. Run `jetbrains_hacker build-config --type {power|url|dns}` to generate the corresponding configurations. Then copy the generated configurations into your `ja-netfilter` configuration files.
3. Run `jetbrains_hacker generate-license --licenseId ${licenseId} --name ${name} --user ${user} --email ${email} --time {2999-01-02}`. Or simplest of all, you can just use `jetbrain_hacker generate-license`.
4. Use the activation code in the `Activation Code` window.
5. Don't care about the activation time, it is a fallback license and will not expire.

Enjoy it~

### License Server

> This approach is used in online scenario.

1. Run `jetbrains_hacker run-server` to run a online license server. By default the server will run on `:80`. If you want to change the address, use the `--addr` argument.
2. Go to `${your server address}/config/{power|url|dns}` to get the corresponding configurations. Then copy the generated configurations into your `ja-netfilter` configuration files.
3. Type your server address in the `License Server` window.
4. Don't care about the activation time, it is a fallback license and will not expire.

Enjoy it~

## Installation

You can choose to use the pre-compiled binary or build it by yourself.

### Release Binaries

[Available for download in releases](https://github.com/LovesAsuna/jetbrains_hacker/releases)

Binaries available for:

#### Linux

- jetbrain-hacker_linux_amd64 (linux musl statically linked)
- jetbrains_hacker-linux-aarch64.tar.gz (linux on 64 bit arm)

All contain a single binary file

#### macOS

- jetbrains_hacker-mac.tar.gz (arm64)
- jetbrain-hacker_darwin_amd64 (intel x86)

#### Windows

- jetbrain-hacker_windows_amd64.exe (single 64bit binary)

## Build

### Requirements

- Minimum supported `go` version: `1.24`
  - See [Install Go](https://go.dev/dl/)

- To build needed dependency (run `go mod tidy`)

### Go Install

The simplest way to start playing around with `jetbrains_hacker` is to have `go` build and install it with `go build .`.
