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

package crawler

import (
	log "github.com/cihub/seelog"
	. "github.com/medcl/gopa/core/env"
	. "github.com/medcl/gopa/core/pipeline"
	"github.com/medcl/gopa/core/types"
	. "github.com/medcl/gopa/modules/crawler/pipe"
	"runtime"
	"time"
	"github.com/medcl/gopa/core/queue"
	"github.com/medcl/gopa/modules/config"
)

var fetchQuitChannels []*chan bool
var started = false

func (this CrawlerModule) Name() string{
	return "Crawler"
}

func (this CrawlerModule) Start(env *Env) {
	if started {
		log.Error("crawler already started, please stop it first.")
	}
	numGoRoutine := env.RuntimeConfig.MaxGoRoutine
	//shutdownSignal signals for each go routing
	fetchQuitChannels = make([]*chan bool, numGoRoutine)
	if env.RuntimeConfig.CrawlerConfig.Enabled {
		go func() {

			//start fetcher
			for i := 0; i < numGoRoutine; i++ {
				log.Trace("start crawler:", i)
				quitC := make(chan bool, 1)
				fetchQuitChannels[i] = &quitC
				go RunPipeline(env, &quitC, i)

			}
		}()
	} else {
		log.Info("crawler currently not enabled")
		return
	}

	started = true
}

func (this CrawlerModule) Stop() error {
	if started {
		log.Debug("start shutting down crawler")
		for i, item := range fetchQuitChannels {
			if item != nil {
				*item <- true
			}
			log.Debug("send exit signal to fetch channel: ", i)
		}

		started = false
	} else {
		log.Error("crawler is not started, please start it first.")
	}

	return nil
}

func RunPipeline(env *Env, quitC *chan bool, shard int) {

	go func() {
		for {
			log.Trace("waiting url to fetch")
			data := queue.Pop(config.FetchChannel)
			url:=types.PageTaskFromBytes(data)
			urlStr := string(url.Url)
			log.Debug("shard:", shard, ",url received:", urlStr)

			execute(url, env)
		}
	}()

	log.Trace("fetch task started.shard:", shard)

	<-*quitC

	log.Trace("fetch task exit.shard:", shard)

}

func execute(seed types.TaskSeed, env *Env) {

	log.Trace("start crawler")

	var pipeline *Pipeline
	defer func() {
		if !env.IsDebug {
			if r := recover(); r != nil {
				if _, ok := r.(runtime.Error); ok {
					err := r.(error)
					log.Error(pipeline.GetID(), ", ", seed.Url, ", ", err)
				}
				log.Error("error in crawler")
			}
		}
	}()

	pipeline = NewPipeline("crawler")

	pipeline.Context(&Context{Env: env}).
		Start(Start{Seed: seed}).
		Join(UrlNormalizationJoint{FollowSubDomain: true}).
		Join(UrlFilterJoint{}).
		Join(LoadMetadataJoint{}).
		Join(IgnoreTimeoutJoint{IgnoreTimeoutAfterCount: 100}).
		Join(FetchJoint{}).
		Join(ParserJoint{DispatchLinks: true, MaxDepth: 3}).
		//Join(SaveToFileSystemJoint{}).
		Join(SaveToDBJoint{CompressBody: true}).
		Join(PublishJoint{}).
		End(End{}).
		Run()

	if env.RuntimeConfig.TaskConfig.FetchDelayThreshold > 0 {
		log.Debug("sleep ", env.RuntimeConfig.TaskConfig.FetchDelayThreshold, "ms to control crawling speed")
		time.Sleep(time.Duration(env.RuntimeConfig.TaskConfig.FetchDelayThreshold) * time.Millisecond)
		log.Debug("wake up now,continue crawing")
	}

	log.Trace("end crawler")
}

type CrawlerModule struct {
}

func CreateSeed(task types.TaskSeed)  {
	queue.Push(config.CheckChannel,task.MustGetBytes())
	log.Trace("end create seed")
}