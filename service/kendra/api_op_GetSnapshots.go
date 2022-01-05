// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves search metrics data. The data provides a snapshot of how your users
// interact with your search application and how effective the application is.
func (c *Client) GetSnapshots(ctx context.Context, params *GetSnapshotsInput, optFns ...func(*Options)) (*GetSnapshotsOutput, error) {
	if params == nil {
		params = &GetSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSnapshots", params, optFns, c.addOperationGetSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSnapshotsInput struct {

	// The identifier of the index to get search metrics data.
	//
	// This member is required.
	IndexId *string

	// The time interval or time window to get search metrics data. The time interval
	// uses the time zone of your index. You can view data in the following time
	// windows:
	//
	// * THIS_WEEK: The current week, starting on the Sunday and ending on
	// the day before the current date.
	//
	// * ONE_WEEK_AGO: The previous week, starting on
	// the Sunday and ending on the following Saturday.
	//
	// * TWO_WEEKS_AGO: The week
	// before the previous week, starting on the Sunday and ending on the following
	// Saturday.
	//
	// * THIS_MONTH: The current month, starting on the first day of the
	// month and ending on the day before the current date.
	//
	// * ONE_MONTH_AGO: The
	// previous month, starting on the first day of the month and ending on the last
	// day of the month.
	//
	// * TWO_MONTHS_AGO: The month before the previous month,
	// starting on the first day of the month and ending on last day of the month.
	//
	// This member is required.
	Interval types.Interval

	// The metric you want to retrieve. You can specify only one metric per call. For
	// more information about the metrics you can view, see Gaining insights with
	// search analytics
	// (https://docs.aws.amazon.com/kendra/latest/dg/search-analytics.html).
	//
	// This member is required.
	MetricType types.MetricType

	// The maximum number of returned data for the metric.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Kendra returns a pagination token in the response. You can use
	// this pagination token to retrieve the next set of search metrics data.
	NextToken *string

	noSmithyDocumentSerde
}

type GetSnapshotsOutput struct {

	// If the response is truncated, Amazon Kendra returns this token, which you can
	// use in a later request to retrieve the next set of search metrics data.
	NextToken *string

	// The date-time for the beginning and end of the time window for the search
	// metrics data.
	SnapShotTimeFilter *types.TimeRange

	// The search metrics data. The data returned depends on the metric type you
	// requested.
	SnapshotsData [][]string

	// The column headers for the search metrics data.
	SnapshotsDataHeader []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetSnapshotsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSnapshots(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "GetSnapshots",
	}
}