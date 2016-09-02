# hecate
Hygieia CLI

This project aims to provide a CLI, utilizing the Hygieia API for common Hygieia tasks.

## Usage:
After [##Building](Building) you should have a binary named `hecate`
Then, set your `hygieia_base` environment variable to point to your hygieia instance.
Then:
```
hecate '{"executionId": "abacadab","jobUrl":"http://foo.bar","jobName": "foo bar deploy","appName": "coolapp","envName": "QA","artifactName": "md5magic","artifactVersion":"sha1","instanceUrl": "http://foo.bar.com","deployStatus": "success","startTime": 1469038554}'
```

Currently, the JSON payload must be constructed and passed manually. This may change to accomodate output from Jenkins (XML) or simple parameters, as needed.

## Building:

```
git clone https://github.com/charleswisniewski/hecate $GOHOME/src/github.com/charleswisniewski/hecate

go install github.com/charleswisniewski/hecate
```
If this completed successfully you should have a binary located in:
$GOHOME/bin/hecate

## Testing
Currently we are only testing that the Hygieia API you are expecting to hit conforms to the swagger spec described below. To confirm this, we are using [Dredd](https://github.com/apiaryio/dredd)

```
npm install -g dredd

dredd hygieia.yaml http://yourhygieiainstance/api
```

You should see no failures, if you do, then your instance of Hygieia does not match the swagger spec and the binary will not work

## Hygieia API
The API of Hygieia is to be defined in [hygieia.yaml](hygieia.yaml)
Currently only the deploy endpoint is described.



## CLI v2
Rough outline for v2 CLI

1. Generate models from swagger - swagger should be passing against version X of hygieia

install go-swagger(OS X, using homebrew):
```
brew tap go-swagger/go-swagger
brew install go-swagger
swagger generate client -f hygieia.yaml
```
This should generate two folders client and models(not included in this repo, yet)
The commands/deploy.go file included is an example of how we can programatically invoke an API request based on the models defined in swagger.

2. CLI generated from those models using cobra (@TODO)

  a) persistent flags = base url
  b) read in swagger models
  c) convert swagger models via loop to cobra syntax

An conceptual implementation is included in commands/deploy.go (commented out)
