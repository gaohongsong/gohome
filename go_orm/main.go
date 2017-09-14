package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
	"time"
)

type dict map[string]interface{}

var DATABASE = dict{
	"name":     "gorm",
	"host":     "127.0.0.1",
	"port":     3306,
	"user":     "root",
	"password": "root",
}

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string  `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
	Num      int     `gorm:"AUTO_INCREMENT"` // 自增

	Emails   []Email   // One-To-Many (拥有多个 - Email表的UserID作外键)
	Products []Product // One-To-Many (拥有多个 - Product表的UserID作外键)

	BillingAddress   Address // One-To-One (属于 - 本表的BillingAddressID作外键)
	BillingAddressID sql.NullInt64
	//BillingAddressId  sql.NullInt64 `sql:"type:bigint REFERENCES address(id)"`

	IgnoreMe  int `gorm:"-"`                                // 忽略这个字段
	Languages []Language `gorm:"many2many:user_languages;"` // Many-To-Many , 'user_languages'是连接表

	//测试发现在结构体标签中指定外键和约束对创建表没有任何指导意义，仅当在Migrate后调用AddForeignKey后执行Alter table操作
	//才能为指定的外键添加约束
	//Profile   Profile
	//ProfileID int	`sql:"REFERENCES gorm_profiles(id)"`	// 外键约束定义
	//Profile   Profile `gorm:"ForeignKey:ProfileID;AssociationForeignKey:ID"`
	Profile   Profile
	ProfileID int
}

type Profile struct {
	ID         int
	//UserID     int
	Interest   string    `gorm:"size:255"`
	Company    string    `gorm:"size:32"`
	CreditCard CreditCard // One-To-One (拥有一个 - CreditCard表的UserID作外键)
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` // 创建索引并命名，如果找到其他相同名称的索引则创建组合索引
	Code string `gorm:"index:idx_name_code"` // `unique_index` also works
}

type Product struct {
	gorm.Model
	UserID int     `gorm:"index"`                  // 外键 (属于), tag `index`是为该列创建索引
	Code   string    `gorm:"not null"`
	Price  uint
	Type   string        `gorm:"default:'django'"` // 默认值
}

type Email struct {
	ID         int
	UserID     int     `gorm:"index"`                          // 外键 (属于), tag `index`是为该列创建索引
	Email      string  `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` // 设置字段为非空并唯一
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

type CreditCard struct {
	gorm.Model
	ProfileID int
	Number    string `gorm:"column:credit_no"` // 设置列名为`credit_no`
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		DATABASE["user"],
		DATABASE["password"],
		DATABASE["host"],
		DATABASE["port"],
		DATABASE["name"])
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	//打印详细日志
	db.LogMode(true)

	defer db.Close()

	// 自定义数据表命名规则
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "gorm_" + defaultTableName
	}

	// 自动迁移模式
	db.AutoMigrate(&User{}, &Profile{}, &Product{}, &Email{}, &Address{}, &Language{}, &CreditCard{})
	// 创建外键约束：User.profile_id(int) => Profile.id(int)
	db.Model(&User{}).AddForeignKey("profile_id", "gorm_profiles(id)", "CASCADE", "RESTRICT")
	// 检查表是否存在
	fmt.Println("product exist? ", db.HasTable("gorm_products"), db.HasTable(&Product{}))
	// 修改列
	//db.Model(&Product{}).ModifyColumn("type", "string")

	//profile := Profile{Interest: "fly", Company: "tencent"}
	//db.Create(&profile)

	user := User{
		Name:           "hongsong1",
		Age:            25,
		Birthday:       time.Now(),
		BillingAddress: Address{
			Address1: "Billing Address - Address 11",
			Address2: "Billing Address - Address 2",
			Post: sql.NullString{String: "Pickin", Valid: true},
		},
		Emails: []Email{
			{Email: "jinzhu@example.com1"},
			{Email: "jinzhu-2@example@example.com1"},
		},
		Languages: []Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
		Products: []Product{
			{Code: "FEIJI", Price: 1},
			{Code: "DAPAO", Price: 12},
		},
		//ProfileID: profile.ID,
		Profile: Profile{Interest: "fly", Company: "tencent"},
	}
	db.Create(&user)

	//profile.UserID = user.ID
	//db.Save(&profile)

	//var user User
	//db.Where("name = ?", "hongsong").First(&user)
	// 要更改它的值, 你需要使用`Update`
	//db.Model(&user).Update("CreatedAt", time.Now())

	//执行Save后UpdateAt自动更新
	//user.Name = "miya"
	//db.Save(&user)

	// 创建
	db.Create(&Product{Code: "dji", Price: 1000})
	// 读取
	var product Product
	db.First(&product, 1)                 // 查询id为1的product
	db.First(&product, "code = ?", "dji") // 查询code为l1212的product

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)
	fmt.Println(product)

	// 删除 - 删除product
	//db.Delete(&product)
}
