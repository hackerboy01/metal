package controllers

//包名并非必须和文件夹名相同，但是按照惯例最后一个路径名和包名一致
import (
	"log"
	. "metal/models"
)

type JobCountController struct {
	BaseController
}

/**
 * 工作数量统计
 */

// @router /job-count [get]
func (c *UserController) JobCount() {
	c.Data["title"] = "报表"
	c.TplName = "admin/job-count.html"
}

//  近一个月数据
// @router /job-count/count-data-recently [get]
func (c *UserController) CountDataRecently() {
	lang := c.GetString("language")
	startDate := c.GetString("startDate") + " 00:00:00"
	endDate := c.GetString("endDate") + " 23:59:59"

	log.Println(lang, startDate, endDate)
	jobCount := new(JobCount)
	list, err := jobCount.GetCountData(lang, startDate, endDate)
	if nil != err {
		log.Print(err)
		c.Data["json"] = ErrorData(err)
	} else {
		c.Data["json"] = SuccessData(list)
	}
	c.ServeJSON()
}

// 所有历史数据，按月平均值统计
// @router /job-count/count-data-all [get]
func (c *UserController) CountDataAll() {
	lang := c.GetString("language")

	jobCount := new(JobCount)
	jobCount.JobTitle = lang
	list, err := jobCount.GetCountDataAll()
	if nil != err {
		log.Print(err)
		c.Data["json"] = ErrorData(err)
	} else {
		c.Data["json"] = SuccessData(list)
	}
	c.ServeJSON()
}
