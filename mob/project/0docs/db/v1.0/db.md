# 数据字典

### 1. 字典表[dds_dictionary_info]

| 字段名     | 类型        | 默认值 | 为空  |   约束   | 描述                 |
| ---------- | ----------- | :----: | :---: | :------: | :------------------- |
| id         | bigint(20)  |        |  否   | PK, SEQI | id                   |
| name       | varchar(64) |        |  否   |    DN    | 名称                 |
| value      | varchar(32) |        |  否   |    DI    | 值                   |
| type       | varchar(32) |        |  否   |    DT    | 类型                 |
| sort_no    | bigint(20)  |   0    |  否   |          | 排序值               |
| group_code | varchar(32) |   *    |  否   |          | 分组编号             |
| status     | tinyint(1)  |        |  否   |          | 状态 1: 禁用 0: 启用 |