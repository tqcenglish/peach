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

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Unknwon/com"
	"github.com/urfave/cli"
	"gopkg.in/ini.v1"

	"k-peach/pkg/bindata"
)

//New 初始化项目
var New = cli.Command{
	Name:   "new",
	Usage:  "Initialize a new Peach project",
	Action: runNew,
	Flags: []cli.Flag{
		stringFlag("target, t", "my.peach", "Directory to save project files"),
		boolFlag("yes, y", "Yes to all confirmations"),
	},
}

func checkYesNo() bool {
	var choice string
	fmt.Scanln(&choice)
	return strings.HasPrefix(strings.ToLower(choice), "y")
}

func toRed(str string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", str)
}

func toGreen(str string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", str)
}

func toYellow(str string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", str)
}

func restoreAssets(target, dir string) {
	if err := bindata.RestoreAssets(target, dir); err != nil {
		fmt.Printf(toRed("✗  %v\n"), err)
		os.Exit(1)
	}
}

func runNew(ctx *cli.Context) {
	target := ctx.String("target")
	if com.IsExist(target) && !ctx.Bool("yes") {
		fmt.Printf(toYellow("Directory '%s' already exists, do you want to overwrite?[N/y] "), target)
		if !checkYesNo() {
			os.Exit(0)
		}
	}

	fmt.Printf("➜  Creating '%s'...\n", target)
	os.MkdirAll(target, os.ModePerm)

	// Create default files.
	// ! templates 不支持内存模版，所以 custom 下必须存在
	// ! public 使用 内存 public
	// dirs := []string{"templates", "public", "docs"}
	dirs := []string{"docs"}
	for _, dir := range dirs {
		fmt.Printf("➜  Creating '%s'...\n", dir)
		os.RemoveAll(filepath.Join(target, dir))
		restoreAssets(target, dir)
	}

	// Create custom templates.
	yes := ctx.Bool("yes")
	if !yes {
		// ! 当前没有复制模版到目录，所以必须有自定义模版
		//fmt.Printf(toYellow("Do you want to use custom templates?[N/y] "))
		//yes = checkYesNo()
		yes = true
	}

	if yes {
		fmt.Println("➜  Creating 'custom/templates'...")
		restoreAssets(filepath.Join(target, "custom"), "templates")
		fmt.Println("➜  Creating 'custom/dict'...")
		restoreAssets(filepath.Join(target, "custom"), "dict.txt")

		// Update configuration to use custom templates.
		fmt.Println("➜  Updating custom configuration...")
		var cfg *ini.File
		var err error
		customPath := filepath.Join(target, "custom/app.ini")
		if com.IsExist(customPath) {
			cfg, err = ini.Load(customPath)
			if err != nil {
				fmt.Printf(toRed("✗  %v\n"), err)
				os.Exit(1)
			}
		} else {
			cfg = ini.Empty()
		}

		cfg.Section("page").Key("USE_CUSTOM_TPL").SetValue("true")
		if err = cfg.SaveTo(customPath); err != nil {
			fmt.Printf(toRed("✗  %v\n"), err)
			os.Exit(1)
		}
	}

	fmt.Println(toGreen("✓  Done!"))
}
