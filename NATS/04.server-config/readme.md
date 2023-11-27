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

#### run local NATS server
```bash
nats-server -c nats-auth.conf
```

you have to see following output:

```bash
[73871] 2023/11/05 18:12:05.376223 [INF] Starting nats-server
[73871] 2023/11/05 18:12:05.376366 [INF]   Version:  2.10.4
[73871] 2023/11/05 18:12:05.376371 [INF]   Git:      [not set]
[73871] 2023/11/05 18:12:05.376374 [INF]   Name:     NDXX2C3PLL6DURUODZ3RELETLCSSOBKX3JKPQV6VCA7THWL4RD4K3LI5
[73871] 2023/11/05 18:12:05.376379 [INF]   ID:       NDXX2C3PLL6DURUODZ3RELETLCSSOBKX3JKPQV6VCA7THWL4RD4K3LI5
[73871] 2023/11/05 18:12:05.376397 [INF] Using configuration file: nats-auth.conf
[73871] 2023/11/05 18:12:05.376783 [INF] Listening for client connections on 0.0.0.0:4222
[73871] 2023/11/05 18:12:05.377041 [INF] Server is ready
```

to configure secured password use following command to generate encrypted password

```bash
nats server passwd
```

it will prompt you a request to enter your password

```bash
[nats_development] ? Enter password [? for help]
```

and then to re-enter your password

```bash
[nats_development] ? Reenter password [? for help]
```

after that you will receive encrypted password:

```bash
$2a$11$0YjzDPlOmRKWk/HqYDOdV.iRuuIb4FTKj8ZD8CPu0F7INeAtZSGtO
```

update ```nats-auth.conf``` file with the encrypted password:

```json
authorization: {
	users: [{
		user: "sa",
		password: "$2a$11$FmswOz/JvDxXeT35g42O3OQa1sFwzgoeNPeqbFOfmnxt1WknafN7e"
	}]
}
```

reload nats server

```bash
nats-server --signal reload
```