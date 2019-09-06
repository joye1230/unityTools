package protocol

import (
	"github.com/name5566/leaf/gate"
	"server/protocol"
	"server/define"
)

type BaseError struct {
	Status   int    `json:"status"`
	ErrorMsg string `json:"errorMsg"`
}
type BaseResponse struct {
	Echo       interface{}    `json:"echo"`
	Data       protocol.PData `json:"data"`
	ModuleData interface{} `json:"moduleData"`
	BaseError
	Uid string `json:"uid"`
}

func (p *BaseError) SetErrorCode(errorCode int) {
	p.Status = errorCode
	p.ErrorMsg = define.ErrorCode[errorCode]
}

func (p *BaseResponse) Init(echo interface{}, uid string) {
	protocol.InitPData(&p.Data)
	p.Echo = echo
	p.Uid = uid
}

func (p *BaseResponse) WriteMsg(agent gate.Agent, msg interface{}) {
	if agent == nil {
		return
	}
	if msg == nil {
		return
	}
	agent.WriteMsg(msg)
}

type PData struct {
	Op     map[string][]string    `json:"op"`
	Data   map[string]interface{} `json:"data"`
	Rmkey  map[string]interface{} `json:"rmkey"`
	inited bool
}

func InitPData(pd *PData) {
	pd.Op = make(map[string][]string)
	pd.Data = make(map[string]interface{})
	pd.Rmkey = make(map[string]interface{})
	pd.inited = true
}

func (p *BaseResponse) SetField(pd *PData, fieldName string, data interface{}) {
	// 自动初始化
	if !pd.inited {
		InitPData(pd)
	}
	p.addFieldOp(pd, fieldName, "set")
	pd.Data[fieldName] = data
}

// 多次调用需要自己处理之前的数据，操作是无状态的
func (p *BaseResponse) UpdateField(pd *PData, fieldName string, data interface{}) {
	// 自动初始化
	if !pd.inited {
		InitPData(pd)
	}
	p.addFieldOp(pd, fieldName, "update")
	pd.Data[fieldName] = data
}

func (p *BaseResponse) RemoveField(pd *PData, fieldName string) {
	// 自动初始化
	if !pd.inited {
		InitPData(pd)
	}
	p.addFieldOp(pd, fieldName, "remove")
}

// 多次调用会累加，操作有状态
func (p *BaseResponse) RemoveFieldItems(pd *PData, fieldName string, keys []string) {
	p.RemoveField(pd, fieldName)
	rmkeys, ok := pd.Rmkey[fieldName]
	if ok {
		rmkeys = append(rmkeys.([]string), keys...)
	} else {
		rmkeys = keys
	}
	pd.Rmkey[fieldName] = rmkeys
}

func (p *BaseResponse) addFieldOp(pd *PData, fieldName, opType string) {
	_, ok := pd.Op[fieldName]
	if !ok {
		pd.Op[fieldName] = []string{opType}
	} else {
		var hasSet bool
		for _, ownType := range pd.Op[fieldName] {
			if opType == ownType {
				hasSet = true
				break
			}
		}
		if !hasSet {
			pd.Op[fieldName] = append(pd.Op[fieldName], opType)
		}
	}
}
