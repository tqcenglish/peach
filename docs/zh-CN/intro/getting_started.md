---
name: 开始使用
---

## 开始使用

二进制文件 k-peach 包含了(templates, public) 等内容. 通过命令 k-peach new  创建知识库目录
```bash
➜  Downloads k-peach new
➜  Creating 'my.peach'...
➜  Creating 'docs'...
➜  Creating 'custom/templates'...
➜  Updating custom configuration...
✓  Done!
```
生成如下目录

```bash
➜  Downloads tree -L 2 my.peach
my.peach
├── custom
│   ├── app.ini
│   └── templates
└── docs
    ├── TOC.ini
    ├── images
    ├── protect.ini
    └── zh-CN

5 directories, 3 files
```
最后执行
```bash
k-peac web
```
访问 http://ip:5556 即可， 知识库内容在 **docs/zh-CN** 目录。

好了，让我们继续学习下一个部分：[创建文档仓库](../howto/documentation)。