# azure-translate
Azure translation SDK

# Example
```
key := "YOUR-SUBSCRIPTION-KEY"
location := "global"
client := translate.NewClient(key, location)

result, err := client.To(&translate.Request{
  Data: []translate.Body{{"Hello, world!"}},
  ToLanguage:   []string{"it"},
})
```