// Copyright (c) 2021, SailPoint Technologies, Inc. All rights reserved.

package transform

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sailpoint-oss/sailpoint-cli/internal/mocks"
)

func TestNewCreateCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mocks.NewMockClient(ctrl)

	client.EXPECT().
		Post(gomock.Any(), gomock.Any(), "application/json", gomock.Any()).
		Return(&http.Response{StatusCode: http.StatusCreated, Body: io.NopCloser(bytes.NewReader([]byte("{}")))}, nil).
		Times(1)

	client.EXPECT().
		Get(gomock.Any(), gomock.Any()).
		Return(&http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader([]byte("[]")))}, nil).
		Times(1)

	cmd := newCreateCmd(client)

	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.Flags().Set("file", "test_data/CreateTest.json")
	cmd.PersistentFlags().StringP("transforms-endpoint", "e", transformsEndpoint, "Override transforms endpoint")

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("TestNewCreateCmd: Unable to execute the command successfully: %v", err)
	}
}
