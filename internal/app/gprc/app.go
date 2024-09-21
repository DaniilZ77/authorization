package gprc

import (
	"context"
	"fmt"
	"net"

	"github.com/DaniilZ77/authorization/internal/lib/logger"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type App struct {
	gRPCServer *grpc.Server
	port       string
}

func New(
	port string,
) *App {
	ctx := context.Background()

	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived,
			logging.PayloadSent,
		),
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p any) (err error) {
			logger.Log().Error(ctx, "recovered from panic")

			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	gRPCServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recoveryOpts...),
			logging.UnaryServerInterceptor(interceptorLogger(logger.Log()), loggingOpts...),
		),
	)

	// Register you grpc services
	//

	return &App{
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func interceptorLogger(l logger.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		switch lvl {
		case logging.LevelDebug:
			l.Debug(ctx, msg, fields...)
		case logging.LevelInfo:
			l.Info(ctx, msg, fields...)
		case logging.LevelWarn:
			l.Warn(ctx, msg, fields...)
		case logging.LevelError:
			l.Error(ctx, msg, fields...)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	ctx := context.Background()

	l, err := net.Listen("tcp", a.port)
	if err != nil {
		return err
	}

	logger.Log().Info(ctx, "grpc server started")

	if err := a.gRPCServer.Serve(l); err != nil {
		return err
	}

	return nil
}

func (a *App) Stop() {
	ctx := context.Background()

	logger.Log().Info(ctx, "stopping grpc server")

	a.gRPCServer.GracefulStop()
}
