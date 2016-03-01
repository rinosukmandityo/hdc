package test

import (
	. "github.com/eaciit/hdc/hive"
	"os"
	"testing"
)

var h *Hive
var e error

type Sample7 struct {
	Code        string `tag_name:"code"`
	Description string `tag_name:"description"`
	Total_emp   string `tag_name:"total_emp"`
	Salary      string `tag_name:"salary"`
}

type Students struct {
	Name    string
	Age     int
	Phone   string
	Address string
}

type SportMatch struct {
	Point       string
	HomeTeam    string
	AwayTeam    string
	MarkerImage string
	Information string
	Fixture     string
	Capacity    string
	Tv          string
}

func killApp(code int) {
	if h != nil {
		h.Conn.Close()
	}
	os.Exit(code)
}

func fatalCheck(t *testing.T, what string, e error) {
	if e != nil {
		t.Fatalf("%s: %s", what, e.Error())
	}
}

func TestHiveConnect(t *testing.T) {
	h = HiveConfig("192.168.0.223:10000", "default", "developer", "b1gD@T@", "")
}

func TestLoadFileWithWorker(t *testing.T) {
	err := h.Conn.Open()
	fatalCheck(t, "Populate", e)

	//test csv
	var Student Students
	retVal, err := h.LoadFileWithWorker("/home/developer/contoh2.txt", "studentworker", "csv", "dd/MM/yyyy", &Student, 3)
	t.Log(retVal)

	//test json
	//var SportMatch SportMatch
	//retValSport, err := h.LoadFileWithWorker("/home/developer/test json.txt", "sportmatchworker", "json", "dd/MM/yyyy", &SportMatch, 3)
	//t.Log(retValSport)

	if err != nil {
		t.Log(err)
	}
	h.Conn.Close()
}
