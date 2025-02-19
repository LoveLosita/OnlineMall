# [接口文档(Api docs)](https://g4kfvgyyq7.apifox.cn)

# 1.使用的技术栈

1. `Go`语言：项目全部由该编程语言编写。
2. `hertz`框架：项目基于该框架构建。
3. `JWT Token`：用于用户的鉴权以及Token的生成。
4. `bcrypt`：用于将密码加密存储，以及明文密码和加密后密码的对比。
5. `MySQL`：是项目的数据库。
6. 中间件：用于获取请求标头中的内容，并写入上下文，用于下一步操作。
7. 还有一些杂七杂八的用于测试项目的东西，我也记不清了，就不在此列出了。

# 2.项目功能（目前版本`1.5.6Beta`）

本项目类似⼀个商品网站，可以实现以下功能：

- [x] 用户的注册与登录
- [x] 加密存储用户密码
- [x] 添加商品分类
- [x] 修改用户的信息
- [x] 查询商品
- [x] 查看商品详情
- [x] 查看分类下的商品
- [x] 给商品进行评论（评论的增删查）
- [x] 商品加入购物车
- [x] 获取购物车所有商品
- [x] 搜索购物车中的商品
- [x] 初级下单（向前端返回⼀个订单order）
- [ ] 终极下单（影响库存，并模拟微信/支付宝的支付回调等）
- [x] 嵌套评论
- [x] 匿名评论
- [x] 浏览的商品的历史记录
- [x] 商家和管理员的出现（可以自由增删商品）
- [ ] 与AI助手的聊天（使用ChatAI接口）
- [ ] 加入验证码登录的操作（成本有限，可能会研究如何实现2FA）
- [ ] 使用缓存（`Redis`缓存）
- [x] 设计⼀套热度算法，使得能让该用户常看的商品分类下的商品出现在首页
- [ ] 用户可以与客服进行聊天（可能会用Python实现商家、管理员和用户的客户端，如果有时间的话）
- [ ] 部署到自己的服务器上，并且可以访问
- [x] 防止sql注入
- [ ] 考虑更多的安全性（xxs和csrf等）

# 3.项目结构

## 3.1.文件结构图

```go
OnlineMall
│
├── api                  // 存放所有的 API 接口定义
│   ├── auth.go           // 用户认证相关接口
│   ├── cart.go           // 购物车相关接口
│   ├── categories.go     // 商品分类相关接口
│   ├── order.go          // 订单管理相关接口
│   ├── product.go        // 商品管理相关接口
│   ├── review.go         // 商品评论相关接口
│   └── user.go           // 用户管理相关接口
│
├── auth                 // 负责身份认证的功能
│   ├── check_permission.go  // 检查用户权限
│   └── jwt_generator.go     // 生成和解析 JWT
│
├── cmd                  // 存放应用启动和初始化的逻辑
│   └── start.go           // 启动应用并初始化服务
│
├── dao                  // 数据库交互层
│   ├── db_connection.go  // 负责连接数据库事宜
│   ├── cart.go           // 购物车数据操作
│   ├── categories.go     // 商品分类数据操作
│   ├── order.go          // 订单数据操作
│   ├── product.go        // 商品数据操作
│   ├── review.go         // 商品评论数据操作
│   └── user.go           // 用户数据操作
│
├── middleware           // 存放中间件，如 token 校验
│   └── token_handler.go    // 处理 token 校验逻辑
│
├── model                // 存放数据模型
│   ├── auth.go           // 用户认证模型
│   ├── cart.go           // 购物车模型
│   ├── order.go          // 订单模型
│   ├── products.go       // 商品模型
│   ├── review.go         // 商品评论模型
│   ├── user.go           // 用户模型
│   └── map_to_slice.go   //类似于映射的键值对模型
│
│
├── respond              // 返回响应的处理
│   └── responses.go       // 定义统一的响应格式
│
├── routers              // 存放路由配置
│   └── router.go          // 配置所有 API 路由
│
├── service              // 业务逻辑层
│   ├── auth.go           // 用户认证相关业务逻辑
│   ├── cart.go           // 购物车相关业务逻辑
│   ├── categories.go     // 商品分类相关业务逻辑
│   ├── order.go          // 订单相关业务逻辑
│   ├── product.go        // 商品相关业务逻辑
│   ├── review.go         // 商品评论相关业务逻辑
│   └── user.go           // 用户相关业务逻辑
│
├── utils                   // 工具函数
│   ├── if_in.go             // 判断元素是否在序列中的工具函数集
│   ├── pwd_encryption.go    // 密码加密工具
│   ├── list_in_rank_out.go  // 传入数值切片，返回每个数值在此切片中的排名
│   └── map_to_int_slice.go  // 传入映射，传出由自定义键值对构成的结构体的切片
│ 
├── go.mod               // Go Modules 配置文件
├── main.go              // 项目的入口文件，在此启动项目
├── readme.md            // 项目的 README 文件的英文版本
└── readme_zh_cn.md      // 项目的 README 文件的中文版本 (本文件)
```

## 3.2.目录详细说明：

1. **api**：该文件夹包含了与 API 相关的接口定义，分别涉及到用户认证、购物车管理、商品分类、商品管理、订单处理和评论功能。每个文件对应一个功能模块的 API。
2. **auth**：处理用户认证相关的逻辑，主要负责 JWT 的生成与验证，确保用户身份的合法性。
3. **cmd**：包含应用的启动逻辑，`start.go` 是应用启动和初始化的入口，负责加载配置并启动服务。
4. **dao**：数据访问层，包含与数据库的交互操作，处理数据的增删改查。例如，`cart.go` 文件负责与购物车相关的数据交互，`order.go` 文件负责订单相关的数据操作等。
5. **middleware**：中间件层，通常用于处理请求前后需要的操作，如验证 token。
6. **model**：数据模型定义文件，通常与数据库中的表结构相关联。每个文件对应一个模块的数据模型，如 `cart.go` 定义了购物车的结构和属性，`products.go` 定义了商品的结构等。
7. **respond**：统一的响应格式处理模块，确保 API 返回的数据结构一致，便于前端处理。
8. **routers**：定义各个 API 路由及其对应的处理函数，管理请求路径和处理逻辑。
9. **service**：业务逻辑层，封装各个功能模块的具体业务操作，提供更高层次的功能接口。例如，`cart.go` 文件在服务层进行购物车相关的操作。
10. **utils**：存放工具类函数，如判断条件和密码加密等。
11. **go.mod**：Go Modules 配置文件，指定项目依赖的外部包和版本。
12. **main.go**：项目的入口文件，启动应用并进行初始化配置。
13. **readme.md**：项目的` README` 文件的英文版本，提供项目的基本信息、使用说明和技术栈等。
14. **readme_zh_cn.md**: 本文件。

## 3.3.`MySql`表单结构

### 3.3.1.结构图

![结构图](mysql_struct.jpg)

### 3.3.2.结构说明

1. **`users`（用户表）**
   - 作为核心表，存储用户的基本信息。
   - 其他表（如 `orders`、`carts`、`reviews`、`product_view_history`）都会与用户关联。
2. **`products`（商品表）**
   - 存储商品的详细信息。
   - 关联 `categories` 表（一个商品属于一个分类）。
   - 关联 `reviews` 表（用户可以对商品进行评论）。
   - 关联 `product_view_history`（记录用户的浏览历史）。
3. **`categories`（商品分类表）**
   - 存储商品的类别信息。
   - 通过 `category_id` 关联到 `products` 表。
4. **`carts`（购物车表）**
   - 关联 `users` 表（一个用户有一个购物车）。
   - 存储用户添加到购物车的商品（通常有 `product_id` 关联到 `products`）。
5. **`orders`（订单表）**
   - 关联 `users` 表（一个用户可以有多个订单）。
   - 关联 `order_items` 表（一个订单包含多个商品）。
6. **`order_items`（订单商品表）**
   - 关联 `orders`（一个订单包含多个商品）。
   - 关联 `products`（记录该订单中的具体商品信息）。
7. **`reviews`（商品评论表）**
   - 关联 `users`（评论是由用户提交的）。
   - 关联 `products`（评论是针对某个商品的）。
8. **`product_view_history`（商品浏览历史表）**
   - 关联 `users`（记录用户的浏览记录）。
   - 关联 `products`（存储用户浏览过的商品）。

### 3.3.3.结构总结

- **一对多**：
  - 一个 `users` 对应多个 `orders`（一个用户可以有多个订单）。
  - 一个 `orders` 对应多个 `order_items`（一个订单可以有多个商品）。
  - 一个 `products` 对应多个 `reviews`（一个商品可以有多个评论）。
  - 一个 `users` 对应多个 `reviews`（一个用户可以写多个评论）。
  - 一个 `categories` 对应多个 `products`（一个分类下有多个商品）。
  - 一个 `users` 对应多个 `product_view_history`（一个用户有多个浏览记录）。
- **多对多**（通过中间表实现）：
  - `orders` 和 `products` 是多对多关系（通过 `order_items` 关联）。
  - `users` 和 `products` 在 `carts` 里形成多对多（一个用户可以添加多个商品，每个商品也可以被多个用户添加到购物车）。

#  4.状态码的定义

| 状态码 | HTTP状态码 |         描述          |                             原因                             |                           解决方案                           |
| :----: | :--------: | :-------------------: | :----------------------------------------------------------: | :----------------------------------------------------------: |
| 20000  |    200     |         成功          |                              -                               |                              -                               |
| 40001  |    401     |      用户名错误       |   登录时，传入的用户名参数错误，在数据库中找不到匹配的用户   |                       传入正确的用户名                       |
| 40002  |    401     |       密码错误        |     登录时，传入的密码参数错误，无法和数据库中现存的匹配     |                        传入正确的密码                        |
| 40003  |    400     |      用户名无效       |           注册时，传入的用户名在数据库中已经存在了           |                       传入唯一的用户名                       |
| 40004  |    400     |       缺少参数        |                传入的参数数量小于所要求的数量                |                         传入足够参数                         |
| 40005  |    400     |     参数类型错误      |           传入的参数类型错误，导致无法和结构体绑定           |                         传入正确参数                         |
| 40006  |    400     |       参数过长        |                   传入的某个参数长度过于长                   |                      缩短过长参数的长度                      |
| 40007  |    400     |   用户名或密码错误    |                   传入的用户名或者密码错误                   |                    传入正确的用户名或密码                    |
| 40008  |    400     |       性别错误        |     传入的性别不属于（"male","female","other"）中的一种      |                        传入其中的一种                        |
| 40009  |    401     |       缺少token       |                    Header中未填写JWT key                     |                        在Header中填写                        |
| 40010  |    401     | jwt token签名方法无效 |                       JWT key格式错误                        |                     传入正确的JWT token                      |
| 40011  |    401     |       无效token       |                        JWT token无效                         |                     传入正确的JWT token                      |
| 40012  |    401     |       无效声明        |                     JWT token的声明无效                      |                     传入正确的JWT token                      |
| 40013  |    400     |   传入的用户id无效    |             在执行通过id查找用户信息时没找到用户             |                       传入正确的用户id                       |
| 40014  |    401     |       权限不够        |                    用户不是管理员或者店主                    |                 让管理员或者店主来执行此操作                 |
| 40015  |    404     |      分类不存在       |         在添加商品的时候，没有找到分类id所对应的分类         |                       传入正确的分类id                       |
| 40016  |    400     |   分类名称已经存在    |            在添加分类时，填写了一个重复的分类名称            |                传入不和现有分类重复的分类名称                |
| 40017  |    404     |      商品不存在       |              尝试通过ID找商品，但是不存在该商品              |                       传入存在商品的ID                       |
| 40018  |    404     |      找不到商品       |                  通过关键字搜索无法找到商品                  |                     传入存在商品的关键字                     |
| 40019  |    404     |     商品列表为空      |      在显示全部商品/显示某个分类的商品时，商品列表为空       |  前者，请先添加商品再进行其他操作；后者，请传入正确的分类id  |
| 40020  |    401     |   Refresh Token无效   | 在尝试通过刷新令牌接口刷新Access Token时，传入了无效的Refresh Token | 传入有效的Refresh Token。如果Refresh Token也过期了，就重新登录 |
| 40021  |    400     |  商品已经在购物车中   | 在尝试将一定数量的某商品加入购物车时，购物车里面已经有相同数量的同一个商品 |   如果是想更新数量，请传入数量不同的该商品；否则就换个商品   |
| 40022  |    400     |       数量太大        |    在尝试下单或者将商品加入购物车时，传入的数量超过了999     |                     请传入小于999的数量                      |
| 50001  |    500     |      订单不存在       | 在检查该用户是否购买了此商品时，订单表单和商品表单不匹配，属于内部错误 |                              -                               |
| 40024  |    400     | 用户没有购买过此商品  |        在用户尝试评论该商品时，发现用户没购买过此商品        |                     请购买此商品后再评论                     |
| 40025  |    400     |   用户已经评论过了    |       在用户尝试评论该商品时，系统发现用户已经评论过了       |                         请勿重复评论                         |
| 40026  |    400     |   用户打分超出范围    |                 用户的打分超出了1-5分的范围                  |                 请将对商品的分数打在此范围内                 |
| 40027  |    400     |   用户评论字数过长    |                用户在评论商品时，评论字数过长                |                      请用户缩短评论字数                      |
| 40028  |    400     |     父评论不存在      |          在用户尝试回复评论时，传入的父评论id不存在          |                     请传入存在的父评论id                     |
| 40029  |    404     |      购物车为空       |        在用户请求展示购物车全部商品时，购物车中无商品        |                      请先添加商品再展示                      |
| 40030  |    404     |     商品评论为空      |            在用户尝试查看某商品下评论时，评论为空            |                      请先评论再获取评论                      |
| 40031  |    404     |      找不到评论       | 在商家通过关键词搜索某商品下的评论时，没有找到符合要求的评论 |                         请更换关键词                         |
| 40032  |    404     |      评论不存在       |   在传入评论id需要进行查询/删除操作时，没有找到该id的评论    |                      请传入正确的评论id                      |

# 5.通用错误的返回示例

项目返回的一些错误是通用性的，所以我仅仅在项目初期的`apifox`接口文档编写中保存了其示例，在后期便没有再保存示例。在此，我将这些错误的示例列出，原因请自行查阅上方表格：

## 5.1.未登录

```json
{
    "status": "40009",
    "info": "missing token"
}
```

## 5.2.jwt token签名方法无效

```json
{
    "status": "40010",
    "info": "invalid signing method"
}
```

## 5.3.无效Token/Token过期

```json
{
    "status": "40011",
    "info": "invalid token"
}
```

## 5.4.Token的声明(claim)无效

```json
{
    "status": "40012",
    "info": "invalid claims"
}
```

## 5.5.权限不足

```json
{
    "status": "40014",
    "info": "unauthorized"
}
```

## 5.6.缺少参数

```json
{
    "status": "40004",
    "info": "missing param"
}
```

## 5.7.参数类型错误

```json
{
    "status": "40005",
    "info": "wrong param type"
}
```

## 5.8.参数过长

```json
{
    "status": "40006",
    "info": "param too long"
}
```

# 6.使用项目

## 6.1.配置数据库

请你创建名为`OnlineMall`的`mysql`数据库，设置密码为`123456`。

然后再创建如上方`2.3.2.结构说明`处的8张表单，使用如下语句，**按顺序**一一创建：

```mysql
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `full_name` varchar(100) DEFAULT NULL,
  `phone_number` varchar(20) DEFAULT NULL,
  `nickname` varchar(50) DEFAULT NULL,
  `qq` varchar(20) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `gender` enum('male','female','other') DEFAULT 'other',
  `bio` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `role` enum('user','merchant','admin') NOT NULL DEFAULT 'user',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
)
```

```mysql
CREATE TABLE `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `description` text,
  PRIMARY KEY (`id`)
)
```

```mysql
CREATE TABLE `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `description` text,
  `price` decimal(10,2) NOT NULL,
  `stock` int NOT NULL DEFAULT '0',
  `category_id` int DEFAULT NULL,
  `popularity` int DEFAULT '0',
  `ave_rating` decimal(3,2) DEFAULT '0.00',
  `product_image` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `products_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
)
```

```mysql
 CREATE TABLE `orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `order_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `status` enum('pending','completed','canceled') NOT NULL DEFAULT 'pending',
  `total_price` decimal(10,2) NOT NULL DEFAULT '0.00',
  `address` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
)
```

```mysql
CREATE TABLE `order_items` (
  `item_id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `product_id` int NOT NULL,
  `quantity` int NOT NULL,
  `price` decimal(10,2) NOT NULL,
  PRIMARY KEY (`item_id`),
  KEY `order_id` (`order_id`),
  KEY `product_id` (`product_id`),
  CONSTRAINT `order_items_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  CONSTRAINT `order_items_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
)
```

```mysql
 CREATE TABLE `reviews` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `product_id` int NOT NULL,
  `parent_id` int DEFAULT NULL,
  `rating` int NOT NULL,
  `comment` text,
  `is_anonymous` tinyint(1) DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `product_id` (`product_id`),
  KEY `parent_id` (`parent_id`),
  CONSTRAINT `reviews_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `reviews_ibfk_3` FOREIGN KEY (`parent_id`) REFERENCES `reviews` (`id`),
  CONSTRAINT `reviews_chk_1` CHECK (((`rating` = -(1)) or ((`rating` >= 1) and (`rating` <= 5))))
)
```

```mysql
CREATE TABLE `carts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `product_id` int NOT NULL,
  `quantity` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `product_id` (`product_id`),
  CONSTRAINT `carts_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `carts_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
)
```

## 6.2.启动项目

将项目部署在本地后，在项目根目录的终端中运行：

```bash
go mod tidy
```

以整理依赖。然后，再运行：

```bash
go run main.go
```

即可启动项目。
