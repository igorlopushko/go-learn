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
nats-server
```

you have to see following output:

```bash
[15234] 2023/11/04 21:36:43.628853 [INF] Starting nats-server
[15234] 2023/11/04 21:36:43.631851 [INF]   Version:  2.10.4
[15234] 2023/11/04 21:36:43.631856 [INF]   Git:      [not set]
[15234] 2023/11/04 21:36:43.631860 [INF]   Name:     NAYZXZ2CSQFR4ZZGFJWD4GH5ZRSBDDRVD6TZVARUEFEREZ4EWWYCQD72
[15234] 2023/11/04 21:36:43.631864 [INF]   ID:       NAYZXZ2CSQFR4ZZGFJWD4GH5ZRSBDDRVD6TZVARUEFEREZ4EWWYCQD72
[15234] 2023/11/04 21:36:43.634592 [INF] Listening for client connections on 0.0.0.0:4222
[15234] 2023/11/04 21:36:43.634766 [INF] Server is ready
```