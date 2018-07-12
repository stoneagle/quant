package php

import (
	"evolution/backend/time/models"
	"fmt"
	"time"

"github.com/go-xorm/xorm"


ype TargetEntityLink struct {
Id       int       `xorm:"not null pk autoincr INT(11)"`
TargetId int       `xorm:"not null default 0 comment('领域对象ID') INT(11)"`
EntityId int       `xorm:"not null default 0 comment('相关实体ID') INT(11)"`
Ctime    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
Utime    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`


ype TargetEntityLinkJoin struct {
TargetEntityLink `xorm:"extends"`
Target           `xorm:"extends"`


unc (c *TargetEntityLink) Transfer(src, des *xorm.Engine) {
olds := make([]TargetEntityLinkJoin, 0)
err := src.Table("target_entity_link").Join("LEFT", "target", "target.id = target_entity_link.target_id").Find(&olds)
if err != nil {
	fmt.Printf("target join get error:%v\r\n", err.Error())
	return
	}

	fieldPhaseMap := map[int]int{}
	insertNum := 0

for _, one := range olds {
	field := one.Target.FieldId
	entityId := one.TargetEntityLink.EntityId

	_, ok := fieldPhaseMap[field]
	if !ok {
			phases := make([]models.Phase, 0)
			des.Where("field_id = ?", field).Asc("level").Find(&phases)
			if len(phases) == 0 {
				fmt.Printf("phase not exist")
				return
			}
		fieldPhaseMap[field] = phases[0].Id
	}

	// 找到新版本entity的Id
	name, err := getResourceName(src, entityId, field)
	if err != nil {
		return
	}

	newResource := models.ResourceJoin{}
	_, err = des.Table("resource").Join("LEFT", "area", "area.id = resource.area_id").Where("area.field_id = ?", field).And("resource.name = ?", name).Get(&newResource)
	if err != nil {
		fmt.Printf("new resource join get error:%v\r\n", err.Error())
		return
	}
	if newResource.Resource.Name == "" {
		fmt.Printf("%v:%v:%v\r\n", field, name, newResource.Resource.Name)
	}

	tmp := models.UserResource{}
	tmp.UserId = 1
	tmp.ResourceId = newResource.Resource.Id
	tmp.CreatedAt = one.TargetEntityLink.Ctime
	tmp.UpdatedAt = one.TargetEntityLink.Utime

	has, err := des.Where("user_id = ?", 1).And("resource_id = ?", newResource.Resource.Id).Get(new(models.Resource))
	if err != nil {
		fmt.Printf("target transfer has check error:%v\r\n", err.Error())
		continue
		}
		if !has {
			_, err = des.Insert(&tmp)
			if err != nil {
				fmt.Printf("target transfer insert error:%v\r\n", err.Error())
				continue
			}
			insertNum++
		}
	}

	fmt.Printf("target and entity transfer success:%v\r\n", insertNum)
}

func getResourceName(src *xorm.Engine, entityId int, field int) (name string, err error) {
	switch field {
	case 1:
		oldEntity := EntitySkill{}
		_, err = src.Id(entityId).Get(&oldEntity)
		if err != nil {
			fmt.Printf("entity skill get error:%v\r\n", err.Error())
			return
		}
		name = oldEntity.Name
	case 2:
		oldEntity := EntityAsset{}
		_, err = src.Id(entityId).Get(&oldEntity)
		if err != nil {
			fmt.Printf("entity asset get error:%v\r\n", err.Error())
			return
		}
		name = oldEntity.Name
	case 3:
		oldEntity := EntityWork{}
		_, err = src.Id(entityId).Get(&oldEntity)
		if err != nil {
			fmt.Printf("entity work get error:%v\r\n", err.Error())
			return
		}
		name = oldEntity.Name
	case 4:
		oldEntity := EntityCircle{}
		_, err = src.Id(entityId).Get(&oldEntity)
		if err != nil {
			fmt.Printf("entity circle get error:%v\r\n", err.Error())
			return
		}
		name = oldEntity.Name
	case 5:
		oldEntity := EntityQuest{}
		_, err = src.Id(entityId).Get(&oldEntity)
		if err != nil {
			fmt.Printf("entity quest get error:%v\r\n", err.Error())
			return
		}
		name = oldEntity.Name
	case 6:
		oldEntity := EntityLife{}
		_, err = src.Id(entityId).Get(&oldEntity)
		if err != nil {
			fmt.Printf("entity life get error:%v\r\n", err.Error())
			return
		}
		name = oldEntity.Name
	}
	return
}
