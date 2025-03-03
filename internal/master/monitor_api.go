// Copyright 2019 The Vearch Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package master

import (
	"github.com/gin-gonic/gin"
	"github.com/vearch/vearch/v3/internal/config"
	"github.com/vearch/vearch/v3/internal/monitor"
)

type monitorApi struct {
	router         *gin.Engine
	monitorService *monitorService
}

func ExportToMonitorHandler(router *gin.Engine, monitorService *monitorService) {
	monitor.Register(monitorService.Client, monitorService.etcdServer, config.Conf().Masters.Self().MonitorPort)
	// monitorService.Register()
}
