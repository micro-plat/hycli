

### 生成前后端代码
命令: hycli mgr code ../db.md

### 生成数据库
命令: hicli db create ../docs/db/db.md .\api\modules\const\db\mysql -g

### 基础参数
| 指令 | 格式                     | 示例                                       | 说明                                    |
| ---- | ------------------------ | ------------------------------------------ | --------------------------------------- |
| l    | l[(format)]              | l,l(2),l(MM:dd HH:mm)                      | 标记字段显示到列表                      |
| le   | le[(format)]             | le,le(2)                                   | 标记字段显示到子列表信息                |
| c    | c[(format)]              | c(2)                                       | 标记字段为新增字段                      |
| u    | u[(format)]              | u                                          | 标记字段为修改字段                      |
| v    | v[(format)]              | v                                          | 标记字段显示到详情页面                  |
| q    | q[(format)]              | q                                          | 标记字段为查询字段                      |
| bq   | bq[(format)]             | bq(pwd)                                    | 标记字段为接口查询字段                  |
| row  | row[(format)]            | row(4)                                     | textarea的行数                          |
| lw   | lw[(format)]             | lw(180)                                    | 列表宽度(列表)                          |
| fl   | fl                       | fl                                         | 显示到首行,用于控制添加，编辑组件的位置 |
| f    | f[(format)])             | f(yyyy-MM-dd HH:mm)                        | 通用格式(列表/详情)                     |
| ps   | ps([类型])               | ps(newline)                                | 控件位置,c-u有效，仅支持newline         |
| sl   | sl([枚举名][,#父级字段]) | sl(area_info,*),sl(area_info,#province_no) | 标记字段为枚举字段，默认为下拉选择      |



### 关联表配置
| 指令 | 示例 | 说明                            |
| ---- | ---- | ------------------------------- |
| DI   |      | 字典编号                        |
| DN   |      | 字典名称                        |
| DT   |      | 字典类型                        |
| DP   |      | 字典扩展参数                    |
| DSN  |      | 字典排序字段                    |
| DS   |      | 字典状态，配置的字段值为0时有效 |


### 组件配置
| 指令                                  | 示例                                                                  | 说明         |
| ------------------------------------- | --------------------------------------------------------------------- | ------------ |
| view(名称，tab:表名[，样式设置])      | view(货架,tab:ots_supplier_shelf,style:{rwName:spp_no;fwName:spp_no}) | 预览页面组件 |
| lst(名称,类型:地址[，样式设置])       | lst(账户加款,dialog:/beanpay/balance,style:{showExpr:@status==0}))    | 列表组件     |
| lstat(名称,url:接口地址[，样式设置])  | lstat(今日交易,url:/ots/trade/order/lstat1day,style:{span:24})        | 字典类型     |
| chart(名称,类型:接口地址[，样式设置]) | chart(支付情况,bar:/ots/trade/order/as7d,style:{height:300px})        | 图表         |

| 样式      | 示例                               | 适用范围    | 说明                 |
| --------- | ---------------------------------- | ----------- | -------------------- |
| rwName    |                                    | 所有        | 本表字段名或标记     |
| fwName    |                                    | 所有        | 外部表字段名或标记   |
| showExpr  | showExpr:@status==0 showExpr:age>0 | lst         | 显示表达式           |
| span      | span:24                            | lstat chart | 分份数量             |
| height    | height:300px                       | lstat chart | 高度                 |
| rType     | rType:rose rType:radius            | chart       | pie图表类型          |
| showLabel | showLabel:false                    | chart       | 显示文字标签         |
| add       | add:true,add:分配                  | view        | 允许添加列表         |
| @字段名   | @pid:2                             | view        | 用于设置字段的固定值 |


### 数据库
| 指令 | 格式               | 示例                                  | 说明           |
| ---- | ------------------ | ------------------------------------- | -------------- |
| PK   | PK                 | PK                                    | 标记为主键     |
| UNQ  | UNQ[(name[,顺序])] | unq(unq_order_no),unq(unq_order_no,1) | 标记为唯一索引 |
| IDX  | IDX[(name[,顺序])] | idx(idx_order_no),idx(idx_order_no,1) | 标记为普通索引 |

### 格式参数
| 指令 | 示例                | 说明                                         |
| ---- | ------------------- | -------------------------------------------- |
| 数字 | f(2),l(1,8)         | 字段类型为数字时有效，用于数字字段的格式显示 |
| 日期 | yyyy-MM-dd HH:mm:ss | 日期格式，用于日期格式显示                   |

### 颜色组件
| 指令                        | 示例             | 说明                 |
| --------------------------- | ---------------- | -------------------- |
| color                       | color            | 颜色配置             |
| color[(字体颜色)]           | color(#33333)    | 自定义字体颜色       |
| color[(字体颜色，背景颜色)] | color(#333,#666) | 自定义字体，背景颜色 |

### 排序
| 指令             | 示例     | 说明           |
| ---------------- | -------- | -------------- |
| lidx(顺序号)     | lidx(1)  | 列表字段排序   |
| leidx(顺序号)    | leidx(1) | 子列表字段排序 |
| vidx(顺序号)     | vidx(1)  | 详情列表排序   |
| cidx(顺序号)     | cidx(1)  | 添加字段排序   |
| uidx(顺序号)     | uidx(1)  | 修改字段排序   |
| vidx(顺序号)     | vidx(1)  | 详情字段排序   |
| didx(顺序号)     | didx(1)  | 删除字段排序   |
| [tag]idx(顺序号) | midx(1)  | 自定义字段排序 |

* 字段默认排序编号为100+字段顺序号(期望排第四个则编号可设置为104)，显示到最前面则排序数字应小于100

### 操作栏按钮组
export(导出,url:/a/b),export(导入,url:/a/b),bcheck(修改,url:/a/b)

### 组件清单
| 指令                         | 示例               | 说明                                  |
| ---------------------------- | ------------------ | ------------------------------------- |
| tp(link[,页面类型-页面类型]) | tp(link,l-c-u)     | 显示为链接(列表/详情)                 |
| tp(image[,页面类型])         | tp(image,v)        | 显示为图片(列表/详情)                 |
| tp(switch[,页面类型])        | tp(switch)         | 显示为switch控件(所有)                |
| tp(pwd[,页面类型])           | tp(pwd)            | 显示为密码控件类型(新增/修改)         |
| tp(mobile[,页面类型])        | tp(mobile)         | 显示手机号(列表/详情下的显示格式)     |
| tp(progress[,页面类型])      | tp(progress)       | 显示为进度条(列表/详情)               |
| tp(tag[,页面类型])           | tp(tag)            | 显示为图章(列表/详情)                 |
| tp(date[,页面类型])          | tp(date)           | 显示日期控件类型(查询/新增/修改)      |
| tp(datetime[,页面类型])      | tp(datetime)       | 显示日期控件类型(查询/新增/修改)      |
| tp(daterange[,页面类型])     | tp(daterange)      | 显示为日期范围(查询)                  |
| tp(year[,页面类型])          | tp(year)           | 显示为年(所有)                        |
| tp(month[,页面类型])         | tp(month)          | 显示为月(所有)                        |
| tp(blist[,页面类型])         | tp(blist)          | 以当单选按钮组方式显示到查询栏        |
| tp(tabs[,页面类型])          | tp(tabs)           | 以tab栏显示到查询栏                   |
| tp(readonly[,页面类型])      | tp(readonly)       | 显示为只读                            |
| tp(cut[,页面类型][,长度])    | tp(cut,le,6)       | 列表详情页进行字符截取，长度为6各字符 |
| tp(ddl[,页面类型])           | tp(ddl,q)          | 搜索页面显示为下拉菜单                |
| tp(radio[,页面类型])         | tp(radio,c-u)      | 新增或修改页面的按钮单选              |
| tp(tabs,q,枚举名)            | tp(tabs,q,td_tabs) | tab标签为指定的枚举值，仅查询页面有效 |


### 组件
chart(交易统计,line:/ots/merchant/info/chart,style:{height:250px;span:12})
![image](https://github.com/micro-plat/hycli/raw/master/zimg/line.jpg)

chart(订单统计,bar:/ots/merchant/info/chart,style:{height:250px;span:12})
![image](https://github.com/micro-plat/hycli/raw/master/zimg/bar.jpg)


chart(订单图表1,pie:/ots/merchant/info/chart2,style:{height:250px;span:12;showLabel:false})
![image](https://github.com/micro-plat/hycli/raw/master/zimg/pie.jpg)



chart(订单图表2,pie:/ots/merchant/info/chart2,style:{height:250px;span:12;rType:radius})
![image](https://github.com/micro-plat/hycli/raw/master/zimg/pie-1.jpg)


chart(订单图3,pie:/ots/merchant/info/chart3,style:{height:250px;span:12})
![image](https://github.com/micro-plat/hycli/raw/master/zimg/pie-2.jpg)


chart(订图4,pie:/ots/merchant/info/chart3,style:{height:250px;span:12;rType:rose;showLabel:false})
![image](https://github.com/micro-plat/hycli/raw/master/zimg/pie-3.jpg)



![image](https://github.com/micro-plat/hycli/raw/master/zimg/1.jpeg)

![image](https://github.com/micro-plat/hycli/raw/master/zimg/2.jpg)

![image](https://github.com/micro-plat/hycli/raw/master/zimg/3.jpg)

![image](https://github.com/micro-plat/hycli/raw/master/zimg/4.jpg)

![image](https://github.com/micro-plat/hycli/raw/master/zimg/5.jpg)