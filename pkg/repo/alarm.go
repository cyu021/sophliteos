package repo

import (
	"sophliteos/global"
	"sophliteos/logger"
	"sophliteos/pkg/dto/response"

	"time"
)

// 保存告警
func SaveAlarm(alarm Alarm) error {
	db := global.DB.Model(&Alarm{}).Save(&alarm)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

// 查询告警
func QueryAlarms(alarm *Alarm, pageNo, pageSize int, startTime, endTime *time.Time) response.Page {
	alarms := make([]Alarm, 0)
	db := global.DB.Model(&Alarm{})
	var vars []interface{}
	var query = ""
	if alarm.Code != 0 {
		query = query + " and code = ? "
		vars = append(vars, alarm.Code)
	}
	if len(alarm.Msg) > 0 {
		query = query + " and msg like ? "
		vars = append(vars, "%"+alarm.Msg+"%")
	}
	if len(alarm.DeviceSn) > 0 {
		query = query + " and device_sn = ? "
		vars = append(vars, alarm.DeviceSn)
	}
	if len(alarm.ComponentType) > 0 {
		query = query + " and component_type = ? "
		vars = append(vars, alarm.ComponentType)
	}
	if len(alarm.ControllerUnitSn) > 0 {
		query = query + " and controller_unit_sn = ? "
		vars = append(vars, alarm.ControllerUnitSn)
	}
	if len(alarm.CoreUnitBoardSn) > 0 {
		query = query + " and core_unit_board_sn = ? "
		vars = append(vars, alarm.CoreUnitBoardSn)
	}
	if len(alarm.CoreUnitBoardChipSn) > 0 {
		query = query + " and core_unit_board_chip_sn = ? "
		vars = append(vars, alarm.CoreUnitBoardChipSn)
	}
	if startTime != nil {
		query = query + " and created_at >= ?  "
		vars = append(vars, &startTime)
	}
	if endTime != nil {
		query = query + " and created_at <= ?  "
		vars = append(vars, &endTime)
	}
	var total int
	if len(query) > 0 {
		db.Where(query[4:], vars...).Count(&total)
		db.Where(query[4:], vars...).Offset((pageNo - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&alarms)
	} else {
		db.Count(&total)
		db.Offset((pageNo - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&alarms)
	}

	return response.Page{
		PageCount: (total + pageSize - 1) / pageSize,
		PageNo:    pageNo,
		PageSize:  pageSize,
		Total:     total,
		Items:     alarms,
	}
}

// 删除日期前数据
func DeleteAlarmByCreatedAt(createdAt time.Time) error {
	db := global.DB.Model(&Alarm{}).Debug().Where(" created_time <= ? ", createdAt).Delete(&Alarm{})
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

// 删除日期前数据，保留最新的200条
func DeleteOldestAlarms() error {
	// 获取最新的200条记录的ID
	var ids []uint
	if err := global.DB.Model(&Alarm{}).Order("created_at DESC").Limit(200).Pluck("id", &ids).Error; err != nil {
		logger.Error("获取最新记录ID错误：%v", err)
		return err
	}

	// 删除除最新的200条记录之外的所有记录
	db := global.DB.Where("id NOT IN (?)", ids).Unscoped().Delete(&Alarm{})
	if err := db.Error; err != nil {
		logger.Error("删除错误：%v", err)
		return err
	}
	return nil
}
