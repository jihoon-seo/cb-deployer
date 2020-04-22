# CB-Deployer
CB-Deployer; the module for deploying Cloud-Barista

- To run Cloud-Barista system:

```
$ git clone https://github.com/jihoon-seo/cb-deployer.git
$ cd cb-deployer
$ docker-compose up
```

## 실행되는 서비스 목록
| Name | 직접 접속 주소 | APIGW 통한 접속 주소 | 비고 |
|---|---|---|---|
| cb-restapigw | http://{{host}}:8000 |   |   |
| cb-restapigw-influxdb | http://{{host}}:8086 |   | 8083: Admin Panel <br>8086: client-server comm. |
| cb-restapigw-grafana | http://{{host}}:3100 |   | ID: admin / PW: admin |
| cb-restapigw-jaeger | http://{{host}}:16686 |   |   |
| --- |   |   |   |
| cb-spider | http://{{host}}:1024/spider | http://{{host}}:8000/spider |   |
| cb-tumblebug | http://{{host}}:1323/tumblebug | http://{{host}}:8000/tumblebug |   |
| cb-webtool | http://{{host}}:1234 |   |   |
| --- |   |   |   |
| cb-dragonfly | http://{{host}}:9090 |   |   |
| cb-dragonfly-influxdb | http://{{host}}:28086 |   |   |
| cb-dragonfly-etcd | http://{{host}}:2379 |   | 2379: client communication <br>2380: server-to-server communication |
