package hosts

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func Test_Parse(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "valid data",
			args:    args{data: []byte("# test data\n127.0.0.1 localhost\n0.0.0.0 blabla.test\n0.0.0.0 example.com")},
			want:    []string{"blabla.test", "example.com"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := bytes.NewReader(tt.args.data)
			got, err := Parse(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Download(t *testing.T) {
	reader := ioutil.NopCloser(bytes.NewReader([]byte("test")))
	type args struct {
		url    string
		doFunc func(*http.Request) (*http.Response, error)
	}
	tests := []struct {
		name    string
		args    args
		want    io.Reader
		wantErr bool
	}{
		{
			name: "download default hosts file",
			args: args{
				url: DefaultURL,
				doFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{Body: reader, StatusCode: 200}, nil
				},
			},
			want:    reader,
			wantErr: false,
		},
		{
			name: "bad status code",
			args: args{
				url: DefaultURL,
				doFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{Body: reader, StatusCode: 300}, nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "do with error",
			args: args{
				url: DefaultURL,
				doFunc: func(req *http.Request) (*http.Response, error) {
					return nil, fmt.Errorf("")
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	client := &MockClient{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client.DoFunc = tt.args.doFunc
			got, err := Download(client, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Download() = %v, want %v", got, tt.want)
			}
		})
	}
}

func DoGetStatus300(req *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 300}, nil
}
func DoGetStatus400(req *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 400}, nil
}
