package classGRPC

// import (
// 	classPostgres "github.com/zchelalo/sa_class_management/internal/modules/class/infrastructure/repositories/postgres"
// 	classProto "github.com/zchelalo/sa_class_management/pkg/proto/class"
// 	"github.com/zchelalo/sa_class_management/pkg/sqlc/db"
// )

// type ClassRouter struct {
// 	useCase *classApplication.ClassUseCases
// 	classProto.UnimplementedClassServiceServer
// }

// func New(store *db.SQLStore) *ClassRouter {
// 	classRepository := classPostgres.New(store)
// 	classUseCases := classApplication.New(classRepository)

// 	return &ClassRouter{
// 		useCase: classUseCases,
// 	}
// }
