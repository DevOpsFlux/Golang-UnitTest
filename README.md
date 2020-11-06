***
## go - Setup
- go vergion [ go1.22.3 -> 업그레이드(의존성이슈), go1.15.2 windows/amd64 ]
- https://golang.org/dl/

## go path - Set
- Path, GoBin, GoRoot, GoPath

## go get
- git package download
```bash
 go get -u github.com/go-redis/redis

 [issue] go get gopkg > x509:certificate.....
 https://github.com/golang/go/issues/18519
 go get -u -v -insecure github.com/go-redis/redis
```
## go mod
- Go modules
https://tutorialedge.net/golang/go-modules-tutorial/
go list -m all : 최종 버전의 직간접적인 디펜더시 리스트
go mod tidy : 불필요한 종속성 제거

## go make 
- https://tutorialedge.net/golang/makefiles-for-go-developers/
```bash
 choco install make
 make -v
```
## godoc
```bash
 godoc -v
 godoc -http=localhost:6060
 http://localhost:6060/pkg/
 Go Documentation Server
```
## go 바이너리 실행 파일 생성
```bash
 go build devgo.go
 devgo.exe
```
## docker-compose
- https://docs.docker.com/compose/
- docker-compose.yml
```bash
docker-compose up --build -d
docker-compose ps
docker-compose stop
docker-compose start
docker-compose down
```

***
## Go Web Framework

- Gin
https://github.com/gin-gonic/gin
https://github.com/gin-gonic/examples

- Echo
https://github.com/labstack/echo
https://github.com/pangpanglabs/echosample

- fasthttp & router
https://github.com/valyala/fasthttp
https://github.com/qiangxue/fasthttp-routing
https://github.com/fasthttp/router


## Go Log Library

- logrus vs zap 
https://go.libhunt.com/compare-logrus-vs-zap
- logrus
https://github.com/sirupsen/logrus
https://godoc.org/github.com/sirupsen/logrus
- zap
https://github.com/uber-go/zap
https://godoc.org/go.uber.org/zap


***
## Windows VSCode를 통한 Go 개발 환경 구축 
https://velog.io/@ilcm96/windows-go-dev-env-setup-with-vscode

## Plugin 정보
GoDoc, GoDocBrowser : 패키지 문서 탐험
GoRename : 이름 변경
GoCoverage : 테스트 커버리지 측정
GoLint : 오류 검사
GoVert : 정적 오류 탐지
GoErrCheck : 오류 검사
GoImplements, GoCallees, GoReferrers : 고급 분석 도구

---