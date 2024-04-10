module github.com/fhke/infrastructure-abstraction/dsl

go 1.22.1

require (
	github.com/fhke/infrastructure-abstraction/sdk v0.0.0-00010101000000-000000000000
	github.com/fhke/infrastructure-abstraction/util v0.0.0-00010101000000-000000000000
	github.com/fhke/tpp v0.0.0-20240410224959-e4dbc25cb990
	github.com/samber/lo v1.39.0
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/ProtonMail/go-crypto v1.0.0 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/chainguard-dev/git-urls v1.0.2 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.5.0 // indirect
	github.com/go-git/go-git/v5 v5.12.0 // indirect
	github.com/go-resty/resty/v2 v2.12.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/hashicorp/hcl/v2 v2.20.1 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/kubescape/go-git-url v0.0.30 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/sergi/go-diff v1.3.2-0.20230802210424-5b0b94c5c0d3 // indirect
	github.com/skeema/knownhosts v1.2.2 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	github.com/zclconf/go-cty v1.13.0 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.20.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	k8s.io/apimachinery v0.29.3 // indirect
)

replace (
	github.com/fhke/infrastructure-abstraction/sdk => ../sdk
	github.com/fhke/infrastructure-abstraction/util => ../util
)
