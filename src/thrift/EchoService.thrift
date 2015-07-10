namespace go rpc

// 测试服务
service EchoService {
	// 发起远程调用
	string echo(1:i64 callTime, 2:string helloCode);
}