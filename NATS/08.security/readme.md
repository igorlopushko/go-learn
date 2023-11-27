Before start do following steps

#### install NATS via brew

```bash
brew tap nats-io/nats-tools
brew install nats-io/nats-tools/nats
```

#### install NATS servcer via brew

```bash
brew install nats-server
```
#### configure ```nats-server.conf``` file

```conf
jetstream: {
	
}

authorization: {
	users: [{
		user: "admin",
		password: "password"
	}]	
}
```

#### run local NATS server

```bash
nats-server -c nats-server.conf
```

#### configure ```user``` and ```password``` in ~/.config/nats/context/local.json

```json
{
  "description": "",
  "url": "nats://127.0.0.1:4222",
  "socks_proxy": "",
  "token": "",
  "user": "admin",
  "password": "password",
  "creds": "",
  "nkey": "",
  "cert": "",
  "key": "",
  "ca": "",
  "nsc": "",
  "jetstream_domain": "",
  "jetstream_api_prefix": "",
  "jetstream_event_prefix": "",
  "inbox_prefix": "",
  "user_jwt": "",
  "color_scheme": ""
}
```

#### check NATS context

```bash
nats context ls
```

#### select NATS context

```bash
nats context select local
```