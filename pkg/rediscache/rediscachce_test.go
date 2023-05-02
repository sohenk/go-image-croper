package rediscache

// import (
// 	"bus-search/app/route_segement_search/service/internal/conf"
// 	"bus-search/app/route_segement_search/service/internal/data"
// 	"bus-search/app/route_segement_search/service/internal/pkg/bus_api_helper"
// 	"context"
// 	"testing"

// 	"github.com/go-kratos/kratos/v2/log"
// 	"github.com/golang/protobuf/ptypes/duration"
// )

// func TestTTT(t *testing.T) {
// 	d := &duration.Duration{Seconds: 0, Nanos: 200000000}
// 	rdb := data.NewRedisClient(&conf.Data{
// 		Redis: &conf.Data_Redis{
// 			Addr:         "r-c8d70def92d700b4pd.redis.rds.aliyuncs.com:6379",
// 			Password:     "#BUS@2021#man@ZH",
// 			Db:           9,
// 			DialTimeout:  d,
// 			WriteTimeout: d,
// 			ReadTimeout:  d,
// 		},
// 	}, log.GetLogger())

// 	if rdb == nil {
// 		t.Fatal("rdb is nil")
// 	}
// 	t.Log("rdb connect")

// 	a := &bus_api_helper.QueryDetailByRouteIDModel{RouteID: "111"}
// 	err := Set(context.Background(), rdb, "test111", *a, 0)
// 	if err != nil {
// 		t.Fatal("set err: ", err)
// 	}
// 	// t.Log(Get[dao.BusRoute](context.Background(), rdb, "test111"))
// }
