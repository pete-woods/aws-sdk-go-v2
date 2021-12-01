// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Explainability exports created using the
// CreateExplainabilityExport operation. This operation returns a summary for each
// Explainability export. You can filter the list using an array of Filter objects.
// To retrieve the complete set of properties for a particular Explainability
// export, use the ARN with the DescribeExplainability operation.
func (c *Client) ListExplainabilityExports(ctx context.Context, params *ListExplainabilityExportsInput, optFns ...func(*Options)) (*ListExplainabilityExportsOutput, error) {
	if params == nil {
		params = &ListExplainabilityExportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExplainabilityExports", params, optFns, c.addOperationListExplainabilityExportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExplainabilityExportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExplainabilityExportsInput struct {

	// An array of filters. For each filter, provide a condition and a match statement.
	// The condition is either IS or IS_NOT, which specifies whether to include or
	// exclude resources that match the statement from the list. The match statement
	// consists of a key and a value. Filter properties
	//
	// * Condition - The condition to
	// apply. Valid values are IS and IS_NOT.
	//
	// * Key - The name of the parameter to
	// filter on. Valid values are PredictorArn and Status.
	//
	// * Value - The value to
	// match.
	Filters []types.Filter

	// The number of items to return in the response.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListExplainabilityExportsOutput struct {

	// An array of objects that summarize the properties of each Explainability export.
	ExplainabilityExports []types.ExplainabilityExportSummary

	// Returns this token if the response is truncated. To retrieve the next set of
	// results, use the token in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExplainabilityExportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListExplainabilityExports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListExplainabilityExports{}, middleware.After)
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
	if err = addOpListExplainabilityExportsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExplainabilityExports(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListExplainabilityExports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "ListExplainabilityExports",
	}
}