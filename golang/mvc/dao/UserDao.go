package dao

import (
	"github.com/astaxie/beego/orm"
	"jamSwap/mvc/po"
)

type UserDao struct {
}

func (dao *UserDao) QueryAdvertisement(userAddr, chainType string) po.UserAdvertisement {
	var user po.UserAdvertisement
	o.QueryTable("user_advertisement").Filter("user_addr", userAddr).
		Filter("chain_type", chainType).One(&user)
	return user
}

// 插入用户
func (dao *UserDao) InsertUser(user *po.UserAdvertisement) (int64, error) {
	return o.Insert(user)
}

func (dao *UserDao) TotalCount() int {
	var count int
	o.Raw("select count(*) as Count from user_advertisement").QueryRow(&count)
	return count
}

func (dao *UserDao) List(chainType string, start, pageSize int) *[]po.UserAdvertisement {
	var users []po.UserAdvertisement
	o.Raw("select id, user_addr, img_url, create_time, update_time, is_delete, chain_type from user_advertisement "+
		"where chain_type = ? and is_delete = 0 order by create_time desc limit ?,?",
		chainType, start, pageSize).QueryRows(&users)
	return &users
}

func (dao *UserDao) QueryValidList() []po.UserAdvertisement {
	var users []po.UserAdvertisement
	o.Raw("select id, user_addr, chain_type from user_advertisement " +
		"where is_delete = 0").QueryRows(&users)
	return users
}

func (dao *UserDao) UpdateAdvertisement(addr, chainType string, isDelete int) (int, error) {
	n, err := o.QueryTable("user_advertisement").Filter("user_addr", addr).
		Filter("chain_type", chainType).Update(orm.Params{
		"is_delete": isDelete,
	})
	return int(n), err
}

func (dao *UserDao) UpdateUser(p *po.UserAdvertisement) (int64, error) {
	user := po.UserAdvertisement{Id: p.Id}
	user.ImgUrl = p.ImgUrl
	user.IsDelete = 0
	return o.Update(&user, "ImgUrl", "IsDelete")
}
