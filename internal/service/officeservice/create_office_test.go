package officeservice

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"mediasoft-customer/internal/model"
	"mediasoft-customer/internal/repository/officerepository"
	"mediasoft-customer/internal/repository/officerepository/mocks"
	"mediasoft-customer/pkg/logger"
	"reflect"
	"testing"
	"time"
)

type OfficeMatcher struct {
	ExpectedOffice model.Office
}

func (m *OfficeMatcher) Matches(x interface{}) bool {
	if eo, ok := x.(*model.Office); ok {
		return eo.Name == m.ExpectedOffice.Name && eo.Address == m.ExpectedOffice.Address
	}
	return false
}

func TestService_CreateOffice(t *testing.T) {
	type fields struct {
		log                              *logger.Logger
		officeRepository                 officerepository.OfficeRepository
		UnimplementedOfficeServiceServer customer.UnimplementedOfficeServiceServer
	}
	type args struct {
		ctx context.Context
		req *customer.CreateOfficeRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customer.CreateOfficeResponse
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				req: &customer.CreateOfficeRequest{
					Name:    "Mediasoft",
					Address: "Ulyanovsk",
				},
			},
			want:    &customer.CreateOfficeResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logger.New()
			officeRepository := mocks.NewOfficeRepository(t)

			expectedOffice := model.Office{
				Uuid:      uuid.UUID{},
				Name:      tt.args.req.Name,
				Address:   tt.args.req.Address,
				CreatedAt: time.Time{},
			}

			officeRepository.
				On("Create", mock.Anything, mock.MatchedBy(func(x interface{}) bool {
					return (&OfficeMatcher{
						ExpectedOffice: expectedOffice,
					}).
						Matches(x)
				})).
				Once().
				Return(nil)

			s := &Service{
				log:                              l,
				officeRepository:                 officeRepository,
				UnimplementedOfficeServiceServer: tt.fields.UnimplementedOfficeServiceServer,
			}

			got, err := s.CreateOffice(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOffice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOffice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateOfficeValidate(t *testing.T) {
	type fields struct {
		log                              *logger.Logger
		officeRepository                 officerepository.OfficeRepository
		UnimplementedOfficeServiceServer customer.UnimplementedOfficeServiceServer
	}
	type args struct {
		ctx context.Context
		req *customer.CreateOfficeRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customer.CreateOfficeResponse
		wantErr bool
	}{
		{
			name: "Validate Name",
			args: args{
				ctx: context.Background(),
				req: &customer.CreateOfficeRequest{
					Name:    "Media",
					Address: "Ulyanovsk",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Validate Address",
			args: args{
				ctx: context.Background(),
				req: &customer.CreateOfficeRequest{
					Name:    "Mediasoft",
					Address: "Ulya",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logger.New()
			officeRepository := mocks.NewOfficeRepository(t)

			s := &Service{
				log:                              l,
				officeRepository:                 officeRepository,
				UnimplementedOfficeServiceServer: tt.fields.UnimplementedOfficeServiceServer,
			}

			got, err := s.CreateOffice(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOffice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOffice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
