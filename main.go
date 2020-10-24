package main

import (
    "github.com/logrusorgru/aurora"
    "strings"
    "flag"
    "bufio"
    "os"
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type (
    cert struct {
        Subject string `json:"subject"`
        Issuer string `json:"issuer"`
        SubjectCN string `json:"subjectCN"`
        SubjectAltName string `json:"subjectAltName"`
    }

    result struct {
        Ip string `json:"ip"`
	CertificateChain []cert `json:"certificateChain"`
}
)

var au aurora.Aurora
var details bool


func init() {
   au = aurora.NewAurora(true)
}

func main() {
    var data result
    domains := []string{}
    flag.BoolVar(&details, "v", false, "输出详情")

    flag.Parse()

    file := flag.Arg(0)


    sc := bufio.NewScanner(os.Stdin)

    for sc.Scan() {
	domain := strings.ToLower(sc.Text())
	domains = append(domains, domain)
    }
    // 读取JSON文件内容 返回字节切片
    bytes, _ := ioutil.ReadFile(file)

    // 将字节切片映射到指定结构上
    json.Unmarshal(bytes, &data)


    var ok bool
    for _, res := range data.CertificateChain {
	if details {
            fmt.Println(au.Yellow("[!] "+res.Subject))
            fmt.Println(au.Yellow("[!] "+res.Issuer))
            fmt.Println(au.Yellow("[!] "+res.SubjectCN))
	    fmt.Println(au.Yellow("[!] "+res.SubjectAltName))
            fmt.Println("========================================================================================================")

	}
	for _, d := range domains {
	    if strings.Contains(res.Subject, d) || strings.Contains(res.Issuer, d) || strings.Contains(res.SubjectCN, d) || strings.Contains(res.SubjectAltName, d){
		    ok = true
	    }
        }

     }


     if ok {
         fmt.Println(data.Ip)
     }

}
