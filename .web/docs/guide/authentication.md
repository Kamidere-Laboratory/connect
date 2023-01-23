# Authentication

_How do servers and players authenticate with the Connect Network?_

## Unique Server Endpoint Name

Connect has the concept of unique endpoint names to identify your server even after restarts.
The Connect Plugin uses a token file to authenticate that you own an endpoint name in the Connect Network.

```json plugins/connect/token.json
{"token":"T-ozinikukmabrpyzogjjl"}
```

The token and endpoint name have a direct relationship, and you can use the same token for multiple endpoint names
as loong as they are available for use.

### Trouble Shooting

If you see an error message like this in your console:

```sh
[connect]: Connection error with WatchService:
  java.net.ProtocolException: Expected HTTP 101 response
    but was '401 Unauthorized'
```

It means that your endpoint name is already taken by another server.
If you lose your `token.json` file your endpoint name is lost, and you will have to use a new name.

## Player Authentication

Players authenticate their Minecraft account when joining the Connect Network.
Connect will only route successfully authenticated players through encrypted tunnel
connections to your server.

> Cracked players are not supported.
> Feel free to request this feature on our [Discord](https://minekube.com/discord) to be added sooner.