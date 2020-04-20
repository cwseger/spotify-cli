# spotify-cli
A simple CLI to interact with Spotify.

## Installation
### Download:
```
go get github.com/cwseger/spotify-cli
```

### Create Spotify Developer Account
In order to use this cli, you'll need a spotify developer account.
To create an account go here: https://developer.spotify.com/dashboard/.
After creating an account you will need to create a new app in order to get credentials necessary for this cli.
When you've created a new application you will have access to your `clientId` and `clientSecret`.

### Add Secrets File
Add a file called client-secrets.json in the root of the directory. It should look like this:
```
{
    "clientId": "<your client id here>",
    "clientSecret": "<your client secret here>"
}
```

### Build/Install
To be able to issue commands you need to either install or build.
```
go install
```
or
```
go build
```

## Issue commands
Via `go install`:
```
spotify-cli [command]
```
Via `go build`:
```
./spotify-cli [command]
```