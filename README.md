# LastFM
## LastFM API Wrapper for Go 
[![Build Status](https://travis-ci.com/ptrkrlsrd/lastfm.svg?token=EC6EZTgzr1WN8mybj2yE&branch=master)](https://travis-ci.com/ptrkrlsrd/lastfm)

## Installation
Add the following environment variables to `~/.profile`:
```
export LASTFM_API="<API>"
export LASTFM_KEY="<KEY>"
```

## Usage
```
LastFM

Usage:
  lastfm [command]

Available Commands:
  help        Help about any command
  info        Get info about an artist or an album
  scrobbles   Get an users top scrobbles
  search      Search the API
  similar     Get similar artists or albums
  top         

Flags:
      --config string   config file (default is $HOME/.lastfm.yml)
  -h, --help            help for lastfm

Use "lastfm [command] --help" for more information about a command.
```
