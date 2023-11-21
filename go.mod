module github.com/Conor-Fleming/CraigsFinder

go 1.21.4

require (
	github.com/ecnepsnai/craigslist v0.0.6
	gopkg.in/yaml.v2 v2.4.0
)

require github.com/google/uuid v1.3.0 // indirect

replace github.com/ecnepsnai/craigslist => github.com/Conor-Fleming/cl v0.1.0
