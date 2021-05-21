package umios

type PushTemplate struct {
	// 必填，应用唯一标识
	Appkey string `json:"appkey"`
	// 必填，时间戳，10位或者13位均可，时间戳有效期为10分钟
	Timestamp string `json:"timestamp"`
	// 必填，消息发送类型,其值可以为:
	//   unicast-单播
	//   listcast-列播，要求不超过500个device_token
	//   filecast-文件播，多个device_token可通过文件形式批量发送
	//   broadcast-广播
	//   groupcast-组播，按照filter筛选用户群, 请参照filter参数
	//   customizedcast，通过alias进行推送，包括以下两种case:
	//     - alias: 对单个或者多个alias进行推送
	//     - file_id: 将alias存放到文件后，根据file_id来推送
	Type string `json:"type"`
	// 当type=unicast时, 必填, 表示指定的单个设备
	//当type=listcast时, 必填, 要求不超过500个, 以英文逗号分隔
	DeviceTokens string `json:"device_tokens,omitempty"`
	// 当type=customizedcast时, 必填
	// alias的类型, alias_type可由开发者自定义, 开发者在SDK中
	// 调用setAlias(alias, alias_type)时所设置的alias_type
	AliasType string `json:"alias_type,omitempty"`
	// 当type=customizedcast时, 选填(此参数和file_id二选一)
	// 开发者填写自己的alias, 要求不超过500个alias, 多个alias以英文逗号间隔
	// 在SDK中调用setAlias(alias, alias_type)时所设置的alias
	Alias string `json:"alias,omitempty"`
	// 当type=filecast时，必填，file内容为多条device_token，以回车符分割
	// 当type=customizedcast时，选填(此参数和alias二选一)
	//   file内容为多条alias，以回车符分隔。注意同一个文件内的alias所对应
	//   的alias_type必须和接口参数alias_type一致。
	// 使用文件播需要先调用文件上传接口获取file_id，参照"文件上传"
	FileID string `json:"file_id,omitempty"`
	// 当type=groupcast时，必填，用户筛选条件，如用户标签、渠道等，参考附录G。
	// filter的内容长度最大为3000B）
	Filter interface{} `json:"filter,omitempty"`
	// 必填，JSON格式，具体消息内容(iOS最大为2012B)
	Payload Payload `json:"payload"`
	// 可选，发送策略
	Policy *Policy `json:"policy,omitempty"`
	// 可选，正式/测试模式。默认为true
	// 测试模式只对“广播”、“组播”类消息生效，其他类型的消息任务（如“文件播”）不会走测试模式
	// 测试模式只会将消息发给测试设备。测试设备需要到web上添加。
	ProductionMode string `json:"production_mode,omitempty"`
	// 可选，发送消息描述，建议填写。
	Description string `json:"description,omitempty"`
}

type Payload struct {
	// 必填，严格按照APNs定义来填写
	Aps Aps `json:"aps"`
	// 可选，JSON格式，用户自定义key-value
	Extra map[string]interface{} `json:"extra,omitempty"`
}

type Aps struct {
	// 当content-available=1时(静默推送)，别填不然就不是静默通知了, 否则必填。
	Alert *Alert `json:"alert,omitempty"`
	// 应用角标,可选
	// 如果不填，表示不改变角标数字，否则把角标数字改为指定的数字；为 0 表示清除,当content-available=1时(静默推送) 别填不然就不是静默通知了
	Badge string `json:"badge,omitempty"`
	// 通知提示声音或警告通知,可选,当content-available=1时(静默推送) 别填不然就不是静默通知了
	Sound string `json:"sound,omitempty"`
	// 可选，1代表静默推送
	ContentAvailable int `json:"content-available,omitempty"`
	// 可选，注意: ios8才支持该字段。
	Category string `json:"category,omitempty"`
}

type Alert struct {
	Title    string `json:"title"`    // 通知标题
	SubTitle string `json:"subtitle"` // 子标题
	Body     string `json:"body"`     //内容
}

type Policy struct {
	// 可选，定时发送时间，若不填写表示立即发送。
	// 定时发送时间不能小于当前时间
	// 格式: "yyyy-MM-dd HH:mm:ss"。
	// 注意，start_time只对任务生效。
	StartTime string `json:"start_time,omitempty"`
	// 可选，消息过期时间，其值不可小于发送时间或者
	// start_time(如果填写了的话),
	// 如果不填写此参数，默认为3天后过期。格式同start_time
	ExpireTime string `json:"expire_time,omitempty"`
	// 可选，消息发送接口对任务类消息的幂等性保证。
	// 强烈建议开发者在发送任务类消息时填写这个字段，友盟服务端会根据这个字段对消息做去重避免重复发送。
	// 同一个appkey下面的多个消息会根据out_biz_no去重，不同发送任务的out_biz_no需要保证不同，否则会出现后发消息被去重过滤的情况。
	// 注意，out_biz_no只对任务类消息有效。
	OutBizNo string `json:"out_biz_no,omitempty"`
	// 可选，多条带有相同apns_collapse_id的消息，iOS设备仅展示
	// 最新的一条，字段长度不得超过64bytes
	ApnsCollapseID string `json:"apns_collapse_id,omitempty"`
}

type PushResp struct {
	// SUCCESS/FAIL
	Ret  string `json:"ret"`
	Data struct {
		// 当"ret"为"SUCCESS"时，包含如下参数:
		// 单播类消息(type为unicast、listcast、customizedcast且不带file_id)返回：
		MsgID string `json:"msg_id"`
		// 任务类消息(type为broadcast、groupcast、filecast、customizedcast且file_id不为空)返回：
		TaskID string `json:"task_id"`
		// 当"ret"为"FAIL"时,包含如下参数:
		// 错误码，详见附录I
		ErrorCode string `json:"error_code"`
		// 错误信息
		ErrorMsg string `json:"error_msg"`
	} `json:"data"`
}
