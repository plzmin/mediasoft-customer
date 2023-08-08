package userservice

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"mediasoft-customer/internal/model"
	"mediasoft-customer/internal/repository/userrepository"
	"mediasoft-customer/internal/repository/userrepository/mocks"
	"mediasoft-customer/pkg/logger"
	"reflect"
	"testing"
)

type UserMatcher struct {
	ExpectedUser model.User
}

func (m *UserMatcher) Matches(x interface{}) bool {
	if eu, ok := x.(*model.User); ok {
		return eu.Name == m.ExpectedUser.Name && eu.OfficeUuid == m.ExpectedUser.OfficeUuid
	}
	return false
}

func TestService_CreateUser(t *testing.T) {
	type fields struct {
		log                            *logger.Logger
		userRepository                 userrepository.UserRepository
		UnimplementedUserServiceServer customer.UnimplementedUserServiceServer
	}
	type args struct {
		ctx context.Context
		req *customer.CreateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customer.CreateUserResponse
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				req: &customer.CreateUserRequest{
					Name:       "Pavel Zimin",
					OfficeUuid: uuid.New().String(),
				},
			},
			want:    &customer.CreateUserResponse{},
			wantErr: false,
		},
		{
			name: "error Validate Name",
			args: args{
				ctx: context.Background(),
				req: &customer.CreateUserRequest{
					Name:       "Pavel",
					OfficeUuid: uuid.New().String(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error Validate Uuid",
			args: args{
				ctx: context.Background(),
				req: &customer.CreateUserRequest{
					Name:       "Pavel1",
					OfficeUuid: "sdaa",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logger.New()
			userRepository := mocks.NewUserRepository(t)

			officeUuid, _ := uuid.Parse(tt.args.req.OfficeUuid)
			expectedUser := model.User{
				Uuid:       uuid.UUID{},
				Name:       tt.args.req.Name,
				OfficeUuid: officeUuid,
				OfficeName: "medaiasoft",
			}

			if tt.wantErr == false {
				userRepository.
					On("Create", mock.Anything, mock.MatchedBy(func(x interface{}) bool {
						return (&UserMatcher{
							ExpectedUser: expectedUser,
						}).
							Matches(x)
					})).
					Once().
					Return(nil)

			}

			s := &Service{
				log:                            l,
				userRepository:                 userRepository,
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			}

			got, err := s.CreateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
