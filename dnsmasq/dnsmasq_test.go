package dnsmasq

import (
	"bytes"
	"testing"

	"git.gautier.ovh/hosts2dnsmasq/tests"
)

func Test_writeDNSMasqConfig(t *testing.T) {
	type args struct {
		hosts []string
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
		wantErr    bool
	}{
		{
			name:       "valid data",
			args:       args{hosts: tests.Hosts},
			wantErr:    false,
			wantWriter: tests.Config,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if err := WriteConfig(writer, DefaultConfigFormat, tt.args.hosts); (err != nil) != tt.wantErr {
				t.Errorf("writeDNSMasqConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("writeDNSMasqConfig() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
