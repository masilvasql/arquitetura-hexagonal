package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/masilvasql/go-hexagonal/adapters/cli"
	mock_application "github.com/masilvasql/go-hexagonal/application/mocks"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 10.0
	productStatus := "enabled"
	productId := uuid.NewV4().String()

	productMock := mock_application.NewMockProductInterface(ctrl)

	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	productServiceMock := mock_application.NewMockProductServiceInterface(ctrl)
	productServiceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s created with the name %s has been created with price %f and status %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)

	result, err := cli.Run(productServiceMock, "create", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s with the name %s has been enabled with price %f and status %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)

	result, err = cli.Run(productServiceMock, "enable", productId, "", 0.0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s with the name %s has been disabled with price %f and status %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)

	result, err = cli.Run(productServiceMock, "disable", productId, "", 0.0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s with the name %s has been disabled with price %f and status %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)

	result, err = cli.Run(productServiceMock, "get", productId, "", 0.0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}
