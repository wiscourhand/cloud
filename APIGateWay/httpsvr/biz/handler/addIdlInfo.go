package handler

//func AddIdlInfo(ctx context.Context, c *app.RequestContext)  {
//	cli, err := idlservice.NewClient("student-server", kclient.WithHostPorts("127.0.0.1:8889"))
//	if err != nil {
//		panic("err new client:" + err.Error())
//	}
//
//	httpReq, err := adaptor.GetCompatRequest(c.GetRequest())
//	if err != nil {
//		panic("get http req failed")
//	}
//
//	customReq, err := generic.FromHTTPRequest(httpReq)
//	if err != nil {
//		panic(err)
//		//panic("get custom req failed")
//	}
//
//	jsonStr, err := json.Marshal(customReq.Body)
//	jsonReq := string(jsonStr)
//	fmt.Println(jsonReq)
//
//	resp, err := cli.GenericCall(ctx, c.Param("methodName"), jsonReq)
//	if err != nil {
//		panic("err query rpc server:" + err.Error())
//	}
//
//	c.JSON(consts.StatusOK, resp)
//}
