module mepserver

go 1.14

replace (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b => github.com/go-chassis/glog v0.0.0-20180920075250-95a09b2413e9
	go.uber.org/zap v1.9.0 => github.com/uber-go/zap v1.9.0
	golang.org/x/crypto v0.0.0-20180904163835-0709b304e793 => github.com/golang/crypto v0.0.0-20180904163835-0709b304e793
	golang.org/x/net v0.0.0-20180824152047-4bcd98cce591 => github.com/golang/net v0.0.0-20180824152047-4bcd98cce591
	golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33 => github.com/golang/sys v0.0.0-20180905080454-ebe1bf3edb33
	golang.org/x/text v0.0.0-20170627122817-6353ef0f9243 => github.com/golang/text v0.0.0-20170627122817-6353ef0f9243
	golang.org/x/time v0.0.0-20170424234030-8be79e1e0910 => github.com/golang/time v0.0.0-20170424234030-8be79e1e0910
	google.golang.org/genproto v0.0.0-20170531203552-aa2eb687b4d3 => github.com/google/go-genproto v0.0.0-20170531203552-aa2eb687b4d3
	google.golang.org/grpc v1.7.5 => github.com/grpc/grpc-go v1.7.5
	k8s.io/api v0.0.0-20180601181742-8b7507fac302 => github.com/kubernetes/api v0.0.0-20180601181742-8b7507fac302
	k8s.io/apimachinery v0.0.0-20180601181227-17529ec7eadb => github.com/kubernetes/apimachinery v0.0.0-20180601181227-17529ec7eadb
	k8s.io/client-go v2.0.0-alpha.0.0.20180817174322-745ca8300397+incompatible => github.com/kubernetes/client-go v0.0.0-20180817174322-745ca8300397
)

require (
	github.com/apache/servicecomb-service-center v0.0.0-20191027084911-c2dc0caef706
	github.com/satori/go.uuid v1.1.0
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
)
