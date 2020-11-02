package snowflake

import (
	"go.uber.org/zap"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(starttime string,machineID int64) (err error) {
	var st time.Time
	st ,err = time.Parse("2006-01-02",starttime)
	if err != nil{
		zap.L().Error("Generate ID error")
		return err
	}
	sf.Epoch = st.UnixNano()/1000000
	node,err = sf.NewNode(machineID)
	return nil
}

func GenID() int64 {
	return node.Generate().Int64()
}
