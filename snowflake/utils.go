package snowflake

import (
	"os"
	"strconv"
)

type Constant struct {
	Epoch    int64
	NodeBits uint8
	StepBits uint8

	// nodeMax   int64
	// nodeMask  int64
	// stepMask  int64
	// timeShift uint8
	// nodeShift uint8
}

func GetEnvInt64(envName string) int64 {
	var res, err = strconv.ParseInt(os.Getenv(envName), 10, 64)
	if err != nil {
		ErrorLogger.Fatal(err)
		return 0
	}
	return res
}

func GetEnvUint8(envName string) uint8 {
	var res, err = strconv.ParseUint(os.Getenv(envName), 10, 8)
	if err != nil {
		ErrorLogger.Fatal(err)
		return 0
	}
	return uint8(res)
}

func InitConstant() (*Constant, error) {
	var constant = Constant{}
	var err error

	constant.Epoch = GetEnvInt64("EPOCH")
	constant.NodeBits = GetEnvUint8("NODE_BITS")
	constant.StepBits = GetEnvUint8("STEP_BITS")

	return &constant, err
}
