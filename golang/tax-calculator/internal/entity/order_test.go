package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_Get_An_Error_If_Id_Is_Blank(t *testing.T) {

	order := Order{}
	//err := order.IsValid()

	// if err == nil {
	// 	t.Error("Expected error, nothing happened feijoada")
	// } without testify

	assert.Error(t, order.IsValid(), "Expected error, nothing happened feijoada")
}
