package repositories

// import (
// 	"database/sql"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/jmoiron/sqlx"

// 	rmock "svc-user-management/mocks/repositories"

// 	. "github.com/onsi/gomega"
// )

// type testUser struct {
// 	module       *Opts
// 	moduleMock   *rmock.UserRepositories
// 	masterdb     *sqlx.DB
// 	masterdbMock sqlmock.Sqlmock
// }

// func doTestUser(t *testing.T, fn func(*GomegaWithT, *testUser)) {
// 	masterDB, masterDBMock, _ := sqlmock.New()
// 	mdb := sqlx.NewDb((*sql.DB)(masterDB), "sqlmock")

// 	o := testUser{
// 		module:       &Opts{PGM: nil},
// 		moduleMock:   &rmock.UserRepositories{},
// 		masterdb:     mdb,
// 		masterdbMock: masterDBMock,
// 	}

// 	g := NewGomegaWithT(t)
// 	fn(g, &o)

// 	if err := o.masterdbMock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func Test_NewUserRepositories(t *testing.T) {
// 	doTestUser(t, func(gwt *GomegaWithT, tu *testUser) {
// 		inst := New(tu.module)
// 		gwt.Expect(inst).ShouldNot(BeNil())
// 	})
// }
