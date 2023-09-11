# TheNextLeg Go Client

TheNextLeg Go Client is a Go library for interacting with TheNextLeg API. It allows you to access various features provided by TheNextLeg, such as image processing, text analysis, and more.

## Installation

You can install TheNextLeg Go Client using `go get`:

```shell
go get github.com/nikoole/thenextleg-go
```

## Usage

To use the client, you'll need to import it and create a new client instance with your API authentication token:

```go
import "github.com/nikoole/thenextleg-go"

client := thenextleg.NewClient("your-auth-token")
```

### WebhookOverride Parameter Notice

Please take note that the `WebhookOverride` parameter is optional in any *Request. When specified, you will receive callbacks regarding the generation progress at the provided URL.

These callbacks will contain data described in the *Progress structures.

### MessageId in *Response

The `MessageId` in the *Response allows you to request the progress of your generation by invoking the corresponding *Progress method.

### Imagine

You can use the `Imagine` method to perform image processing:

```go
request := thenextleg.ImagineRequest{
	Msg: "beautiful day on the edge of the forest",
	Ref: "any data you want to receive back once your generation complete",
	WebhookOverride: "https://example.com/callback/here/once/job/done",
	IgnorePrefilter: false, // your request will be premoderated
}

response, err := client.Imagine(request)
if err != nil {
    // Handle error
}

// Use the response data to receive current progress after a while
progress, err := client.ImagineProgress(response.MessageId)
```

### Describe

You can use the `Describe` method to describe an image that you upload :

```go
request := thenextleg.DescribeRequest{
	Url: "https://example.com/link-to-image-to-be-described",
	// other parameters have the same significance
}

response, err := client.Describe(request)
if err != nil {
    // Handle error
}
// Use the response data
```

### Blend

You can use the `Blend` method to perform blending:

```go
request := thenextleg.BlendRequest{
	Urls: []string{<URLs to images that will be blended together. Up to 5.>},
	// other parameters have the same significance
}

response, err := client.Blend(request)
if err != nil {
    // Handle error
}
// Use the response data
```

### FaceSwap

You can use the `FaceSwap` method to perform face swapping:

```go
request := thenextleg.FaceSwapRequest{
	SourceImg: "https://...",
	TargetImg: "https://...",
}

imageReader, err := client.FaceSwap(request)
if err != nil {
    // Handle error
}
// Use the imageReader to access the resulting image
```

### GetImage

Retriving images from the Midjourney CDN from your server will result in a 403. This endpoint allows you to retrieve the image data safely.

```go
request := thenextleg.GetImageRequest{
	ImgUrl: "https://...",
}

imageReader, err := client.GetImage(request)
if err != nil {
    // Handle error
}
// Use the imageReader to access the retrieved image
```

### IsThisNaughty

You can use the `IsThisNaughty` method to check if text is potentially inappropriate. If you use inappropriate words, you risk getting banned by Midjourney!

```go
request := thenextleg.IsThisNaughtyRequest{
	Msg: "Is this an acceptable image generation request"
}
response, err := client.IsThisNaughty(request)
```

### Info

Retrieve the fast hours and other information on your Midjourney account.


```go
request := thenextleg.InfoRequest{
	Ref: "some data",
	WebhookOverride: "https://example.com/callback/here/once/job/done"
}

response, err := client.Info(request) // request optional!
if err != nil {
    // Handle error
}
// Use the response data
```

## Contributing

Contributions are welcome! If you would like to contribute to this project, please follow the standard Go development practices.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
