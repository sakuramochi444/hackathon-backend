package usecase

import (
	"kaizen/model"
	"strings"
	"testing"
)

func Test_validateUser(t *testing.T) {
	type args struct {
		req model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "名前が空", args: args{req: model.User{Name: "", Age: 30}}, wantErr: true},
		{name: "名前の長さ制限1", args: args{req: model.User{Name: strings.Repeat("a", 50), Age: 30}}, wantErr: false},
		{name: "名前の長さ制限1", args: args{req: model.User{Name: strings.Repeat("a", 51), Age: 30}}, wantErr: true},
		{name: "年齢が範囲外1", args: args{req: model.User{Name: "Tester", Age: 19}}, wantErr: true},
		{name: "年齢が範囲外2", args: args{req: model.User{Name: "Tester", Age: 81}}, wantErr: true},
		{name: "年齢が範囲外3", args: args{req: model.User{Name: "Tester", Age: 20}}, wantErr: false},
		{name: "年齢が範囲外4", args: args{req: model.User{Name: "Tester", Age: 80}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateUser(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("validateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
