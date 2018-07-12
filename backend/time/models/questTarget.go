package models

import (
	es "evolution/backend/common/structs"

	"github.com/fatih/structs"
	"github.com/go-xorm/builder"
)

type QuestTarget struct {
	es.ModelWithId `xorm:"extends"`
	QuestId        int      `xorm:"unique not null default 0 comment('所属quest') INT(11)" structs:"quest_id,omitempty"`
	ResourceId     int      `xorm:"unique(quest_id) not null default 0 comment('目标资源id') INT(11)" structs:"resource_id,omitempty"`
	Desc           string   `xorm:"not null default '' comment('描述') VARCHAR(255)" structs:"desc,omitempty"`
	Status         int      `xorm:"not null default 1 comment('当前状态:1未完成2已完成') INT(11)" structs:"status,omitempty"`
	Resource       Resource `xorm:"-"`
}

type QuestTargetJoin struct {
	QuestTarget `xorm:"extends" json:"-"`
	Resource    `xorm:"extends" json:"-"`
	Area        `xorm:"extends" json:"-"`
}

var (
	QuestTargetStatusWait   = 1
	QuestTargetStatusFinish = 2
)

func (m *QuestTarget) TableName() string {
	return "quest_target"
}

func (m *QuestTarget) BuildCondition() (condition builder.Eq) {
	keyPrefix := m.TableName() + "."
	params := structs.Map(m)
	condition = m.Model.BuildCondition(params, keyPrefix)
	return condition
}
