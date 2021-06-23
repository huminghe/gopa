/*
Copyright 2016 Medcl (m AT medcl.net)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	_ "expvar"
	"github.com/huminghe/framework"
	"github.com/huminghe/framework/core/module"
	"github.com/huminghe/framework/core/orm"
	pipeline2 "github.com/huminghe/framework/core/pipeline"
	"github.com/huminghe/framework/modules"
	"github.com/huminghe/gopa/api"
	"github.com/huminghe/gopa/config"
	"github.com/huminghe/gopa/model"
	"github.com/huminghe/gopa/pipeline"
	"github.com/huminghe/gopa/plugins"
	"github.com/huminghe/gopa/ui"
)

func main() {

	terminalHeader := ("  ________ ________ __________  _____   \n")
	terminalHeader += (" /  _____/ \\_____  \\\\______   \\/  _  \\  \n")
	terminalHeader += ("/   \\  ___  /   |   \\|     ___/  /_\\  \\ \n")
	terminalHeader += ("\\    \\_\\  \\/    |    \\    |  /    |    \\\n")
	terminalHeader += (" \\______  /\\_______  /____|  \\____|__  /\n")
	terminalHeader += ("        \\/         \\/                \\/ \n")

	terminalFooter := ("                         |    |                \n")
	terminalFooter += ("   _` |   _ \\   _ \\   _` |     _ \\  |  |   -_) \n")
	terminalFooter += (" \\__, | \\___/ \\___/ \\__,_|   _.__/ \\_, | \\___| \n")
	terminalFooter += (" ____/                             ___/        \n")

	app := framework.NewApp("gopa", "A Spider Written in Go.",
		config.Version, config.LastCommitLog, config.BuildDate, terminalHeader, terminalFooter)

	app.Init(nil)
	defer app.Shutdown()

	app.Start(func() {

		//load core modules first
		modules.Register()

		//register API
		api.InitAPI()

		//register UI
		ui.InitUI()

		//register joints
		pipeline.InitJoints()

		//load plugins
		plugins.Register()

		//start each module, with enabled provider
		module.Start()

	}, func() {
		orm.RegisterSchema(pipeline2.PipelineConfig{})
		orm.RegisterSchema(model.Host{})
		orm.RegisterSchema(model.Task{})
		orm.RegisterSchema(model.Snapshot{})
		orm.RegisterSchema(model.HostConfig{})
		orm.RegisterSchema(model.Project{})
		orm.RegisterSchema(model.Index{})
	})

}
