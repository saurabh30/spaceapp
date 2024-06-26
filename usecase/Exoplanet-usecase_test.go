package usecase

import (
	"spaceapp/domain"
	"testing"
)

func TestExoplanetUsecase_FuelEstimation(t *testing.T) {
	type fields struct {
		Repository domain.ExoplanetRepository
	}
	type args struct {
		p        domain.Exoplanet
		crewSize int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantF   float64
		wantErr bool
	}{
		{
			name:   "Test should give fuel estimate when given valid input for gas planet",
			fields: fields{},
			args: args{
				p: domain.Exoplanet{
					Distance: 5,
					Radius:   5,
					Mass:     5,
					Type:     1,
				},
				crewSize: 5,
			},
			wantF:   2500,
			wantErr: false,
		},
		{
			name:   "Test should give fuel estimate when given valid input for terrestial planet",
			fields: fields{},
			args: args{
				p: domain.Exoplanet{
					Distance: 125,
					Radius:   1,
					Mass:     5,
					Type:     2,
				},
				crewSize: 5,
			},
			wantF:   1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &ExoplanetUsecase{
				Repository: tt.fields.Repository,
			}
			gotF, err := u.FuelEstimation(tt.args.p, tt.args.crewSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExoplanetUsecase.FuelEstimation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotF != tt.wantF {
				t.Errorf("ExoplanetUsecase.FuelEstimation() = %v, want %v", gotF, tt.wantF)
			}
		})
	}
}
