package main

import (
	// "github.com/alibaba/ioc-golang/extension/registry/nacos"

	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	nacosClient "github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	nacosConst "github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v3"
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

	// 配置中心請求
	configClient, err := nacosClient.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  cc,
		ServerConfigs: sc,
	})
	if err != nil {
		panic(err)
	}

	configStr, err := configClient.GetConfig(vo.ConfigParam{
		DataId:  "aone",
		Group:   "payV2",
		Content: "content",
		CasMd5:  "44f99849f4d918fd209e05730033a348",
	})
	if err != nil {
		panic(err)
	}

	payConfig := &PayConfig{}
	err = yaml.Unmarshal([]byte(configStr), payConfig)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(payConfig)
	}

	// 服務註冊客戶端請求
	client, err := nacosClient.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  cc,
		ServerConfigs: sc,
	})
	if err != nil {
		panic(err)
	}

	// 服務註冊
	ExampleServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "localhost",
		Port:        uint64(payConfig.Port),
		ServiceName: "aone",
		GroupName:   "payV2",
		ClusterName: "cluster-a",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
	})

	// 優雅關機
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGHUP)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	<-ctx.Done()
}

// 註冊服務
func ExampleServiceClient_RegisterServiceInstance(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	success, err := client.RegisterInstance(param)
	if !success || err != nil {
		panic("RegisterServiceInstance failed!" + err.Error())
	}
	fmt.Printf("RegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

// 註銷服務
func ExampleServiceClient_DeRegisterServiceInstance(client naming_client.INamingClient, param vo.DeregisterInstanceParam) {
	success, err := client.DeregisterInstance(param)
	if !success || err != nil {
		panic("DeRegisterServiceInstance failed!")
	}
	fmt.Printf("DeRegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

// 取得服務設定檔
func ExampleServiceClient_GetServiceConfig(client naming_client.INamingClient, param vo.DeregisterInstanceParam) {

}

type PayConfig struct {
	SectionName string            `yaml:"section_name"`
	Name        string            `yaml:"name"`
	Port        int               `yaml:"port"`
	PayURL      string            `yaml:"pay_url"`
	PayQRURL    string            `yaml:"pay_qrurl"`
	PayH5URL    string            `yaml:"pay_h5url"`
	QueryURL    string            `yaml:"query_url"`
	OrderURL    string            `yaml:"order_url"`
	BankURL     string            `yaml:"bank_url"`
	TokenURL    string            `yaml:"token_url"`
	APIRoute    map[string]string `yaml:"api_route"`
	WhiteList   []string          `yaml:"white_list"`
	Channel     map[string]string `yaml:"channel"`
	ChanNum     string            `yaml:"chan_num"`
	PayMode     string            `yaml:"pay_mode"`
}
