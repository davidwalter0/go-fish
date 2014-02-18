// copyright 2013-2014 rocky bernstein.
// whatis command

package fishcmd

import (
	"go/parser"
	"github.com/rocky/go-fish"
	"github.com/0xfaded/eval"
)

func init() {
	name := "whatis"
	repl.Cmds[name] = &repl.CmdInfo{
		Fn: WhatisCommand,
		Help: `whatis expression

Shows the type checker information for an expression. As a special
case, if expression is a package name, we'll confirm that.
 `,

		Min_args: 0,
		Max_args: -1,
	}
	repl.AddToCategory("data", name)
}

func WhatisCommand(args []string) {
	if len(args) == 2 {
		if _, ok := repl.Env.Pkg(args[1]).(*eval.SimpleEnv); ok  {
			repl.Msg("`%s' is a package", args[1])
			return
		}
	}
	line := repl.CmdLine[len(args[0]):len(repl.CmdLine)]
	if expr, err := parser.ParseExpr(line); err != nil {
		if pair := eval.FormatErrorPos(line, err.Error()); len(pair) == 2 {
			repl.Msg(pair[0])
			repl.Msg(pair[1])
		}
		repl.Errmsg("parse error: %s\n", err)
	} else if cexpr, errs := eval.CheckExpr(expr, repl.Env); len(errs) != 0 {
		for _, cerr := range errs {
			repl.Msg("%v", cerr)
		}
	} else {
		repl.Section(cexpr.String())
		if cexpr.IsConst() {
			repl.Msg("constant:\t%s", cexpr.Const())
		}
		knownTypes := cexpr.KnownType()
		if len(knownTypes) == 1{
			repl.Msg("type:\t%s", knownTypes[0])
		} else {
			for i, v := range knownTypes {
				repl.Msg("type[%d]:\t%s", i, v)
			}
		}
	}
}
