package main()
import fmt
import database/sql
import _"github.com/mattn/go-sqlite3"


func main()
{
con, e1:= sql.open("sqlite3", "st1.db")
if e1!= nil{
	fmt.Println("issue", e1)
	return
}
fmt.Println("connected")
defer con.Close()
}

sql:= "create table student(rno int primary key, name text)"
res,e2 := con.Exec(sql)
if e2!= nil{
	fmt.Println("creating issue",e2)
	return
}
fmt.Println("connected")
defer con.Close()

sql:="insert into values( ?,?)"
stmt ,_:=con.Prepare(sql);

var rno int
var name string
fmt.Println("Enter roll no")
fmt.Scanln("&rno")
fmt.Println("Enter name")
fmt.Scanln("&name")
res ,_ :=stmt.Exec(rno,name)
rc ,_:=res.RowsAffected()
fmt.Prinln(rc,"Rows inserted")
}
