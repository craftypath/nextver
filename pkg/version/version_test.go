package version

import (
	"testing"
)

func TestNext(t *testing.T) {
	type args struct {
		current          string
		incrementPattern string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "reject non semver",
			args: args{
				current:          "0.0",
				incrementPattern: "0.x.0",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "reject non semver increment",
			args: args{
				current:          "0.0.0",
				incrementPattern: "0.x",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "increment major",
			args: args{
				current:          "0.0.0",
				incrementPattern: "x.0.0",
			},
			want:    "1.0.0",
			wantErr: false,
		},
		{
			name: "increment minor",
			args: args{
				current:          "0.0.0",
				incrementPattern: "0.x.0",
			},
			want:    "0.1.0",
			wantErr: false,
		},
		{
			name: "explicit bump",
			args: args{
				current:          "0.1.0",
				incrementPattern: "1.x.0",
			},
			want:    "1.0.0",
			wantErr: false,
		},
		{
			name: "increment patch",
			args: args{
				current:          "0.0.0",
				incrementPattern: "0.0.x",
			},
			want:    "0.0.1",
			wantErr: false,
		},
		{
			name: "too many 'x's",
			args: args{
				current:          "0.0.0",
				incrementPattern: "0.x.x",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "too many 'x's, one bordering build",
			args: args{
				current:          "0.0.0",
				incrementPattern: "0.x.x+build",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "too many 'x's, one bordering pre-release",
			args: args{
				current:          "0.0.0",
				incrementPattern: "0.x.x-alpha",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "missing x",
			args: args{
				current:          "0.0.0",
				incrementPattern: "0.0.0",
			},
			want:    "0.0.0",
			wantErr: false,
		},
		{
			name: "missing x with x in build",
			args: args{
				current:          "0.0.0",
				incrementPattern: "0.0.0+build.x",
			},
			want:    "0.0.0+build.x",
			wantErr: false,
		},
		{
			name: "starting pre-release",
			args: args{
				current:          "0.0.0",
				incrementPattern: "0.0.x-alpha.x",
			},
			want:    "0.0.1-alpha.x",
			wantErr: false,
		},
		{
			name: "pre-release and build",
			args: args{
				current:          "0.0.0-alpha+build",
				incrementPattern: "0.0.x-alpha+build",
			},
			want:    "0.0.1-alpha+build",
			wantErr: false,
		},
		{
			name: "x only in pre-release",
			args: args{
				current:          "0.0.1-alpha.0",
				incrementPattern: "0.0.1-alpha.x",
			},
			want:    "0.0.1-alpha.x",
			wantErr: false,
		},
		{
			name: "incrementing and changing pre-release",
			args: args{
				current:          "0.0.0-alpha.0",
				incrementPattern: "0.0.x-alpha.1",
			},
			want:    "0.0.1-alpha.1",
			wantErr: false,
		},
		{
			name: "ignore extra x in pre-release",
			args: args{
				current:          "0.0.0-alpha.0",
				incrementPattern: "0.0.x-alpha.x",
			},
			want:    "0.0.1-alpha.x",
			wantErr: false,
		},
		{
			name: "handle build",
			args: args{
				current:          "0.0.0+build",
				incrementPattern: "0.0.x+build",
			},
			want:    "0.0.1+build",
			wantErr: false,
		},
		{
			name: "ignore extra x in build",
			args: args{
				current:          "0.0.0+build.x",
				incrementPattern: "0.0.x+build.x",
			},
			want:    "0.0.1+build.x",
			wantErr: false,
		},
		{
			name: "ignore x in build",
			args: args{
				current:          "0.0.0+build.x",
				incrementPattern: "0.0.0+build.x",
			},
			want:    "0.0.0+build.x",
			wantErr: false,
		},
		{
			name: "major neither a number nor x",
			args: args{
				current:          "0.0.0",
				incrementPattern: "y.0.0",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "minor neither a number nor x",
			args: args{
				current:          "0.0.0",
				incrementPattern: "0.y.0",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "patch neither a number nor x",
			args: args{
				current:          "0.0.0",
				incrementPattern: "0.0.y",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "support v prefix",
			args: args{
				current:          "v1.0.0",
				incrementPattern: "v1.x.0",
			},
			want:    "v1.1.0",
			wantErr: false,
		},
		{
			name: "add v prefix",
			args: args{
				current:          "1.0.0",
				incrementPattern: "v1.x.0",
			},
			want:    "v1.1.0",
			wantErr: false,
		},
		{
			name: "remove v prefix",
			args: args{
				current:          "v1.0.0",
				incrementPattern: "1.x.0",
			},
			want:    "1.1.0",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Next(tt.args.current, tt.args.incrementPattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("Next() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Next() got = %v, want %v", got, tt.want)
			}
		})
	}
}
