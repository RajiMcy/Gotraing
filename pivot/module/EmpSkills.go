package module

import (
	"fmt"
	"log"

	"github.com/raji802/gotraining/pivot/config"
)

type EmpSkill struct {
	Skill_nm    string
	Proficiency string
	Version     string
	IsPrimary   string
	Last_used   string
	Total_exp   int
	Attuid      string
	First_nm    string
	Last_nm     string
}

type Data struct {
	Items []EmpSkill
}

func Skills(attuid string) Data {

	Rowdata := Data{}
	db, err := config.GetDB()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	var Skill EmpSkill

	stmt := "SELECT b.skill_nm,a.proficiency,a.version,a.is_primary,a.total_exp_in_month ,strftime('%Y-%m-%d',a.last_used),a.attuid,c.first_nm,c.last_nm FROM employeeskill a INNER JOIN skill b ON a.skill_id=b.skill_id INNER JOIN employee c on a.attuid=c.attuid where a.attuid=? ;"
	rows, err1 := db.Query(stmt, attuid)
	if err1 != nil {
		fmt.Println("Failed to run the select query.")
		fmt.Println(err1)
	} else {
		//fmt.Println("row", row)
		for rows.Next() {
			err2 := rows.Scan(&Skill.Skill_nm, &Skill.Proficiency, &Skill.Version, &Skill.IsPrimary, &Skill.Total_exp, &Skill.Last_used, &Skill.Attuid, &Skill.First_nm, &Skill.Last_nm)
			if err2 != nil {
				fmt.Println("Error in fetching the skill details")
			}
			//fmt.Println(Skill.Skill_nm, Skill.Proficiency, Skill.Version, Skill.IsPrimary, Skill.Last_used, Skill.Total_exp)
			Rowdata.Items = append(Rowdata.Items, Skill)
			//fmt.Println(Rowdata)
		}
	}
	return Rowdata
}
