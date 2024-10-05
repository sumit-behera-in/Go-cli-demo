package main

import (
	"net/http"
	"reflect"
	"testing"
)

func TestDefaultClient_Get(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name     string
		c        *DefaultClient
		args     args
		wantResp *http.Response
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		c := &DefaultClient{}
		gotResp, err := c.Get(tt.args.url)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. DefaultClient.Get(%v) error = %v, wantErr %v", tt.name, tt.args.url, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(gotResp, tt.wantResp) {
			t.Errorf("%q. DefaultClient.Get(%v) = %v, want %v", tt.name, tt.args.url, gotResp, tt.wantResp)
		}
	}
}

func Test_check(t *testing.T) {
	type args struct {
		config  SiteConfig
		client  HttpClient
		results chan<- Result
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		check(tt.args.config, tt.args.client, tt.args.results)
	}
}

func Test_scheduleCheck(t *testing.T) {
	type args struct {
		config  SiteConfig
		client  HttpClient
		results chan<- Result
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		scheduleCheck(tt.args.config, tt.args.client, tt.args.results)
	}
}
