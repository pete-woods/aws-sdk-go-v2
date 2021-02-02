// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of the import jobs.
func (c *Client) ListImportJobs(ctx context.Context, params *ListImportJobsInput, optFns ...func(*Options)) (*ListImportJobsOutput, error) {
	if params == nil {
		params = &ListImportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImportJobs", params, optFns, addOperationListImportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to list all of the import jobs for a data destination
// within the specified maximum number of import jobs.
type ListImportJobsInput struct {

	// The destination of the import job, which can be used to list import jobs that
	// have a certain ImportDestinationType.
	ImportDestinationType types.ImportDestinationType

	// A string token indicating that there might be additional import jobs available
	// to be listed. Copy this token to a subsequent call to ListImportJobs with the
	// same parameters to retrieve the next page of import jobs.
	NextToken *string

	// Maximum number of import jobs to return at once. Use this parameter to paginate
	// results. If additional import jobs exist beyond the specified limit, the
	// NextToken element is sent in the response. Use the NextToken value in subsequent
	// requests to retrieve additional addresses.
	PageSize *int32
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type ListImportJobsOutput struct {

	// A list of the import job summaries.
	ImportJobs []types.ImportJobSummary

	// A string token indicating that there might be additional import jobs available
	// to be listed. Copy this token to a subsequent call to ListImportJobs with the
	// same parameters to retrieve the next page of import jobs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListImportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImportJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImportJobs(options.Region), middleware.Before); err != nil {
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

// ListImportJobsAPIClient is a client that implements the ListImportJobs
// operation.
type ListImportJobsAPIClient interface {
	ListImportJobs(context.Context, *ListImportJobsInput, ...func(*Options)) (*ListImportJobsOutput, error)
}

var _ ListImportJobsAPIClient = (*Client)(nil)

// ListImportJobsPaginatorOptions is the paginator options for ListImportJobs
type ListImportJobsPaginatorOptions struct {
	// Maximum number of import jobs to return at once. Use this parameter to paginate
	// results. If additional import jobs exist beyond the specified limit, the
	// NextToken element is sent in the response. Use the NextToken value in subsequent
	// requests to retrieve additional addresses.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImportJobsPaginator is a paginator for ListImportJobs
type ListImportJobsPaginator struct {
	options   ListImportJobsPaginatorOptions
	client    ListImportJobsAPIClient
	params    *ListImportJobsInput
	nextToken *string
	firstPage bool
}

// NewListImportJobsPaginator returns a new ListImportJobsPaginator
func NewListImportJobsPaginator(client ListImportJobsAPIClient, params *ListImportJobsInput, optFns ...func(*ListImportJobsPaginatorOptions)) *ListImportJobsPaginator {
	options := ListImportJobsPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListImportJobsInput{}
	}

	return &ListImportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImportJobsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListImportJobs page.
func (p *ListImportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImportJobsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListImportJobs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListImportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListImportJobs",
	}
}