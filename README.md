# Distributed Unique ID Generation - GRPC Golang

The algorithm of Unique id generation is Snowflake. You can refer this repo to get the detail explanation for the algorithm

* https://github.com/bwmarrin/snowflake

## How to use

Firstly, we must compile this module to binary program by running this command

```
cd distributed-uid
go build .
```

We will use Docker to build a nginx as a load balancing.

```
docker build . -t uid:1.0
docker run -d --name uid -p 8000:8000 uid:1.0
```

Now, the Unique ID Generation is available on port `8000`
