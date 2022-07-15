package entities

import "testing"

func TestValidateCPF(t *testing.T) {
	tests := []struct {
		name string
		cpf  string
		want bool
	}{
		{
			name: "should return CPF valid",
			cpf:  "22234455544",
			want: true,
		},
		{
			name: "should return error CPF invalid",
			cpf:  "224455544",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateCPF(tt.cpf); got != tt.want {
				t.Errorf("ValidateCPF() = %v, want %v", got, tt.want)
			}
		})
	}
}
