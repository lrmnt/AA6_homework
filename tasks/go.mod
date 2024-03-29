module github.com/lrmnt/AA6_homework/tasks

go 1.22

toolchain go1.22.0

replace github.com/lrmnt/AA6_homework/lib => ../lib // do not load from github

require (
	entgo.io/ent v0.13.1
	github.com/go-chi/chi/v5 v5.0.12
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
	github.com/lrmnt/AA6_homework/lib v0.0.0-00010101000000-000000000000
	github.com/segmentio/kafka-go v0.4.47
	go.uber.org/zap v1.27.0
	golang.org/x/sync v0.6.0
)

require (
	ariga.io/atlas v0.19.1 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/hashicorp/hcl/v2 v2.19.1 // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/zclconf/go-cty v1.14.2 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/mod v0.15.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
