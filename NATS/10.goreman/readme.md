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

#### install goreman

```bash
go install github.com/mattn/goreman@latest
```

#### crate ```Procfile```

```
nats1: nats-server -c nats-server-1.conf
nats2: nats-server -c nats-server-2.conf
nats3: nats-server -c nats-server-3.conf
```

#### run cluster with goreman

```bash
goreman start
```

#### stop servers with goreman

```bash
goreman run stop nats-1
```