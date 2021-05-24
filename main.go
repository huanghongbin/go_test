package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

// BaseModel 基础模型
type BaseModel struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

// User 用户模型
type User struct {
	BaseModel
	UserID   int    `json:"userId" gorm:"primary_key"`
	NickName string `gorm:"type:varchar(128)" json:"nickName"` // 昵称
	Phone    string `gorm:"type:varchar(11)" json:"phone"`     // 手机号
	RoleID   int    `gorm:"type:int(11)" json:"roleId"`        // 角色编码
	Avatar   string `gorm:"type:varchar(255)" json:"avatar"`   //头像
	Sex      string `gorm:"type:varchar(255)" json:"sex"`      //性别
	Email    string `gorm:"type:varchar(128)" json:"email"`    //邮箱
	DeptID   int    `gorm:"type:int(11)" json:"deptId"`        //部门编码
	CreateBy string `gorm:"type:varchar(128)" json:"createBy"` //
	UpdateBy string `gorm:"type:varchar(128)" json:"updateBy"` //
	Remark   string `gorm:"type:varchar(255)" json:"remark"`   //备注
	Status   string `gorm:"type:int(1);" json:"status"`
}

// TableName 指定用户表名称
func (User) TableName() string {
	return "user"
}

// Dept 部门模型
type Dept struct {
	BaseModel
	DeptID   int    `json:"deptId" gorm:"primary_key;AUTO_INCREMENT"` //部门编码
	ParentID int    `json:"parentId" gorm:"type:int(11);"`            //上级部门
	DeptPath string `json:"deptPath" gorm:"type:varchar(255);"`       //部门路径（以ID为标识）
	DeptName string `json:"deptName"  gorm:"type:varchar(128);"`      //部门名称
	Sort     int    `json:"sort" gorm:"type:int(4);"`                 //排序
	Leader   string `json:"leader" gorm:"type:varchar(128);"`         //负责人
	Phone    string `json:"phone" gorm:"type:varchar(11);"`           //手机
	Email    string `json:"email" gorm:"type:varchar(64);"`           //邮箱
	Status   string `json:"status" gorm:"type:int(1);"`               //状态
	CreateBy string `json:"createBy" gorm:"type:varchar(64);"`
	UpdateBy string `json:"updateBy" gorm:"type:varchar(64);"`
}

// TableName 指定部门表名称
func (Dept) TableName() string {
	return "dept"
}

// Role 角色模型
type Role struct {
	BaseModel
	RoleID    int    `json:"roleId" gorm:"primary_key;AUTO_INCREMENT"` // 角色编码
	RoleName  string `json:"roleName" gorm:"type:varchar(128);"`       // 角色名称
	Status    string `json:"status" gorm:"type:int(1);"`               //
	RoleKey   string `json:"roleKey" gorm:"type:varchar(128);"`        //角色代码
	RoleSort  int    `json:"roleSort" gorm:"type:int(4);"`             //角色排序
	Flag      string `json:"flag" gorm:"type:varchar(128);"`           //
	CreateBy  string `json:"createBy" gorm:"type:varchar(128);"`       //
	UpdateBy  string `json:"updateBy" gorm:"type:varchar(128);"`       //
	Remark    string `json:"remark" gorm:"type:varchar(255);"`         //备注
	Admin     bool   `json:"admin" gorm:"type:char(1);"`
	DataScope string `json:"dataScope" gorm:"type:varchar(128);"`
}

// TableName 指定角色表名称
func (Role) TableName() string {
	return "role"
}

// InitDB 初始化DB
func InitDB() {
	var err error
	db, err = gorm.Open("mysql", "root:root@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	// 2, 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Dept{})
	db.AutoMigrate(&Role{})
}

// CreateDept 创建部门
func CreateDept() {
	Dept1 := &Dept{BaseModel: BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, DeptID: 1, ParentID: 0, DeptPath: "/0/1", DeptName: "二丫讲梵科技", Sort: 0, Leader: "zhangsan", Phone: "15638887180", Email: "zhangsan@eryajf.net", Status: "1", CreateBy: "", UpdateBy: ""}
	Dept2 := &Dept{BaseModel: BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, DeptID: 2, ParentID: 1, DeptPath: "/0/1/2", DeptName: "技术部", Sort: 0, Leader: "zhangsan1", Phone: "15638887180", Email: "zhangsan1@eryajf.net", Status: "1", CreateBy: "", UpdateBy: ""}
	Dept3 := &Dept{BaseModel: BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, DeptID: 3, ParentID: 1, DeptPath: "/0/1/3", DeptName: "客服部", Sort: 0, Leader: "zhangsan1", Phone: "15638887180", Email: "zhangsan1@eryajf.net", Status: "1", CreateBy: "", UpdateBy: ""}
	Dept4 := &Dept{BaseModel: BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, DeptID: 4, ParentID: 1, DeptPath: "/0/1/2/4", DeptName: "运维部", Sort: 0, Leader: "lisi", Phone: "15638887180", Email: "lisi@eryajf.net", Status: "1", CreateBy: "", UpdateBy: ""}
	defer db.Close()
	db.Debug().Model(&Dept{}).Create(&Dept1) // INSERT INTO `dept` (`created_at`,`updated_at`,`deleted_at`,`dept_id`,`parent_id`,`dept_path`,`dept_name`,`sort`,`leader`,`phone`,`email`,`status`,`create_by`,`update_by`) VALUES ('2020-06-15 15:57:23','2020-06-15 15:57:23',NULL,1,0,'/0/1','二丫讲梵科技',0,'zhangsan','15638887180','zhangsan@eryajf.net','1','','')
	db.Debug().Model(&Dept{}).Create(&Dept2) // INSERT INTO `dept` (`created_at`,`updated_at`,`deleted_at`,`dept_id`,`parent_id`,`dept_path`,`dept_name`,`sort`,`leader`,`phone`,`email`,`status`,`create_by`,`update_by`) VALUES ('2020-06-15 15:57:23','2020-06-15 15:57:23',NULL,2,1,'/0/1/2','技术部',0,'zhangsan1','15638887180','zhangsan1@eryajf.net','1','','')
	db.Debug().Model(&Dept{}).Create(&Dept3) // INSERT INTO `dept` (`created_at`,`updated_at`,`deleted_at`,`dept_id`,`parent_id`,`dept_path`,`dept_name`,`sort`,`leader`,`phone`,`email`,`status`,`create_by`,`update_by`) VALUES ('2020-06-15 15:57:23','2020-06-15 15:57:23',NULL,3,1,'/0/1/3','客服部',0,'zhangsan1','15638887180','zhangsan1@eryajf.net','1','','')
	db.Debug().Model(&Dept{}).Create(&Dept4) // INSERT INTO `dept` (`created_at`,`updated_at`,`deleted_at`,`dept_id`,`parent_id`,`dept_path`,`dept_name`,`sort`,`leader`,`phone`,`email`,`status`,`create_by`,`update_by`) VALUES ('2020-06-15 15:57:23','2020-06-15 15:57:23',NULL,4,1,'/0/1/2/4','运维部',0,'lisi','15638887180','lisi@eryajf.net','1','','')
}

// CreateRole 创建角色
func CreateRole() {
	Role1 := &Role{BaseModel: BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, RoleID: 1, RoleName: "系统管理员", Status: "1", RoleKey: "admin", RoleSort: 1, Flag: "", CreateBy: "", UpdateBy: "", Remark: "", Admin: true, DataScope: "1"}
	Role2 := &Role{BaseModel: BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, RoleID: 2, RoleName: "普通角色", Status: "1", RoleKey: "user", RoleSort: 1, Flag: "", CreateBy: "", UpdateBy: "", Remark: "", Admin: false, DataScope: "2"}
	Role3 := &Role{BaseModel: BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, RoleID: 3, RoleName: "测试角色", Status: "1", RoleKey: "test", RoleSort: 1, Flag: "", CreateBy: "", UpdateBy: "", Remark: "", Admin: false, DataScope: "3"}
	defer db.Close()
	db.Debug().Model(&Role{}).Create(&Role1) // INSERT INTO `role` (`created_at`,`updated_at`,`deleted_at`,`role_id`,`role_name`,`status`,`role_key`,`role_sort`,`flag`,`create_by`,`update_by`,`remark`,`admin`,`data_scope`) VALUES ('2020-06-15 16:02:17','2020-06-15 16:02:17',NULL,1,'系统管理员','1','admin',1,'','','','',true,'1')
	db.Debug().Model(&Role{}).Create(&Role2) // INSERT INTO `role` (`created_at`,`updated_at`,`deleted_at`,`role_id`,`role_name`,`status`,`role_key`,`role_sort`,`flag`,`create_by`,`update_by`,`remark`,`admin`,`data_scope`) VALUES ('2020-06-15 16:02:17','2020-06-15 16:02:17',NULL,2,'普通角色','1','user',1,'','','','',false,'2')
	db.Debug().Model(&Role{}).Create(&Role3) // INSERT INTO `role` (`created_at`,`updated_at`,`deleted_at`,`role_id`,`role_name`,`status`,`role_key`,`role_sort`,`flag`,`create_by`,`update_by`,`remark`,`admin`,`data_scope`) VALUES ('2020-06-15 16:02:17','2020-06-15 16:02:17',NULL,3,'测试角色','1','test',1,'','','','',false,'3')
}

// CreateUser 创建用户
func CreateUser() {
	User1 := &User{BaseModel: BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, NickName: "liql", Phone: "15638887180", RoleID: 1, Avatar: "", Sex: "man", Email: "liql@qq.com", DeptID: 4, CreateBy: "", UpdateBy: "", Remark: "", Status: "1"}
	User2 := &User{BaseModel: BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, NickName: "liql1", Phone: "15638887180", RoleID: 1, Avatar: "", Sex: "man", Email: "liql@qq.com", DeptID: 2, CreateBy: "", UpdateBy: "", Remark: "", Status: "1"}
	User3 := &User{BaseModel: BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, NickName: "liuzb", Phone: "15638887180", RoleID: 2, Avatar: "", Sex: "man", Email: "liql@qq.com", DeptID: 3, CreateBy: "", UpdateBy: "", Remark: "", Status: "1"}
	User4 := &User{BaseModel: BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, NickName: "liuzb1", Phone: "15638887180", RoleID: 2, Avatar: "", Sex: "man", Email: "liql@qq.com", DeptID: 4, CreateBy: "", UpdateBy: "", Remark: "", Status: "1"}
	User5 := &User{BaseModel: BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, NickName: "jiangby", Phone: "15638887180", RoleID: 3, Avatar: "", Sex: "man", Email: "liql@qq.com", DeptID: 2, CreateBy: "", UpdateBy: "", Remark: "", Status: "1"}
	User6 := &User{BaseModel: BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}, NickName: "jiangby1", Phone: "15638887180", RoleID: 3, Avatar: "", Sex: "man", Email: "liql@qq.com", DeptID: 3, CreateBy: "", UpdateBy: "", Remark: "", Status: "1"}
	defer db.Close()
	db.Debug().Model(&User{}).Create(&User1) // INSERT INTO `user` (`created_at`,`updated_at`,`deleted_at`,`nick_name`,`phone`,`role_id`,`avatar`,`sex`,`email`,`dept_id`,`create_by`,`update_by`,`remark`,`status`) VALUES ('2020-06-15 16:07:43','2020-06-15 16:07:43',NULL,'liql','15638887180',1,'','man','liql@qq.com',4,'','','','1')
	db.Debug().Model(&User{}).Create(&User2)
	db.Debug().Model(&User{}).Create(&User3)
	db.Debug().Model(&User{}).Create(&User4)
	db.Debug().Model(&User{}).Create(&User5)
	db.Debug().Model(&User{}).Create(&User6) // INSERT INTO `user` (`created_at`,`updated_at`,`deleted_at`,`nick_name`,`phone`,`role_id`,`avatar`,`sex`,`email`,`dept_id`,`create_by`,`update_by`,`remark`,`status`) VALUES ('2020-06-15 16:07:43','2020-06-15 16:07:43',NULL,'jiangby1','15638887180',3,'','man','liql@qq.com',3,'','','','1')
}

// UserDept 用户以及部门联合查询============
type UserDept struct {
	User
	DeptName string `json:"deptName"`
}

// SelectUserDept 查看单个用户以及部门
func (x User) SelectUserDept(id int) ([]UserDept, error) {
	var userdept []UserDept
	if err := db.Debug().Table(x.TableName()).Select("user.*,dept.dept_name").Joins("left join dept on dept.dept_id=user.dept_id").Find(&userdept, id).Error; err != nil {
		return nil, err
	}
	return userdept, nil
}

// SelectUserDepts 查看所有用户以及部门，当然还可以结合limit进行查询分页
func (x User) SelectUserDepts() ([]UserDept, error) {
	var userdepts []UserDept
	if err := db.Debug().Table(x.TableName()).Select("user.*,dept.dept_name").Joins("left join dept on dept.dept_id=user.dept_id").Find(&userdepts).Error; err != nil {
		return nil, err
	}
	return userdepts, nil
}

// UserRole 用户以及角色联合查询============
type UserRole struct {
	User
	RoleName string `json:"roleName"`
}

// SelectUserRole 查看单个用户以及角色
func (x User) SelectUserRole(id int) ([]UserRole, error) {
	var userrole []UserRole
	if err := db.Debug().Table(x.TableName()).Select("user.*,role.role_name").Joins("left join role on role.role_id=user.role_id").Find(&userrole, id).Error; err != nil {
		return nil, err
	}
	return userrole, nil
}

// SelectUserRoles 查看所有用户以及角色，当然还可以结合limit进行查询分页
func (x User) SelectUserRoles() ([]UserRole, error) {
	var userroles []UserRole
	if err := db.Debug().Table(x.TableName()).Select("user.*,role.role_name").Joins("left join role on role.role_id=user.role_id").Find(&userroles).Error; err != nil {
		return nil, err
	}
	return userroles, nil
}

// UserView 用户以及所属角色和部门============
type UserView struct {
	User
	DeptName string `json:"deptName"`
	RoleName string `json:"roleName"`
}

// SelectUserDeptRole 查询单个用户以及其所属角色和部门
func (x User) SelectUserDeptRole(id int) ([]UserView, error) {
	var userview []UserView
	if err := db.Debug().Table(x.TableName()).Select("user.*,dept.dept_name,role.role_name").Joins("left join dept on dept.dept_id=user.dept_id").Joins("left join role on role.role_id=user.role_id").Find(&userview, id).Error; err != nil {
		return nil, err
	}
	return userview, nil
}

// SelectUserDeptRoles 查询多个用户以及其所属角色和部门
func (x User) SelectUserDeptRoles() ([]UserView, error) {
	var userviews []UserView
	if err := db.Debug().Table(x.TableName()).Select("user.*,dept.dept_name,role.role_name").Joins("left join dept on dept.dept_id=user.dept_id").Joins("left join role on role.role_id=user.role_id").Find(&userviews).Error; err != nil {
		return nil, err
	}
	return userviews, nil
}

// DeptUser 部门下的用户============
type DeptUser struct {
	Dept
	NickName string `json:"nickName"`
}

// SelectDeptUser 查询某部门下所有用户
func (x Dept) SelectDeptUser(id int) ([]DeptUser, error) {
	var deptuser []DeptUser
	if err := db.Debug().Table(x.TableName()).Select("dept.*,user.nick_name").Joins("left join user on user.dept_id=dept.dept_id").Find(&deptuser, id).Error; err != nil {
		return nil, err
	}
	return deptuser, nil
}

// RoleUser 角色下的用户============
type RoleUser struct {
	Role
	NickName string `json:"nickName"`
}

// SelectRoleUser 查询某个角色中的所有用户
func (x Role) SelectRoleUser(id int) ([]RoleUser, error) {
	var roleuser []RoleUser
	if err := db.Debug().Table(x.TableName()).Select("role.*,user.nick_name").Joins("left join user on user.role_id=role.role_id").Find(&roleuser, id).Error; err != nil {
		return nil, err
	}
	return roleuser, nil
}
func main() {
	// 1,初始化
	InitDB()
	defer db.Close()
	// 2,创建
	CreateDept()
	CreateRole()
	CreateUser()
	// 3,关联查询
	// 查看单个用户以及部门
	var s1 UserDept
	a1, _ := s1.SelectUserDept(1)
	fmt.Printf("用户ID为1的用户信息以及部门信息为：%v\n", a1)
	a2, _ := s1.SelectUserDepts()
	fmt.Printf("所有用户的信息以及部门信息为:%v\n", a2)
	var s2 UserRole
	b1, _ := s2.SelectUserRole(1)
	fmt.Printf("用户ID为1的用户信息以及角色为：%v\n", b1)
	b2, _ := s2.SelectUserRoles()
	fmt.Printf("所有用户的信息以及角色为:%v\n", b2)
	var s3 UserView
	c1, _ := s3.SelectUserDeptRole(1)
	fmt.Printf("用户ID为1的用户全部信息为：%v\n", c1)
	c2, _ := s3.SelectUserDeptRoles()
	fmt.Printf("所有用户的全部信息为：%v\n", c2)
	var s4 DeptUser
	d1, _ := s4.SelectDeptUser(2)
	fmt.Printf("部门ID为2的用户有：%v\n", d1)
	var s5 RoleUser
	e1, _ := s5.SelectRoleUser(1)
	fmt.Printf("角色ID为1的用户有：%v\n", e1)
}
