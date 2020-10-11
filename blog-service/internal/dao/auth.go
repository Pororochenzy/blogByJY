package dao

import "github.com/go-programming-tour-book/blog-service/internal/model"

func (d *Dao) GetAUth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
