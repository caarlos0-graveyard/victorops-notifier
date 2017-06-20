# victorops-notifier

An unofficial VictorOps desktop notifier for macOS.

![screenshot](https://user-images.githubusercontent.com/245435/27312110-5d166c46-553c-11e7-9133-49c70df34f9a.png)


## Install

```console
$ brew install caarlos0/tap/victorops-notifier
```

## Usage

The idea is to keep it running. You can do that with any shell loop or
by putting it in the crontab.

```crontab
PATH="/usr/local/bin/"

* * * * * victorops-notifier --client client-name --id api-id --key api-key
```

You can get the client name from the URL when you're logged in, something like
`https://portal.victorops.com/client/<< CLIENT NAME HERE >>`.

As for the API keys, you can get them by going to **Settings** -> **API**.
