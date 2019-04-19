// Copyright 2015 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package routes

import (
	"k-peach/models"
	"k-peach/pkg/context"
	"k-peach/pkg/setting"

	"github.com/Unknwon/com"
	log "github.com/sirupsen/logrus"
	"github.com/wangbin/jiebago"
)

var seg jiebago.Segmenter
var useJieBa = false

func init() {
	if com.IsFile("custom/dict.txt") {
		seg.LoadDictionary("custom/dict.txt")
		log.Info("load custom/dict.txt")
		useJieBa = true
	}
}

// Search 搜索处理
func Search(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("search")

	toc := models.Tocs[ctx.Locale.Language()]
	if toc == nil {
		toc = models.Tocs[setting.Docs.Langs[0]]
	}

	q := ctx.Query("q")
	if len(q) == 0 {
		ctx.Redirect(setting.Page.DocsBaseURL)
		return
	}

	ctx.Data["Keyword"] = q

	if useJieBa {
		log.Info("use JieBa")
		results := make([]*models.SearchResult, 0)
		qs := seg.CutForSearch(q, true)
		for word := range qs {
			for _, result := range toc.Search(word) {
				results = append(results, result)
			}
		}
		ctx.Data["Results"] = results
	} else {
		log.Info("not use JieBa")
		ctx.Data["Results"] = toc.Search(q)
	}

	ctx.HTML(200, "search")
}
