package snowflake

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

type ID int64

// Int64 returns an int64 of the snowflake ID
func (f ID) Int64() int64 {
	return int64(f)
}

type ISnowFlake interface {
	Generate() ID
}

type SnowFlake struct {
	mu    sync.Mutex
	epoch time.Time
	time  int64
	node  int64
	step  int64

	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift int64
	nodeShift int64
}

func NewSnowFlake(node int64, constant *Constant) (*SnowFlake, error) {
	n := SnowFlake{}
	n.node = node
	n.nodeMax = -1 ^ (-1 << constant.NodeBits)
	n.nodeMask = n.nodeMax << int64(constant.StepBits)
	n.stepMask = -1 ^ (-1 << constant.StepBits)
	n.timeShift = int64(constant.StepBits) + int64(constant.NodeBits)
	n.nodeShift = int64(constant.StepBits)

	if n.node < 0 || n.node > n.nodeMax {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(n.nodeMax, 10))
	}

	var curTime = time.Now()
	n.epoch = curTime.Add(time.Unix(constant.Epoch/1000, (constant.Epoch%1000)*1000000).Sub(curTime))

	InfoLogger.Printf("Successfully initialize new UUID Distributed SnowFlake's Generator with NodeID: %d", n.node)
	return &n, nil
}

func (n *SnowFlake) Generate() ID {
	n.mu.Lock()

	now := time.Since(n.epoch).Milliseconds()
	if now == n.time {
		n.step = (n.step + 1) & n.stepMask
		if n.step == 0 {
			for now <= n.time {
				now = time.Since(n.epoch).Milliseconds()
			}
		}
	} else {
		n.step = 1
	}

	n.time = now
	var res = (now << n.timeShift) | (n.node << n.nodeShift) | (n.step)

	n.mu.Unlock()
	InfoLogger.Printf("Successful generate %d", res)
	return ID(res)
}
