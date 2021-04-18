package bank

import "testing"

func TestCardInfo(t *testing.T) {
	tests := []struct {
		name    string
		entry   string
		want    string
		wantErr bool
	}{
		{name: "not valid 1", entry: "62198610", want: "", wantErr: true},
		{name: "not valid 2", entry: "603770asdfgbvcfg", want: "", wantErr: true},
		{name: "not valid 3", entry: "dfgsdg", want: "", wantErr: true},
		{name: "not valid 4", entry: "6219861034529008", want: "", wantErr: true},
		{name: "keshavarzi", entry: "6037701689095443", want: "keshavarzi", wantErr: false},
		{name: "saman", entry: "6219861034529007", want: "saman", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CardInfo(tt.entry)
			if (err != nil) != tt.wantErr {
				t.Errorf("CardInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CardInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
