package domain_test

import (
	"context"
	"fmt"

	"github.com/kloudlite/operator/pkg/kubectl"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"kloudlite.io/apps/accounts/internal/domain"
	"kloudlite.io/apps/accounts/internal/entities"
	authMock "kloudlite.io/grpc-interfaces/kloudlite.io/rpc/auth/mocks"
	commsMock "kloudlite.io/grpc-interfaces/kloudlite.io/rpc/comms/mocks"
	consoleMock "kloudlite.io/grpc-interfaces/kloudlite.io/rpc/console/mocks"
	_ "kloudlite.io/grpc-interfaces/kloudlite.io/rpc/container_registry"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/iam"
	iamMock "kloudlite.io/grpc-interfaces/kloudlite.io/rpc/iam/mocks"
	fn "kloudlite.io/pkg/functions"
	k8sMock "kloudlite.io/pkg/k8s/mocks"
	"kloudlite.io/pkg/logging"
	"kloudlite.io/pkg/repos"

	// "kloudlite.io/pkg/repos"
	reposMock "kloudlite.io/pkg/repos/mocks"
)

var _ = Describe("domain.ActivateAccount says", func() {
	// Given("an account", func() {})
	var authClient *authMock.AuthClient
	var iamClient *iamMock.IAMClient
	var consoleClient *consoleMock.ConsoleClient
	// var containerRegistryClient container_registry.ContainerRegistryClient
	var commsClient *commsMock.CommsClient
	var accountRepo *reposMock.DbRepo[*entities.Account]
	var invitationRepo *reposMock.DbRepo[*entities.Invitation]
	var k8sYamlClient *kubectl.YAMLClient
	var k8sExtendedClient *k8sMock.ExtendedK8sClient
	var logger logging.Logger

	BeforeEach(func() {
		authClient = authMock.NewAuthClient()
		iamClient = iamMock.NewIAMClient()
		consoleClient = consoleMock.NewConsoleClient()
		// containerRegistryClient = container_registry.NewContainerRegistryClient()
		commsClient = commsMock.NewCommsClient()
		accountRepo = reposMock.NewDbRepo[*entities.Account]()
		invitationRepo = reposMock.NewDbRepo[*entities.Invitation]()
		// k8sYamlClient = kubectl.NewYAMLClient()
		k8sExtendedClient = k8sMock.NewExtendedK8sClient()
		// logger = logging.NewLogger()
	})

	getDomain := func() domain.Domain {
		return domain.NewDomain(
			iamClient,
			consoleClient,
			//f.containerRegistryClient,
			authClient,
			commsClient,
			k8sYamlClient,
			k8sExtendedClient,

			accountRepo,
			invitationRepo,

			logger,
		)
	}

	When("user has no IAM permission to activate account", func() {
		It("account activation should fail", func() {
			d := getDomain()

			iamClient.MockCan = func(ctx context.Context, in *iam.CanIn, opts ...grpc.CallOption) (*iam.CanOut, error) {
				return &iam.CanOut{Status: false}, nil
			}

			accountRepo.MockFindOne = func(ctx context.Context, filter repos.Filter) (*entities.Account, error) {
				return nil, fmt.Errorf("not found")
			}

			_, err := d.ActivateAccount(domain.UserContext{}, "sample")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("unauthorized"))
		})
	})

	When("user has IAM permission to activate account", func() {
		Context("but account does not exist", func() {
			It("it should fail", func() {
				d := getDomain()

				iamClient.MockCan = func(ctx context.Context, in *iam.CanIn, opts ...grpc.CallOption) (*iam.CanOut, error) {
					return &iam.CanOut{Status: true}, nil
				}

				accountRepo.MockFindOne = func(ctx context.Context, filter repos.Filter) (*entities.Account, error) {
					return nil, fmt.Errorf("mock: account not found")
				}

				_, err := d.ActivateAccount(domain.UserContext{}, "sample")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("mock: account not found"))
			})
		})

		Context("but account is already active", func() {
			It("it should fail", func() {
				d := getDomain()

				iamClient.MockCan = func(ctx context.Context, in *iam.CanIn, opts ...grpc.CallOption) (*iam.CanOut, error) {
					return &iam.CanOut{Status: true}, nil
				}

				accountRepo.MockFindOne = func(ctx context.Context, filter repos.Filter) (*entities.Account, error) {
					return &entities.Account{IsActive: fn.New(true)}, nil
				}

				_, err := d.ActivateAccount(domain.UserContext{}, "sample")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("already active"))
			})
		})

		Context("and account exists and it is currently disabled", func() {
			It("then, account should get activated", func() {
				d := getDomain()

				iamClient.MockCan = func(ctx context.Context, in *iam.CanIn, opts ...grpc.CallOption) (*iam.CanOut, error) {
					return &iam.CanOut{Status: true}, nil
				}

				accountRepo.MockFindOne = func(ctx context.Context, filter repos.Filter) (*entities.Account, error) {
					return &entities.Account{IsActive: fn.New(false)}, nil
				}

				accountRepo.MockUpdateById = func(ctx context.Context, id repos.ID, updatedData *entities.Account, opts ...repos.UpdateOpts) (*entities.Account, error) {
					return &entities.Account{IsActive: fn.New(true)}, nil
				}
				_, err := d.ActivateAccount(domain.UserContext{}, "sample")
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
