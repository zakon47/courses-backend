package service

import (
	"context"
	"github.com/zhashkevych/courses-backend/internal/domain"
	"github.com/zhashkevych/courses-backend/internal/repository"
	"github.com/zhashkevych/courses-backend/pkg/auth"
	"github.com/zhashkevych/courses-backend/pkg/cache"
	"github.com/zhashkevych/courses-backend/pkg/email"
	"github.com/zhashkevych/courses-backend/pkg/hash"
	"github.com/zhashkevych/courses-backend/pkg/payment"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

// TODO handle "not found" errors

type Schools interface {
	GetByDomain(ctx context.Context, domainName string) (domain.School, error)
}

type StudentSignUpInput struct {
	Name           string
	Email          string
	Password       string
	SchoolID       primitive.ObjectID
	RegisterSource string
}

type SignInInput struct {
	Email    string
	Password string
	SchoolID primitive.ObjectID
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Students interface {
	SignUp(ctx context.Context, input StudentSignUpInput) error
	SignIn(ctx context.Context, input SignInInput) (Tokens, error)
	RefreshTokens(ctx context.Context, schoolId primitive.ObjectID, refreshToken string) (Tokens, error)
	Verify(ctx context.Context, hash string) error
	GetModuleLessons(ctx context.Context, schoolId, studentId, moduleId primitive.ObjectID) ([]domain.Lesson, error)
	GiveAccessToModules(ctx context.Context, studentId primitive.ObjectID, moduleIds []primitive.ObjectID) error
	GiveAccessToPackages(ctx context.Context, studentId primitive.ObjectID, packageIds []primitive.ObjectID) error
}

type Admins interface {
	SignIn(ctx context.Context, input SignInInput) (Tokens, error)
	RefreshTokens(ctx context.Context, schoolId primitive.ObjectID, refreshToken string) (Tokens, error)
	GetCourses(ctx context.Context, schoolId primitive.ObjectID) ([]domain.Course, error)
	GetCourseById(ctx context.Context, schoolId, courseId primitive.ObjectID) (domain.Course, error)
}

type AddToListInput struct {
	Email            string
	Name             string
	RegisterSource   string
	VerificationCode string
}

type Emails interface {
	AddToList(AddToListInput) error
}

type UpdateCourseInput struct {
	CourseID    string
	Name        string
	Code        string
	Description string
	Published   *bool
}

type Courses interface {
	Create(ctx context.Context, schoolId primitive.ObjectID, name string) (primitive.ObjectID, error)
	Update(ctx context.Context, schoolId primitive.ObjectID, inp UpdateCourseInput) error
}

type PromoCodes interface {
	GetByCode(ctx context.Context, schoolId primitive.ObjectID, code string) (domain.PromoCode, error)
	GetById(ctx context.Context, id primitive.ObjectID) (domain.PromoCode, error)
}

type Offers interface {
	GetById(ctx context.Context, id primitive.ObjectID) (domain.Offer, error)
	GetByModule(ctx context.Context, schoolId, moduleId primitive.ObjectID) ([]domain.Offer, error)
	GetByPackage(ctx context.Context, schoolId, packageId primitive.ObjectID) ([]domain.Offer, error)
}

type Modules interface {
	GetByCourse(ctx context.Context, courseId primitive.ObjectID) ([]domain.Module, error)
	GetById(ctx context.Context, moduleId primitive.ObjectID) (domain.Module, error)
	GetByPackages(ctx context.Context, packageIds []primitive.ObjectID) ([]domain.Module, error)
	GetWithContent(ctx context.Context, moduleId primitive.ObjectID) (domain.Module, error)
	Create(ctx context.Context, courseId primitive.ObjectID, name string, position int) (primitive.ObjectID, error)
}

type Orders interface {
	Create(ctx context.Context, studentId, offerId, promocodeId primitive.ObjectID) (string, error)
	AddTransaction(ctx context.Context, id primitive.ObjectID, transaction domain.Transaction) (domain.Order, error)
}

type Payments interface {
	ProcessTransaction(ctx context.Context, callbackData payment.Callback) error
}

type Services struct {
	Schools    Schools
	Students   Students
	Courses    Courses
	PromoCodes PromoCodes
	Offers     Offers
	Modules    Modules
	Payments   Payments
	Orders     Orders
	Admins     Admins
}

type ServicesDeps struct {
	Repos              *repository.Repositories
	Cache              cache.Cache
	Hasher             hash.PasswordHasher
	TokenManager       auth.TokenManager
	EmailProvider      email.Provider
	EmailListId        string
	PaymentProvider    payment.FondyProvider
	AccessTokenTTL     time.Duration
	RefreshTokenTTL    time.Duration
	PaymentCallbackURL string
	PaymentResponseURL string
	CacheTTL           int64
}

func NewServices(deps ServicesDeps) *Services {
	emailsService := NewEmailsService(deps.EmailProvider, deps.EmailListId)
	coursesService := NewCoursesService(deps.Repos.Courses)
	modulesService := NewModulesService(deps.Repos.Modules, deps.Repos.LessonContent)
	offersService := NewOffersService(deps.Repos.Offers, modulesService)
	promoCodesService := NewPromoCodeService(deps.Repos.PromoCodes)
	ordersService := NewOrdersService(deps.Repos.Orders, offersService, promoCodesService, deps.PaymentProvider, deps.PaymentCallbackURL, deps.PaymentResponseURL)
	studentsService := NewStudentsService(deps.Repos.Students, modulesService, offersService, deps.Hasher,
		deps.TokenManager, emailsService, deps.AccessTokenTTL, deps.RefreshTokenTTL)

	return &Services{
		Schools:    NewSchoolsService(deps.Repos.Schools, deps.Cache, deps.CacheTTL),
		Students:   studentsService,
		Courses:    coursesService,
		PromoCodes: promoCodesService,
		Offers:     offersService,
		Modules:    modulesService,
		Payments:   NewPaymentsService(deps.PaymentProvider, ordersService, offersService, studentsService),
		Orders:     ordersService,
		Admins:     NewAdminsService(deps.Hasher, deps.TokenManager, deps.Repos.Admins, deps.Repos.Schools, deps.AccessTokenTTL, deps.RefreshTokenTTL),
	}
}
