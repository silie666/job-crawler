package response

type BossJobs struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	ZpData  struct {
		ResCount     int         `json:"resCount"`
		FilterString string      `json:"filterString"`
		Lid          string      `json:"lid"`
		HasMore      bool        `json:"hasMore"`
		JobList      []JobList   `json:"jobList"`
		TotalCount   int         `json:"totalCount"`
		BrandCard    interface{} `json:"brandCard"`
	} `json:"zpData"`
}

type JobList struct {
	SecurityId       string        `json:"securityId"`
	BossAvatar       string        `json:"bossAvatar"`
	BossCert         int           `json:"bossCert"`
	EncryptBossId    string        `json:"encryptBossId"`
	BossName         string        `json:"bossName"`
	BossTitle        string        `json:"bossTitle"`
	GoldHunter       int           `json:"goldHunter"`
	BossOnline       bool          `json:"bossOnline"`
	EncryptJobId     string        `json:"encryptJobId"`
	ExpectId         int           `json:"expectId"`
	JobName          string        `json:"jobName"`
	Lid              string        `json:"lid"`
	SalaryDesc       string        `json:"salaryDesc"`
	JobLabels        []string      `json:"jobLabels"`
	JobValidStatus   int           `json:"jobValidStatus"`
	IconWord         string        `json:"iconWord"`
	Skills           []string      `json:"skills"`
	JobExperience    string        `json:"jobExperience"`
	DaysPerWeekDesc  string        `json:"daysPerWeekDesc"`
	LeastMonthDesc   string        `json:"leastMonthDesc"`
	JobDegree        string        `json:"jobDegree"`
	CityName         string        `json:"cityName"`
	AreaDistrict     string        `json:"areaDistrict"`
	BusinessDistrict string        `json:"businessDistrict"`
	JobType          int           `json:"jobType"`
	ProxyJob         int           `json:"proxyJob"`
	ProxyType        int           `json:"proxyType"`
	Anonymous        int           `json:"anonymous"`
	Outland          int           `json:"outland"`
	Optimal          int           `json:"optimal"`
	IconFlagList     []interface{} `json:"iconFlagList"`
	ItemId           int           `json:"itemId"`
	City             int           `json:"city"`
	IsShield         int           `json:"isShield"`
	AtsDirectPost    bool          `json:"atsDirectPost"`
	Gps              interface{}   `json:"gps"`
	LastModifyTime   int64         `json:"lastModifyTime"`
	EncryptBrandId   string        `json:"encryptBrandId"`
	BrandName        string        `json:"brandName"`
	BrandLogo        string        `json:"brandLogo"`
	BrandStageName   string        `json:"brandStageName"`
	BrandIndustry    string        `json:"brandIndustry"`
	BrandScaleName   string        `json:"brandScaleName"`
	WelfareList      []string      `json:"welfareList"`
	Industry         int           `json:"industry"`
	Contact          bool          `json:"contact"`
}

type LiePinJobs struct {
	Flag int `json:"flag"`
	Data struct {
		Data struct {
			JobCardList      []JobCardList `json:"jobCardList"`
			CompList         []interface{} `json:"compList"`
			JobSubscribeInfo struct {
				Keyword string `json:"keyword"`
				QrUrl   string `json:"qrUrl"`
				DqName  string `json:"dqName"`
				DqCode  string `json:"dqCode"`
				Id      string `json:"id"`
			} `json:"jobSubscribeInfo"`
		} `json:"data"`
		PassThroughData struct {
			Sfrom string `json:"sfrom"`
			Scene string `json:"scene"`
			CkId  string `json:"ckId"`
			SkId  string `json:"skId"`
			FkId  string `json:"fkId"`
		} `json:"passThroughData"`
		Pagination struct {
			CurrentPage int  `json:"currentPage"`
			PageSize    int  `json:"pageSize"`
			TotalCounts int  `json:"totalCounts"`
			TotalPage   int  `json:"totalPage"`
			HasNext     bool `json:"hasNext"`
		} `json:"pagination"`
	} `json:"data"`
}

type JobCardList struct {
	Comp struct {
		CompScale    string `json:"compScale"`
		Link         string `json:"link"`
		CompName     string `json:"compName"`
		CompLogo     string `json:"compLogo"`
		CompId       int    `json:"compId,omitempty"`
		CompIndustry string `json:"compIndustry"`
		CompStage    string `json:"compStage,omitempty"`
	} `json:"comp"`
	Job struct {
		Labels           []string `json:"labels"`
		G                string   `json:"g"`
		RequireWorkYears string   `json:"requireWorkYears"`
		Salary           string   `json:"salary"`
		Link             string   `json:"link"`
		RefreshTime      string   `json:"refreshTime"`
		TopJob           bool     `json:"topJob"`
		J                string   `json:"j"`
		Dq               string   `json:"dq"`
		JobKind          string   `json:"jobKind"`
		JobId            string   `json:"jobId"`
		Title            string   `json:"title"`
		RequireEduLevel  string   `json:"requireEduLevel"`
		AdvViewFlag      bool     `json:"advViewFlag"`
		PcOuterLink      string   `json:"pcOuterLink,omitempty"`
		H5OuterLink      string   `json:"h5OuterLink,omitempty"`
		DataPromId       string   `json:"dataPromId"`
	} `json:"job"`
	DataInfo  string `json:"dataInfo"`
	Recruiter struct {
		ImId           string `json:"imId"`
		ImUserType     string `json:"imUserType"`
		Chatted        bool   `json:"chatted"`
		RecruiterId    string `json:"recruiterId"`
		RecruiterName  string `json:"recruiterName"`
		RecruiterTitle string `json:"recruiterTitle"`
		RecruiterPhoto string `json:"recruiterPhoto"`
	} `json:"recruiter"`
	DataParams string `json:"dataParams"`
}

type ZhiLianJobs struct {
	Code int `json:"code"`
	Data struct {
		AbGroupType string `json:"abGroupType"`
		Attachments struct {
			SearchNullMethodGroup string `json:"searchNullMethodGroup"`
			SearchNullMethod      string `json:"searchNullMethod"`
		} `json:"attachments"`
		Count                int  `json:"count"`
		Filter               bool `json:"filter"`
		IsApprovedInvestment int  `json:"isApprovedInvestment"`
		IsEndPage            int  `json:"isEndPage"`
		IsVerification       int  `json:"isVerification"`
		JobTypes             struct {
			Main []interface{} `json:"main"`
			Sub  []interface{} `json:"sub"`
		} `json:"jobTypes"`
		Kw   string `json:"kw"`
		List []struct {
			AbroadFlag    int `json:"abroadFlag"`
			AbroadTipInfo struct {
				AbroadTips []interface{} `json:"abroadTips"`
				Icon       string        `json:"icon"`
				Title      string        `json:"title"`
			} `json:"abroadTipInfo"`
			AlreadyCallPhone  bool   `json:"alreadyCallPhone"`
			ApplyType         string `json:"applyType"`
			CampusBestCompany struct {
				BestCompanyUrl string `json:"bestCompanyUrl"`
				HomepageType   int    `json:"homepageType"`
				LogoTagUrl     string `json:"logoTagUrl"`
				State          int    `json:"state"`
			} `json:"campusBestCompany"`
			CampusRootOrgInfo   interface{} `json:"campusRootOrgInfo"`
			CanBeRegular        bool        `json:"canBeRegular"`
			CanRemoteInternship bool        `json:"canRemoteInternship"`
			CardCustomJson      string      `json:"cardCustomJson"`
			CardType            int         `json:"cardType"`
			ChatWindow          int         `json:"chatWindow"`
			CityDistrict        string      `json:"cityDistrict"`
			CityId              string      `json:"cityId"`
			CommercialLabel     []struct {
				Type          int    `json:"type"`
				TypeName      string `json:"typeName"`
				TypeShowLabel string `json:"typeShowLabel"`
			} `json:"commercialLabel"`
			CompanyId               int         `json:"companyId"`
			CompanyLogo             string      `json:"companyLogo"`
			CompanyName             string      `json:"companyName"`
			CompanyNumber           string      `json:"companyNumber"`
			CompanyRootId           int         `json:"companyRootId"`
			CompanyScaleTypeTagsNew []string    `json:"companyScaleTypeTagsNew"`
			CompanySize             string      `json:"companySize"`
			CompanyUrl              string      `json:"companyUrl"`
			ComplainFlag            bool        `json:"complainFlag"`
			DeliveryPath            string      `json:"deliveryPath"`
			DisplayPhoneNumber      bool        `json:"displayPhoneNumber"`
			Distance                int         `json:"distance"`
			DistanceFormat          string      `json:"distanceFormat"`
			DistanceText            string      `json:"distanceText"`
			Education               string      `json:"education"`
			ExperimentInfo          interface{} `json:"experimentInfo"`
			Extensions              interface{} `json:"extensions"`
			FeatureServer           struct {
				LastReplyTime               int64   `json:"lastReplyTime,omitempty"`
				LastReplyTimeText           string  `json:"lastReplyTimeText"`
				Reply24H                    int     `json:"reply24h,omitempty"`
				StaffAvgFirstResponseTime7D int     `json:"staffAvgFirstResponseTime7d,omitempty"`
				StaffAvgHandleResumeTime30D int     `json:"staffAvgHandleResumeTime30d,omitempty"`
				StaffHandleResumeCnts30D    int     `json:"staffHandleResumeCnts30d,omitempty"`
				StaffReplyRate30D           float64 `json:"staffReplyRate30d,omitempty"`
				TodayReplyNum               int     `json:"todayReplyNum"`
				TodayReplyNumText           string  `json:"todayReplyNumText"`
			} `json:"featureServer"`
			FeedOperation  interface{} `json:"feedOperation"`
			FeedPosition   interface{} `json:"feedPosition"`
			FinancingStage struct {
				Code int    `json:"code,omitempty"`
				Name string `json:"name"`
			} `json:"financingStage"`
			FirstPublishTime            string        `json:"firstPublishTime"`
			HasAppliedPosition          bool          `json:"hasAppliedPosition"`
			IndustryCompanyTags         []string      `json:"industryCompanyTags"`
			IndustryName                string        `json:"industryName"`
			InternshipMonths            int           `json:"internshipMonths"`
			IsNewPosition               int           `json:"isNewPosition"`
			JdCardType                  int           `json:"jdCardType"`
			JobHitReason                string        `json:"jobHitReason"`
			JobHitReasonHighlights      []interface{} `json:"jobHitReasonHighlights"`
			JobId                       int64         `json:"jobId"`
			JobKnowledgeWelfareFeatures []string      `json:"jobKnowledgeWelfareFeatures"`
			JobSummary                  string        `json:"jobSummary"`
			LiveCard                    interface{}   `json:"liveCard"`
			MatchInfo                   struct {
				Icon     string `json:"icon"`
				Matched  int    `json:"matched"`
				TagState int    `json:"tagState"`
			} `json:"matchInfo"`
			MenVipLevel             int           `json:"menVipLevel"`
			Name                    string        `json:"name"`
			NeedMajor               []interface{} `json:"needMajor"`
			Number                  string        `json:"number"`
			OrgBestEmployerFlag     int           `json:"orgBestEmployerFlag"`
			OrgCommercialTags       []interface{} `json:"orgCommercialTags"`
			OrgPayedFlag            int           `json:"orgPayedFlag"`
			PositionCommercialLabel []struct {
				Type          int    `json:"type"`
				TypeName      string `json:"typeName"`
				TypeShowLabel string `json:"typeShowLabel"`
			} `json:"positionCommercialLabel"`
			PositionExpandCardData       string `json:"positionExpandCardData"`
			PositionExpandCardType       int    `json:"positionExpandCardType"`
			PositionHighlight            string `json:"positionHighlight"`
			PositionOfNlp                int    `json:"positionOfNlp"`
			PositionSourceType           int    `json:"positionSourceType"`
			PositionSourceTypeUrl        string `json:"positionSourceTypeUrl"`
			PositionURL                  string `json:"positionURL"`
			PositionUrl                  string `json:"positionUrl"`
			Property                     string `json:"property"`
			PropertyCode                 string `json:"propertyCode"`
			PropertyType                 string `json:"propertyType"`
			PropertyTypeUrl              string `json:"propertyTypeUrl"`
			ProvideInternshipCertificate bool   `json:"provideInternshipCertificate"`
			ProxyModel                   struct {
				ProxiedOrgName  string `json:"proxiedOrgName"`
				ProxiedOrgSize  string `json:"proxiedOrgSize"`
				RecruitPosition int    `json:"recruitPosition"`
			} `json:"proxyModel"`
			PublishTime string `json:"publishTime"`
			RecallSign  struct {
				GMethod string `json:"gMethod"`
				GParam  string `json:"gParam"`
				GQuery  string `json:"gQuery"`
				GSort   string `json:"gSort"`
				GSource string `json:"gSource"`
				GWeight int    `json:"gWeight"`
			} `json:"recallSign"`
			RecruitNumber     int           `json:"recruitNumber"`
			RedirectUrl       string        `json:"redirectUrl"`
			Redirectable      bool          `json:"redirectable"`
			RootCompanyNumber string        `json:"rootCompanyNumber"`
			RpoProxied        bool          `json:"rpoProxied"`
			RpoProxy          bool          `json:"rpoProxy"`
			Salary60          string        `json:"salary60"`
			SalaryCount       string        `json:"salaryCount"`
			SalaryReal        string        `json:"salaryReal"`
			SalaryType        int           `json:"salaryType"`
			SearchTagList     []interface{} `json:"searchTagList"`
			SettlementType    string        `json:"settlementType"`
			ShowDistance      int           `json:"showDistance"`
			SkillLabel        []struct {
				State int    `json:"state"`
				Value string `json:"value"`
			} `json:"skillLabel"`
			SkillLabelPersonality string `json:"skillLabelPersonality"`
			StaffCard             struct {
				AuthenticationState  int    `json:"authenticationState"`
				Avatar               string `json:"avatar"`
				GoldMedalInterviewer struct {
					GoldMedalInterviewer bool   `json:"goldMedalInterviewer"`
					InterviewerImageUrl  string `json:"interviewerImageUrl"`
					InterviewerTitle     string `json:"interviewerTitle"`
				} `json:"goldMedalInterviewer"`
				HrJob              string `json:"hrJob"`
				HrOnlineIocState   int    `json:"hrOnlineIocState"`
				HrOnlineState      string `json:"hrOnlineState"`
				HrStateInfo        string `json:"hrStateInfo"`
				Id                 int    `json:"id"`
				LastOnlineTime     int64  `json:"lastOnlineTime"`
				LastOnlineTimeText string `json:"lastOnlineTimeText"`
				StaffName          string `json:"staffName"`
			} `json:"staffCard"`
			StreetId            int    `json:"streetId"`
			StreetName          string `json:"streetName"`
			SubJobTypeLevel     string `json:"subJobTypeLevel"`
			SubJobTypeLevelName string `json:"subJobTypeLevelName"`
			Subways             []struct {
				Distance    int     `json:"distance"`
				Latitude    float64 `json:"latitude"`
				LineId      int     `json:"lineId"`
				LineName    string  `json:"lineName"`
				Longitude   float64 `json:"longitude"`
				StationId   int     `json:"stationId"`
				StationName string  `json:"stationName"`
			} `json:"subways"`
			TagABC                 string        `json:"tagABC"`
			TagList                []interface{} `json:"tagList"`
			TodayInterview         bool          `json:"todayInterview"`
			TodayInterviewImageUrl string        `json:"todayInterviewImageUrl"`
			TopLabel               interface{}   `json:"topLabel"`
			TradingArea            string        `json:"tradingArea"`
			VolcanoMeterial        interface{}   `json:"volcanoMeterial"`
			WeeklyInternshipDays   int           `json:"weeklyInternshipDays"`
			WelfareLabel           []interface{} `json:"welfareLabel"`
			WelfareTagList         []string      `json:"welfareTagList"`
			WorkCity               string        `json:"workCity"`
			WorkDateType           string        `json:"workDateType"`
			WorkMode               string        `json:"workMode"`
			WorkType               string        `json:"workType"`
			WorkingExp             string        `json:"workingExp"`
			IndustryTags           []string      `json:"industryTags,omitempty"`
		} `json:"list"`
		MajorHaveJd   bool   `json:"majorHaveJd"`
		Method        string `json:"method"`
		MethodGroup   string `json:"methodGroup"`
		ModularState  int    `json:"modularState"`
		NullException string `json:"nullException"`
		PromptCopy    struct {
			Filter struct {
				SubTitle string `json:"subTitle"`
				Title    string `json:"title"`
			} `json:"filter"`
			General struct {
				SubTitle string `json:"subTitle"`
				Title    string `json:"title"`
			} `json:"general"`
			GeneralNull struct {
				SubTitle string `json:"subTitle"`
				Title    string `json:"title"`
			} `json:"generalNull"`
			Line struct {
				HasRecommend string `json:"hasRecommend"`
				NoRecommend  string `json:"noRecommend"`
			} `json:"line"`
			Purpose struct {
				SubTitle string `json:"subTitle"`
				Title    string `json:"title"`
			} `json:"purpose"`
			SearchLineAccuracy struct {
				HasAccuracy string `json:"hasAccuracy"`
				NoAccuracy  string `json:"noAccuracy"`
				Recommend   string `json:"recommend"`
			} `json:"searchLineAccuracy"`
		} `json:"promptCopy"`
		PurposeOptimize              interface{} `json:"purposeOptimize"`
		RelatedExperiment            string      `json:"relatedExperiment"`
		RelatedPosition              interface{} `json:"relatedPosition"`
		RelatedTrace                 string      `json:"relatedTrace"`
		SearchLineAccuracyExperiment string      `json:"searchLineAccuracyExperiment"`
		SortTypes                    []struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"sortTypes"`
		StudyMajor        string `json:"studyMajor"`
		TypeSearch        int    `json:"typeSearch"`
		TaskId            string `json:"taskId"`
		StatusCode        int    `json:"statusCode"`
		StatusDescription string `json:"statusDescription"`
	} `json:"data"`
}

type WYJob struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	Resultbody struct {
		SearchType        int    `json:"searchType"`
		EngineKeywordType int    `json:"engineKeywordType"`
		RequestId         string `json:"requestId"`
		Job               struct {
			Items []struct {
				Property           string   `json:"property"`
				JobId              string   `json:"jobId"`
				JobType            string   `json:"jobType"`
				JobName            string   `json:"jobName"`
				JobTags            []string `json:"jobTags"`
				JobNumString       string   `json:"jobNumString"`
				WorkAreaCode       string   `json:"workAreaCode"`
				JobAreaCode        string   `json:"jobAreaCode"`
				JobAreaString      string   `json:"jobAreaString"`
				HrefAreaPinYin     string   `json:"hrefAreaPinYin"`
				JobAreaLevelDetail struct {
					ProvinceCode   string `json:"provinceCode"`
					ProvinceString string `json:"provinceString"`
					CityCode       string `json:"cityCode"`
					CityString     string `json:"cityString"`
					LandMarkString string `json:"landMarkString"`
					DistrictString string `json:"districtString,omitempty"`
				} `json:"jobAreaLevelDetail,omitempty"`
				ProvideSalaryString     string   `json:"provideSalaryString"`
				IssueDateString         string   `json:"issueDateString"`
				ConfirmDateString       string   `json:"confirmDateString"`
				WorkYear                string   `json:"workYear"`
				WorkYearString          string   `json:"workYearString"`
				DegreeString            string   `json:"degreeString"`
				IndustryType1           string   `json:"industryType1"`
				IndustryType2           string   `json:"industryType2"`
				IndustryType1Str        string   `json:"industryType1Str"`
				IndustryType2Str        string   `json:"industryType2Str"`
				FuncType1Code           string   `json:"funcType1Code"`
				FuncType2Code           string   `json:"funcType2Code"`
				Major1Str               string   `json:"major1Str"`
				Major2Str               string   `json:"major2Str"`
				EncCoId                 string   `json:"encCoId"`
				CompanyName             string   `json:"companyName"`
				FullCompanyName         string   `json:"fullCompanyName"`
				CompanyLogo             string   `json:"companyLogo"`
				CompanyTypeString       string   `json:"companyTypeString"`
				CompanySizeString       string   `json:"companySizeString"`
				CompanySizeCode         string   `json:"companySizeCode"`
				CompanyIndustryType1Str string   `json:"companyIndustryType1Str"`
				CompanyIndustryType2Str string   `json:"companyIndustryType2Str,omitempty"`
				HrUid                   string   `json:"hrUid"`
				HrName                  string   `json:"hrName"`
				SmallHrLogoUrl          string   `json:"smallHrLogoUrl"`
				HrPosition              string   `json:"hrPosition"`
				HrMedalTitle            string   `json:"hrMedalTitle"`
				HrMedalLevel            string   `json:"hrMedalLevel"`
				ShowHrMedalTitle        bool     `json:"showHrMedalTitle"`
				HrIsOnline              bool     `json:"hrIsOnline"`
				IsOnline                bool     `json:"isOnline"`
				HrLabels                []string `json:"hrLabels,omitempty"`
				UpdateDateTime          string   `json:"updateDateTime"`
				Lon                     string   `json:"lon"`
				Lat                     string   `json:"lat"`
				IsCommunicate           bool     `json:"isCommunicate"`
				IsFromXyx               bool     `json:"isFromXyx"`
				IsIntern                bool     `json:"isIntern"`
				IsModelEmployer         bool     `json:"isModelEmployer"`
				IsQuickFeedback         bool     `json:"isQuickFeedback"`
				IsPromotion             bool     `json:"isPromotion"`
				IsApply                 bool     `json:"isApply"`
				IsExpire                bool     `json:"isExpire"`
				JobHref                 string   `json:"jobHref"`
				JobDescribe             string   `json:"jobDescribe"`
				CompanyHref             string   `json:"companyHref"`
				AllowChatOnline         bool     `json:"allowChatOnline"`
				CtmId                   int      `json:"ctmId"`
				Term                    string   `json:"term"`
				TermStr                 string   `json:"termStr"`
				LandmarkId              string   `json:"landmarkId"`
				LandmarkString          string   `json:"landmarkString"`
				RetrieverName           string   `json:"retrieverName"`
				ExrInfo02               string   `json:"exrInfo02"`
				HrInfoType              int      `json:"hrInfoType"`
				IsRemoteWork            bool     `json:"isRemoteWork"`
				ContactAllowed          string   `json:"contactAllowed"`
				ContactDay              string   `json:"contactDay"`
				ContactTime             string   `json:"contactTime"`
				HasHrMobile             bool     `json:"hasHrMobile"`
				JobTagsForOrder         []string `json:"jobTagsForOrder"`
				IsAllowChat             bool     `json:"isAllowChat"`
				SesameLabelList         []struct {
					LabelName       string `json:"labelName"`
					LabelCode       string `json:"labelCode"`
					LabelDefinition string `json:"labelDefinition"`
				} `json:"sesameLabelList,omitempty"`
				JobWelfareCodeDataList []struct {
					Code         string `json:"code"`
					ChineseTitle string `json:"chineseTitle"`
					EnglishTitle string `json:"englishTitle"`
					TypeCode     string `json:"typeCode"`
					TypeTitle    string `json:"typeTitle"`
				} `json:"jobWelfareCodeDataList"`
				JobSalaryMax        string `json:"jobSalaryMax"`
				JobSalaryMin        string `json:"jobSalaryMin"`
				JobScheme           string `json:"jobScheme"`
				CoId                string `json:"coId"`
				HrActiveStatusGreen string `json:"hrActiveStatusGreen,omitempty"`
				IsAd                bool   `json:"isAd,omitempty"`
				AdId                string `json:"adId,omitempty"`
			} `json:"items"`
			TotalCount int    `json:"totalCount"`
			RequestId  string `json:"requestId"`
			Policies   string `json:"policies"`
			Totalcount int    `json:"totalcount"`
		} `json:"job"`
	} `json:"resultbody"`
}
