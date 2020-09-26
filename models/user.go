package models

/**
 * 构建用户结构体，用于表示数据结构的定义
 */
type User struct {
	User string `json:"name"`
	Birthday string `json:"birthday"`
	Address string	`json:"address"`
	Nick string	`json:"nick"`
	Password string `json:"password"`
}

//type User struct {
//	User string `form:"name"`
//	Birthday string `form:"birthday"`
//	Address string `form:"address"`
//	Nick string `form:"nick"`
//}
