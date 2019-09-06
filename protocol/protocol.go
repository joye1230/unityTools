package protocol

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

func SetField(pd *PData, fieldName string, data interface{}) {
	// 自动初始化
	if !pd.inited {
		InitPData(pd)
	}
	addFieldOp(pd, fieldName, "set")
	pd.Data[fieldName] = data
}

// 多次调用需要自己处理之前的数据，操作是无状态的
func UpdateField(pd *PData, fieldName string, data map[string]interface{}) {
	// 自动初始化
	if !pd.inited {
		InitPData(pd)
	}
	addFieldOp(pd, fieldName, "update")
	if pd.Data[fieldName] == nil {
		pd.Data[fieldName] = make(map[string]interface{})
	}
	for key, value := range data {
		pd.Data[fieldName].(map[string]interface{})[key] = value
	}
}

func RemoveField(pd *PData, fieldName string) {
	// 自动初始化
	if !pd.inited {
		InitPData(pd)
	}
	addFieldOp(pd, fieldName, "remove")
}

// 多次调用会累加，操作有状态
func RemoveFieldItems(pd *PData, fieldName string, keys []string) {
	RemoveField(pd, fieldName)
	rmkeys, ok := pd.Rmkey[fieldName]
	if ok {
		rmkeys = append(rmkeys.([]string), keys...)
	} else {
		rmkeys = keys
	}
	pd.Rmkey[fieldName] = rmkeys
}

func addFieldOp(pd *PData, fieldName, opType string) {
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
