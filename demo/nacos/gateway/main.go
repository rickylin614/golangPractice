package main

import (
	"fmt"

	nacosClient "github.com/nacos-group/nacos-sdk-go/v2/clients"
	nacosConst "github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main() {
	sc := []nacosConst.ServerConfig{
		*nacosConst.NewServerConfig("127.0.0.1", 8848, nacosConst.WithContextPath("/nacos")),
	}

	cc := nacosConst.NewClientConfig(
		nacosConst.WithNamespaceId("eb360ca6-ad4b-4524-99ad-eb9c95e72e45"),
		nacosConst.WithTimeoutMs(5000),
		nacosConst.WithNotLoadCacheAtStart(true),
		// nacosConst.WithLogDir("/tmp/nacos/log"),
		// nacosConst.WithCacheDir("/tmp/nacos/cache"),
		// nacosConst.WithLogLevel("debug"),
	)

	// 服務註冊客戶端請求
	client, err := nacosClient.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  cc,
		ServerConfigs: sc,
	})
	if err != nil {
		panic(err)
	}

	// 取得現有payV2所有服務名稱
	serviceInfo, err := client.GetAllServicesInfo(vo.GetAllServiceInfoParam{
		NameSpace: "eb360ca6-ad4b-4524-99ad-eb9c95e72e45",
		GroupName: "payV2",
		PageNo:    1,
		PageSize:  20,
	})
	if err != nil {
		panic(err)
	}
	for i, v := range serviceInfo.Doms { // 印出服務名稱
		srv, err := client.GetService(vo.GetServiceParam{
			ServiceName: v,
			GroupName:   "payV2",
		})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%d: %+v\n", i+1, srv.Hosts)
		}
	}

}
