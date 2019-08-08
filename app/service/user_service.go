package service

import "github.com/gogf/gf/g/crypto/gmd5"

// EncryptPassword 加密密码
//
// createTime:2019年04月25日 10:19:13
// author:hailaz
func EncryptPassword(data string) string {
	return gmd5.EncryptString(data + ENCRYPTMD5)
}

// {
//     roles: ['admin'],
//     introduction: 'I am a super administrator',
//     avatar: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',
//     name: 'Super Admin'
//   }
// GetUserInfo 获取用户信息
//
// createTime:2019年05月08日 16:53:24
// author:hailaz
func (u *User) GetUserInfo() UserInfo {
	info := UserInfo{}
	if u.UserName == ADMIN_NAME {
		info.Roles = []string{ADMIN_NAME}
	} else {
		info.Roles = Enforcer.GetRolesForUser(u.UserName)
	}

	info.Avatar = u.Avatar
	info.Introduction = u.Introduction
	info.Name = u.NickName

	return info
}
