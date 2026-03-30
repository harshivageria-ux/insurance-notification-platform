package main
import (
  "context"
  "fmt"
  "github.com/jackc/pgx/v5/pgxpool"
)
func main(){
  conn,err:=pgxpool.New(context.Background(),"postgres://admin_user:N96wS%3E80%3FM%2AR@15.206.192.126:5432/postgres")
  if err!=nil{panic(err)}
  defer conn.Close()
  rows,err:=conn.Query(context.Background(),"select id, code, name, description from notification_categories_master limit 5")
  if err!=nil{panic(err)}
  defer rows.Close()
  for rows.Next(){var id,code,name,desc string; err=rows.Scan(&id,&code,&name,&desc); if err!=nil{panic(err)}; fmt.Printf("id=%s code=%s name=%s desc=%s\n",id,code,name,desc)}
  if rows.Err()!=nil{panic(rows.Err())}
}
