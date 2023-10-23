module github.com/ahmetson/today-do/cli

go 1.21rc2

replace github.com/ahmetson/datatype-lib => D:/sds/datatype-lib

replace github.com/ahmetson/handler-lib => D:/sds/handler-lib

replace github.com/ahmetson/client-lib => D:/sds/client-lib

replace github.com/ahmetson/config-lib => D:/sds/config-lib

replace github.com/ahmetson/dev-lib => D:/sds/dev-lib

replace github.com/ahmetson/os-lib => D:/sds/os-lib

replace github.com/ahmetson/log-lib => D:/sds/log-lib

replace github.com/ahmetson/service-lib => D:/sds/service-lib

require (
	github.com/ahmetson/client-lib v0.0.0-20230908110757-5f62078bd7bd
	github.com/ahmetson/config-lib v0.0.0-00010101000000-000000000000
	github.com/ahmetson/datatype-lib v0.0.0-20230927201942-0cc58292a7a3
	github.com/ahmetson/os-lib v0.0.0-20230908110839-83535270d872
	github.com/pebbe/zmq4 v1.2.10
)

require (
	github.com/ahmetson/handler-lib v0.0.0-20230908055435-ceab4155ee16 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/phayes/freeport v0.0.0-20220201140144-74d24b5ae9f5 // indirect
)
