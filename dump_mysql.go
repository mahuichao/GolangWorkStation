package main
import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var MYSQL *sql.DB

func init(){
	fmt.Printf("init start\n")
	var err error
	MYSQL,err=sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=false", "user", "password", "120.25.177.114", "3306", "test"))
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("init end\n")
}

// 导出课程信息
func DumpCourseData(){
	file,_:=os.OpenFile("course",os.O_APPEND|os.O_RDWR|os.O_CREATE,0666)
	
	defer file.Close()
	fmt.Printf("start\n")
	
	rows,_:=MYSQL.Query(fmt.Sprintf("select "+
			"course_id, course_name "+
			"from t_course_active"))
	fmt.Printf("end\n")
	
	var course_id string
	var course_name string
	for rows.Next(){
		fmt.Printf("Scan In\n")
		rows.Scan(
			&course_id,
			&course_name,
		)
		fmt.Printf("Scan Out\n")
	fmt.Printf("begin to write\n")
		file.WriteString(course_id+"|||"+course_name+"\n")
	}
}

//导出文件课程信息
func DumpFiles(){
		file,_:=os.OpenFile("files",os.O_APPEND|os.O_RDWR|os.O_CREATE,0666)
		defer file.Close()
		rows,_:=MYSQL.Query(fmt.Sprintf("select "+
			"file_id, course_id, category, title "+
			"from t_course_pub_file"))
		var file_id string
		var course_id string
		var category string
		var title string
		for rows.Next(){
			rows.Scan(
				&file_id,
				&course_id,
				&category,
				&title,
		)
		file.WriteString(file_id+"|||"+course_id+"|||"+category+"|||"+title+"\n")
		}
}

//导出学生信息
func DumpStu(){
		file,_:=os.OpenFile("students",os.O_APPEND|os.O_RDWR|os.O_CREATE,0666)
		defer file.Close()
		rows,_:=MYSQL.Query(fmt.Sprintf("select "+
		"id, name, major_id, major_name, folk "+
		"from wlxt_student_info"))
		var id string
		var name string
		var major_id string
		var major_name string
		var gender string
		for rows.Next(){
			rows.Scan(
				&id,
				&name,
				&major_id,
				&major_name,
				&gender,
			)
		file.WriteString(id+"|||"+name+"|||"+major_id+"|||"+major_name+"|||"+gender+"\n")
		}
}
func main(){
	DumpCourseData()
	DumpFiles()
	DumpStu()
}
