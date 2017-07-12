package patterns

// Grok is a collection of common grok patterns that extend the common defaults.
var Grok = map[string]string{
	"USERNAME":  `[a-zA-Z0-9._-]+`,
	"USER":      `%{USERNAME}`,
	"INT":       `(?:[+-]?(?:[0-9]+))`,
	"BASE10NUM": `([+-]?(?:[0-9]+(?:\.[0-9]+)?)|\.[0-9]+)`,
	"NUMBER":    `(?:%{BASE10NUM})`,
	"BASE16NUM": `(0[xX]?[0-9a-fA-F]+)`,
	//BASE16FLOAT \b(?<![0-9A-Fa-f.])(?:[+-]?(?:0x)?(?:(?:[0-9A-Fa-f]+(?:\.[0-9A-Fa-f]*)?)|(?:\.[0-9A-Fa-f]+)))\b

	"POSINT":       `\b(?:[1-9][0-9]*)\b`,
	"NONNEGINT":    `\b(?:[0-9]+)\b`,
	"WORD":         `\b\w+\b`,
	"NOTSPACE":     `\S+`,
	"SPACE":        `\s*`,
	"DATA":         `.*?`,
	"GREEDYDATA":   `.*`,
	"QUOTEDSTRING": `"([^"\\]*(\\.[^"\\]*)*)"|\'([^\'\\]*(\\.[^\'\\]*)*)\'`,
	"UUID":         `[A-Fa-f0-9]{8}-(?:[A-Fa-f0-9]{4}-){3}[A-Fa-f0-9]{12}`,

	// Networking
	"MAC":        `(?:%{CISCOMAC}|%{WINDOWSMAC}|%{COMMONMAC})`,
	"CISCOMAC":   `(?:(?:[A-Fa-f0-9]{4}\.){2}[A-Fa-f0-9]{4})`,
	"WINDOWSMAC": `(?:(?:[A-Fa-f0-9]{2}-){5}[A-Fa-f0-9]{2})`,
	"COMMONMAC":  `(?:(?:[A-Fa-f0-9]{2}:){5}[A-Fa-f0-9]{2})`,
	"IPV6":       `((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?`,
	"IPV4":       `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`,
	"IP":         `(?:%{IPV6}|%{IPV4})`,
	"HOSTNAME":   `\b(?:[0-9A-Za-z][0-9A-Za-z-]{0,62})(?:\.(?:[0-9A-Za-z][0-9A-Za-z-]{0,62}))*(\.?|\b)`,
	"HOST":       `%{HOSTNAME}`,
	"IPORHOST":   `(?:%{HOSTNAME}|%{IP})`,
	"HOSTPORT":   `%{IPORHOST}:%{POSINT}`,

	// paths
	"PATH": `(?:%{UNIXPATH}|%{WINPATH})`,
	//UNIXPATH (?>/(?>[\w_%!$@:.,-]+|\\.)*)+
	"UNIXPATH": `(/[\w_%!$@:.,-]?/?)(\S+)?`,
	"TTY":      `(?:/dev/(pts|tty([pq])?)(\w+)?/?(?:[0-9]+))`,
	//WINPATH (?>[A-Za-z]+:|\\)(?:\\[^\\?*]*)+
	"WINPATH": `([A-Za-z]:|\\)(?:\\[^\\?*]*)+`,

	"URIPROTO": `[A-Za-z]+(\+[A-Za-z+]+)?`,
	"URIHOST":  `%{IPORHOST}(?::%{POSINT:port})?`,
	// uripath comes loosely from RFC1738, but mostly from what Firefox
	// doesn't turn into %XX
	"URIPATH": `(?:/[A-Za-z0-9$.+!*'(){},~:;=@#%_\-]*)+`,
	//URIPARAM \?(?:[A-Za-z0-9]+(?:=(?:[^&]*))?(?:&(?:[A-Za-z0-9]+(?:=(?:[^&]*))?)?)*)?
	"URIPARAM":     `\?[A-Za-z0-9$.+!*'|(){},~@#%&/=:;_?\-\[\]]*`,
	"URIPATHPARAM": `%{URIPATH}(?:%{URIPARAM})?`,
	"URI":          `%{URIPROTO}://(?:%{USER}(?::[^@]*)?@)?(?:%{URIHOST})?(?:%{URIPATHPARAM})?`,

	// Months: January, Feb, 3, 03, 12, December
	"MONTH":    `\b(?:Jan(?:uary)?|Feb(?:ruary)?|Mar(?:ch)?|Apr(?:il)?|May|Jun(?:e)?|Jul(?:y)?|Aug(?:ust)?|Sep(?:tember)?|Oct(?:ober)?|Nov(?:ember)?|Dec(?:ember)?)\b`,
	"MONTHNUM": `(?:0?[1-9]|1[0-2])`,
	"MONTHDAY": `(?:(?:0[1-9])|(?:[12][0-9])|(?:3[01])|[1-9])`,

	// Days: Monday, Tue, Thu, etc...
	"DAY": `(?:Mon(?:day)?|Tue(?:sday)?|Wed(?:nesday)?|Thu(?:rsday)?|Fri(?:day)?|Sat(?:urday)?|Sun(?:day)?)`,

	// Years?
	//YEAR (?>\d\d){1,2}
	//c
	"YEAR": `(\d\d){1,2}`,

	"HOUR":   `(?:2[0123]|[01]?[0-9])`,
	"MINUTE": `(?:[0-5][0-9])`,
	// '60' is a leap second in most time standards and thus is valid.
	"SECOND": `(?:(?:[0-5][0-9]|60)(?:[:.,][0-9]+)?)`,
	//TIME (?!<[0-9])%{HOUR}:%{MINUTE}(?::%{SECOND})(?![0-9])
	//c
	"TIME": `([^0-9]?)%{HOUR}:%{MINUTE}(?::%{SECOND})([^0-9]?)`,
	// datestamp is YYYY/MM/DD-HH:MM:SS.UUUU (or something like it)
	"DATE_US":           `%{MONTHNUM}[/-]%{MONTHDAY}[/-]%{YEAR}`,
	"DATE_EU":           `%{MONTHDAY}[./-]%{MONTHNUM}[./-]%{YEAR}`,
	"ISO8601_TIMEZONE":  `(?:Z|[+-]%{HOUR}(?::?%{MINUTE}))`,
	"ISO8601_SECOND":    `(?:%{SECOND}|60)`,
	"TIMESTAMP_ISO8601": `%{YEAR}-%{MONTHNUM}-%{MONTHDAY}[T ]%{HOUR}:?%{MINUTE}(?::?%{SECOND})?%{ISO8601_TIMEZONE}?`,
	"DATE":              `%{DATE_US}|%{DATE_EU}`,
	"DATESTAMP":         `%{DATE}[- ]%{TIME}`,
	"TZ":                `(?:[PMCE][SD]T|UTC|GMT)`,
	"DATESTAMP_RFC822":  `%{DAY} %{MONTH} %{MONTHDAY} %{YEAR} %{TIME} %{TZ}`,
	"DATESTAMP_OTHER":   `%{DAY} %{MONTH} %{MONTHDAY} %{TIME} %{TZ} %{YEAR}`,

	// Syslog Dates: Month Day HH:MM:SS
	"SYSLOGTIMESTAMP": `%{MONTH} +%{MONTHDAY} %{TIME}`,
	"PROG":            `(?:[\w._/%-]+)`,
	"SYSLOGPROG":      `%{PROG:program}(?:\[%{POSINT:pid}\])?`,
	"SYSLOGHOST":      `%{IPORHOST}`,
	"SYSLOGFACILITY":  `<%{NONNEGINT:facility}.%{NONNEGINT:priority}>`,
	"HTTPDATE":        `%{MONTHDAY}/%{MONTH}/%{YEAR}:%{TIME} %{INT}`,

	// Shortcuts
	"QS": `%{QUOTEDSTRING}`,

	// Log formats
	"SYSLOGBASE":        `%{SYSLOGTIMESTAMP:timestamp} (?:%{SYSLOGFACILITY} )?%{SYSLOGHOST:logsource} %{SYSLOGPROG}:`,
	"COMMONAPACHELOG":   `%{IPORHOST:clientip} %{USER:ident} %{USER:auth} \[%{HTTPDATE:timestamp}\] "(?:%{WORD:verb} %{NOTSPACE:request}(?: HTTP/%{NUMBER:httpversion})?|%{DATA:rawrequest})" %{NUMBER:response} (?:%{NUMBER:bytes}|-)`,
	"COMBINEDAPACHELOG": `%{COMMONAPACHELOG} %{QS:referrer} %{QS:agent}`,

	// Log Levels
	"LOGLEVEL": `([A-a]lert|ALERT|[T|t]race|TRACE|[D|d]ebug|DEBUG|[N|n]otice|NOTICE|[I|i]nfo|INFO|[W|w]arn?(?:ing)?|WARN?(?:ING)?|[E|e]rr?(?:or)?|ERR?(?:OR)?|[C|c]rit?(?:ical)?|CRIT?(?:ICAL)?|[F|f]atal|FATAL|[S|s]evere|SEVERE|EMERG(?:ENCY)?|[Ee]merg(?:ency)?)`,
}