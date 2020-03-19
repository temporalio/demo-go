package workflows

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"go.temporal.io/temporal/activity"

	"github.com/stretchr/testify/assert"
	"github.com/temporalio/temporal-go-demo/activities"
	_ "github.com/temporalio/temporal-go-demo/activities"
	"go.temporal.io/temporal/testsuite"
)

func TestWorkflowWithRealActivities(t *testing.T) {
	suite := testsuite.WorkflowTestSuite{}
	env := suite.NewTestWorkflowEnvironment()
	env.RegisterActivityWithOptions(activities.Deposit, activity.RegisterOptions{Name: "deposit"})
	env.RegisterActivityWithOptions(activities.Withdraw, activity.RegisterOptions{Name: "withdraw"})
	request := AccountTransferRequest{
		FromAccountId: "account1",
		ToAccountId:   "account2",
		ReferenceId:   "reference1",
		Amount:        1000,
	}
	env.ExecuteWorkflow(TransferWorkflow, request)
	assert.NoError(t, env.GetWorkflowError())
}

func TestWorkflowWithdrawFailure(t *testing.T) {
	suite := testsuite.WorkflowTestSuite{}
	env := suite.NewTestWorkflowEnvironment()
	env.RegisterActivityWithOptions(activities.Deposit, activity.RegisterOptions{Name: "deposit"})
	env.RegisterActivityWithOptions(activities.Withdraw, activity.RegisterOptions{Name: "withdraw"})

	env.OnActivity("withdraw", mock.Anything, "account1", "reference1", 1000).
		Return(errors.New("simulated failure")).
		Times(100)
	request := AccountTransferRequest{
		FromAccountId: "account1",
		ToAccountId:   "account2",
		ReferenceId:   "reference1",
		Amount:        1000,
	}
	env.ExecuteWorkflow(TransferWorkflow, request)
	assert.Error(t, env.GetWorkflowError())
	assert.Equal(t, "simulated failure", env.GetWorkflowError().Error())
}
