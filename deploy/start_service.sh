#!/bin/bash
/data/src/distributed-uid --port=8001 --node=1 & 
/data/src/distributed-uid --port=8002 --node=2 &
/data/src/distributed-uid --port=8003 --node=3 & 
/data/src/distributed-uid --port=8004 --node=4