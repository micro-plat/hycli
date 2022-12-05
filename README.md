

### 生成前后端代码
命令: hycli mgr code ../db.md

### 生成数据库
命令: hicli db create ../docs/db/db.md .\api\modules\const\db\mysql -g


### 命令清单
| 指令          | 示例                | 说明                            |
| ------------- | ------------------- | ------------------------------- |
| row           | row(4)              | textarea的行数                  |
| fl            | fl                  | 显示到首行                      |
| tp(link)      | tp(link)            | 链接类型                        |
| tp(image)     | tp(image)           | 显示未图片                      |
| tp(switch)    | tp(switch)          | switch控件                      |
| tp(date)      | tp(date)            | 日期控件类型                    |
| tp(pwd)       | tp(pwd)             | 密码控件类型                    |
| tp(mobile)    | tp(mobile)          | 手机号                          |
| tp(progress)  | tp(progress)        | 查看时显示为进度条              |
| tp(daterange) | tp(daterange)       | 查看时显示为日期范围            |
| f             | f(yyyy-MM-dd HH:mm) | 字段格式                        |
| lw            | lw(180)             | 列表宽度                        |
| ds            | ds                  | 字典状态字段，值为0才查询出数据 |


###  控件类型说明
| 名称     | 标签       | 查询 | 新增/编辑 | 列表 | 详情 |
| -------- | ---------- | ---- | --------- | ---- | ---- |
| 链接     | link       | √    | ☓         | √    | √    |
| 图片     | image      | X    | X         | √    | √    |
| 切换     | switch     | √    | √         | √    | √    |
| 日期     | date       | √    | √         | X    | X    |
| 密码     | pwd        | X    | √         | X    | X    |
| 手机号   | mobile     | X    | √         | X    | X    |
| 进度条   | progress   | X    | X         | √    | √    |
| 日期范围 | daterange  | √    | X         | X    | X    |
| 月份范围 | monthrange | √    | X         | X    | X    |
| 年       | year       | √    | X         | X    | X    |
| 月       | month      | √    | X         | X    | X    |