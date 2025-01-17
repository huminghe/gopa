// Generated by ego.
// DO NOT EDIT

package handler

import (
	"fmt"
	"github.com/huminghe/framework/core/elastic"
	"github.com/huminghe/framework/core/util"
	"github.com/huminghe/gopa/ui/search/common"
	"html"
	"io"
	"net/http"
	"strings"
)

var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
func CommonFooter(w io.Writer) error {
	_, _ = io.WriteString(w, "\n\n<script type=\"text/javascript\"  src=\"/static/assets/uikit-2.27.1/js/uikit.min.js\"></script>\n<script type=\"text/javascript\" charset=\"ISO-8859-1\" src=\"/static/assets/js/jquery.autocomplete.js\"></script>\n<script type=\"text/javascript\" charset=\"utf-8\"  src=\"/static/assets/js/footer.js?v=2\"></script>\n<script src=\"/static/assets/js/loadmore.js\"></script>\n<script src=\"/static/assets/js/ie_detect.js\"></script>\n")
	return nil
}
func CommonHeader(w io.Writer, config *common.UIConfig) error {
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n\n<meta content=IE=7 http-equiv=X-UA-Compatible>\n<meta content=text/html;charset=utf-8 http-equiv=content-type>\n\n<meta name=\"robots\" content=\"all\">\n<meta name=\"license\" content=\"keep-copyright-footprint,no-KPI-shit,respect-first\">\n<meta name=\"creator\" content=\"medcl\">\n<meta name=\"generator\" content=\"https://github.com/huminghe/gopa\">\n<meta name=\"copyright\" content=\"Apache License, Version 2.0\">\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1, user-scalable=no, minimal-ui\">\n\n<link rel=\"icon\" href=\"")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(config.SiteFavicon)))
	_, _ = io.WriteString(w, "\" type=\"image/x-icon\" />\n\n<link rel=\"stylesheet\" href=\"/static/assets/uikit-2.27.1/css/uikit.min.css\" />\n<link rel=\"stylesheet\" href=\"/static/assets/css/style.css?v=5\" rel=\"stylesheet\" type=\"text/css\"/>\n<link rel=\"stylesheet\" href=\"/static/assets/css/search_style.css?v=1\" rel=\"stylesheet\" type=\"text/css\"/>\n<link rel=\"stylesheet\" href=\"/static/assets/css/loadmore.css\" rel=\"stylesheet\" type=\"text/css\"/>\n\n<script type=\"text/javascript\"  src=\"/static/assets/js/jquery.min.js\"></script>\n<script>\n    if(/Android|Windows Phone|webOS|iPhone|iPad|iPod|BlackBerry/i.test(navigator.userAgent)){\n        window.location.href = location.origin + \"/m/\" + location.search;\n    }\n</script>\n")
	return nil
}
func Index(w io.Writer, config *common.UIConfig) error {
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n\n<!DOCTYPE html>\n<!--[if lt IE 7 ]> <html class=\"no-js ie6\" lang=\"en-US\"> <![endif]-->\n<!--[if IE 7 ]>    <html class=\"no-js ie7\" lang=\"en-US\"> <![endif]-->\n<!--[if IE 8 ]>    <html class=\"no-js ie8\" lang=\"en-US\"> <![endif]-->\n<!--[if (gte IE 9)|!(IE)]><!--> <html lang=\"en-US\"> <!--<![endif]-->\n<head>\n    <title>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(config.SiteName)))
	_, _ = io.WriteString(w, "</title>\n    ")
	CommonHeader(w, config)
	_, _ = io.WriteString(w, "\n    <link rel=\"stylesheet\" href=\"/static/assets/css/search/index.css\" rel=\"stylesheet\" type=\"text/css\"/>\n</head>\n\n<body link=#0000cc>\n\n<div style=\"max-width:600px;clear: both;display: block; padding-top: -50px;\">\n\n    <div align=\"center\" style=\"padding-bottom: 10px;\">\n        <a  href=\"/\">\n            <img alt=\"")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(config.SiteName)))
	_, _ = io.WriteString(w, "\" border=0 style=\"width: 100%\"  src=\"")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(config.SiteLogo)))
	_, _ = io.WriteString(w, "\">\n        </a>\n\n    </div>\n\n    <div id=\"toolbar_wrapper\">\n        <svg xmlns=\"http://www.w3.org/2000/svg\" style=\"display:none\">\n            <symbol xmlns=\"http://www.w3.org/2000/svg\" id=\"sbx-icon-search-6\" viewBox=\"0 0 40 40\">\n                <path d=\"M28.295 32.517c-2.93 2.086-6.51 3.312-10.38 3.312C8.02 35.83 0 27.81 0 17.914 0 8.02 8.02 0 17.915 0 27.81 0 35.83 8.02 35.83 17.915c0 3.87-1.227 7.45-3.313 10.38l6.61 6.61c1.166 1.165 1.163 3.057 0 4.22-1.167 1.167-3.057 1.167-4.226-.002l-6.605-6.606zm-10.38.326c8.245 0 14.928-6.683 14.928-14.928 0-8.245-6.683-14.93-14.928-14.93-8.245 0-14.93 6.685-14.93 14.93 0 8.245 6.685 14.928 14.93 14.928zm0-26.573c-6.43 0-11.645 5.214-11.645 11.645 0 .494.4.895.896.895.495 0 .896-.4.896-.895 0-5.442 4.41-9.853 9.853-9.853.494 0 .895-.4.895-.896 0-.495-.4-.896-.895-.896z\"\n                      fill-rule=\"evenodd\" />\n            </symbol>\n            <symbol xmlns=\"http://www.w3.org/2000/svg\" id=\"sbx-icon-clear-3\" viewBox=\"0 0 20 20\">\n                <path d=\"M8.114 10L.944 2.83 0 1.885 1.886 0l.943.943L10 8.113l7.17-7.17.944-.943L20 1.886l-.943.943-7.17 7.17 7.17 7.17.943.944L18.114 20l-.943-.943-7.17-7.17-7.17 7.17-.944.943L0 18.114l.943-.943L8.113 10z\" fill-rule=\"evenodd\" />\n            </symbol>\n        </svg>\n\n        <form name=\"f\" id=\"f\" novalidate=\"novalidate\" class=\"searchbox sbx-custom\" action=\"/\" onSubmit=\"submitSearch()\">\n            <div role=\"search\" class=\"sbx-custom__wrapper\">\n                <input name=\"q\" placeholder=\"Please type to search ...\"\n                autocomplete=\"off\" required=\"required\"\n                id=\"query\"  maxLength=100 role=\"textbox\"\n                accesskey=\"s\" type=\"text\" x-webkit-speech=\"\" speech=\"\" spellcheck=\"false\"\n                class=\"sbx-custom__input\" >\n                <button type=\"submit\" title=\"Submit your search query.\" class=\"sbx-custom__submit\">\n                    <svg role=\"img\" aria-label=\"Search\">\n                        <use xmlns:xlink=\"http://www.w3.org/1999/xlink\" xlink:href=\"#sbx-icon-search-6\"></use>\n                    </svg>\n                </button>\n                <button type=\"reset\" title=\"Clear the search query.\" class=\"sbx-custom__reset\">\n                    <svg role=\"img\" aria-label=\"Reset\">\n                        <use xmlns:xlink=\"http://www.w3.org/1999/xlink\" xlink:href=\"#sbx-icon-clear-3\"></use>\n                    </svg>\n                </button>\n            </div>\n\n        </form>\n\n\n    </div>\n    <div class=\"copyright\" style=\"text-align: center;margin-top: 20px\">\n        ")
	common.Copyright(w, config)
	_, _ = io.WriteString(w, "\n    </div>\n</div>\n\n")
	CommonFooter(w)
	_, _ = io.WriteString(w, "\n\n</body>\n</html>\n")
	return nil
}
func Search(w io.Writer, r *http.Request, q string, filter string, from int, size int, config *common.UIConfig, response *elastic.SearchResponse) error {
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n\n<!DOCTYPE html>\n<!--[if lt IE 7 ]> <html class=\"no-js ie6\" lang=\"en-US\"> <![endif]-->\n<!--[if IE 7 ]>    <html class=\"no-js ie7\" lang=\"en-US\"> <![endif]-->\n<!--[if IE 8 ]>    <html class=\"no-js ie8\" lang=\"en-US\"> <![endif]-->\n<!--[if (gte IE 9)|!(IE)]><!--> <html lang=\"en-US\"> <!--<![endif]-->\n\n<head>\n    <title>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(q)))
	_, _ = io.WriteString(w, " | ")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(config.SiteName)))
	_, _ = io.WriteString(w, "</title>\n    ")
	CommonHeader(w, config)
	_, _ = io.WriteString(w, "\n    <link rel=\"stylesheet\" href=\"/static/assets/css/search/search.css\" rel=\"stylesheet\" type=\"text/css\"/>\n</head>\n\n<body link=#0000cc>\n\n<div id=out>\n    <div id=in>\n        <div id=wrapper >\n\n            <div id=head>\n                <div class=\"global_left\">\n\n                    <div class=nv>\n                        <a class=logo  href=\"/\">\n                            <img border=0  style=\"margin-bottom:-55px;height:45px\"   src=\"")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(config.SiteLogo)))
	_, _ = io.WriteString(w, "\">\n                        </a>\n\n                    </div>\n\n\n                    <div id=\"toolbar_wrapper\">\n                        <div id=\"search-context\">\n                            <a href=\"#\" class=\"active\" context=\"default\" suggest_url=\"/suggest/\"></a>\n                        </div>\n                        <svg xmlns=\"http://www.w3.org/2000/svg\" style=\"display:none\">\n                            <symbol xmlns=\"http://www.w3.org/2000/svg\" id=\"sbx-icon-search-6\" viewBox=\"0 0 40 40\">\n                                <path d=\"M28.295 32.517c-2.93 2.086-6.51 3.312-10.38 3.312C8.02 35.83 0 27.81 0 17.914 0 8.02 8.02 0 17.915 0 27.81 0 35.83 8.02 35.83 17.915c0 3.87-1.227 7.45-3.313 10.38l6.61 6.61c1.166 1.165 1.163 3.057 0 4.22-1.167 1.167-3.057 1.167-4.226-.002l-6.605-6.606zm-10.38.326c8.245 0 14.928-6.683 14.928-14.928 0-8.245-6.683-14.93-14.928-14.93-8.245 0-14.93 6.685-14.93 14.93 0 8.245 6.685 14.928 14.93 14.928zm0-26.573c-6.43 0-11.645 5.214-11.645 11.645 0 .494.4.895.896.895.495 0 .896-.4.896-.895 0-5.442 4.41-9.853 9.853-9.853.494 0 .895-.4.895-.896 0-.495-.4-.896-.895-.896z\"\n                                      fill-rule=\"evenodd\" />\n                            </symbol>\n                            <symbol xmlns=\"http://www.w3.org/2000/svg\" id=\"sbx-icon-clear-3\" viewBox=\"0 0 20 20\">\n                                <path d=\"M8.114 10L.944 2.83 0 1.885 1.886 0l.943.943L10 8.113l7.17-7.17.944-.943L20 1.886l-.943.943-7.17 7.17 7.17 7.17.943.944L18.114 20l-.943-.943-7.17-7.17-7.17 7.17-.944.943L0 18.114l.943-.943L8.113 10z\" fill-rule=\"evenodd\" />\n                            </symbol>\n                        </svg>\n                        <form name=\"f\" id=\"f\" novalidate=\"novalidate\" class=\"searchbox sbx-custom\" action=\"/\" onSubmit=\"submitSearch()\">\n                            <div role=\"search\" class=\"sbx-custom__wrapper\">\n                                <input name=\"q\"\n                                       placeholder=\"Please type to search ...\"\n                                       value=\"")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(q)))
	_, _ = io.WriteString(w, "\"   autocomplete=\"off\" required=\"required\"\n                                       id=\"query\"  maxLength=100 role=\"textbox\"\n                                       accesskey=\"s\" type=\"text\" x-webkit-speech=\"\" speech=\"\" spellcheck=\"false\"\n                                       class=\"sbx-custom__input\" >\n                                <button type=\"submit\" title=\"Submit your search query.\" class=\"sbx-custom__submit\">\n                                    <svg role=\"img\" aria-label=\"Search\">\n                                        <use xmlns:xlink=\"http://www.w3.org/1999/xlink\" xlink:href=\"#sbx-icon-search-6\"></use>\n                                    </svg>\n                                </button>\n                                <button type=\"reset\" title=\"Clear the search query.\" class=\"sbx-custom__reset\">\n                                    <svg role=\"img\" aria-label=\"Reset\">\n                                        <use xmlns:xlink=\"http://www.w3.org/1999/xlink\" xlink:href=\"#sbx-icon-clear-3\"></use>\n                                    </svg>\n                                </button>\n                            </div>\n                        </form>\n                    </div>\n                    <div id=\"toolbar_baseline\" class=\"toolbar_baseline\"></div>\n                    <div style=\"padding-top:5px;padding-bottom: 5px;\">\n                        <SPAN style=\"float: left;\" class=nums>About ")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(response.GetTotal())))
	_, _ = io.WriteString(w, " results (<span>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(response.Took)))
	_, _ = io.WriteString(w, "ms</span>)</SPAN>\n                    </div>\n                </div>\n            </div>\n            <div style=\" clear: both;height: 10px;\"></div>\n            <div class=\"global_left\">\n\n                <div id=container >\n\n                    <div style=\" width: 600px;clear: both;\">\n\n\n                        <div style=\"height:0px; float:left; padding-left:585px;\">\n                            <div style=\"font-size: 12px;width:300px;WORD-BREAK: break-all; WORD-WRAP: break-word; overflow:hidden\">\n                                ")
	if len(response.Aggregations) > 0 {
		_, _ = io.WriteString(w, "\n\n                                <ul class=\"middle_column_ul\">\n                                ")

		for k, v := range response.Aggregations {
			if len(v.Buckets) > 0 {

				_, _ = io.WriteString(w, "\n                                    <li style=\"background: #eeeeee; \" >\n                                        <div style=\"padding:5px 10px; font-size: 14px;\">")
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(common.GetBucketLabel(k))))
				_, _ = io.WriteString(w, "</div>\n\n                                        <div style=\"padding:0px 0px; background:#f8f8f8;text-decoration: none;\">\n                                            <ul >\n                                ")

				for _, bkt := range v.Buckets {
					key := bkt.Key
					c := bkt.DocCount

					_, _ = io.WriteString(w, "\n                                <li><a href=")
					_, _ = fmt.Fprint(w, fmt.Sprintf("?q=%s&filter=%s|%v&from=0", q, common.GetBucketKey(k), util.UrlEncode(key)))
					_, _ = io.WriteString(w, " >")
					_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(key)))
					_, _ = io.WriteString(w, "<count>(")
					_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(c)))
					_, _ = io.WriteString(w, ")</count></a></li>\n\n                                ")
				}
				_, _ = io.WriteString(w, "\n                                            </ul>\n                                        </div>\n                                    </li>\n                                ")

			}
		}

		_, _ = io.WriteString(w, "\n\n                                </ul>\n                                ")

	}

	_, _ = io.WriteString(w, "\n                            </div>\n                        </div>\n\n                        ")

	hasResult := len(response.Hits.Hits) > 0

	if hasResult {
		_, _ = io.WriteString(w, "\n                        <div class=\"item-view\">\n                        ")

		for seq, hit := range response.Hits.Hits {
			url := common.SafeGetField(hit.Source["snapshot"].(map[string]interface{})["url"], "N/A")
			snapshotId := common.SafeGetField(hit.Source["snapshot"].(map[string]interface{})["id"], "")
			screenshot := common.SafeGetField(hit.Source["task"].(map[string]interface{})["last_screenshot_id"], "")
			title := common.SmartGetField(hit.Highlight["snapshot.title"], hit.Source["snapshot"].(map[string]interface{})["title"], "N/A")
			summary := util.SubStringWithSuffix(common.SmartGetField(hit.Highlight["snapshot.text"], hit.Source["snapshot"].(map[string]interface{})["text"], "N/A"), 500, "...")

			_, _ = io.WriteString(w, "\n                            <TABLE id=result_")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(seq)))
			_, _ = io.WriteString(w, " class=result cellSpacing=0 cellPadding=0> <TBODY><TR>\n                                <TD class=f>\n                                    <H3 class=t><A  title='")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(hit.Source["snapshot"].(map[string]interface{})["title"])))
			_, _ = io.WriteString(w, "' href=\"")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(url)))
			_, _ = io.WriteString(w, "\"  target=_blank>\n                                        ")
			_, _ = fmt.Fprint(w, util.SubStringWithSuffix(title, 500, "..."))
			_, _ = io.WriteString(w, "</A> </H3>\n                                    <FONT size=-1>")
			if screenshot != "" {
				_, _ = io.WriteString(w, "\n                                        <img class=\"screenshot\" src=\"/screenshot/")
				_, _ = fmt.Fprint(w, screenshot)
				_, _ = io.WriteString(w, "\" class=\"uk-icon-hover uk-icon-history\" />\n                                        ")
			}
			_, _ = io.WriteString(w, "\n                                        ")
			_, _ = fmt.Fprint(w, summary)
			_, _ = io.WriteString(w, " <BR>\n                                        <div  class=g><a target=_blank href=\"")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(url)))
			_, _ = io.WriteString(w, "\">")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(util.SubStringWithSuffix(url, 80, "..."))))
			_, _ = io.WriteString(w, "</a>\n                                            ")
			if snapshotId != "" {
				_, _ = io.WriteString(w, "\n                                            <a target=\"_blank\" href=\"/snapshot/")
				_, _ = fmt.Fprint(w, snapshotId)
				_, _ = io.WriteString(w, "\" class=\"uk-icon-hover uk-icon-history\"></a>\n                                            ")
			}
			_, _ = io.WriteString(w, "\n                                        </div>\n                                    </FONT>\n                                </TD>\n                            </TR></TBODY></TABLE>\n                        ")

		}
		_, _ = io.WriteString(w, "\n                        </div>\n                        ")

		paras := map[string]interface{}{}
		paras["q"] = q

		if filter != "" && strings.Contains(filter, "|") {
			paras["filter"] = filter
		}

		_, _ = io.WriteString(w, "\n\n                        ")
	} else {
		_, _ = io.WriteString(w, "\n                        <div class=\"uk-alert uk-alert-warning\"> Nothing found.</div>\n                        ")
	}
	_, _ = io.WriteString(w, "\n\n\n\n                        ")

	if hasResult && response.GetTotal() > 10 {

		_, _ = io.WriteString(w, "\n                        <div class=\"loadmore\">\n                            <div class=\"pnnext\" data-total=\"")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(response.GetTotal())))
		_, _ = io.WriteString(w, "\" data-from=\"")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(from)))
		_, _ = io.WriteString(w, "\" data-size=\"")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(size)))
		_, _ = io.WriteString(w, "\" data-load-text=\"Loading ...\">\n                                <span class=\"load-icon\"></span><span class=\"load-text\">Load More</span>\n                            </div>\n                            <p class=\"load-tips\"></p>\n                        </div>\n                        ")

	}
	_, _ = io.WriteString(w, "\n\n                        <div class=\"c-back\">\n                            <div class=\"back-top\">\n                                <svg class=\"icon\" viewBox=\"0 0 1024 1024\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" p-id=\"519\">\n                                    <path d=\"M873.6 419.2l-355.2-361.6c-9.6-9.6-22.4-9.6-32 0l-355.2 368c-9.6 9.6-9.6 22.4 0 32 9.6 9.6 22.4 9.6 32 0l316.8-329.6 0 828.8c0 12.8 9.6 22.4 22.4 22.4s22.4-9.6 22.4-22.4l0-822.4 310.4 316.8c9.6 9.6 22.4 9.6 32 0C883.2 441.6 883.2 425.6 873.6 419.2z\" p-id=\"520\">\n                                    </path>\n                                </svg>\n                            </div>\n                        </div>\n\n                    </div>\n\n                    <div class=\"copyright\">\n                        <br/>\n                        ")
	common.Copyright(w, config)
	_, _ = io.WriteString(w, "\n                    </div>\n\n                </div>\n\n            </div>\n\n        </div>\n    </div>\n</div>\n\n")
	CommonFooter(w)
	_, _ = io.WriteString(w, "\n\n<script src=\"/static/assets/uikit-2.27.1/js/components/pagination.min.js\"></script>\n\n</body>\n</html>\n")
	return nil
}
