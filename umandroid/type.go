package umandroid

// AndroidTemplate 安卓推送模板
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
	// 必填，JSON格式，具体消息内容(Android最大为1840B)
	Payload Payload `json:"payload"`
	// 可选，发送策略
	Policy *Policy `json:"policy,omitempty"`
	// 可选，正式/测试模式。默认为true
	// 测试模式只对“广播”、“组播”类消息生效，其他类型的消息任务（如“文件播”）不会走测试模式
	// 测试模式只会将消息发给测试设备。测试设备需要到web上添加。
	// Android: 测试设备属于正式设备的一个子集。
	ProductionMode string `json:"production_mode,omitempty"`
	// 可选，发送消息描述，建议填写。
	Description string `json:"description,omitempty"`
	//系统弹窗，只有display_type=notification生效
	// 可选，默认为false。当为true时，表示MIUI、EMUI、Flyme 系统设备离线转为系统下发
	MiPush string `json:"mipush,omitempty"`
	// 可选，mipush 值为true时生效，表示走系统通道时打开指定页面 acitivity 的完整包路径。
	MiActivity string `json:"mi_activity,omitempty"`
	//可选，厂商通道相关的特殊配置
	ChannelProperties *ChannelProperties `json:"channel_properties,omitempty"`
}
type Payload struct {
	// 必填，消息类型: notification(通知)、message(消息)
	DisplayType string `json:"display_type"`
	// 必填，消息体。
	// 当display_type=message时，body的内容只需填写custom字段。
	// 当display_type=notification时，body包含如下参数:
	// 通知展现内容:
	Body Body `json:"body"`
	// 可选，JSON格式，用户自定义key-value。
	//只对"通知" (display_type=notification)生效。
	// 可以配合通知到达后，打开App/URL/Activity使用。
	Extra map[string]interface{} `json:"extra,omitempty"`
}
type Body struct {
	// 通知栏提示文字
	Ticker string `json:"ticker,omitempty"`
	// 必填，通知标题
	Title string `json:"title,omitempty"`
	// 必填，通知文字描述
	Text string `json:"text,omitempty"`
	// 自定义通知图标:
	// 可选，状态栏图标ID，R.drawable.[smallIcon]
	// 如果没有，默认使用应用图标。
	// 图片要求为24*24dp的图标，或24*24px放在drawable-mdpi下。
	// 注意四周各留1个dp的空白像素
	Icon string `json:"icon,omitempty"`
	// 可选，通知栏拉开后左侧图标ID，R.drawable.[largeIcon]，
	// 图片要求为64*64dp的图标，
	// 可设计一张64*64px放在drawable-mdpi下，
	// 注意图片四周留空，不至于显示太拥挤
	LargeIcon string `json:"largeIcon,omitempty"`
	// 可选，通知栏大图标的URL链接。该字段的优先级大于largeIcon。
	// 该字段要求以http或者https开头，图片建议不大于100KB。
	Img string `json:"img,omitempty"`
	// 可选，通知声音，R.raw.[sound]。
	// 如果该字段为空，采用SDK默认的声音，即res/raw/下的
	// umeng_push_notification_default_sound声音文件。如果
	// SDK默认声音文件不存在，则使用系统默认Notification提示音。
	Sound string `json:"sound,omitempty"`
	// 可选，默认为0，用于标识该通知采用的样式。使用该参数时，
	// 开发者必须在SDK里面实现自定义通知栏样式。
	BuilderID int `json:"builder_id,omitempty"`
	// 通知到达设备后的提醒方式，注意，"true/false"为字符串
	// 可选，收到通知是否震动，默认为"true"
	PlayVibrate string `json:"play_vibrate,omitempty"`
	// 可选，收到通知是否闪灯，默认为"true"
	PlayLights string `json:"play_lights,omitempty"`
	// 可选，收到通知是否发出声音，默认为"true"
	PlaySound string `json:"play_sound,omitempty"`
	// 点击"通知"的后续行为，默认为打开app。
	// 可选，默认为"go_app"，值可以为:
	//   "go_app": 打开应用
	//   "go_url": 跳转到URL
	//   "go_activity": 打开特定的activity
	//   "go_custom": 用户自定义内容。
	AfterOpen string `json:"after_open,omitempty"`
	// 当after_open=go_url时，必填。
	// 通知栏点击后跳转的URL，要求以http或者https开头
	URL string `json:"url,omitempty"`
	// 当after_open=go_activity时，必填。
	// 通知栏点击后打开的Activity
	Activity string `json:"activity,omitempty"`
	// 当display_type=message时, 必填
	// 当display_type=notification且 after_open=go_custom时，必填
	// 用户自定义内容，可以为字符串或者JSON格式。
	Custom interface{} `json:"custom,omitempty"`
}
type ChannelProperties struct {
	//小米channel_id，具体使用及限制请参考小米推送文档 https://dev.mi.com/console/doc/detail?pId=2086
	XiaomiChannelID string `json:"xiaomi_channel_id,omitempty"`
	//vivo消息分类：0 运营消息，1 系统消息， 需要到vivo申请，具体使用及限制参考[vivo消息推送分类功能说明]https://dev.vivo.com.cn/documentCenter/doc/359
	VivoClassification string `json:"vivo_classification,omitempty"`
	//可选， android8以上推送消息需要新建通道，否则消息无法触达用户。
	//push sdk 6.0.5及以上创建了默认的通道:upush_default，
	//消息提交厂商通道时默认添加该通道。如果要自定义通道名称或使用私信，
	//请自行创建通道，推送消息时携带该参数 具体可参考 [oppo通知通道适配] https://open.oppomobile.com/wiki/doc#id=10289
	OppoChannelID string `json:"oppo_channel_id,omitempty"`
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
