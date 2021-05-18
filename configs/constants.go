package configs

var (
	ServerAddr              = "0.0.0.0:8889" //匹配服务器监听地址
	MaxEventQueueSize int32 = 500            //消息队列大小
	DgsAddr           string
	DgsRPCPort        string = "9000"
	NameSpace  string = "default"
	DgsSvcName string = "dgs-srv"
)
