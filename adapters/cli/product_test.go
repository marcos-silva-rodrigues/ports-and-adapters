package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/marcos-silva-rodrigues/go-hexagonal/adapters/cli"
	mock_application "github.com/marcos-silva-rodrigues/go-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	// Product Create Command
	resultExpect := fmt.Sprintf("Produc ID %s with the name %s has been created with price %f and status %s",
		productId, productName, productPrice, productStatus)

	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpect, result)

	// Product Enable Command
	resultExpect = fmt.Sprintf("Product %s has been enabled.", productName)

	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpect, result)

	// Product Disable Command
	resultExpect = fmt.Sprintf("Product %s has been disabled.", productName)

	result, err = cli.Run(service, "disabled", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpect, result)

	// Product Default Command
	resultExpect = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productId, productName, productPrice, productStatus)

	result, err = cli.Run(service, "", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpect, result)

}
