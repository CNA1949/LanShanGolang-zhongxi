package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Storehouse struct {
	StoreCode string
	Capacity  int
}

type ClothingInfo struct {
	ClothingCode string
	Size         string
	Price        int
	ClothingType string
}

type Supplier struct {
	SupplierCode string
	SupplierName string
}

type SupplySituation struct {
	ClothingCode string
	SupplierCode string
	Quality      string
}

//数据库配置
const (
	driverName   = "mysql"
	userName     = "root"
	password     = "123456"
	ip           = "127.0.0.1"
	port         = "3306"
	databaseName = "third_week"
)

var DB *sql.DB // DB数据库连接池

func ConnectToDatabase() error {
	var err error
	// 连接数据库
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", userName, password, ip, port, databaseName)
	//打开数据库
	DB, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		return errors.New(fmt.Sprintf("连接数据库失败：%v", err))
	}
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(10)
	//设置数据库最大闲置连接数
	DB.SetMaxIdleConns(5)
	//验证连接
	if err = DB.Ping(); err != nil {
		return errors.New(fmt.Sprintf("连接失败：%v", err))
	}
	return nil
}

func InsertData(sqlStr string) error {
	//向数据表中插入数据
	if DB == nil {
		return errors.New("数据库未连接")
	}
	transactions, err := DB.Begin() //开启事务
	if err != nil {
		transactions.Rollback() // 回滚
		return errors.New(fmt.Sprintf("开启事务失败%v", err))
	}

	_, err = DB.Exec(sqlStr)
	if err != nil {
		transactions.Rollback() // 回滚
		return errors.New(fmt.Sprintf("插入失败：%v", err))
	}
	transactions.Commit() //提交事务
	return nil
}

func readRows(rows *sql.Rows, tableName string) {
	// 读取数据
	defer rows.Close()
	switch tableName {
	case "storehouse":
		info := Storehouse{}
		for rows.Next() {
			err := rows.Scan(&info.StoreCode, &info.Capacity)
			if err != nil {
				fmt.Println("读取失败！", err)
			}
			fmt.Println(info)
		}
	case "clothing_info":
		info := ClothingInfo{}
		for rows.Next() {
			err := rows.Scan(&info.ClothingCode, &info.Size, &info.Price, &info.ClothingType)
			if err != nil {
				fmt.Println("读取失败！", err)
			}
			fmt.Println(info)
		}
	case "supplier":
		info := Supplier{}
		for rows.Next() {
			err := rows.Scan(&info.SupplierCode, &info.SupplierName)
			if err != nil {
				fmt.Println("读取失败！", err)
			}
			fmt.Println(info)
		}
	case "supply_situation":
		info := SupplySituation{}
		for rows.Next() {
			err := rows.Scan(&info.ClothingCode, &info.SupplierCode, &info.Quality)
			if err != nil {
				fmt.Println("读取失败！", err)
			}
			fmt.Println(info)
		}
	}

}

func SelectData(tableName string, sqlStr string) error {
	//查询
	if DB == nil {
		return errors.New("数据库未连接")
	}

	rows, err := DB.Query(sqlStr)
	if err != nil {
		return errors.New(fmt.Sprintf("查询失败:%v", err))
	}
	readRows(rows, tableName)
	return nil
}

func UpdateData(sqlStr string) error {
	// 更新
	if DB == nil {
		return errors.New("数据库未连接")
	}
	transactions, err := DB.Begin() //开启事务
	if err != nil {
		transactions.Rollback() // 回滚
		return errors.New(fmt.Sprintf("开启事务失败%v", err))
	}

	_, err = DB.Exec(sqlStr)
	if err != nil {
		transactions.Rollback() // 回滚
		return errors.New(fmt.Sprintf("更新失败：%v", err))
	}
	transactions.Commit() //提交事务
	return nil
}

func DeleteData(sqlStr string) error {
	// 删除
	if DB == nil {
		return errors.New("数据库未连接")
	}
	transactions, err := DB.Begin() //开启事务
	if err != nil {
		transactions.Rollback() // 回滚
		return errors.New(fmt.Sprintf("开启事务失败%v", err))
	}

	_, err = DB.Exec(sqlStr)
	if err != nil {
		transactions.Rollback() // 回滚
		return errors.New(fmt.Sprintf("更新失败：%v", err))
	}
	transactions.Commit() //提交事务

	return nil
}

func main() {
	// 连接数据库
	err := ConnectToDatabase()
	if err != nil {
		fmt.Println(err)
	}

	//// 查询服装尺码为'S'且销售价格在100以下的服装信息
	//err = SelectData("clothing_info", "SELECT * FROM clothing_info WHERE size='S' and price<100")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//// 查询仓库容量最大的仓库信息。
	//err = SelectData("storehouse", "SELECT * FROM storehouse WHERE capacity = (SELECT MAX(capacity) FROM storehouse)")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//// 查询服装编码以‘A’开始开头的服装
	//err = SelectData("clothing_info", "SELECT * FROM clothing_info WHERE clothing_code LIKE 'A%'")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//// 查询服装质量等级有不合格的供应商信息。
	//err = SelectData("supplier", "SELECT * FROM supplier WHERE supplier_code IN (SELECT supplier_code FROM supply_situation WHERE quality = '不合格')")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//// 把服装尺寸为'S'的服装的销售价格均在原来基础上提高10%。
	//err = UpdateData("UPDATE clothing_info SET price=price*(1+0.1) WHERE size='S'")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//// 删除所有服装质量等级不合格的供应情况。
	//err = DeleteData("DELETE FROM supply_situation WHERE quality='不合格'")
	//if err != nil {
	//	fmt.Println(err)
	//}

	// 向每张表插入一条记录。
	// # 向storehouse表中插入一条数据
	//err = InsertData("INSERT INTO storehouse (store_code,capacity) VALUES ('CK1006',6000)")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//// # 向clothing_info表中插入一条数据
	//err = InsertData("INSERT INTO clothing_info (clothing_code,size,price,clothing_type) VALUES ('CFZ00006','M',90,'C')")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//// # 向supplier表中插入一条数据
	//err = InsertData("INSERT INTO supplier (supplier_code,supplier_name) VALUES ('GYS1006','供应商F')")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//// # 向supply_situation表中插入一条数据
	//err = InsertData("INSERT INTO supply_situation (clothing_code,supplier_code,quality) VALUES ('CFZ00006','GYS1006','不合格')")
	//if err != nil {
	//	fmt.Println(err)
	//}
}
