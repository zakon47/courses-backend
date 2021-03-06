// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	domain "github.com/zhashkevych/courses-backend/internal/domain"
	service "github.com/zhashkevych/courses-backend/internal/service"
	payment "github.com/zhashkevych/courses-backend/pkg/payment"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	reflect "reflect"
)

// MockSchools is a mock of Schools interface
type MockSchools struct {
	ctrl     *gomock.Controller
	recorder *MockSchoolsMockRecorder
}

// MockSchoolsMockRecorder is the mock recorder for MockSchools
type MockSchoolsMockRecorder struct {
	mock *MockSchools
}

// NewMockSchools creates a new mock instance
func NewMockSchools(ctrl *gomock.Controller) *MockSchools {
	mock := &MockSchools{ctrl: ctrl}
	mock.recorder = &MockSchoolsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSchools) EXPECT() *MockSchoolsMockRecorder {
	return m.recorder
}

// GetByDomain mocks base method
func (m *MockSchools) GetByDomain(ctx context.Context, domainName string) (domain.School, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDomain", ctx, domainName)
	ret0, _ := ret[0].(domain.School)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDomain indicates an expected call of GetByDomain
func (mr *MockSchoolsMockRecorder) GetByDomain(ctx, domainName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDomain", reflect.TypeOf((*MockSchools)(nil).GetByDomain), ctx, domainName)
}

// MockStudents is a mock of Students interface
type MockStudents struct {
	ctrl     *gomock.Controller
	recorder *MockStudentsMockRecorder
}

// MockStudentsMockRecorder is the mock recorder for MockStudents
type MockStudentsMockRecorder struct {
	mock *MockStudents
}

// NewMockStudents creates a new mock instance
func NewMockStudents(ctrl *gomock.Controller) *MockStudents {
	mock := &MockStudents{ctrl: ctrl}
	mock.recorder = &MockStudentsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStudents) EXPECT() *MockStudentsMockRecorder {
	return m.recorder
}

// SignUp mocks base method
func (m *MockStudents) SignUp(ctx context.Context, input service.StudentSignUpInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp
func (mr *MockStudentsMockRecorder) SignUp(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockStudents)(nil).SignUp), ctx, input)
}

// SignIn mocks base method
func (m *MockStudents) SignIn(ctx context.Context, input service.SignInInput) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, input)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn
func (mr *MockStudentsMockRecorder) SignIn(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockStudents)(nil).SignIn), ctx, input)
}

// RefreshTokens mocks base method
func (m *MockStudents) RefreshTokens(ctx context.Context, schoolId primitive.ObjectID, refreshToken string) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTokens", ctx, schoolId, refreshToken)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshTokens indicates an expected call of RefreshTokens
func (mr *MockStudentsMockRecorder) RefreshTokens(ctx, schoolId, refreshToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTokens", reflect.TypeOf((*MockStudents)(nil).RefreshTokens), ctx, schoolId, refreshToken)
}

// Verify mocks base method
func (m *MockStudents) Verify(ctx context.Context, hash string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, hash)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify
func (mr *MockStudentsMockRecorder) Verify(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockStudents)(nil).Verify), ctx, hash)
}

// GetModuleLessons mocks base method
func (m *MockStudents) GetModuleLessons(ctx context.Context, schoolId, studentId, moduleId primitive.ObjectID) ([]domain.Lesson, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleLessons", ctx, schoolId, studentId, moduleId)
	ret0, _ := ret[0].([]domain.Lesson)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModuleLessons indicates an expected call of GetModuleLessons
func (mr *MockStudentsMockRecorder) GetModuleLessons(ctx, schoolId, studentId, moduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleLessons", reflect.TypeOf((*MockStudents)(nil).GetModuleLessons), ctx, schoolId, studentId, moduleId)
}

// GiveAccessToModules mocks base method
func (m *MockStudents) GiveAccessToModules(ctx context.Context, studentId primitive.ObjectID, moduleIds []primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GiveAccessToModules", ctx, studentId, moduleIds)
	ret0, _ := ret[0].(error)
	return ret0
}

// GiveAccessToModules indicates an expected call of GiveAccessToModules
func (mr *MockStudentsMockRecorder) GiveAccessToModules(ctx, studentId, moduleIds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GiveAccessToModules", reflect.TypeOf((*MockStudents)(nil).GiveAccessToModules), ctx, studentId, moduleIds)
}

// GiveAccessToPackages mocks base method
func (m *MockStudents) GiveAccessToPackages(ctx context.Context, studentId primitive.ObjectID, packageIds []primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GiveAccessToPackages", ctx, studentId, packageIds)
	ret0, _ := ret[0].(error)
	return ret0
}

// GiveAccessToPackages indicates an expected call of GiveAccessToPackages
func (mr *MockStudentsMockRecorder) GiveAccessToPackages(ctx, studentId, packageIds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GiveAccessToPackages", reflect.TypeOf((*MockStudents)(nil).GiveAccessToPackages), ctx, studentId, packageIds)
}

// MockAdmins is a mock of Admins interface
type MockAdmins struct {
	ctrl     *gomock.Controller
	recorder *MockAdminsMockRecorder
}

// MockAdminsMockRecorder is the mock recorder for MockAdmins
type MockAdminsMockRecorder struct {
	mock *MockAdmins
}

// NewMockAdmins creates a new mock instance
func NewMockAdmins(ctrl *gomock.Controller) *MockAdmins {
	mock := &MockAdmins{ctrl: ctrl}
	mock.recorder = &MockAdminsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdmins) EXPECT() *MockAdminsMockRecorder {
	return m.recorder
}

// SignIn mocks base method
func (m *MockAdmins) SignIn(ctx context.Context, input service.SignInInput) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, input)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn
func (mr *MockAdminsMockRecorder) SignIn(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockAdmins)(nil).SignIn), ctx, input)
}

// RefreshTokens mocks base method
func (m *MockAdmins) RefreshTokens(ctx context.Context, schoolId primitive.ObjectID, refreshToken string) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTokens", ctx, schoolId, refreshToken)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshTokens indicates an expected call of RefreshTokens
func (mr *MockAdminsMockRecorder) RefreshTokens(ctx, schoolId, refreshToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTokens", reflect.TypeOf((*MockAdmins)(nil).RefreshTokens), ctx, schoolId, refreshToken)
}

// GetCourses mocks base method
func (m *MockAdmins) GetCourses(ctx context.Context, schoolId primitive.ObjectID) ([]domain.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourses", ctx, schoolId)
	ret0, _ := ret[0].([]domain.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourses indicates an expected call of GetCourses
func (mr *MockAdminsMockRecorder) GetCourses(ctx, schoolId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourses", reflect.TypeOf((*MockAdmins)(nil).GetCourses), ctx, schoolId)
}

// GetCourseById mocks base method
func (m *MockAdmins) GetCourseById(ctx context.Context, schoolId, courseId primitive.ObjectID) (domain.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourseById", ctx, schoolId, courseId)
	ret0, _ := ret[0].(domain.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourseById indicates an expected call of GetCourseById
func (mr *MockAdminsMockRecorder) GetCourseById(ctx, schoolId, courseId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourseById", reflect.TypeOf((*MockAdmins)(nil).GetCourseById), ctx, schoolId, courseId)
}

// MockEmails is a mock of Emails interface
type MockEmails struct {
	ctrl     *gomock.Controller
	recorder *MockEmailsMockRecorder
}

// MockEmailsMockRecorder is the mock recorder for MockEmails
type MockEmailsMockRecorder struct {
	mock *MockEmails
}

// NewMockEmails creates a new mock instance
func NewMockEmails(ctrl *gomock.Controller) *MockEmails {
	mock := &MockEmails{ctrl: ctrl}
	mock.recorder = &MockEmailsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEmails) EXPECT() *MockEmailsMockRecorder {
	return m.recorder
}

// AddToList mocks base method
func (m *MockEmails) AddToList(arg0 service.AddToListInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToList", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToList indicates an expected call of AddToList
func (mr *MockEmailsMockRecorder) AddToList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToList", reflect.TypeOf((*MockEmails)(nil).AddToList), arg0)
}

// MockCourses is a mock of Courses interface
type MockCourses struct {
	ctrl     *gomock.Controller
	recorder *MockCoursesMockRecorder
}

// MockCoursesMockRecorder is the mock recorder for MockCourses
type MockCoursesMockRecorder struct {
	mock *MockCourses
}

// NewMockCourses creates a new mock instance
func NewMockCourses(ctrl *gomock.Controller) *MockCourses {
	mock := &MockCourses{ctrl: ctrl}
	mock.recorder = &MockCoursesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCourses) EXPECT() *MockCoursesMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCourses) Create(ctx context.Context, schoolId primitive.ObjectID, name string) (primitive.ObjectID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, schoolId, name)
	ret0, _ := ret[0].(primitive.ObjectID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockCoursesMockRecorder) Create(ctx, schoolId, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCourses)(nil).Create), ctx, schoolId, name)
}

// Update mocks base method
func (m *MockCourses) Update(ctx context.Context, schoolId primitive.ObjectID, inp service.UpdateCourseInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, schoolId, inp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockCoursesMockRecorder) Update(ctx, schoolId, inp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCourses)(nil).Update), ctx, schoolId, inp)
}

// MockPromoCodes is a mock of PromoCodes interface
type MockPromoCodes struct {
	ctrl     *gomock.Controller
	recorder *MockPromoCodesMockRecorder
}

// MockPromoCodesMockRecorder is the mock recorder for MockPromoCodes
type MockPromoCodesMockRecorder struct {
	mock *MockPromoCodes
}

// NewMockPromoCodes creates a new mock instance
func NewMockPromoCodes(ctrl *gomock.Controller) *MockPromoCodes {
	mock := &MockPromoCodes{ctrl: ctrl}
	mock.recorder = &MockPromoCodesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPromoCodes) EXPECT() *MockPromoCodesMockRecorder {
	return m.recorder
}

// GetByCode mocks base method
func (m *MockPromoCodes) GetByCode(ctx context.Context, schoolId primitive.ObjectID, code string) (domain.PromoCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCode", ctx, schoolId, code)
	ret0, _ := ret[0].(domain.PromoCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCode indicates an expected call of GetByCode
func (mr *MockPromoCodesMockRecorder) GetByCode(ctx, schoolId, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCode", reflect.TypeOf((*MockPromoCodes)(nil).GetByCode), ctx, schoolId, code)
}

// GetById mocks base method
func (m *MockPromoCodes) GetById(ctx context.Context, id primitive.ObjectID) (domain.PromoCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(domain.PromoCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById
func (mr *MockPromoCodesMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockPromoCodes)(nil).GetById), ctx, id)
}

// MockOffers is a mock of Offers interface
type MockOffers struct {
	ctrl     *gomock.Controller
	recorder *MockOffersMockRecorder
}

// MockOffersMockRecorder is the mock recorder for MockOffers
type MockOffersMockRecorder struct {
	mock *MockOffers
}

// NewMockOffers creates a new mock instance
func NewMockOffers(ctrl *gomock.Controller) *MockOffers {
	mock := &MockOffers{ctrl: ctrl}
	mock.recorder = &MockOffersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOffers) EXPECT() *MockOffersMockRecorder {
	return m.recorder
}

// GetById mocks base method
func (m *MockOffers) GetById(ctx context.Context, id primitive.ObjectID) (domain.Offer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(domain.Offer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById
func (mr *MockOffersMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockOffers)(nil).GetById), ctx, id)
}

// GetByModule mocks base method
func (m *MockOffers) GetByModule(ctx context.Context, schoolId, moduleId primitive.ObjectID) ([]domain.Offer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByModule", ctx, schoolId, moduleId)
	ret0, _ := ret[0].([]domain.Offer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByModule indicates an expected call of GetByModule
func (mr *MockOffersMockRecorder) GetByModule(ctx, schoolId, moduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByModule", reflect.TypeOf((*MockOffers)(nil).GetByModule), ctx, schoolId, moduleId)
}

// GetByPackage mocks base method
func (m *MockOffers) GetByPackage(ctx context.Context, schoolId, packageId primitive.ObjectID) ([]domain.Offer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPackage", ctx, schoolId, packageId)
	ret0, _ := ret[0].([]domain.Offer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPackage indicates an expected call of GetByPackage
func (mr *MockOffersMockRecorder) GetByPackage(ctx, schoolId, packageId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPackage", reflect.TypeOf((*MockOffers)(nil).GetByPackage), ctx, schoolId, packageId)
}

// MockModules is a mock of Modules interface
type MockModules struct {
	ctrl     *gomock.Controller
	recorder *MockModulesMockRecorder
}

// MockModulesMockRecorder is the mock recorder for MockModules
type MockModulesMockRecorder struct {
	mock *MockModules
}

// NewMockModules creates a new mock instance
func NewMockModules(ctrl *gomock.Controller) *MockModules {
	mock := &MockModules{ctrl: ctrl}
	mock.recorder = &MockModulesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModules) EXPECT() *MockModulesMockRecorder {
	return m.recorder
}

// GetByCourse mocks base method
func (m *MockModules) GetByCourse(ctx context.Context, courseId primitive.ObjectID) ([]domain.Module, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCourse", ctx, courseId)
	ret0, _ := ret[0].([]domain.Module)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCourse indicates an expected call of GetByCourse
func (mr *MockModulesMockRecorder) GetByCourse(ctx, courseId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCourse", reflect.TypeOf((*MockModules)(nil).GetByCourse), ctx, courseId)
}

// GetById mocks base method
func (m *MockModules) GetById(ctx context.Context, moduleId primitive.ObjectID) (domain.Module, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, moduleId)
	ret0, _ := ret[0].(domain.Module)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById
func (mr *MockModulesMockRecorder) GetById(ctx, moduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockModules)(nil).GetById), ctx, moduleId)
}

// GetByPackages mocks base method
func (m *MockModules) GetByPackages(ctx context.Context, packageIds []primitive.ObjectID) ([]domain.Module, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPackages", ctx, packageIds)
	ret0, _ := ret[0].([]domain.Module)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPackages indicates an expected call of GetByPackages
func (mr *MockModulesMockRecorder) GetByPackages(ctx, packageIds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPackages", reflect.TypeOf((*MockModules)(nil).GetByPackages), ctx, packageIds)
}

// GetWithContent mocks base method
func (m *MockModules) GetWithContent(ctx context.Context, moduleId primitive.ObjectID) (domain.Module, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithContent", ctx, moduleId)
	ret0, _ := ret[0].(domain.Module)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithContent indicates an expected call of GetWithContent
func (mr *MockModulesMockRecorder) GetWithContent(ctx, moduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithContent", reflect.TypeOf((*MockModules)(nil).GetWithContent), ctx, moduleId)
}

// Create mocks base method
func (m *MockModules) Create(ctx context.Context, courseId primitive.ObjectID, name string, position int) (primitive.ObjectID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, courseId, name, position)
	ret0, _ := ret[0].(primitive.ObjectID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockModulesMockRecorder) Create(ctx, courseId, name, position interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockModules)(nil).Create), ctx, courseId, name, position)
}

// MockOrders is a mock of Orders interface
type MockOrders struct {
	ctrl     *gomock.Controller
	recorder *MockOrdersMockRecorder
}

// MockOrdersMockRecorder is the mock recorder for MockOrders
type MockOrdersMockRecorder struct {
	mock *MockOrders
}

// NewMockOrders creates a new mock instance
func NewMockOrders(ctrl *gomock.Controller) *MockOrders {
	mock := &MockOrders{ctrl: ctrl}
	mock.recorder = &MockOrdersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrders) EXPECT() *MockOrdersMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockOrders) Create(ctx context.Context, studentId, offerId, promocodeId primitive.ObjectID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, studentId, offerId, promocodeId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOrdersMockRecorder) Create(ctx, studentId, offerId, promocodeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrders)(nil).Create), ctx, studentId, offerId, promocodeId)
}

// AddTransaction mocks base method
func (m *MockOrders) AddTransaction(ctx context.Context, id primitive.ObjectID, transaction domain.Transaction) (domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransaction", ctx, id, transaction)
	ret0, _ := ret[0].(domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTransaction indicates an expected call of AddTransaction
func (mr *MockOrdersMockRecorder) AddTransaction(ctx, id, transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransaction", reflect.TypeOf((*MockOrders)(nil).AddTransaction), ctx, id, transaction)
}

// MockPayments is a mock of Payments interface
type MockPayments struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentsMockRecorder
}

// MockPaymentsMockRecorder is the mock recorder for MockPayments
type MockPaymentsMockRecorder struct {
	mock *MockPayments
}

// NewMockPayments creates a new mock instance
func NewMockPayments(ctrl *gomock.Controller) *MockPayments {
	mock := &MockPayments{ctrl: ctrl}
	mock.recorder = &MockPaymentsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPayments) EXPECT() *MockPaymentsMockRecorder {
	return m.recorder
}

// ProcessTransaction mocks base method
func (m *MockPayments) ProcessTransaction(ctx context.Context, callbackData payment.Callback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTransaction", ctx, callbackData)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessTransaction indicates an expected call of ProcessTransaction
func (mr *MockPaymentsMockRecorder) ProcessTransaction(ctx, callbackData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTransaction", reflect.TypeOf((*MockPayments)(nil).ProcessTransaction), ctx, callbackData)
}
