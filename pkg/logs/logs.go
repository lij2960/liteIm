/************************************************************
 * Author:        jackey
 * Date:        2022/10/18
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package logs

import (
	"fmt"
	"liteIm/pkg/common"
	"liteIm/pkg/config"
	"liteIm/pkg/utils"
	"strings"
)

func Info(f interface{}, v ...interface{}) {
	// 判断如果不是线上版本，则打印info信息
	if config.Env != common.RunModeProd {
		fmt.Println(utils.GetNowDateTime(), "[I] ", formatLog(f, v))
	}
}

func Alert(f interface{}, v ...interface{}) {
	fmt.Println(utils.GetNowDateTime(), "[A] ", formatLog(f, v))
}

func Error(f interface{}, v ...interface{}) {
	fmt.Println(utils.GetNowDateTime(), "\u001b[31m [E] \u001b[0m", formatLog(f, v))
}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if !strings.Contains(msg, "%") {
			// do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
