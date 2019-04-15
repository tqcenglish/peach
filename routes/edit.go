/*
 * @Author: tqcenglish
 * @LastEditors: tqcenglish
 * @Email: tqcenglish#gmail.com
 * @Description: 一梦如是，总归虚无
 * @LastEditTime: 2019-04-15 14:14:19
 */

package routes

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
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

// UploadPage 上传页面
func UploadPage(ctx *context.Context) {
	ctx.HTML(200, "upload")
}

// Upload 上传
func Upload(ctx *context.Context) {
	ctx.SaveToFile("upload", "docs.zip")
	defer os.Remove("docs.zip")
	unzip("docs.zip", ".")
	setting.NewContext()
	models.NewContext()

	ctx.Redirect("/docs", 301)
}

// Download 下载
func Download(ctx *context.Context) {
	zipfile := "docs.zip"
	zipit("docs/", zipfile)
	defer os.Remove(zipfile)
	ctx.ServeFile(zipfile)
}

// http://blog.ralch.com/tutorial/golang-working-with-zip/
func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

//http://blog.ralch.com/tutorial/golang-working-with-zip/
func zipit(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			fmt.Printf(baseDir)
			fmt.Printf(source)
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
