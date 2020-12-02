package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

/*
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
应该Wrap。 dao层一般为我们应用程序的最底层，和sql Driver交互，这里出错，需要将堆栈信息包含在里面。网上抛，在最上层打印这个错误信息。
*/
func main() {
	serviceRun(context.TODO())
}

// service层
func serviceRun(ctx context.Context) {
	err := biz(ctx)
	if err != nil {
		fmt.Printf("serviceRun error. err=%+v \n", err)
	}
	return
}

// 业务biz层
func biz(ctx context.Context) error {
	_, err := dao(ctx, 100)
	if err != nil {
		return errors.WithMessage(err, "biz")
	}
	return nil
}

// dao层
func dao(ctx context.Context, id int) (interface{}, error) {
	data, err := getRow(ctx, id)
	return data, errors.Wrapf(err, "id=%+v", id)
}

// 底层数据库driver层
func getRow(ctx context.Context, id int) (interface{}, error) {
	//模拟出错
	return nil, sql.ErrNoRows
}
