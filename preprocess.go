// Copyright 2013-2014 Rocky Bernstein.
// REPL Command processing

package repl

import (
	"strings"
)

func wasProcessed(line string) bool {
	line = strings.Trim(line, " \t\n")
	args  := strings.Split(line, " ")
	if len(args) == 0 || len(args[0]) == 0 {
		Msg("Empty line skipped")
		// gnureadline.RemoveHistory(gnureadline.HistoryLength()-1)
		return true
	}
	if args[0][0] == '/' && len(args) > 1 && args[0][1] == '/' {
		// gnureadline.RemoveHistory(gnureadline.HistoryLength()-1)
		Msg(line) // echo line but do nothing
		return true
	}

	name := args[0]
	if newname := LookupCmd(name); newname != "" {
		name = newname
	}
	cmd := Cmds[name];

	if cmd != nil {
		if ArgCountOK(cmd.Min_args, cmd.Max_args, args) {
			Cmds[name].Fn(args)
		}
		return true
	}
	return false
}
