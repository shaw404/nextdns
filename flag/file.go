package flag

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

// ParseFile reads file and append each line as flags to the command line. Lines
// starting with a # or empty are ignored.
func ParseFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		panic(err.Error())
	}

	s := bufio.NewScanner(f)
	var args []string
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		arg := line
		value := ""
		if idx := strings.IndexByte(line, ' '); idx != -1 {
			arg = line[:idx]
			value = strings.TrimSpace(line[idx+1:])
		}
		if value != "" {
			// Accept yes/no as boolean values
			switch value {
			case "yes":
				value = "true"
			case "no":
				value = "false"
			}
			arg += "=" + value
		}
		args = append(args, "-"+arg)
	}
	if err := s.Err(); err != nil {
		panic(err.Error())
	}

	_ = flag.CommandLine.Parse(args)
}
