module logger-test

go 1.19

require github.com/krls08/go_logger_ex v0.0.0-20221107161405-ba9b2300478b
require github.com/krls08/example-go-submodule/logger v0.0.0

replace github.com/krls08/example-go-submodule/logger => ../logger
