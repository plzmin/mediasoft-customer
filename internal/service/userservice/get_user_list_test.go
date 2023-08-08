package userservice

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/protobuf/types/known/timestamppb"
	"mediasoft-customer/internal/model"
	"mediasoft-customer/internal/repository/userrepository"
	"mediasoft-customer/internal/repository/userrepository/mocks"
	"mediasoft-customer/pkg/logger"
	"reflect"
	"testing"
	"time"
)

func TestService_GetUserList(t *testing.T) {
	expectedModelUsers := []*model.User{
		{
			Uuid:       uuid.New(),
			Name:       "PAVEL1",
			OfficeUuid: uuid.New(),
			OfficeName: "MEDIASOFT",
			CreatedAt:  time.Now(),
		},
		{
			Uuid:       uuid.New(),
			Name:       "PAVEL2",
			OfficeUuid: uuid.New(),
			OfficeName: "MEDIASOFT1",
			CreatedAt:  time.Now(),
		},
	}
	expectedCustomerUsers := []*customer.User{
		{
			Uuid:       expectedModelUsers[0].Uuid.String(),
			Name:       expectedModelUsers[0].Name,
			OfficeUuid: expectedModelUsers[0].OfficeUuid.String(),
			OfficeName: expectedModelUsers[0].OfficeName,
			CreatedAt:  timestamppb.New(expectedModelUsers[0].CreatedAt),
		},
		{
			Uuid:       expectedModelUsers[1].Uuid.String(),
			Name:       expectedModelUsers[1].Name,
			OfficeUuid: expectedModelUsers[1].OfficeUuid.String(),
			OfficeName: expectedModelUsers[1].OfficeName,
			CreatedAt:  timestamppb.New(expectedModelUsers[1].CreatedAt),
		},
	}
	expectedCustomerUsersByOfficeUuid := []*customer.User{
		{
			Uuid:       expectedModelUsers[0].Uuid.String(),
			Name:       expectedModelUsers[0].Name,
			OfficeUuid: expectedModelUsers[0].OfficeUuid.String(),
			OfficeName: expectedModelUsers[0].OfficeName,
			CreatedAt:  timestamppb.New(expectedModelUsers[0].CreatedAt),
		},
	}
	type fields struct {
		log                            *logger.Logger
		userRepository                 userrepository.UserRepository
		UnimplementedUserServiceServer customer.UnimplementedUserServiceServer
	}
	type args struct {
		ctx context.Context
		req *customer.GetUserListRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customer.GetUserListResponse
		wantErr bool
	}{
		{
			name: "OK list",
			args: args{
				ctx: context.Background(),
				req: &customer.GetUserListRequest{},
			},
			want: &customer.GetUserListResponse{
				Result: expectedCustomerUsers,
			},
			wantErr: false,
		},
		{
			name: "OK list by office uuid",
			args: args{
				ctx: context.Background(),
				req: &customer.GetUserListRequest{
					OfficeUuid: expectedModelUsers[0].OfficeUuid.String(),
				},
			},
			want: &customer.GetUserListResponse{
				Result: expectedCustomerUsersByOfficeUuid,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logger.New()
			userRepository := mocks.NewUserRepository(t)

			if tt.wantErr == false {
				userRepository.
					On("List", mock.Anything, tt.args.req.OfficeUuid).
					Once().
					Return(returnModelUser(tt.args.req.OfficeUuid, expectedModelUsers), nil)
			}

			s := &Service{
				log:                            l,
				userRepository:                 userRepository,
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			}
			got, err := s.GetUserList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserList() got = %v \n, want %v", got, tt.want)
			}
		})
	}
}

func returnModelUser(uuid string, expectedModelUsers []*model.User) []*model.User {
	if uuid == "" {
		return expectedModelUsers
	}
	var users []*model.User
	for _, user := range expectedModelUsers {
		if user.OfficeUuid.String() == uuid {
			users = append(users, user)
		}
	}
	return users
}
