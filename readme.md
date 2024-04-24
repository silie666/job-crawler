## boss,51job,智联，猎聘 招聘爬虫

> 做这个的目的是为了方便自己查看各大招聘网站的新岗位，
> 因为这些招聘网站都不会去显示岗位的发布日期，
> 有些岗位一直挂着，或者已读不回，
> 想找份合适的工作找了快两个月，岗位来来去去就那几个，一个个手动屏蔽太麻烦，一个个打开app太麻烦，
> 所以先爬一边把之前看过的岗位先记录起来，后续对其进行判断，
> 有新岗位出现再进行相应的提示

## 用法

1、修改`config.example`为`config.ini`

2、修改对应的配置文件

| 配置项                        | 说明                                      |
|----------------------------|-----------------------------------------|
| **params**                 |
| boss_scene                 | 不用管                                     |
| boss_city                  | 城市code   （例如：101280100=广州）              |
| boss_position              | 职位code    （例如：100103=php）               |
| boss_multiBusinessDistrict | 区域code （例如：440118,440112 = 黄埔+增城）空代表全广州 |
| boss_page                  | 页码                                      |
| boss_page_size             | 数量                                      |
| liepin_city                | 城市code  （例如：050020=广州）                  |
| liepin_dq                  | 区域code  （例如：050020050=黄埔）               |
| liepin_current_page        | 页码                                      |
| liepin_page_size           | 数量                                      |
| liepin_key                 | 关键词     （例如：php）                        |
| zhilian_kw                 | 关键词     （例如：php）                        |
| zhilian_p                  | 页码                                      |
| wyjob_job_area             | 城市code   （例如：030200=广州）                 |
| wyjob_keyword              | 关键词     （例如：php）                        |
| wyjob_search_type          | 不用管                                     |
| wyjob_page_num             | 页码                                      |
| wyjob_page_size            | 数量                                      |
| **cookie**                 | 
| boss_wt2                   | boss的wt2                                |
| zhilian_x_zp_client_id     | 智联的x-zp-client-id                       |
| zhilian_acw_tc             | 智联的acw_tc                               |
| zhilian_at                 | 智联的at                                   |
| zhilian_rt                 | 智联的rt                                   |

> cookie部分是因为这些boss和智联需要登录才能看到被隐藏的岗位


3、运行命令
```
go run main.go
```
