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

create a new stream 'ORDERS' and add subjects wildcard

```bash
nats stream add ORDERS --subjects 'orders.*' --server=localhost:4222
```

create producers for the subjects
```bash
nats pub --server=localhost:4222 "orders.us" --count=-1 --sleep=5s "US Order #{{Count}} at {{Time}}"
```

```bash
nats pub --server=localhost:4222 "orders.eu" --count=-1 --sleep=5s "EU Order #{{Count}} at {{Time}}"
```

monitor stream

```bash
watch -d 'nats stream ls --server=localhost:4222'
```

monitor consumers for the stream

```bash
watch -d 'nats consumer ls ORDERS  --server=localhost:4222'
```

run ```subsciber/main.go``` to test program

clean up consumers

```bash
nats consumer rm ORDER my-consumer-1 --server=localhost:4222
```

clean up streams

```bash
nats stream rm ORDER --server=localhost:4222
```