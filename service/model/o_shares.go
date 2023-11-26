/*
 * @Author: LinkLeong link@icewhale.org
 * @Date: 2022-07-26 11:17:17
 * @LastEditors: LinkLeong
 * @LastEditTime: 2022-07-27 15:25:07
 * @FilePath: /CasaOS/service/model/o_shares.go
 * @Description:
 * @Website: https://www.casaos.io
 * Copyright (c) 2022 by icewhale, All Rights Reserved.
 */
package model

type SharesDBModel struct {
	ID          uint   `gorm:"column:id;primary_key" json:"id"`
	Anonymous   bool   `json:"anonymous"`
	Path        string `json:"path"`
	Name        string `json:"name"`
	Updated     int64  `gorm:"autoUpdateTime"`
	Created     int64  `gorm:"autoCreateTime"`
	TimeMachine bool   `gorm:"column:time_machine;default:false" json:"time_machine"`
}

func (p *SharesDBModel) TableName() string {
	return "o_shares"
}
