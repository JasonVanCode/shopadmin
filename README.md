# 微信小程序商城后台接口以及后台商城管理系统接口

[![Build Status][1]][2] [![Go Report Card][3]][4] [![MIT licensed][5]][6]

[1]: https://travis-ci.org/harlanc/moshopserver.svg?branch=master
[2]: https://travis-ci.org/harlanc/moshopserver
[3]: https://goreportcard.com/badge/github.com/harlanc/moshopserver
[4]: https://goreportcard.com/report/github.com/harlanc/moshopserver
[5]: https://img.shields.io/badge/license-MIT-blue.svg
[6]: LICENSE

## 介绍

- [shopadmin](https://github.com/JasonVanCode/shopadmin.git)的golang实现
- 基于[beego 2.0](https://github.com/beego/beego)开发



本项目需要配合微信小程序端使用，GitHub: [https://github.com/JasonVanCode/wechatshop](https://github.com/JasonVanCode/wechatshop)

本项目需要配合商城管理端，GitHub: [https://github.com/JasonVanCode/shop_vue](https://github.com/JasonVanCode/shop_vue)

## 测试环境搭建

- 克隆源码

        git clone https://github.com/JasonVanCode/shopadmin.git

- 下载所有依赖包

       go mod vendor

- 创建数据库nideshop并导入项目根目录下的nideshop.sql

        CREATE DATABASE `nideshop_new` DEFAULT CHARACTER SET utf8mb4 ;

- 配置好小程序相关字段

        [weixin] 
        #小程序 appid
        appid=""
        #小程序密钥
        secret="" 
  
- 配置好数据库相关字段
  
      [mysql]
      host = ""
      name = ""
      password = ""
      port = 3306
      database = ""
      charset = ""



- 配置好相关的图片服务器

       [attachment]
       #最大文件
       validate_size = 52428800
       validate_ext = "bmp,ico,psd,jpg,jpeg,png,gif,doc,docx,xls,xlsx,pdf,zip,rar,7z,tz,mp3,mp4,mov,swf,flv,avi,mpg,ogg,wav,flac,ape"
       #下面地址随便自定义，图片上传到自定义nginx 虚拟主机当中
       real_temp_path = ""
       real_save_path = ""

       #图片服务器实际存放的目录
       save_path = ""
       #图片服务器临时存放目录
       temp_path = ""
       #图片服务器地址 nginx配置的
       file_host = ""

- 运行以下命令（默认为开启8080端口）

        go run main.go






