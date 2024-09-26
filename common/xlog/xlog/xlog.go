package xlog

import (
	gLocal "github.com/mooncake9527/orange-core/common/xlog/g_local"
	"log/slog"
	"runtime"
	"strconv"
)

func appendArgs(args ...any) []any {
	_, file, line, _ := runtime.Caller(2)
	args = append([]interface{}{"file", file + ":" + strconv.Itoa(line)}, args...)
	reqId, userId, companyId, appKey := gLocal.GetIds()
	if reqId != "" {
		args = append(args, "reqId", reqId)
	}
	if userId != "" {
		args = append(args, "userId", userId)
	}
	if companyId != "" {
		args = append(args, "companyId", companyId)
	}
	if appKey != "" {
		args = append(args, "appKey", appKey)
	}
	return args
}

func Debug(msg string, args ...any) {
	slog.Debug(msg, appendArgs(args...)...)
}

func Info(msg string, args ...any) {
	slog.Info(msg, appendArgs(args...)...)
}

func Warn(msg string, args ...any) {
	slog.Warn(msg, appendArgs(args...)...)
}

func Error(msg string, args ...any) {
	slog.Error(msg, appendArgs(args...)...)
}

func With(args ...any) *slog.Logger {
	return slog.With(appendArgs(args...)...)
}
