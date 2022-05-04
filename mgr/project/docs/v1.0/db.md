# 数据字典

### 1. 表名称[数据库表名]

| 字段名      | 类型        | 默认值  | 为空  |              约束              | 描述     |
| ----------- | ----------- | :-----: | :---: | :----------------------------: | :------- |
| id          | bigint(20)  |         |  否   |        PK, SEQ,L,DI,LE         | 编号     |
| name        | varchar(64) |         |  否   |     UNQ,Q,L,DN,C,U,LE,D,x      | 标题     |
| status      | tinyint(1)  |    0    |  否   | L,Q,sl,C,U,color,LE,tp(switch) | 状态     |
| create_time | datetime    | sysdate |  否   |   Q,L,LE,f(yyyy-MM-dd HH:mm)   | 创建时间 |
