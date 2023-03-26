package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestPassword_PasswordInit(t *testing.T) {
	type fields struct {
		length   int
		letters  string
		LETTERS  string
		numbers  string
		symbols  string
		password string
	}
	type args struct {
		input int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test 1 - pass",
			fields:  fields{},
			args:    args{20},
			wantErr: false,
		},
		{
			name:    "test 2 - too high",
			fields:  fields{},
			args:    args{65},
			wantErr: true,
		},
		{
			name:    "test 3 - too low",
			fields:  fields{},
			args:    args{2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pw := &Password{
				length:   tt.fields.length,
				letters:  tt.fields.letters,
				LETTERS:  tt.fields.LETTERS,
				numbers:  tt.fields.numbers,
				symbols:  tt.fields.symbols,
				password: tt.fields.password,
			}
			if _, err := pw.PasswordInit(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("Password.PasswordInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPassword_GeneratePassword(t *testing.T) {
	type fields struct {
		length   int
		letters  string
		LETTERS  string
		numbers  string
		symbols  string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "passed test",
			fields: fields{
				numbers: "1",
				symbols: "!",
				letters: "a",
				LETTERS: "Z",
			},
			wantErr: false,
		},
		{
			name: "no values",
			fields: fields{
				numbers: "",
				symbols: "",
				letters: "",
				LETTERS: "",
			},
			wantErr: true,
		},
		{
			name: "incorrect values",
			fields: fields{
				numbers: "a",
				symbols: "b",
				letters: "c",
				LETTERS: "d",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pw := &Password{
				length:   tt.fields.length,
				letters:  tt.fields.letters,
				LETTERS:  tt.fields.LETTERS,
				numbers:  tt.fields.numbers,
				symbols:  tt.fields.symbols,
				password: tt.fields.password,
			}
			value := pw.GeneratePassword()
			if (!strings.Contains(value, "1") || !strings.Contains(value, "!") || !strings.Contains(value, "a") ||
				!strings.Contains(value, "Z")) && !tt.wantErr {
				t.Errorf("Password.GeneratePassword() error = Invalid Values, wantErr %v", tt.wantErr)
			}
		})
	}
}

func TestPassword_GenerateLetters(t *testing.T) {
	type fields struct {
		length   int
		letters  string
		LETTERS  string
		numbers  string
		symbols  string
		password string
	}
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test 1 - pass",
			fields:  fields{},
			args:    args{20},
			wantErr: false,
		},
		{
			name:    "test 2 - fail",
			fields:  fields{},
			args:    args{19},
			wantErr: true,
		},
		{
			name:    "test 3 - pre-existing values",
			fields:  fields{letters: "a"},
			args:    args{20},
			wantErr: true,
		},
		{
			name:    "test 4 - contains number",
			fields:  fields{letters: "4"},
			args:    args{19},
			wantErr: true,
		},
		{
			name:    "test 5 - contains upper case",
			fields:  fields{letters: "A"},
			args:    args{19},
			wantErr: true,
		},
		{
			name:    "test 6 - contains symbol",
			fields:  fields{letters: "!"},
			args:    args{19},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pw := &Password{
				length:   tt.fields.length,
				letters:  tt.fields.letters,
				LETTERS:  tt.fields.LETTERS,
				numbers:  tt.fields.numbers,
				symbols:  tt.fields.symbols,
				password: tt.fields.password,
			}
			pw.GenerateLetters(tt.args.length)
			if len(pw.letters) != 20 && !tt.wantErr {
				t.Errorf("Password.GenerateLetters() error = incorrect length of string returned, wantErr %v", tt.wantErr)
			}
			for _, char := range pw.letters {
				if (char > 122 || char < 97) && !tt.wantErr {
					t.Errorf("Password.GenerateLetters() error = contains invalid chars, wantErr %v", tt.wantErr)
				}
			}
		})
	}
}

func TestPassword_GenerateCapitals(t *testing.T) {
	type fields struct {
		length   int
		letters  string
		LETTERS  string
		numbers  string
		symbols  string
		password string
	}
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    int
	}{
		{
			name:    "test 1 - pass",
			fields:  fields{},
			args:    args{20},
			wantErr: false,
			want:    20,
		},
		{
			name:    "test 2 - wrong length returned",
			fields:  fields{},
			args:    args{19},
			wantErr: true,
			want:    20,
		},
		{
			name:    "test 3 - pre-existing values",
			fields:  fields{LETTERS: "A"},
			args:    args{20},
			wantErr: true,
			want:    20,
		},
		{
			name:    "test 4 - contains number",
			fields:  fields{LETTERS: "4"},
			args:    args{19},
			wantErr: true,
			want:    20,
		},
		{
			name:    "test 5 - contains lower case",
			fields:  fields{LETTERS: "a"},
			args:    args{19},
			wantErr: true,
			want:    20,
		},
		{
			name:    "test 6 - contains symbol",
			fields:  fields{LETTERS: "!"},
			args:    args{19},
			wantErr: true,
			want:    20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pw := &Password{
				length:   tt.fields.length,
				letters:  tt.fields.letters,
				LETTERS:  tt.fields.LETTERS,
				numbers:  tt.fields.numbers,
				symbols:  tt.fields.symbols,
				password: tt.fields.password,
			}
			if got := pw.GenerateCapitals(tt.args.length); got != tt.want && !tt.wantErr {
				t.Errorf("Password.GenerateCapitals() = %v, want %v", got, tt.want)
			}
			for _, char := range pw.LETTERS {
				if (char > 90 || char < 65) && !tt.wantErr {
					t.Errorf("Password.GenerateCapitals() error = contains invalid chars, wantErr %v", tt.wantErr)
				}
			}
		})
	}
}

func TestPassword_GenerateNumbers(t *testing.T) {
	type fields struct {
		length   int
		letters  string
		LETTERS  string
		numbers  string
		symbols  string
		password string
	}
	type args struct {
		length int
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantLenErr  bool
		wantConvErr bool
		want        int
	}{
		{
			name:        "test 1 - pass",
			fields:      fields{},
			args:        args{10},
			wantLenErr:  false,
			wantConvErr: false,
			want:        10,
		},
		{
			name:        "test 2 - wrong length returned",
			fields:      fields{},
			args:        args{9},
			wantLenErr:  true,
			wantConvErr: false,
			want:        10,
		},
		{
			name:        "test 3 - pre-existing values",
			fields:      fields{numbers: "1"},
			args:        args{10},
			wantLenErr:  true,
			wantConvErr: false,
			want:        10,
		},
		{
			name:        "test 4 - contains capitals",
			fields:      fields{numbers: "A"},
			args:        args{9},
			wantLenErr:  false,
			wantConvErr: true,
			want:        10,
		},
		{
			name:        "test 5 - contains lower case",
			fields:      fields{numbers: "a"},
			args:        args{9},
			wantLenErr:  false,
			wantConvErr: true,
			want:        10,
		},
		{
			name:        "test 6 - contains symbol",
			fields:      fields{numbers: "!"},
			args:        args{9},
			wantLenErr:  false,
			wantConvErr: true,
			want:        10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pw := &Password{
				length:   tt.fields.length,
				letters:  tt.fields.letters,
				LETTERS:  tt.fields.LETTERS,
				numbers:  tt.fields.numbers,
				symbols:  tt.fields.symbols,
				password: tt.fields.password,
			}
			if got := pw.GenerateNumbers(tt.args.length); got != tt.want && !tt.wantLenErr {
				t.Errorf("Password.GenerateNumbers() = %v, want %v", got, tt.want)
			}
			if _, err := strconv.Atoi(pw.numbers); (err != nil) != tt.wantConvErr {
				t.Errorf("Password.GenerateNumbers() error = %v contains invalid chars, wantConvErr %v", err, tt.wantConvErr)
			}
		})
	}
}

// test by not contains, maybe convert to sets?
func TestPassword_GenerateSymbols(t *testing.T) {
	type fields struct {
		length   int
		letters  string
		LETTERS  string
		numbers  string
		symbols  string
		password string
	}
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "test 1 - pass",
			fields:  fields{},
			args:    args{10},
			wantErr: false,
			want:    10,
		},
		{
			name:    "test 2 - pre-existing",
			fields:  fields{symbols: "!"},
			args:    args{10},
			wantErr: true,
			want:    10,
		},
		{
			name:    "test 3 - contains letter",
			fields:  fields{symbols: "a"},
			args:    args{9},
			wantErr: true,
			want:    10,
		},
		{
			name:    "test 3 - contains number",
			fields:  fields{symbols: "2"},
			args:    args{9},
			wantErr: true,
			want:    10,
		},
		{
			name:    "test 3 - contains invalid symbol",
			fields:  fields{symbols: "`"},
			args:    args{9},
			wantErr: true,
			want:    10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pw := &Password{
				length:   tt.fields.length,
				letters:  tt.fields.letters,
				LETTERS:  tt.fields.LETTERS,
				numbers:  tt.fields.numbers,
				symbols:  tt.fields.symbols,
				password: tt.fields.password,
			}
			if got := pw.GenerateSymbols(tt.args.length); got != tt.want && !tt.wantErr {
				t.Errorf("Password.GenerateSymbols() = %v, want %v", got, tt.want)
			}
			for _, sym := range pw.symbols {
				if !strings.Contains(symbols, string(sym)) && !tt.wantErr {
					t.Errorf("Password.GenerateSymbols() = %v contains invalid characters", pw.symbols)
				}
			}
		})
	}
}

func TestPassword_Iterator(t *testing.T) {
	type fields struct {
		length   int
		letters  string
		LETTERS  string
		numbers  string
		symbols  string
		password string
	}
	type args struct {
		values string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   int
		wantErr bool
	}{
		{
			name:    "test one - pass",
			fields:  fields{},
			args:    args{values: "abc"},
			want:    "bc",
			want1:   0,
			wantErr: false,
		},
		{
			name:    "test two - failed slice",
			fields:  fields{},
			args:    args{values: "abc"},
			want:    "abc",
			want1:   0,
			wantErr: true,
		},
		{
			name:    "test three - empty string",
			fields:  fields{},
			args:    args{values: ""},
			want:    "",
			want1:   -1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pw := &Password{
				length:   tt.fields.length,
				letters:  tt.fields.letters,
				LETTERS:  tt.fields.LETTERS,
				numbers:  tt.fields.numbers,
				symbols:  tt.fields.symbols,
				password: tt.fields.password,
			}
			got, got1 := pw.Iterator(tt.args.values)
			if got != tt.want && !tt.wantErr {
				t.Errorf("Password.Iterator() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 && !tt.wantErr {
				t.Errorf("Password.Iterator() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
