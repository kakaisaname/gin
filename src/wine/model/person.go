package model

import (
	db "wine/database"
	//"time"
	"fmt"
	"time"
)
type Fans struct {
	Id 			int `json:"id" form:"id"`
	Name 		string `json:"name" form:"name"`
	Address 	string `json:"address" form:"address"`
	Ctime		int `json:"ctime" form:"ctime"`
}

//db.Query()表示向数据库发送一个query，defer rows.Close()非常重要，遍历rows使用rows.Next()，
// 把遍历到的数据存入变量使用rows.Scan()
//结果集(rows)未关闭前，底层的连接处于繁忙状态。当遍历读到最后一条记录时，会发生一个内部EOF错误，自动调用rows.Close()，
// 但是如果提前退出循环，rows不会关闭，连接不会回到连接池中，连接也不会关闭。所以手动关闭非常重要
func (p *Fans) AddFans() (id int64,err error) {
	exec, err := db.Db.Exec("INSERT INTO fans(name,address,ctime) VALUES (?,?,?)",p.Name,p.Address,time.Now().Unix())
	fmt.Println(exec)
	if err != nil {
		fmt.Println(err)
	}
	id,err = exec.LastInsertId()
	return
}

func (p *Fans) GetFans() (fan Fans) {
	row := db.Db.QueryRow("SELECT name,address,ctime FROM fans WHERE id = ?", p.Id)
	row.Scan(&fan.Name,&fan.Address,&fan.Ctime)   //必须去地址
	return
}

//获取全部数据
func (p *Fans) GetAllFan() (fans []Fans,err error)  {  //返回对个Fans
	fans = make([]Fans,0) //创建fan切片  因为返回值中为fan
	rows, err := db.Db.Query("SELECT id,name,ctime FROM fans")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var fan Fans   //Fans类型
		rows.Scan(&fan.Id,&fan.Name,&fan.Ctime)
		fans = append(fans,fan)
	}
	//如果循环中发生错误会自动运行rows.Close()，用rows.Err()接收这个错误，Close方法可以多次调用。
	// 循环之后判断error是非常必要的。
	if err = rows.Err();err != nil {
		return
	}
	return
}

func (p *Fans) DelFans() (ra int64,err error) {
	result, err := db.Db.Exec("DELETE FROM fans WHERE id=?", p.Id)
	if err != nil {
		return
	}
	//受影响数
	ra, err = result.RowsAffected()
	return
}

func (p *Fans) UpdateFan() (ra int64,err error)  {
	result,err := db.Db.Exec("UPDATE fans SET name=?,address=? WHERE id=?",p.Name,p.Address,p.Id)
	if err != nil {
		return
	}
	ra,err = result.RowsAffected()
	return
}