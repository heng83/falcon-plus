// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"github.com/open-falcon/falcon-plus/modules/agent/g"
	"github.com/toolkits/file"
	"net/http"
	"os"
	"time"
)

func configAdminRoutes() {
	bindRoutes("/exit", func(w http.ResponseWriter, r *http.Request) {
		if g.IsAdministrator(r.RemoteAddr) {
			w.Write([]byte("exiting..."))
			go func() {
				time.Sleep(time.Second)
				os.Exit(0)
			}()
		} else {
			w.Write([]byte("no privilege"))
		}
	})

	bindRoutes("/config/reload", func(w http.ResponseWriter, r *http.Request) {
		if g.IsAdministrator(r.RemoteAddr) {
			g.ParseConfig(g.ConfigFile)
			RenderDataJson(w, g.Config())
		} else {
			w.Write([]byte("no privilege"))
		}
	})

	bindRoutes("/workdir", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w, file.SelfDir())
	})

	bindRoutes("/ips", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w, g.TrustableIps(false))
	})
}
