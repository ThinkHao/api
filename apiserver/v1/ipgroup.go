package v1

import (
	"time"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

type IPGroup struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Id                int       `json:"id" gorm:"column:id" validate:"omitempty"`
	CreateTime        time.Time `json:"create_time" gorm:"column:create_time" validate:"omitempty"`
	UpdateTime        time.Time `json:"update_time" gorm:"column:update_time" validate:"omitempty"`
	IPGroupId         string    `json:"ipgroup_id" gorm:"column:ipgroup_id" validate:"omitempty"`
	IPGroupName       string    `json:"ipgroup_name" gorm:"column:ipgroup_name" validate:"omitempty"`
	Saler             string    `json:"saler" gorm:"column:saler" validate:"omitempty"`
	SalerGroup        string    `json:"saler_group" gorm:"column:saler_group" validate:"omitempty"`
	Region            string    `json:"region" gorm:"column:region" validate:"omitempty"`
	Tag               string    `json:"tag" gorm:"column:tag" validate:"omitempty"`
}

type IPGroupList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*IPGroup `json:"items"`
}

func (i *IPGroup) TableName() string {
	return "nfa_ip_group"
}
