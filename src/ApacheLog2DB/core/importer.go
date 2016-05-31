package core
import (
	"time"
	"strings"
	"strconv"
)


func ImportLog(log string) error{
	for _,line := range APACHE_COMBINED.FindAllStringSubmatch(log,-1) {
		source := line[1]
		ident := line[2]
		username := line[3]
		occurred,err := time.Parse(APACHE_TIME_LAYOUT,line[4])
		if err != nil {return err}
		request := strings.Split(line[5], " ")

		var verb string
		var uri string
		var protocol string
		switch len(request) {
		case 3:
			verb = request[0]
			uri = request[1]
			protocol = request[2]
		case 2:
			uri = request[0]
			protocol = request[1]
		case 1:
			uri = request[0]
		}

		status,err := strconv.Atoi(line[6])
		if err != nil {return err}

		size,err := strconv.Atoi(line[7])
		if err != nil {return err}

		referrer := line[8]

		agentname := line[9]

	}
}