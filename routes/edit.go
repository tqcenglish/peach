/*
 * @Author: tqcenglish
 * @LastEditors: tqcenglish
 * @Email: tqcenglish#gmail.com
 * @Description: 一梦如是，总归虚无
 * @LastEditTime: 2019-04-11 18:01:58
 */

package routes

import (
	"fmt"
	"io/ioutil"
	"strings"

	"k-peach/models"
	"k-peach/pkg/context"
	"k-peach/pkg/setting"

	log "github.com/sirupsen/logrus"
)

//Edit 支持简单编辑
func Edit(ctx *context.Context) {
	// Check if auth is presented.
	authHead := ctx.Req.Header.Get("Authorization")
	if len(authHead) == 0 {
		authRequired(ctx)
		return
	}

	auths := strings.Fields(authHead)
	if len(auths) != 2 || auths[0] != "Basic" {
		ctx.Error(401)
		return
	}

	uname, passwd, err := basicAuthDecode(auths[1])
	if err != nil {
		ctx.Error(401)
		log.Error(err)
		return
	}

	// Check if auth is valid.
	if uname != "admin" || models.Protector.Users[uname] != encodeMd5(passwd) {
		ctx.Error(401)
		return
	}

	lang := ctx.Params("lang")
	dir := ctx.Params("dir")
	file := ctx.Params("filename")

	ctx.Data["FileName"] = file
	ctx.Data["Dir"] = dir
	data, err := ioutil.ReadFile(setting.Docs.Target + "/" + lang + "/" + dir + "/" + file)
	if err != nil {
		log.Error(err)
		ctx.Error(500)
		return
	}
	ctx.Data["Context"] = string(data)
	ctx.HTML(200, "edit")
}

// Update 更新文档
func Update(ctx *context.Context) {
	// Check if auth is presented.
	authHead := ctx.Req.Header.Get("Authorization")
	if len(authHead) == 0 {
		authRequired(ctx)
		return
	}

	auths := strings.Fields(authHead)
	if len(auths) != 2 || auths[0] != "Basic" {
		ctx.Error(401)
		return
	}

	uname, passwd, err := basicAuthDecode(auths[1])
	if err != nil {
		ctx.Error(401)
		log.Error(err)
		return
	}

	// Check if auth is valid.
	if uname != "admin" || models.Protector.Users[uname] != encodeMd5(passwd) {
		ctx.Error(401)
		return
	}

	lang := ctx.Params("lang")
	dir := ctx.Params("dir")
	file := ctx.Params("filename")
	context := ctx.Req.PostFormValue("context")

	log.Debug(setting.Docs.Target + "/" + lang + "/" + dir + "/" + file)
	err = ioutil.WriteFile(setting.Docs.Target+"/"+lang+"/"+dir+"/"+file, []byte(strings.TrimSpace(context)), 0644)
	if err != nil {
		log.Error(err)
		ctx.Error(500)
		return
	}
	// remove .md
	fileName := file[:len(file)-3]
	if fileName == "README" {
		ctx.JSON(200, map[string]string{"path": fmt.Sprintf("/docs/%s", dir)})
		return
	}
	ctx.JSON(200, map[string]string{"path": fmt.Sprintf("/docs/%s/%s", dir, fileName)})
}
