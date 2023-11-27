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
#### configure ```nats-server-N.conf``` files

#### run 3 local NATS servers in cluster

```bash
nats-server -c nats-server-1.conf
```
```bash
nats-server -c nats-server-2.conf
```
```bash
nats-server -c nats-server-3.conf
```

#### configure ```user``` and ```password``` in ~/.config/nats/context/local_cluster.json

```json
{
  "description": "",
  "url": "nats://127.0.0.1:4222,nats://127.0.0.1:4223,nats://127.0.0.1:4224",
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

#### monitor NATS cluster

```bash
watch -d 'nats server ls'
```

#### download nats-box docker image

```bash
docker pull natsio/nats-box
```

#### run nats-box container

```bash
docker run rm -it natsio/nats-box
```