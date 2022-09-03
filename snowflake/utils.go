package snowflake

import (
	"os"
	"strconv"
)

type Constant struct {
	Epoch    int64
	NodeBits uint8
	StepBits uint8

	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift uint8
	nodeShift uint8
}

func getEnvInt64(envName string) int64 {
	var res, err = strconv.ParseInt(os.Getenv(envName), 10, 64)
	if err != nil {
		return 0
	}
	return res
}

func getEnvUint8(envName string) uint8 {
	var res, err = strconv.ParseUint(os.Getenv(envName), 10, 8)
	if err != nil {
		return 0
	}
	return uint8(res)
}

func InitConstant() (*Constant, error) {
	var constant = Constant{}
	var err error

	constant.Epoch = getEnvInt64("EPOCH")
	constant.NodeBits = getEnvUint8("NODE_BITS")

	return &constant, err
}
