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

#### run docker compose file

```bash
docker compose up
```

#### monitor NATS servers

```bash
watch -d 'nats server ls --context=local_cluster'
```