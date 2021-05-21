# 友盟免费版消息推送

```
//安卓
push := umandroid.NewPush("androidAppKey", "androidAppMasterSecret")
	resp, _ := push.Push(PushTemplate{
		Appkey:       "androidAppKey",
		Timestamp:    strconv.FormatInt(time.Now().Unix(), 10),
		Type:         "unicast",
		DeviceTokens: "设备token",
		Payload: Payload{
			DisplayType: "notification",
			Body: Body{
				Ticker: "Ticker",
				Title:  "Title",
				Text:   "Text",
			},
		},
		MiPush:     "true",
		MiActivity: "com.mango.sqt.android.NotifyClickActivity",
	})
	prettyJSON, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Printf("%s", string(prettyJSON))
	
//ios
push := umios.NewPush("androidAppKey", "androidAppMasterSecret")
	resp, err := push.Push(PushTemplate{
		Appkey:       "androidAppKey",
		Timestamp:    strconv.FormatInt(time.Now().Unix(), 10),
		Type:         "unicast",
		DeviceTokens: "设备token",
		Payload: Payload{
			Aps: Aps{
				Alert: &Alert{
					Title:    "标题啊",
					SubTitle: "子标题",
					Body:     "内容",
				},
			},
			Extra: nil,
		},
		ProductionMode: "true",
		Description:    "测试推送",
	})
	prettyJSON, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Printf("%s", string(prettyJSON))
```