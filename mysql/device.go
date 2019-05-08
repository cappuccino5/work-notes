package mysql

import (
	"fmt"
	"go_hisens_admin/server/model"
	"strings"
)





type Device struct {
	Id          uint64 `gorm:"primary_key" json:"id"`
	BaseAboveSn string `sql:"type:varchar(100) DEFAULT NULL comment '硬件上座sn'" json:"base_above_sn"`
	BaseBelowSn string `gorm:"unique_index" sql:"type:varchar(100) DEFAULT NULL comment '硬件底座sn'"json:"base_below_sn"`
	BlueTooth   string `sql:"type:varchar(100) DEFAULT NULL comment '蓝牙名字'"json:"blue_tooth"`
	AgentId     uint64 `sql:"type:varchar(100) DEFAULT NULL comment '代理ID'"json:"agent_id"`
	UserId      uint64 `sql:"type:bigint(20) DEFAULT NULL comment '用户id'" json:"user_id"`
	CreateTime  int64  `json:"create_time"`
	UpdateTime  int64  `json:"update_time"`
}


// 存在则更新不存在则新增sql语句
// INSERT INTO device (base_below_sn,agent_id) VALUES("s1",1),('2',2) 	on  DUPLICATE key update agent_id=VALUES(agent_id),base_below_sn=VALUES(base_below_sn)
func CreateOrAppendDeviceInfo(data []Device) error {
	insertSql := "INSERT INTO device (base_below_sn,agent_id) VALUES"
	values := "(?,?)"
	updateSql := " on  DUPLICATE key update base_below_sn=VALUES(base_below_sn),agent_id=VALUES(agent_id)"
	var strValue []string
	var insertParam []interface{}
	for i := range data {
		strValue = append(strValue, values)
		insertParam = append(insertParam, data[i].BaseBelowSn, data[i].AgentId)
	}
	values = strings.Join(strValue, ",")
	insertSql += values + updateSql

	return db.Exec(insertSql,insertParam...).Error
}
