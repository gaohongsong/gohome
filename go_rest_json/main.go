package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "log"
    "net/http"
    "time"
)

type Reminder struct {
    Id        int64     `json:"id"`
    Message   string    `sql:"size:1024" json:"message"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    DeletedAt time.Time `json:"-"`
}

type Impl struct {
    DB *gorm.DB
}

// db连接初始化
func (i *Impl) InitDB() {
    var err error
    i.DB, err = gorm.Open("mysql", "root:bk@321@tcp(db:3306)/gorm?charset=utf8&parseTime=True")
    if err != nil {
        log.Fatalf("Got error when connect database, the error is '%v'", err)
    }
    i.DB.LogMode(true)
}

// db结构初始化
func (i *Impl) InitSchema() {
    i.DB.AutoMigrate(&Reminder{})
}

// 查reminders(:all)
func (i *Impl) GetAllReminders(w rest.ResponseWriter, r *rest.Request) {
    reminders := []Reminder{}
    i.DB.Find(&reminders)
    w.WriteJson(&reminders)
}

// 查reminder(:id)
func (i *Impl) GetReminder(w rest.ResponseWriter, r *rest.Request) {
    id := r.PathParam("id")
    reminder := Reminder{}
    if i.DB.First(&reminder, id).Error != nil {
        rest.NotFound(w, r)
        return
    }
    w.WriteJson(&reminder)
}
// 增reminder
func (i *Impl) PostReminder(w rest.ResponseWriter, r *rest.Request) {
    reminder := Reminder{}
    if err := r.DecodeJsonPayload(&reminder); err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := i.DB.Save(&reminder).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteJson(&reminder)
}

// 改reminder
func (i *Impl) PutReminder(w rest.ResponseWriter, r *rest.Request) {

    id := r.PathParam("id")
    reminder := Reminder{}
    if i.DB.First(&reminder, id).Error != nil {
        rest.NotFound(w, r)
        return
    }

    updated := Reminder{}
    if err := r.DecodeJsonPayload(&updated); err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    reminder.Message = updated.Message

    if err := i.DB.Save(&reminder).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteJson(&reminder)
}

// 删reminder
func (i *Impl) DeleteReminder(w rest.ResponseWriter, r *rest.Request) {
    id := r.PathParam("id")
    reminder := Reminder{}
    if i.DB.First(&reminder, id).Error != nil {
        rest.NotFound(w, r)
        return
    }
    if err := i.DB.Delete(&reminder).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func main() {

    i := Impl{}
    i.InitDB()
    i.InitSchema()

    api := rest.NewApi()

	//定义中间件 rest.DefaultDevStack|rest.DefaultProdStack|rest.DefaultCommonStack
    api.Use(rest.DefaultDevStack...)
    router, err := rest.MakeRouter(
        rest.Get("/reminders", i.GetAllReminders),
        rest.Post("/reminders", i.PostReminder),
        rest.Get("/reminders/:id", i.GetReminder),
        rest.Put("/reminders/:id", i.PutReminder),
        rest.Delete("/reminders/:id", i.DeleteReminder),
    )
    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)
    log.Fatal(http.ListenAndServe(":8000", api.MakeHandler()))
}

