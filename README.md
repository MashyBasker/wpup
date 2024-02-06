# wpup

I like to curate wallpapers, but they take up a lot of disk space. So I started storing them in my discord server and created this tool to automate the whole process

## Quickstart

You need to have the Go compiler and Git installed


### Installation

```console
$ git clone https://github.com/MashyBasker/wpup
$ cd wpup
$ go build
$ cp wpup ~/.local/bin
```

### Setting the discord webhook URL

To get the discord channel webhook URL. [Follow this](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-WebhooksME)

Copy the webhook URL of the discord channel you want to send the files and store it in your `.bashrc`, `.zshrc`, etc. file

```bash
export DISCORD_WEBHOOK="--COMPLETE DISCORD WEBHOOOK URL--"
```

### Usage

```console
$ wpup /path/to/directory
```
