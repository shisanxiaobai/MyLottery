package prize

import (
	"context"
	"fmt"
	"lottery_gateway/proto/prize"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PrizeList(c *gin.Context) {
	rep, _ := PrizeService.PrizeList(context.TODO(), &prize.ListRequest{})
	if rep.Code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
		"data": rep.Prizes,
	})
}

func PrizeAdd(c *gin.Context) {
	name := c.PostForm("name")
	num, _ := strconv.Atoi(c.PostForm("num"))
	file, err := c.FormFile("pic")
	if err != nil {
		fmt.Println(err)
		data := &prize.Prize{
			Name:       name,
			Num:        int32(num),
			CreateTime: time.Now().Format("2006-01-02 15:06:07"),
		}
		rep, _ := PrizeService.PrizeAdd(context.TODO(), data)
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	unix_int64 := time.Now().Unix()
	unix_str := strconv.FormatInt(unix_int64, 10)
	file_path := "upload/" + unix_str + file.Filename
	err = c.SaveUploadedFile(file, file_path)
	data := &prize.Prize{
		Name:       name,
		Num:        int32(num),
		Pic:        "http://localhost:9001/" + file_path,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	rep, _ := PrizeService.PrizeAdd(context.TODO(), data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func PrizeDel(c *gin.Context) {
	name := c.Query("name")
	data := &prize.NameRequest{
		Name: name,
	}
	rep, _ := PrizeService.PrizeDel(context.TODO(), data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func PrizeEdit(c *gin.Context) {
	name := c.PostForm("name")
	num, _ := strconv.Atoi(c.PostForm("num"))
	file, err := c.FormFile("pic")
	if err != nil {
		data := &prize.Prize{
			Name: name,
			Num:  int32(num),
		}
		rep, _ := PrizeService.PrizeEdit(context.TODO(), data)
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	unix_int64 := time.Now().Unix()
	unix_str := strconv.FormatInt(unix_int64, 10)
	file_path := "upload/" + unix_str + file.Filename
	c.SaveUploadedFile(file, file_path)
	data := &prize.Prize{
		Name: name,
		Num:  int32(num),
		Pic:  "http://localhost:9002/" + file_path,
	}
	rep, _ := PrizeService.PrizeEdit(context.TODO(), data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func PrizeInfo(c *gin.Context) {
	name := c.Query("name")
	data := &prize.NameRequest{
		Name: name,
	}
	rep, _ := PrizeService.PrizeInfo(context.TODO(), data)
	if rep.Code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
		"data": rep.Prize,
	})
}
