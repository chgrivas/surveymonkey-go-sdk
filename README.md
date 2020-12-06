# surveymonkey-go-sdk
The SurveyMonkey SDK for Go programming language.

## Install

```bash
go get -u github.com/chgrivas/surveymonkey-go-sdk
```

## Use

```bash
import "github.com/chgrivas/surveymonkey-go-sdk"

// Replace ACCESS_TOKEN with your real key
client := surveymonkey.NewClient("ACCESS_TOKEN")
```
## Test

Run integration tests with real API Key.
```bash
export SURVEYMONKEY_ACCESS_TOKEN=
go test -v -tags=integration
```