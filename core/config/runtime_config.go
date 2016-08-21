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

package config

import (
	. "github.com/medcl/gopa/core/types"
)

type PathConfig struct {
	Data     string  `data`
	Log      string  `log`
	TaskData string  `task`
	WebData  string  `web`

	SavedFileLog    string //path of saved files
	PendingFetchLog string //path of pending fetch
	FetchFailedLog  string //path of failed fetch
}

type ClusterConfig struct {
	Name string
}

type LoggingConfig struct {
	Level     string `level`

	//config string of seelog
	ConfigStr string
}

type IndexingConfig struct {
	Host string `host`
	Index string `index`
}

type SaveConfig struct {
	DefaultExtension string
}

type RuledFetchConfig struct {
	UrlTemplate        string
	From               int
	To                 int
	Step               int
	LinkExtractPattern string
	LinkTemplate       string
}

type RuntimeConfig struct {
	//cluster
	ClusterConfig *ClusterConfig `cluster`

	//logging related
	LoggingConfig *LoggingConfig `logging`
	IndexingConfig *IndexingConfig `indexing`


	//path related
	PathConfig *PathConfig `path`


	//task
	TaskConfig *TaskConfig

	RuledFetchConfig *RuledFetchConfig

	//splitter of joined array string
	ArrayStringSplitter string


	StoreWebPageTogether bool

	MaxGoRoutine int

	//switch config
	ParseUrlsFromSavedFileLog      bool
	LoadTemplatedFetchJob          bool
	LoadPendingFetchJobs           bool //fetch url parse and extracted from saved page,load data from:"pending_fetch.urls"
	ParseUrlsFromPreviousSavedPage bool //extract urls from previous saved page
	LoadRuledFetchJob              bool //extract urls from previous saved page
	HttpEnabled                    bool //extract urls from previous saved page

	//runtime variables
	Storage Store

	WalkBloomFilterFileName         string
	FetchBloomFilterFileName        string
	ParseBloomFilterFileName        string
	PendingFetchBloomFilterFileName string
}
