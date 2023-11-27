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
configure ```nats-server.conf``` file

```json
jetstream: {

}
```

#### run local NATS server
```bash
nats-server -c nats-server.conf
```

you have to see following output:

```bash
[75016] 2023/11/05 18:30:40.620674 [INF] Starting nats-server
[75016] 2023/11/05 18:30:40.620960 [INF]   Version:  2.10.4
[75016] 2023/11/05 18:30:40.620968 [INF]   Git:      [not set]
[75016] 2023/11/05 18:30:40.620975 [INF]   Name:     NBQTQWBRBWINH55C3XCJGM7E5ESHDCA7GPGMR5KO7IE2BX63RQHBEQPQ
[75016] 2023/11/05 18:30:40.620982 [INF]   Node:     K39rD2no
[75016] 2023/11/05 18:30:40.620988 [INF]   ID:       NBQTQWBRBWINH55C3XCJGM7E5ESHDCA7GPGMR5KO7IE2BX63RQHBEQPQ
[75016] 2023/11/05 18:30:40.621003 [INF] Using configuration file: nats-auth.conf
[75016] 2023/11/05 18:30:40.621370 [INF] Starting JetStream
[75016] 2023/11/05 18:30:40.621892 [INF]     _ ___ _____ ___ _____ ___ ___   _   __  __
[75016] 2023/11/05 18:30:40.621905 [INF]  _ | | __|_   _/ __|_   _| _ \ __| /_\ |  \/  |
[75016] 2023/11/05 18:30:40.621913 [INF] | || | _|  | | \__ \ | | |   / _| / _ \| |\/| |
[75016] 2023/11/05 18:30:40.621922 [INF]  \__/|___| |_| |___/ |_| |_|_\___/_/ \_\_|  |_|
[75016] 2023/11/05 18:30:40.621928 [INF]
[75016] 2023/11/05 18:30:40.621935 [INF]          https://docs.nats.io/jetstream
[75016] 2023/11/05 18:30:40.621941 [INF]
[75016] 2023/11/05 18:30:40.621947 [INF] ---------------- JETSTREAM ----------------
[75016] 2023/11/05 18:30:40.621955 [INF]   Max Memory:      24.00 GB
[75016] 2023/11/05 18:30:40.621962 [INF]   Max Storage:     475.15 GB
[75016] 2023/11/05 18:30:40.621969 [INF]   Store Directory: "/var/folders/vw/67wc2mb93kg4ht5v9bdvxr940000gn/T/nats/jetstream"
[75016] 2023/11/05 18:30:40.621980 [INF] -------------------------------------------
[75016] 2023/11/05 18:30:40.622739 [INF] Listening for client connections on 0.0.0.0:4222
[75016] 2023/11/05 18:30:40.623047 [INF] Server is ready
```

#### monitor streams

```bash
watch -d 'nats stream ls -a --server=localhost:4222'
```

#### monitor key-value buckets

```bash
watch -d 'nats kv ls --server=localhost:4222'
```

#### create a new key-value bucket

```bash
nats kv add music --server=localhost:4222
```

#### monitor newly created key-value bucket

```bash
nats kv watch music --server=localhost:4222
```