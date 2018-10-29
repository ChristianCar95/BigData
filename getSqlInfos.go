package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "os"
//import "os/exec"

func main() {

	type Person struct {
		Name string
		Vorname string
	}

	db,err:= sql.Open("mysql", "hadoop:hadoophadoophadoop@/test")
	if err != nil {panic(err)}
	
	defer db.Close()

	err = db.Ping()
	if err != nil {panic(err)}

	rows, err := db.Query("SELECT * FROM test")
	if err != nil {panic(err.Error())}
	
	var people []Person

	for rows.Next(){

		var person Person

		err = rows.Scan(&person.Name, &person.Vorname)
		if err != nil {
			panic(err.Error())
		}
		people = append(people, person)
	}
	
	for i:=0; i<len(people); i++ {
		fmt.Printf(people[i].Name + " " + people[i].Vorname + "\n")
	}

	os.Remove("/var/www/html/index.html")
	os.Create("/var/www/html/index.html")

	homepageTrackInfos, _ := os.OpenFile("/var/www/html/index.html", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer homepageTrackInfos.Close()

	homepageTrackInfos.WriteString("<HTML>\n	<head>\n		<title>Track information</title>\n	</head>\n	<body>\n")

	for i:=0; i<len(people); i++{
		homepageTrackInfos.WriteString("\n			<p>" + people[i].Name + "	" + people[i].Vorname + "</p>")
	}

	homepageTrackInfos.WriteString("\n	</body>\n</HTML>")
	
	

//	cmd:=exec.Command("bash", "move.sh")
//	cmd.Run()



}

