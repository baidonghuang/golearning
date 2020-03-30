package consul

import (
	"fmt"
	//"git.dev.chelizitech.com/cloud-apigw-controller/cmd/apigw"
	"github.com/hashicorp/consul/api"
	"os"
	"testing"
)

func TestConsul001(t *testing.T) {
	sigs := make(chan os.Signal, 1)

	NewConsulClient("10.9.50.4:8500", "http", "e9138d5b-c037-e88b-5cea-a381ae7be43e", "", "", "")

	Services(func(services map[string][]string) {
		//fmt.Printf("%s\n", services)

		for k, v := range services {
			fmt.Printf("k=%v, v=%v\n", k, v)
		}
	})

	<-sigs
}

func TestConsul002(t *testing.T) {
	sigs := make(chan os.Signal, 1)

	NewConsulClient("10.9.50.4:8500", "http", "", "", "", "")

	ServiceChange("sss", func(k string) bool {
		return true
	}, func(name string, se []*api.ServiceEntry) {
		//fmt.Printf("%s\n", services)

		for s := range se {
			fmt.Println(s)
		}
	})

	<-sigs
}

/*func TestConsul005(t *testing.T) {
	sigs := make(chan os.Signal, 1)
	v := `{"anonymousPath":["/ui/ops","/api/ops"],"deniedPath":[""],"domain":"ops.chetailian.net","routes":[{"path":"/api/ops","service":"ops-service"},
	{"path":"/ui/ops","service":"ops-web"}]}`
	ingress := apigw.Ingress{}
	err := json.Unmarshal([]byte(v), &ingress)
	fmt.Println("%s", err)
	<-sigs
}
*/
