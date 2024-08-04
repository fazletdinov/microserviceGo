package auth

import (
	"auth/config"
	authgrpc "auth/protogen/auth"
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	oteltrace "go.opentelemetry.io/otel/trace"

	"auth/internal/domain/service/grpc_service"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCController struct {
	authgrpc.UnimplementedGatewayAuthServer
	UserGRPCService grpcservice.UserGRPCService
	Env             *config.Config
}

func (gc *GRPCController) CreateUser(
	ctx context.Context,
	authRequest *authgrpc.CreateUserRequest,
) (*authgrpc.CreateUserResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if authRequest.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле email обязательно")
	}
	if authRequest.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле password обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"CreateUser",
		oteltrace.WithAttributes(attribute.String("Email", authRequest.Email)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в auth для создания пользователя")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"CreateUser_counter",
		metric.WithDescription("Сколько раз вызывалась функция CreateUser"),
	)
	counter.Add(
		ctx,
		1,
		metric.WithAttributes(attribute.String("registration", "additional")),
	)

	userID, err := gc.UserGRPCService.CreateUser(traceCtx, authRequest.GetEmail(), authRequest.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authgrpc.CreateUserResponse{
		UserId: userID.String(),
	}, nil
}

func (gc *GRPCController) GetUserByID(
	ctx context.Context,
	authRequest *authgrpc.GetUserRequest,
) (*authgrpc.GetUserResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)
	if authRequest.GetUserId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле user_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"GetUserByID",
		oteltrace.WithAttributes(attribute.String("UserID", authRequest.UserId)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в auth для получения пользователя по ID")
	defer span.End()

	counter, err := meter.Int64Counter(
		"GetUserByID_counter",
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "ошибка при создании счетчика")
	}

	counter.Add(
		ctx,
		1,
	)

	user, err := gc.UserGRPCService.GetUserByID(traceCtx, uuid.MustParse(authRequest.GetUserId()))
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authgrpc.GetUserResponse{
		UserId:    user.ID.String(),
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
	}, nil
}

func (gc *GRPCController) UpdateUser(
	ctx context.Context,
	authRequest *authgrpc.UpdateUserRequest,
) (*authgrpc.UpdateUserResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if authRequest.GetUserId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле user_id обязательно")
	}
	if authRequest.GetFirstName() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле first_name обязательно")
	}
	if authRequest.GetLastName() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле last_name обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"UpdateUser",
		oteltrace.WithAttributes(attribute.String("UserID", authRequest.UserId)),
		oteltrace.WithAttributes(attribute.String("FirstName", authRequest.FirstName)),
		oteltrace.WithAttributes(attribute.String("LastName", authRequest.LastName)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в auth для обновления пользователя по ID")
	defer span.End()

	counter, err := meter.Int64Counter(
		"UpdateUser_counter",
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "ошибка при создании счетчика")
	}

	counter.Add(
		ctx,
		1,
	)

	err = gc.UserGRPCService.UpdateUser(traceCtx, uuid.MustParse(authRequest.GetUserId()), authRequest.GetFirstName(), authRequest.GetLastName())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authgrpc.UpdateUserResponse{
		SuccessMessage: "Пользователь успешно обновлен",
	}, nil
}

func (gc *GRPCController) DeleteUser(
	ctx context.Context,
	authRequest *authgrpc.DeleteUserRequest,
) (*authgrpc.DeleteUserResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if authRequest.GetUserId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле user_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"DeleteUser",
		oteltrace.WithAttributes(attribute.String("UserID", authRequest.UserId)),
	)

	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в auth для удаления пользователя по ID")
	defer span.End()

	counter, err := meter.Int64Counter(
		"DeleteUser_counter",
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "ошибка при создании счетчика")
	}

	counter.Add(
		ctx,
		1,
	)

	err = gc.UserGRPCService.DeleteUser(traceCtx, uuid.MustParse(authRequest.GetUserId()))
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authgrpc.DeleteUserResponse{
		SuccessMessage: "Успешно удалено",
	}, nil
}

func (gc *GRPCController) GetUserByEmail(
	ctx context.Context,
	authRequest *authgrpc.GetUserByEmailRequest,
) (*authgrpc.GetUserResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if authRequest.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле email обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"GetUserByEmail",
		oteltrace.WithAttributes(attribute.String("Email", authRequest.Email)),
	)

	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в auth для получения пользователя по Email")
	defer span.End()

	counter, err := meter.Int64Counter(
		"GetUserByEmail_counter",
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "ошибка при создании счетчика")
	}

	counter.Add(
		ctx,
		1,
	)

	user, err := gc.UserGRPCService.GetUserByEmail(traceCtx, authRequest.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authgrpc.GetUserResponse{
		UserId:    user.ID.String(),
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.String(),
	}, nil
}

func (gc *GRPCController) GetUserByEmailIsActive(
	ctx context.Context,
	authRequest *authgrpc.GetUserByEmailRequest,
) (*authgrpc.GetUserResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if authRequest.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле email обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"GetUserByEmailIsActive",
		oteltrace.WithAttributes(attribute.String("Email", authRequest.Email)),
	)

	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в auth для получения активного пользователя по Email")
	defer span.End()

	counter, err := meter.Int64Counter(
		"GetUserByEmailIsActive_counter",
		metric.WithDescription("Сколько раз вызывалась функция GetUserByEmailIsActive"),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "ошибка при создании счетчика")
	}

	counter.Add(
		ctx,
		1,
	)

	user, err := gc.UserGRPCService.GetUserByEmailIsActive(traceCtx, authRequest.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authgrpc.GetUserResponse{
		UserId:    user.ID.String(),
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.String(),
	}, nil
}
