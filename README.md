# door-api

door-api is a rest server to provide Deep Packet Inspection obtained by DOOR.


See wiki for more details e.g database scheme, endpoint and so forth.

## Development

1. Create below directory hierarchy.

```
$GOROOT/
└── src
    └── github.com
        └── westlab
            └── door-api
```

```
mkdir -p src/github.com/westlab; cd src/github.com/westlab
git clone git@github.com:westlab/door-api.git
```

2. Download dependencies

```
glide install
```

### Testing

Test are executed in circleci.

#### Run test locallly
You may need to export GOROOT and ignore vendor directory. Go does not automatically ignore vendor directory nor provides flags for ignoring this directory.
See discussion: https://github.com/golang/go/issues/11659

```
GOROOT=YOUR_GOROOT go test $(GOROOT=YOUR_GOROOT go list ./... | grep -v vendor)
```


### Code convention
Please follow go standard and check `gofmt` command.


### Workflow
Please follow github flow. https://guides.github.com/introduction/flow/


### Others
If you have a question, please make an issue or buzz me(@giwa) at slack.
