package configuration

import (
	"crave/hub/cmd/model"
	"crave/hub/cmd/target/cmd/api/domain/service"
	"crave/hub/cmd/target/cmd/api/infrastructure/repository"
	"crave/shared/database"
	"net/http"
)

type HubWorkTargetContainer struct {
	Variable         *Variable
	MysqlWrapper     *database.MysqlWrapper
	TargetRepository repository.IRepository
	TargetService    service.IService
}

// defineDatabase implements configuration.IContainer.
func (ctnr *HubWorkTargetContainer) DefineDatabase() error {
	ctnr.MysqlWrapper = database.ConnectMysqlDatabase(ctnr.Variable.Database)
	if err := ctnr.MysqlWrapper.Driver.AutoMigrate(&model.Work{}); err != nil {
		return err
	}
	return nil
}

// getHttpHandler implements configuration.IContainer.
func (ctnr *HubWorkTargetContainer) GetHttpHandler() http.Handler {
	return nil
}

// initVariable implements configuration.IContainer.
func (ctnr *HubWorkTargetContainer) InitVariable() error {
	ctnr.Variable = NewVariable()
	return nil
}

// setRouter implements configuration.IContainer.
func (ctnr *HubWorkTargetContainer) SetRouter(router any) {
	return
}

// DefineGrpc implements configuration.IContainer.
func (ctnr *HubWorkTargetContainer) DefineGrpc() error {
	// lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ctnr.Variable.ctnrubApiIp, ctnr.Variable.ctnrubApiPort))
	// if err != nil {
	// 	return fmt.Errorf("failed to listen : %d, %w", ctnr.Variable.ctnrubApiPort, err)
	// }
	// s := grpc.NewServer()

	// if servErr := s.Serve(lis); servErr != nil {
	// 	return fmt.Errorf("failed to create server: %w", err)
	// }
	return nil
}

func (ctnr *HubWorkTargetContainer) DefineRoute() error {
	return nil
}

func (ctnr *HubWorkTargetContainer) InitDependency(dependency any) error {
	ctnr.DefineDatabase()
	ctnr.TargetRepository = repository.NewRepository(ctnr.MysqlWrapper)
	ctnr.TargetService = service.NewService(ctnr.TargetRepository)
	return nil
}

func NewHubWorkTargetContainer() *HubWorkTargetContainer {
	ctnr := &HubWorkTargetContainer{}
	ctnr.InitVariable()
	ctnr.InitDependency(nil)
	return ctnr
}
