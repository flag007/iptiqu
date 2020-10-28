别问 问就是不知道用来干啥的
用法:

```
for i in $(cat ssl-alive.in); do tls-scan -c $i --port=443,8443 --cacert=/root/db/ca-bundle.crt 2>/dev/null -o $i.json; done
```

```
获取ip
▶ for i in $(cat ssl-alive.in); do cat root.txt | iptiqu 1.json | tee -a ips.2;done
```

安装:

```
▶ go get -u -v github.com/flag007/iptiqu
```
