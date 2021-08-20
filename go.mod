module examples

go 1.14

require (
	github.com/Shopify/sarama v1.29.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/evanphx/json-patch v4.9.0+incompatible
	github.com/golang-collections/collections v0.0.0-20130729185459-604e922904d3
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.10.0
	github.com/mattbaird/jsonpatch v0.0.0-20200820163806-098863c1fc24
	github.com/nats-io/nats-server/v2 v2.2.0
	github.com/nats-io/nats.go v1.10.1-0.20210228004050-ed743748acac
	github.com/prometheus/client_golang v1.9.0
	github.com/rs/zerolog v1.20.0
	github.com/soheilhy/cmux v0.1.4
	github.com/tidwall/gjson v1.6.8
	go.etcd.io/etcd v0.0.0-20200520232829-54ba9589114f
	golang.org/x/net v0.0.0-20210427231257-85d9c07bbe3a
	google.golang.org/protobuf v1.25.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	k8s.io/api v0.20.3
	k8s.io/apimachinery v0.20.3
	k8s.io/client-go v0.20.3
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
