<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# rules

```go
import "github.com/brittonhayes/pillager/pkg/rules"
```

Package rules enables the parsing of Gitleaks rulesets

## Index

- [Constants](<#constants>)
- [Variables](<#variables>)
- [type Loader](<#type-loader>)
  - [func NewLoader(opts ...LoaderOption) *Loader](<#func-newloader>)
  - [func (l *Loader) Load() config.Config](<#func-loader-load>)
  - [func (l *Loader) WithStrict() LoaderOption](<#func-loader-withstrict>)
- [type LoaderOption](<#type-loaderoption>)
  - [func FromFile(file string) LoaderOption](<#func-fromfile>)


## Constants

```go
const (
    ErrReadConfig = "Failed to read config"
)
```

## Variables

```go
var (
    //go:embed rules_simple.toml
    RulesDefault string

    //go:embed rules_strict.toml
    RulesStrict string
)
```

## type [Loader](<https://github.com/brittonhayes/pillager/blob/main/pkg/rules/rules.go#L25-L27>)

```go
type Loader struct {
    // contains filtered or unexported fields
}
```

### func [NewLoader](<https://github.com/brittonhayes/pillager/blob/main/pkg/rules/rules.go#L33>)

```go
func NewLoader(opts ...LoaderOption) *Loader
```

NewLoader creates a configuration loader\.

### func \(\*Loader\) [Load](<https://github.com/brittonhayes/pillager/blob/main/pkg/rules/rules.go#L58>)

```go
func (l *Loader) Load() config.Config
```

Load parses the gitleaks configuration\.

### func \(\*Loader\) [WithStrict](<https://github.com/brittonhayes/pillager/blob/main/pkg/rules/rules.go#L48>)

```go
func (l *Loader) WithStrict() LoaderOption
```

WithStrict enables more strict pillager scanning\.

## type [LoaderOption](<https://github.com/brittonhayes/pillager/blob/main/pkg/rules/rules.go#L29>)

```go
type LoaderOption func(*Loader)
```

### func [FromFile](<https://github.com/brittonhayes/pillager/blob/main/pkg/rules/rules.go#L69>)

```go
func FromFile(file string) LoaderOption
```

FromFile decodes the configuration from a local file\.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)