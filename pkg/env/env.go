package env

import (
	"os"

	"github.com/c4pt0r/cfg"
	log "github.com/ngaut/logging"
	"github.com/ngaut/zkhelper"
	"github.com/wandoulabs/codis/pkg/utils"
)

type ZkConnFactoryFunc func() (zkhelper.Conn, error)

type CodisEnv struct {
	ZkAddr        string
	ZkConn        zkhelper.Conn
	ZkLock        zkhelper.ZLocker
	ZkConnFactory ZkConnFactoryFunc

	DashboardAddr string
	ProductName   string
	cfg           *cfg.Cfg
}

func LoadCodisEnv(cfg *cfg.Cfg) *CodisEnv {
	if cfg == nil {
		log.Fatal("config error")
	}

	productName, err := cfg.ReadString("product", "test")
	if err != nil {
		log.Fatal(err)
	}

	zkAddr, err := cfg.ReadString("zk", "localhost:2181")
	if err != nil {
		log.Fatal(err)
	}

	hostname, _ := os.Hostname()
	dashboardAddr, err := cfg.ReadString("dashboard_addr", hostname+":18087")
	if err != nil {
		log.Fatal(err)
	}

	zkConn, err := zkhelper.ConnectToZk(zkAddr)
	if err != nil {
		log.Fatal(err)
	}

	zkLock := utils.GetZkLock(zkConn, productName)

	return &CodisEnv{
		ZkAddr:        zkAddr,
		ZkConn:        zkConn,
		ZkLock:        zkLock,
		DashboardAddr: dashboardAddr,
		ProductName:   productName,
		ZkConnFactory: func() (zkhelper.Conn, error) {
			return zkhelper.ConnectToZk(zkAddr)
		},
		cfg: cfg,
	}
}
