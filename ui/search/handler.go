package search

import (
	"github.com/huminghe/framework/core/api/router"

	"fmt"
	log "github.com/cihub/seelog"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/huminghe/framework/core/api"
	core "github.com/huminghe/framework/core/elastic"
	"github.com/huminghe/framework/core/errors"
	"github.com/huminghe/framework/core/global"
	"github.com/huminghe/framework/core/kv"
	"github.com/huminghe/framework/core/ui"
	"github.com/huminghe/framework/core/util"
	"github.com/huminghe/gopa/config"
	"github.com/huminghe/gopa/model"
	"github.com/huminghe/gopa/ui/search/common"
	"github.com/huminghe/gopa/ui/search/handler"
	mobileHandler "github.com/huminghe/gopa/ui/search/m/handler"
	"net/http"
	"strings"
)

// UserUI is the user namespace, public web
type UserUI struct {
	api.Handler
	Config *common.IndexConfig
}

func (h *UserUI) GetSearchClient() core.API {
	return core.GetClient(h.Config.Elasticsearch)
}

// IndexPageAction index page is for PC
func (h *UserUI) IndexPageAction(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	h.searchPageAction(w, req, ps, false)
}

// MobileIndexPageAction is for mobile
func (h *UserUI) MobileIndexPageAction(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	h.searchPageAction(w, req, ps, true)
}

func (h *UserUI) AJAXMoreItemAction(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rawQuery := h.GetParameter(req, "q")
	query := util.FilterSpecialChar(rawQuery)
	query = util.XSSHandle(query)
	if strings.TrimSpace(query) != "" {

		size := h.GetIntOrDefault(req, "size", 5)
		from := h.GetIntOrDefault(req, "from", 0)
		filter := h.GetParameterOrDefault(req, "filter", "")
		filterQuery := ""
		if filter != "" && strings.Contains(filter, "|") {
			arr := strings.Split(filter, "|")
			filterQuery = fmt.Sprintf(`{
				"match": {
			"%s": "%s"
			}
			},`, arr[0], util.UrlDecode(arr[1]))
		}

		response, err := h.execute(filterQuery, query, false, from, size)
		if err != nil {
			h.Error(w, err)
			return
		}

		if len(response.Hits.Hits) > 0 {
			mobileHandler.Block(w, req, rawQuery, filter, from, size, h.Config.UIConfig, response)
		}
	}
}

func (h *UserUI) execute(filterQuery, query string, agg bool, from, size int) (*core.SearchResponse, error) {

	aggStr := ""
	if agg {
		aggStr = `
	"aggs": {
        "host|Host": {
            "terms": {
                "field": "host",
                "size": 10
            }
        },
        "snapshot.lang|Language": {
            "terms": {
                "field": "snapshot.lang",
                "size": 10
            }
        },
        "task.schema|Protocol": {
            "terms": {
                "field": "task.schema",
                "size": 10
            }
        },
        "snapshot.content_type|Content Type": {
            "terms": {
                "field": "snapshot.content_type",
                "size": 10
            }
        },
        "snapshot.ext|File Ext": {
            "terms": {
                "field": "snapshot.ext",
                "size": 10
            }
        }
    },
	`
	}

	format := `
		{

  "query": {
    "bool": {
      "must": [
       %s
        {
          "query_string": {
            "query": "%s",
            "default_operator": "and",
            "fields": [
              "snapshot.title^100",
              "snapshot.summary",
              "snapshot.text"
            ],
            "allow_leading_wildcard": false
          }
        }
      ]
    }
  },
  "collapse": {
    "field": "snapshot.title.keyword",
    "inner_hits": {
      "name": "collpased_docs",
      "size": 5
    }
  },
    "highlight": {
        "pre_tags": [
            "<span class=highlight_snippet >"
        ],
        "post_tags": [
            "</span>"
        ],
        "fields": {
            "snapshot.title": {
                "fragment_size": 100,
                "number_of_fragments": 5
            },
            "snapshot.text": {
                "fragment_size": 150,
                "number_of_fragments": 5
            }
        }
    },
    %s
    "from": %v,
    "size": %v
}
		`

	q := fmt.Sprintf(format, filterQuery, query, aggStr, from, size)

	if global.Env().IsDebug {
		log.Trace(q)
	}

	return h.GetSearchClient().SearchWithRawQueryDSL("index", []byte(q))
}

func (h *UserUI) searchPageAction(w http.ResponseWriter, req *http.Request, ps httprouter.Params, mobile bool) {
	rawQuery := h.GetParameter(req, "q")
	query := util.FilterSpecialChar(rawQuery)
	query = util.XSSHandle(query)
	if strings.TrimSpace(query) == "" {
		if mobile {
			mobileHandler.Index(w, h.Config.UIConfig)
		} else {
			handler.Index(w, h.Config.UIConfig)
		}
	} else {

		size := h.GetIntOrDefault(req, "size", 10)
		from := h.GetIntOrDefault(req, "from", 0)
		filter := h.GetParameterOrDefault(req, "filter", "")
		filterQuery := ""
		if filter != "" && strings.Contains(filter, "|") {
			arr := strings.Split(filter, "|")
			filterQuery = fmt.Sprintf(`{
				"match": {
			"%s": "%s"
			}
			},`, arr[0], util.UrlDecode(arr[1]))
		}

		response, err := h.execute(filterQuery, query, true, from, size)
		if err != nil {
			h.Error(w, err)
			return
		}

		if mobile {
			mobileHandler.Search(w, req, rawQuery, filter, from, size, h.Config.UIConfig, response)
		} else {
			handler.Search(w, req, rawQuery, filter, from, size, h.Config.UIConfig, response)
		}

	}
}

func (h *UserUI) SuggestAction(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	q := h.GetParameter(req, "query")
	//t := h.GetParameter(req, "type")
	//v := h.GetParameter(req, "version")

	field := "snapshot.title"

	q = util.FilterSpecialChar(q)
	q = util.XSSHandle(q)
	if strings.TrimSpace(q) == "" {
		h.Error(w, errors.New("empty query"))
	} else {

		template := `
		{
    "from": 0,
    "size": 10,
    "query": {

     "bool": {
      "should": [
        {
          "query_string": {
            "query":  "%s",
            "default_operator": "OR",
             "fields" : ["%s"],
            "allow_leading_wildcard": false,
            "boost": 1
          }
        },
        {
          "query_string": {
            "query":  "%s",
            "default_operator": "AND",
            "fields" : ["%s"],
            "allow_leading_wildcard": false,
            "boost": 10
          }
        }
      ]
    }
    },
    "_source": [
    "%s"
    ]
}
		`

		query := fmt.Sprintf(template, q, field, q, field, field)

		response, err := h.GetSearchClient().SearchWithRawQueryDSL("index", []byte(query))
		if err != nil {
			h.Error(w, err)
			return
		}

		if response.GetTotal() > 0 {
			terms := []string{}
			docs := []interface{}{}
			hash := hashset.New()
			for _, hit := range response.Hits.Hits {
				term := hit.Source["snapshot"].(map[string]interface{})["title"]
				text, ok := term.(string)
				text = strings.TrimSpace(text)
				if ok && text != "" {
					if !hash.Contains(text) {
						terms = append(terms, text)
						docs = append(docs, hit.Source)
						hash.Add(text)
						if hash.Size() >= 7 {
							break
						}
					}
				}
			}
			result := map[string]interface{}{}
			result["query"] = q
			result["suggestions"] = terms
			//result["docs"] = docs
			result["data"] = terms
			h.WriteJSON(w, result, 200)
		}

	}
}

func (h *UserUI) GetSnapshotPayloadAction(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	snapshot, err := model.GetSnapshot(id)
	if err != nil {
		h.Error(w, err)
		return
	}

	compressed := h.GetParameterOrDefault(req, "compressed", "true")
	var bytes []byte
	if compressed == "true" {
		bytes, err = kv.GetCompressedValue(config.SnapshotBucketKey, []byte(id))
	} else {
		bytes, err = kv.GetValue(config.SnapshotBucketKey, []byte(id))
	}

	if err != nil {
		h.Error(w, err)
		return
	}

	if len(bytes) > 0 {
		h.Write(w, bytes)

		//add link rewrite
		if util.ContainStr(snapshot.ContentType, "text/html") {

			siteLogo := h.Config.UIConfig.SiteLogo
			if !util.PrefixStr(siteLogo, "http") {
				if ui.GetUIConfig().TLSConfig.TLSEnabled {
					siteLogo = util.JoinPath("https://", ui.GetBindAddress(), siteLogo)
				} else {
					siteLogo = util.JoinPath("http://", ui.GetBindAddress(), siteLogo)
				}

			}

			h.Write(w, []byte("<script language='JavaScript' type='text/javascript'>"))
			h.Write(w, []byte(`var dom=document.createElement("div");dom.innerHTML='<div style="overflow: hidden;z-index: 99999999999999999;width:100%;height:18px;position: absolute top:1px;background:#ebebeb;font-size: 12px;text-align:center;">`))
			h.Write(w, []byte(fmt.Sprintf(`<a href="/"><img border=0 style="float:left;height:18px" src="%s"></a><span style="font-size: 12px;">Saved by Gopa, %v, <a title="%v" href="%v">View original</a></span></div>';var first=document.body.firstChild;  document.body.insertBefore(dom,first);`, siteLogo, snapshot.Created, snapshot.Url, snapshot.Url)))
			h.Write(w, []byte("</script>"))
			h.Write(w, []byte("<script src=\"/static/assets/js/snapshot_footprint.js?v=1\"></script> "))
		}
		return
	}

	h.Error404(w)

}
