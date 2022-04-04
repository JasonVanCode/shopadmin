# moshopserver小程序商城后台系统

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



本项目需要配合微信小程序端使用，GitHub: [https://github.com/tumobi/nideshop-mini-program](https://github.com/tumobi/nideshop-mini-program)

## 测试环境搭建

- 克隆源码

        git clone https://github.com/JasonVanCode/shopadmin.git

- 下载所有依赖包

       go mod vendor

- 创建数据库nideshop并导入项目根目录下的nideshop.sql

        CREATE SCHEMA `nideshop_new` DEFAULT CHARACTER SET utf8mb4 ;

- 配置好小程序相关字段

        [weixin] 
        #小程序 appid
        appid=""
        #小程序密钥
        secret="" 
             

-  运行以下命令（默认为开启8080端口）

        go run main.go



## 功能列表
+ 首页
+ 分类首页、分类商品、新品首发、人气推荐商品页面
+ 商品详情页面，包含加入购物车、收藏商品、商品评论功能
+ 搜索功能
+ 专题功能
+ 品牌功能
+ 完整的购物流程，商品的加入、编辑、删除、批量选择，收货地址的选择，下单支付
+ 会员中心（订单、收藏、足迹、收货地址、意见反馈）
  ....






