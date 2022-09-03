package snowflake

import (
	"sync"
	"time"
)

type Node struct {
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

type ID int64

// func main() {
// 	fmt.Println(os.Getenv("EPOCH"))
// }
