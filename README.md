# Tatsu API Go

Official Go wrapper for Tatsu's API

## Install

```shell
go get github.com/tatsuworks/tatsu-api-go
```

## Example

```go
import tatsuapi "github.com/tatsuworks/tatsu-api-go"

client := tatsuapi.New("YOUR_API_KEY")

user, err := client.getUserProfile("273261090404696074")
if err == nil {
    fmt.Println("Rep: " + user.Reputation)
}
```

## License

```text
Copyright ©2021 Tatsu Works.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
