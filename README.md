# Gollum

Gollum is a n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations.

There are a few basic terms used throughout Gollum:

* A "consumer" is a plugin that reads from an external source
* A "producer" is a plugin that writes to an external source
* A "stream" is a message channel between consumer(s) and producer(s)
* A "formatter" is a plugin that adds information to a message
* A "distributor" is a plugin that routes/filters messages on a given stream

Writing a custom plugin does not require you to change any additional code besides your new plugin file.

## Consumers (reading data)

* `Console` read from stdin.
* `LoopBack` Process routed (e.g. dropped) messages.
* `File` read from a file (like tail).
* `Kafka` read from a [Kafka](http://kafka.apache.org/) topic.
* `Socket` read from a socket (gollum specfic protocol).
* `Syslogd` read from a socket (syslogd protocol).

## Producers (writing data)

* `Console` write to stdin or stdout.
* `ElasticSearch` write to [elasticsearch](http://www.elasticsearch.org/) via http/bulk.
* `File` write to a file. Supports log rotation and compression.
* `Kafka` write to a [Kafka](http://kafka.apache.org/) topic.
* `Null` like /dev/null.
* `Scribe` send messages to a [Facebook scribe](https://github.com/facebookarchive/scribe) server.
* `Socket` send messages to a socket (gollum specfic protocol).

## Formatters (modifying data)

* `Forward` write the message without modifying it.
* `JSON` write the message as a JSON object. Messages can be parsed to generate fields.
* `Runlength` prepend the length of themessage. Other formatters can be nested.
* `Sequence` prepend the sequence number of the message. Other formatters can be nested.
* `Delimiter` add a delimiter string after the message.
* `Timestamp` add a timestamp before the message. Other formatters can be nested.

## Distributors (multiplexing)

* `Broadcast` send to all producers in a stream.
* `Random` send to a random roducers in a stream.
* `RoundRobin` switch the producer after each send in a round robin fashion.

## Installation

### From source

Installation from source requires the installation of the [Go toolchain](http://golang.org/).  
Gollum has [Godeps](https://github.com/tools/godep) support but this is considered optional.

```
$ go get .
$ go build
$ gollum --help
```

You can use the supplied make file to trigger cross platform builds.  
Make will produce ready to deploy .tar.gz files with the corresponding platform builds.  
This does require a cross platform golang build.  
Valid make targets (besides all and clean) are:
 * freebsd
 * linux
 * max
 * pi
 * win

## Usage

To test gollum you can make a local profiler run with a predefined configuration:

```
$ gollum -c gollum_profile.conf -ps -ll 3
```

By default this test profiles the theoretic maximum throughput of 256 Byte messages.  
You can enable different producers to test the write performance of these producers, too.

## Configuration

Configuration files are written in the YAML format and have to be loaded via command line switch.
Each plugin has a different set of configuration options which are currently described in the plugin itself, i.e. you can find examples in the GoDocs.

### Commandline

#### `-c` or `--config` [file]

Use a given configuration file.

#### `-h` or `--help`

Print this help message.

#### `-m` or `--metrics` [port]

Port to use for metric queries. Set 0 to disable.

#### `-n` or `--numcpu` [number]

Number of CPUs to use. Set 0 for all CPUs.

#### `-p` or `--pidfile` [file]

Write the process id into a given file.

#### `-pc` or `--profilecpu` [file]

Write CPU profiler results to a given file.

#### `-pm` or `--profilemem` [file]

Write heap profile results to a given file.

#### `-ps` or `--profilespeed`

Write msg/sec measurements to log.

#### `-tc` or `--testconfig` [file]

Test a given configuration file and exit.

#### `-v` or `--version`

Print version information and quit.

### Configuration file

TODO

## Use cases

TODO

### Nginx logs to Kafka

To aggregate logs by a [nginx](http://nginx.org/) web server *Gollum* can be used.
Configure a *Gollum* syslogd consumer like **

```
...
- "consumer.Syslogd":
    Enable: true
    Channel: 1024
    Stream: "profile"
    Format: "RFC3164"
    Address: 0.0.0.0:5880
...

# TODO Add Kafka producer
```
This consumer will listen to 0.0.0.0:5880 and follow the [RFC 3164](http://tools.ietf.org/html/rfc3164) for the *profile* stream and a buffer of 1024 messages.

An example *nginx.conf* can look like
```
http {
  ...
  log_format syslogd "$remote_addr - $remote_user [$time_local] '$request' $status $body_bytes_sent '$http_referer' '$http_user_agent'\n";
  access_log syslog:server=192.168.7.52:5880 syslogd;
  ...
}
```
Important note: A syslog message will be delimited by a newline. The *\n* at the end of *log_format* is important here.

References:
* [Logging to syslog @ Nginx docs](http://nginx.org/en/docs/syslog.html)
* [Module ngx_http_log_module @ Nginx docs](http://nginx.org/en/docs/http/ngx_http_log_module.html)

### Business logging (by PHP) to Kafka

TODO

### Accesslog parsing for Elasticsearch

TODO

### Log aggregation by many servers to files

TODO

## License

This project is released under the terms of the [Apache 2.0 license](http://www.apache.org/licenses/LICENSE-2.0).
