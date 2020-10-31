// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a configuration for running an Amazon SageMaker image as a KernelGateway
// app.
func (c *Client) CreateAppImageConfig(ctx context.Context, params *CreateAppImageConfigInput, optFns ...func(*Options)) (*CreateAppImageConfigOutput, error) {
	if params == nil {
		params = &CreateAppImageConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAppImageConfig", params, optFns, addOperationCreateAppImageConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppImageConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppImageConfigInput struct {

	// The name of the AppImageConfig. Must be unique to your account.
	//
	// This member is required.
	AppImageConfigName *string

	// The KernelGatewayImageConfig.
	KernelGatewayImageConfig *types.KernelGatewayImageConfig

	// A list of tags to apply to the AppImageConfig.
	Tags []*types.Tag
}

type CreateAppImageConfigOutput struct {

	// The Amazon Resource Name (ARN) of the AppImageConfig.
	AppImageConfigArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateAppImageConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAppImageConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAppImageConfig{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateAppImageConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAppImageConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateAppImageConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateAppImageConfig",
	}
}