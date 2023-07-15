package officeservice

import (
	"context"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"mediasoft-customer/internal/repository/officerepository"
	"mediasoft-customer/internal/repository/officerepository/mocks"
	"mediasoft-customer/pkg/logger"
	"reflect"
	"testing"
)

func TestService_GetOfficeList(t *testing.T) {
	//expectedModelOffices := []*customer.Office{
	//	{
	//		Uuid:      uuid.New(),
	//		Name:      "Mediasoft 1",
	//		Address:   "address 1",
	//		CreatedAt: time.Now(),
	//	},
	//	{
	//		Uuid:      uuid.New(),
	//		Name:      "Mediasoft 2",
	//		Address:   "address 2",
	//		CreatedAt: time.Now(),
	//	},
	//}
	//expectedCustomerOffices := []*customer.Office{
	//	{
	//		Uuid:      expectedModelOffices[0].Uuid.String(),
	//		Name:      expectedModelOffices[0].Name,
	//		Address:   expectedModelOffices[0].Address,
	//		CreatedAt: timestamppb.New(expectedModelOffices[0].CreatedAt),
	//	},
	//	{
	//		Uuid:      expectedModelOffices[1].Uuid.String(),
	//		Name:      expectedModelOffices[1].Name,
	//		Address:   expectedModelOffices[1].Address,
	//		CreatedAt: timestamppb.New(expectedModelOffices[1].CreatedAt),
	//	},
	//}

	type fields struct {
		log                              *logger.Logger
		officeRepository                 officerepository.OfficeRepository
		UnimplementedOfficeServiceServer customer.UnimplementedOfficeServiceServer
	}
	type args struct {
		ctx context.Context
		req *customer.GetOfficeListRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customer.GetOfficeListResponse
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				req: &customer.GetOfficeListRequest{},
			},
			//want: &customer.GetOfficeListResponse{
			//	Result: expectedCustomerOffices,
			//},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logger.New()
			officeRepository := mocks.NewOfficeRepository(t)

			//officeRepository.
			//	On("List", mock.Anything).
			//	Once().
			//	Return(expectedModelOffices, nil)

			s := &Service{
				log:                              l,
				officeRepository:                 officeRepository,
				UnimplementedOfficeServiceServer: tt.fields.UnimplementedOfficeServiceServer,
			}

			got, err := s.GetOfficeList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOfficeList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOfficeList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
