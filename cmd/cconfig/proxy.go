// Copyright 2014 Wandoujia Inc. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package main

import (
	"fmt"

	"github.com/wandoulabs/codis/pkg/models"

	"github.com/docopt/docopt-go"
	log "github.com/ngaut/logging"
)

func cmdProxy(argv []string) (err error) {
	usage := `usage:
	codis-config proxy list
	codis-config proxy offline <proxy_name>
	codis-config proxy online <proxy_name>
`
	args, err := docopt.Parse(usage, argv, true, "", false)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debug(args)

	globalEnv.ZkLock.Lock(fmt.Sprintf("proxy, %+v", argv))
	defer func() {
		err := globalEnv.ZkLock.Unlock()
		if err != nil {
			log.Error(err)
		}
	}()

	if args["list"].(bool) {
		log.Warning(err)
		return runProxyList()
	}

	proxyName := args["<proxy_name>"].(string)
	if args["online"].(bool) {
		return runSetProxyStatus(proxyName, models.PROXY_STATE_ONLINE)
	}
	if args["offline"].(bool) {
		return runSetProxyStatus(proxyName, models.PROXY_STATE_MARK_OFFLINE)
	}
	return nil
}

func runProxyList() error {
	return nil
}

func runSetProxyStatus(proxyName, status string) error {
	return nil
}
