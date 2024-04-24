package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"net/http"
	"newjobs/config"
	"newjobs/response"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	db, _ := leveldb.OpenFile("./data/db", nil)
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"

	// boss直聘
	bossJob(db, userAgent)
	// 猎聘
	liePinJob(db, userAgent)
	// 智联招聘
	zhiLianJob(db, userAgent)
	// 51job
	wyJob(db, userAgent)

	defer db.Close()
}

func contains(s []string, target string) bool {
	for _, v := range s {
		if strings.ToLower(v) == target {
			return true
		}
	}
	return false
}

func bossJob(db *leveldb.DB, userAgent string) {
	// 获取boss直聘的新岗位
	var networkCookies []*network.Cookie
	var opts []chromedp.ExecAllocatorOption
	flags := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // 是否隐藏浏览器窗口，如果是win10要测试改为false
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.UserAgent(userAgent),
	}
	opts = chromedp.DefaultExecAllocatorOptions[:]
	opts = append(opts, flags...)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	if err := chromedp.Run(ctx,
		network.ClearBrowserCache(),
		chromedp.Navigate("https://www.zhipin.com/web/geek/job"),
		chromedp.Sleep(5*time.Second),
		chromedp.MouseClickXY(100, 100),
		chromedp.WaitVisible(`//*[@id="wrap"]/div[2]/div[2]/div/div[1]/div[1]`, chromedp.BySearch),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			networkCookies, err = network.GetCookies().Do(ctx)
			return err
		}),
	); err != nil {
		fmt.Println(err)
	}

	bossClient := req.C().
		SetUserAgent(userAgent).
		SetTLSFingerprintChrome()
	bossRequest := bossClient.R().
		SetQueryParams(map[string]string{
			"scene":                 config.Env.GetString("params.boss_scene"),
			"city":                  config.Env.GetString("params.boss_city"),
			"position":              config.Env.GetString("params.boss_position"),
			"multiBusinessDistrict": config.Env.GetString("params.boss_multiBusinessDistrict"),
			"page":                  config.Env.GetString("params.boss_page"),
			"pageSize":              config.Env.GetString("params.boss_page_size"),
		})
	cookies := make(map[string]string)
	for _, cookie := range networkCookies {
		cookies[cookie.Name] = cookie.Value
	}
	cookies["wt2"] = config.Env.GetString("cookie.boss_wt2")
	for name, value := range cookies {
		bossRequest = bossRequest.SetCookies(&http.Cookie{
			Name:  name,
			Value: value,
		})
	}
	resp, err := bossRequest.Get("https://www.zhipin.com/wapi/zpgeek/search/joblist.json")
	if err != nil {
		return
	}
	if resp.IsSuccessState() {
		var list response.BossJobs
		json.Unmarshal(resp.Bytes(), &list)

		// 访问成功
		if list.Code == 0 {
			for _, job := range list.ZpData.JobList {
				value, _ := db.Get([]byte("boss-"+job.EncryptJobId), nil)
				if value == nil {
					jobJson, _ := json.Marshal(job)
					_ = db.Put([]byte("boss-"+job.EncryptJobId), jobJson, nil)
					output := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
						color.Green.Sprint("boss新岗位"),
						color.Red.Sprint("新增职位："+job.JobName),
						color.Red.Sprint("福利待遇："+strings.Join(job.WelfareList, "、")),
						color.Yellow.Sprint("公司："+job.BrandName),
						color.Yellow.Sprint("薪资："+job.SalaryDesc),
						color.Yellow.Sprint("最后更新时间："+time.Unix(job.LastModifyTime/1000, 0).Format("2006-01-02 15:04:05")),
						color.BgWhite.Sprint("学历要求："+job.JobDegree),
						color.BgWhite.Sprint("经验要求："+job.JobExperience),
						color.BgWhite.Sprint("地址："+job.CityName+job.BusinessDistrict),
					)
					fmt.Println(output)
				}
			}
			color.Green.Println("boss的新岗位已查询完毕~")
		} else {
			color.Red.Println("boss的新岗位查询失败！原因：" + list.Message)
		}
	}
}

func liePinJob(db *leveldb.DB, userAgent string) {
	liePinClient := req.C().SetUserAgent(userAgent).
		SetTLSFingerprintChrome()
	liePinRequest := liePinClient.R().
		SetHeaders(map[string]string{
			"X-Client-Type":    "web",
			"X-Fscp-Std-Info":  "{\"client_id\": \"40108\"}",
			"X-Fscp-Trace-Id":  "0ce99156-360a-4a18-87d7-9342ae09e901",
			"X-Fscp-Version":   "1.1",
			"X-Requested-With": "XMLHttpRequest",
		}).
		SetBody(map[string]interface{}{
			"data": map[string]interface{}{
				"mainSearchPcConditionForm": map[string]interface{}{
					"city":        config.Env.GetString("params.liepin_city"),
					"dq":          config.Env.GetString("params.liepin_dq"),
					"currentPage": config.Env.GetString("params.liepin_current_page"),
					"pageSize":    config.Env.GetInt("params.liepin_page_size"),
					"key":         config.Env.GetString("params.liepin_key"),
				},
				"passThroughForm": map[string]interface{}{
					"scene": "condition",
					"skId":  "r8ndbjhi8ffzlueyynuh0bd2znv6ia2a",
					"fkId":  "r8ndbjhi8ffzlueyynuh0bd2znv6ia2a",
					"ckId":  "e0cxy0mf2izboophfeij4j8lcgjylrv8",
				},
			},
		})
	resp, err := liePinRequest.Post("https://api-c.liepin.com/api/com.liepin.searchfront4c.pc-search-job")
	if err != nil {
		return
	}
	if resp.IsSuccessState() {
		var list response.LiePinJobs
		json.Unmarshal(resp.Bytes(), &list)

		if list.Flag == 1 {
			for _, job := range list.Data.Data.JobCardList {
				// 猎聘没有专门的搜索标签条件，只能通过判断筛选
				if contains(job.Job.Labels, config.Env.GetString("params.liepin_key")) {
					//_ = db.Delete([]byte("liepin-"+jobId), nil)
					value, _ := db.Get([]byte("liepin-"+job.Job.JobId), nil)
					if value == nil {
						jobJson, _ := json.Marshal(job)
						err = db.Put([]byte("liepin-"+job.Job.JobId), jobJson, nil)

						formattedTime, _ := time.Parse("20060102150405", job.Job.RefreshTime)
						output := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
							color.Green.Sprint("猎聘新岗位"),
							color.Red.Sprint("新增职位："+job.Job.Title),
							color.Red.Sprint("职位标签："+strings.Join(job.Job.Labels, "、")),
							color.Yellow.Sprint("公司："+job.Comp.CompName),
							color.Yellow.Sprint("薪资："+job.Job.Salary),
							color.Yellow.Sprint("最后更新时间："+formattedTime.Format("2006-01-02 15:04:05")),
							color.BgWhite.Sprint("学历要求："+job.Job.RequireEduLevel),
							color.BgWhite.Sprint("经验要求："+job.Job.RequireWorkYears),
							color.BgWhite.Sprint("地址："+job.Job.Dq),
						)
						fmt.Println(output)
					}
				}
			}
			color.Green.Println("猎聘的新岗位已查询完毕~")
		} else {
			color.Red.Println("猎聘的新岗位查询失败！")
		}
	}
}

func zhiLianJob(db *leveldb.DB, userAgent string) {
	var opts []chromedp.ExecAllocatorOption
	flags := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // 是否隐藏浏览器窗口，如果是win10要测试改为false
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.UserAgent(userAgent),
	}
	opts = chromedp.DefaultExecAllocatorOptions[:]
	opts = append(opts, flags...)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var evens []network.RequestID
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventResponseReceived:
			if strings.Contains(ev.Response.URL, "https://fe-api.zhaopin.com/c/i/search/positions") {
				evens = append(evens, ev.RequestID)
			}
		case *network.EventLoadingFinished:
			i := slices.Index(evens, ev.RequestID)
			if i < 0 {
				break
			}
			evens = slices.Delete(evens, i, i+1)
			go func() {
				c := chromedp.FromContext(ctx)
				rbp := network.GetResponseBody(ev.RequestID)
				body, err := rbp.Do(cdp.WithExecutor(ctx, c.Target))
				if err != nil {
					fmt.Println(err)
				}
				var list response.ZhiLianJobs
				_ = json.Unmarshal(body, &list)
				if list.Code == 200 {
					for _, job := range list.Data.List {
						jobId := strconv.FormatInt(job.JobId, 10)
						//_ = db.Delete([]byte("zhilian-"+jobId), nil)
						value, _ := db.Get([]byte("zhilian-"+jobId), nil)
						if value == nil || len(value) == 0 {
							jobJson, _ := json.Marshal(job)
							_ = db.Put([]byte("zhilian-"+jobId), jobJson, nil)
							output := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
								color.Green.Sprint("智联招聘新岗位"),
								color.Red.Sprint("新增职位："+job.Name),
								color.Red.Sprint("福利待遇："+strings.Join(job.JobKnowledgeWelfareFeatures, "、")),
								color.Yellow.Sprint("公司："+job.CompanyName),
								color.Yellow.Sprint("薪资："+job.Salary60),
								color.Yellow.Sprint("最后更新时间："+job.PublishTime),
								color.BgWhite.Sprint("学历要求："+job.Education),
								color.BgWhite.Sprint("经验要求："+job.WorkingExp),
								color.BgWhite.Sprint("地址："+job.WorkCity+job.CityDistrict+job.StreetName),
							)
							fmt.Println(output)
						}
					}
					color.Green.Println("智联招聘的新岗位已查询完毕~")
				}
			}()
		}
	})

	if err := chromedp.Run(ctx,
		network.ClearBrowserCache(),
		network.Enable(),
		network.SetCookies([]*network.CookieParam{
			{Name: "x-zp-client-id", Value: config.Env.GetString("cookie.zhilian_x_zp_client_id"), Domain: ".zhaopin.com", Path: "/"},
			{Name: "acw_tc", Value: config.Env.GetString("cookie.zhilian_acw_tc"), Domain: ".zhaopin.com", Path: "/"},
			{Name: "at", Value: config.Env.GetString("cookie.zhilian_at"), Domain: ".zhaopin.com", Path: "/"},
			{Name: "rt", Value: config.Env.GetString("cookie.zhilian_rt"), Domain: ".zhaopin.com", Path: "/"},
		}),
		chromedp.Navigate("https://sou.zhaopin.com/?kw="+config.Env.GetString("params.zhilian_key")+"&p="+config.Env.GetString("params.zhilian_p")),
		chromedp.Sleep(5*time.Second),
		chromedp.WaitVisible("/html/body/div/div[4]/div[2]/div[1]/ul/li[3]/a", chromedp.BySearch),
		chromedp.Click("/html/body/div/div[4]/div[2]/div[1]/ul/li[3]/a", chromedp.BySearch),
		chromedp.Sleep(5*time.Second),
	); err != nil {
		fmt.Println(err)
	}
}

func wyJob(db *leveldb.DB, userAgent string) {

	var networkCookies []*network.Cookie
	var opts []chromedp.ExecAllocatorOption
	flags := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // 是否隐藏浏览器窗口，如果是win10要测试改为false
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.UserAgent(userAgent),
	}
	opts = chromedp.DefaultExecAllocatorOptions[:]
	opts = append(opts, flags...)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	if err := chromedp.Run(ctx,
		network.ClearBrowserCache(),
		chromedp.Navigate("https://we.51job.com/pc/search?"+
			"jobArea="+
			config.Env.GetString("params.wyjob_job_area")+
			"&keyword="+
			config.Env.GetString("params.wyjob_keyword")+
			"&searchType="+config.Env.GetString("params.wyjob_search_type")),
		chromedp.Sleep(5*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			networkCookies, err = network.GetCookies().Do(ctx)
			return err
		}),
	); err != nil {
		fmt.Println(err)
	}
	cookies := make(map[string]string)
	for _, cookie := range networkCookies {
		cookies[cookie.Name] = cookie.Value
	}

	wyJobClient := req.C().
		SetUserAgent(userAgent).
		SetTLSFingerprintChrome()
	wyJobRequest := wyJobClient.R().
		SetHeaders(map[string]string{
			"Referer":    "https://we.51job.com/",
			"Host":       "we.51job.com",
			"Accept":     "*/*",
			"Connection": "keep-alive",
		}).
		SetQueryParams(map[string]string{
			"keyword":    config.Env.GetString("params.wyjob_keyword"),
			"searchType": config.Env.GetString("params.wyjob_search_type"),
			"jobArea":    config.Env.GetString("params.wyjob_job_area"),
			"pageNum":    config.Env.GetString("params.wyjob_page_num"),
			"pageSize":   config.Env.GetString("params.wyjob_page_size"),
		},
		)
	for name, value := range cookies {
		wyJobRequest = wyJobRequest.SetCookies(&http.Cookie{
			Name:  name,
			Value: value,
		})
	}
	resp, err := wyJobRequest.Get("https://we.51job.com/api/job/search-pc")
	if err != nil {
		return
	}
	if resp.IsSuccessState() {
		var list response.WYJob
		json.Unmarshal(resp.Bytes(), &list)

		if list.Status == "1" {
			for _, job := range list.Resultbody.Job.Items {
				if strings.Contains(job.JobName, config.Env.GetString("params.wyjob_keyword")) {
					//_ = db.Delete([]byte("wyjob-"+job.jobId), nil)
					value, _ := db.Get([]byte("wyjob-"+job.JobId), nil)
					if value == nil {
						jobJson, _ := json.Marshal(job)
						err = db.Put([]byte("wyjob-"+job.JobId), jobJson, nil)

						output := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
							color.Green.Sprint("51job新岗位"),
							color.Red.Sprint("新增职位："+job.JobName),
							color.Red.Sprint("职位标签："+strings.Join(job.JobTags, "、")),
							color.Yellow.Sprint("公司："+job.CompanyName),
							color.Yellow.Sprint("薪资："+job.ProvideSalaryString),
							color.Yellow.Sprint("最后更新时间："+job.IssueDateString),
							color.BgWhite.Sprint("学历要求："+job.DegreeString),
							color.BgWhite.Sprint("经验要求："+job.WorkYearString),
							color.BgWhite.Sprint("地址："+job.JobAreaLevelDetail.ProvinceString+job.JobAreaLevelDetail.CityString+job.JobAreaLevelDetail.LandMarkString),
						)
						fmt.Println(output)
					}
				}
			}
			color.Green.Println("51job的新岗位已查询完毕~")
		} else {
			color.Red.Println("51job的新岗位查询失败！")
		}
	}
}
