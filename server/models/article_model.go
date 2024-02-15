package models

// ArticleModel 用户表
type ArticleModel struct {
	MODEL
	NickName      string    `gorm:"size:42" json:"nick_name"`   //昵称
	Title         string    `gorm:"size:32" json:"title"`       //用户名
	Abstract      string    `json:"abstract"`                   //密码
	Content       string    ` json:"content"`                   //头像
	LookCount     int       `json:"look_count"`                 //邮箱
	CommentCount  int       `json:"comment_count"`              //手机号
	DiggCount     int       `json:"digg_count"`                 //地址
	CollectsCount int       `json:"collects_count"`             //其他平台唯一id
	UserModel     UserModel `gorm:"foreignKey:UserID" json:"-"` //ip地址
	UserID        uint      `json:"user_id"`                    //权限 1：管理员 2：普通用户 3：游客
	Category      string    `gorm:"size: 20" json:"category"`   //注册来源
	Source        string    `json:"source"`                     //发布的文章列表
	Link          string    `json:"link"`                       //收藏的文章列表
}
