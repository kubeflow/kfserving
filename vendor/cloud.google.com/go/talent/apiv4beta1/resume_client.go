// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gapic-generator. DO NOT EDIT.

package talent

import (
	"context"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ResumeCallOptions contains the retry settings for each method of ResumeClient.
type ResumeCallOptions struct {
	ParseResume []gax.CallOption
}

func defaultResumeClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("jobs.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultResumeCallOptions() *ResumeCallOptions {
	retry := map[[2]string][]gax.CallOption{}
	return &ResumeCallOptions{
		ParseResume: retry[[2]string{"default", "non_idempotent"}],
	}
}

// ResumeClient is a client for interacting with Cloud Talent Solution API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ResumeClient struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	resumeClient talentpb.ResumeServiceClient

	// The call options for this service.
	CallOptions *ResumeCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewResumeClient creates a new resume service client.
//
// A service that handles resume parsing.
func NewResumeClient(ctx context.Context, opts ...option.ClientOption) (*ResumeClient, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultResumeClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &ResumeClient{
		conn:        conn,
		CallOptions: defaultResumeCallOptions(),

		resumeClient: talentpb.NewResumeServiceClient(conn),
	}
	c.setGoogleClientInfo()
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *ResumeClient) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ResumeClient) Close() error {
	return c.conn.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ResumeClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ParseResume parses a resume into a [Profile][google.cloud.talent.v4beta1.Profile]. The
// API attempts to fill out the following profile fields if present within the
// resume:
//
//   personNames
//
//   addresses
//
//   emailAddress
//
//   phoneNumbers
//
//   personalUris
//
//   employmentRecords
//
//   educationRecords
//
//   skills
//
// Note that some attributes in these fields may not be populated if they're
// not present within the resume or unrecognizable by the resume parser.
//
// This API does not save the resume or profile. To create a profile from this
// resume, clients need to call the CreateProfile method again with the
// profile returned.
//
// The following list of formats are supported:
//
//   PDF
//
//   TXT
//
//   DOC
//
//   RTF
//
//   DOCX
//
//   PNG (only when [ParseResumeRequest.enable_ocr][] is set to true,
//   otherwise an error is thrown)
func (c *ResumeClient) ParseResume(ctx context.Context, req *talentpb.ParseResumeRequest, opts ...gax.CallOption) (*talentpb.ParseResumeResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ParseResume[0:len(c.CallOptions.ParseResume):len(c.CallOptions.ParseResume)], opts...)
	var resp *talentpb.ParseResumeResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.resumeClient.ParseResume(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
