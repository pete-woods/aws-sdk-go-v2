// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeUsageLimitsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the cluster for which you want to describe usage limits.
	ClusterIdentifier *string `type:"string"`

	// The feature type for which you want to describe usage limits.
	FeatureType UsageLimitFeatureType `type:"string" enum:"true"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeUsageLimits request exceed
	// the value specified in MaxRecords, AWS returns a value in the Marker field
	// of the response. You can retrieve the next set of response records by providing
	// the returned marker value in the Marker parameter and retrying the request.
	Marker *string `type:"string"`

	// The maximum number of response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// A tag key or keys for which you want to return all matching usage limit objects
	// that are associated with the specified key or keys. For example, suppose
	// that you have parameter groups that are tagged with keys called owner and
	// environment. If you specify both of these tag keys in the request, Amazon
	// Redshift returns a response with the usage limit objects have either or both
	// of these tag keys associated with them.
	TagKeys []string `locationNameList:"TagKey" type:"list"`

	// A tag value or values for which you want to return all matching usage limit
	// objects that are associated with the specified tag value or values. For example,
	// suppose that you have parameter groups that are tagged with values called
	// admin and test. If you specify both of these tag values in the request, Amazon
	// Redshift returns a response with the usage limit objects that have either
	// or both of these tag values associated with them.
	TagValues []string `locationNameList:"TagValue" type:"list"`

	// The identifier of the usage limit to describe.
	UsageLimitId *string `type:"string"`
}

// String returns the string representation
func (s DescribeUsageLimitsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeUsageLimitsOutput struct {
	_ struct{} `type:"structure"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`

	// Contains the output from the DescribeUsageLimits action.
	UsageLimits []UsageLimit `type:"list"`
}

// String returns the string representation
func (s DescribeUsageLimitsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeUsageLimits = "DescribeUsageLimits"

// DescribeUsageLimitsRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Shows usage limits on a cluster. Results are filtered based on the combination
// of input usage limit identifier, cluster identifier, and feature type parameters:
//
//    * If usage limit identifier, cluster identifier, and feature type are
//    not provided, then all usage limit objects for the current account in
//    the current region are returned.
//
//    * If usage limit identifier is provided, then the corresponding usage
//    limit object is returned.
//
//    * If cluster identifier is provided, then all usage limit objects for
//    the specified cluster are returned.
//
//    * If cluster identifier and feature type are provided, then all usage
//    limit objects for the combination of cluster and feature are returned.
//
//    // Example sending a request using DescribeUsageLimitsRequest.
//    req := client.DescribeUsageLimitsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeUsageLimits
func (c *Client) DescribeUsageLimitsRequest(input *DescribeUsageLimitsInput) DescribeUsageLimitsRequest {
	op := &aws.Operation{
		Name:       opDescribeUsageLimits,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeUsageLimitsInput{}
	}

	req := c.newRequest(op, input, &DescribeUsageLimitsOutput{})

	return DescribeUsageLimitsRequest{Request: req, Input: input, Copy: c.DescribeUsageLimitsRequest}
}

// DescribeUsageLimitsRequest is the request type for the
// DescribeUsageLimits API operation.
type DescribeUsageLimitsRequest struct {
	*aws.Request
	Input *DescribeUsageLimitsInput
	Copy  func(*DescribeUsageLimitsInput) DescribeUsageLimitsRequest
}

// Send marshals and sends the DescribeUsageLimits API request.
func (r DescribeUsageLimitsRequest) Send(ctx context.Context) (*DescribeUsageLimitsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUsageLimitsResponse{
		DescribeUsageLimitsOutput: r.Request.Data.(*DescribeUsageLimitsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeUsageLimitsRequestPaginator returns a paginator for DescribeUsageLimits.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeUsageLimitsRequest(input)
//   p := redshift.NewDescribeUsageLimitsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeUsageLimitsPaginator(req DescribeUsageLimitsRequest) DescribeUsageLimitsPaginator {
	return DescribeUsageLimitsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeUsageLimitsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeUsageLimitsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeUsageLimitsPaginator struct {
	aws.Pager
}

func (p *DescribeUsageLimitsPaginator) CurrentPage() *DescribeUsageLimitsOutput {
	return p.Pager.CurrentPage().(*DescribeUsageLimitsOutput)
}

// DescribeUsageLimitsResponse is the response type for the
// DescribeUsageLimits API operation.
type DescribeUsageLimitsResponse struct {
	*DescribeUsageLimitsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUsageLimits request.
func (r *DescribeUsageLimitsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}