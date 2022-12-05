# hycli


数据字典

### 1. 用户表[sso_user_info]

| 字段名          | 类型         | 默认值  | 为空  |       约束       | 描述                         |
| --------------- | ------------ | :-----: | :---: | :--------------: | :--------------------------- |
| user_id         | bigint(20)   |         |  否   |   PK, SEQ,L,DI   | 用户编号                     |
| full_name       | varchar(32)  |         |  否   |      L,Q,DN      | 姓名                         |
| user_name       | varchar(64)  |         |  否   |     UNQ  ,L      | 登录名                       |
| password        | varchar(32)  |         |  否   |                  | 密码                         |
| email           | varchar(32)  |         |  是   |        L         | 邮件地址                     |
| mobile          | varchar(12)  |         |  否   |        L         | 电话号码                     |
| create_time     | datetime     | sysdate |  否   |  L,Q,tp(daterange)  | 创建时间                     |
| status          | tinyint(1)   |    1    |  否   | L,Q(bl),sl,color | 状态 0: 正常 1: 锁定 2: 禁用 |
| source_id       | varchar(128) |    0    |  否   |                  | 来源id                       |
| source          | varchar(36)  |   ''    |  否   |      Q,sl,L      | 用户来源                     |
| last_login_time | datetime     |         |  是   |    L,Q(date)     | 最近登录                     |



生成页面代码
```sh
 hycli ui page ..\0docs\2.设计\db.mysql.md web --table "user_info"

```

