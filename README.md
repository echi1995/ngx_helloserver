# ngx_helloserver
用于nginx反向代理和负载均衡的一个demo服务


启动服务
```
go run main.go -port 8001 
```

有一个接口  
- /hello  
会返回当前这个服务的port