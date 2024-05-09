Below is a basic configuration of this server.

## `main.go`  

```go
var bsky_host = "https://bsky.social"
var bsky_username = "jordanreger.com"

var mastodon_host = "mastodon.sdf.org"
```

## `.env`  
```
smtpuser=user@email.com
smtppass={smtp password or app password}
bskypass={[bluesky app password](https://bsky.app/settings/app-passwords)}
mastodontoken={[access token](#access-token)}
```

### Access Token
1. In your Mastodon settings, go to the *Development* tab
2. Hit *New application*
3. Fill in the fields as necessary, but select `read:notifications` and `write:notifications` (uncheck everything else)
4. Submit
5. Copy *Your access token*
6. Paste it in your `.env`
