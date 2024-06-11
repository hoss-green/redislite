package stringparser 

import (
	"errors"
	"net"
	"strings"

	"redislite/app/commands/errormessages"
	"redislite/app/commands/parsers/parserentities"
	"redislite/app/data"
	"redislite/app/setup"
)

// STRING COMMAND PARSER
func ParseStringCommand(connpointer *net.Conn, redisCommand data.RedisCommand, server *setup.Server) parserentities.ParserInfo {
	conn := *connpointer
	var err error = nil
	instruction := strings.ToUpper(redisCommand.Command)
	switch instruction {
	case "SET":
		if redisCommand.ParamLength < 2 {
			err = errors.New(errormessages.IncorrectArgumentsError)
			break
		}
		err = set(conn, server, redisCommand)
	case "GET":
		if redisCommand.ParamLength != 1 {
			err = errors.New(errormessages.IncorrectArgumentsError)
			break
		}
		err = get(conn, server, redisCommand.Params[0])
	case "APPEND":
		if redisCommand.ParamLength != 2 {
			err = errors.New(errormessages.IncorrectArgumentsError)
			break
		}
		err = append(conn, server, redisCommand)
	default:
		return parserentities.ParserInfo{Executed: false, Err: nil}
	}

	if err != nil {
		return parserentities.ParserInfo{Executed: false, Err: err}
	}

	return parserentities.ParserInfo{Executed: true, Err: nil}
}
