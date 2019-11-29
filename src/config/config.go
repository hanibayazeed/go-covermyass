package config

var GeneralLogs = []string{
	"/var/log/messages",
	"/var/log/kern.log",
	"/var/log/cron.log",
	"/var/log/maillog",
	"/var/log/boot.log",
	"/var/log/mysqld.log",
	"/var/log/qmail",
	"/var/log/httpd",
	"/var/log/lighttpd",
	"/var/log/secure",
	"/var/log/utmp",
	"/var/log/wtmp",
	"/var/log/yum.log",
	"/var/log/system.log",
	"/var/log/DiagnosticMessages",
	"/Library/Logs",
	"/Library/Logs/DiagnosticReports",
	"~/Library/Logs",
	"~/Library/Logs/DiagnosticReports",
}

var AuthLogs = []string{
	"/var/log/auth.log",
}

var HistoryLogs = []string{
	"~/.bash_history",
	"~/.zsh_history",
}
