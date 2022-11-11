package order

import (
	"context"
	"lottery_gateway/proto/order"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OrderList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "6"))
	data := &order.ListRequest{
		PageNum:  int32(pageNum),
		PageSize: int32(pageSize),
	}
	rep, _ := OrderService.OrderList(context.TODO(), data)
	if rep.Code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  rep.Code,
		"msg":   rep.Msg,
		"data":  rep.Orders,
		"count": rep.Count,
	})
}

func OrderByUser(c *gin.Context) {
	email, exit := c.Get("user")
	if !exit {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "请先登录",
		})
		return
	}
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	data := &order.UserRequest{
		Email:    email.(string),
		PageSize: int32(pageSize),
		PageNum:  int32(pageNum),
	}
	rep, _ := OrderService.OrderUser(context.TODO(), data)
	if rep.Code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  rep.Code,
		"msg":   rep.Msg,
		"data":  rep.Orders,
		"count": rep.Count,
	})
}

func OrderDel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data := &order.IdRequest{
		Id: int32(id),
	}
	rep, _ := OrderService.OrderDel(context.TODO(), data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}
