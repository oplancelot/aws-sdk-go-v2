// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotjobsdataplane

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// IoTJobsDataPlane provides the API operation methods for making requests to
// AWS IoT Jobs Data Plane. See this package's package overview docs
// for details on the service.
//
// IoTJobsDataPlane methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type IoTJobsDataPlane struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*IoTJobsDataPlane)

// Used for custom request initialization logic
var initRequest func(*IoTJobsDataPlane, *aws.Request)

// Service information constants
const (
	ServiceName = "data.jobs.iot" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName     // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the IoTJobsDataPlane client with a config.
// If additional configuration is needed for the client instance use the
// optional aws.Config parameter to add your extra config.
//
// Example:
//     // Create a IoTJobsDataPlane client from just a config.
//     svc := iotjobsdataplane.New(myConfig)
//
//     // Create a IoTJobsDataPlane client with additional configuration
//     svc := iotjobsdataplane.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *IoTJobsDataPlane {
	var signingName string
	signingName = "iot-jobs-data"
	signingRegion := config.Region

	svc := &IoTJobsDataPlane{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2017-09-29",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a IoTJobsDataPlane operation and runs any
// custom request initialization.
func (c *IoTJobsDataPlane) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}