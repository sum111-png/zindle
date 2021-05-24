// Code generated by goctl. DO NOT EDIT.
package types

type BookPostReq struct {
	Id          int64  `db:"id" form:"id,optional"`                        // ID
	Title       string `db:"title" form:"title,optional"`                  // 标题
	Description string `db:"description" form:"description,default="`      // 简介
	Cover       string `db:"cover" form:"cover,optional,default="`         // 封面
	Sort        int64  `db:"sort" form:"sort,default=0"`                   // 排序
	Quantity    int64  `db:"quantity" form:"quantity,optional,default=0"`  // 数量
	CreateBy    string `db:"create_by" form:"create_by,optional,default="` // 创建者
	UpdateBy    string `db:"update_by" form:"update_by,optional,default="` // 更新者
}

type BookReply struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type BookListReq struct {
	Page     int64  `json:"page,optional,default=1" form:"page,optional,default=1"`
	PageSize int64  `json:"pageSize,optional,default=10" form:"pageSize,optional,default=10"`
	Keyword  string `json:"keyword,optional" form:"pageSize,optional"`
}

type BookDelReq struct {
	Id int64 `json:"id"`
}

type BookDelBatchReq struct {
	Ids string `json:"ids"`
}

type BookOrderPostReq struct {
	Id         int64  `db:"id" json:"id,optional"`                                      // ID
	MemberId   int64  `db:"member_id" json:"member_id,default=0"`                       // 用户ID
	BookId     int64  `db:"book_id" json:"book_id,default=0"`                           // 书ID
	ReturnDate string `db:"return_date" json:"return_date,default=2006-01-02 15:04:05"` // 还书日期
}

type BookOrderReply struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type BookOrderListReq struct {
	Page     int64  `json:"page,optional,default=1" form:"page,optional,default=1"`
	PageSize int64  `json:"pageSize,optional,default=10" form:"pageSize,optional,default=10"`
	Keyword  string `json:"keyword,optional" form:"pageSize,optional"`
}

type BookOrderDelReq struct {
	Id int64 `json:"id"`
}

type BookOrderDelBatchReq struct {
	Ids string `json:"ids"`
}

type ReaderPostReq struct {
	Id        int64  `db:"id" json:"id,optional"`                                             // 用户ID
	Username  string `db:"username" json:"username,optional"`                                 // 用户账号
	Password  string `db:"password" json:"password,default="`                                 // 密码
	Nickname  string `db:"nickname" json:"nickname,optional"`                                 // 用户昵称
	Phone     string `db:"phone" json:"phone,default=0"`                                      // 手机
	Email     string `db:"email" json:"email,default="`                                       // 用户邮箱
	Status    int64  `db:"status" json:"status,default=0"`                                    // 帐号状态（0正常 1停用）
	LoginIp   string `db:"login_ip" json:"login_ip,default="`                                 // 最后登录IP
	LoginDate string `db:"login_date" json:"login_date,optional,default=2006-01-02 15:04:05"` // 最后登录时间
	Remark    string `db:"remark" json:"remark,default=Empty string"`                         // 备注
}

type RegisterReq struct {
	Username         string `db:"username" json:"username"`                                   // 用户账号
	Password         string `db:"password" json:"password,default="`                          // 密码
	Email            string `db:"email" json:"email,default="`                                // 用户邮箱
	VerificationCode string `db:"verificationCode" json:"verificationCode,optional,default="` // 验证码
}

type LoginReq struct {
	Username string `db:"username" json:"username"`          // 用户账号
	Password string `db:"password" json:"password,default="` // 密码
}

type UpdatePwdReq struct {
	OldPassword string `json:"oldPassword"`
	Password    string `json:"password"`
}

type VerifyReq struct {
	Account string `db:"account" json:"account"`
	Type    int64  `db:"type" json:"type"`
}

type FindPasswordReq struct {
	Password         string `db:"password" json:"password,default="`                          // 密码
	Email            string `db:"email" json:"email,default="`                                // 用户邮箱
	VerificationCode string `db:"verificationCode" json:"verificationCode,optional,default="` // 验证码
}

type ReaderReply struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ReaderListReq struct {
	Page     int64  `json:"page,optional,default=1" form:"page,optional,default=1"`
	PageSize int64  `json:"pageSize,optional,default=10" form:"pageSize,optional,default=10"`
	Keyword  string `json:"keyword,optional" form:"pageSize,optional"`
}

type ReaderDelReq struct {
	Id int64 `json:"id"`
}

type ReaderDelBatchReq struct {
	Ids string `json:"ids"`
}

type VerifyCodePostReq struct {
	Id      int64  `db:"id" json:"id,optional"`           //
	Account string `db:"account" json:"account,optional"` // 号码(手机或邮箱)
	Code    string `db:"code" json:"code,optional"`       // 验证码
	Type    int64  `db:"type" json:"type,default=0"`      // 类型0手机1邮箱
	Status  int64  `db:"status" json:"status,default=0"`  // 状态0未验证1已验证2验证错误
}

type VerifyCodeReply struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type VerifyCodeListReq struct {
	Page     int64  `json:"page,optional,default=1" form:"page,optional,default=1"`
	PageSize int64  `json:"pageSize,optional,default=10" form:"pageSize,optional,default=10"`
	Keyword  string `json:"keyword,optional" form:"pageSize,optional"`
}

type VerifyCodeDelReq struct {
	Id int64 `json:"id"`
}

type VerifyCodeDelBatchReq struct {
	Ids string `json:"ids"`
}
