package consts

import "testing"

func TestBuildErrorCode(t *testing.T) { //
	expectedCode := int64(123)
	expectedMsg := "Test error message"
	err := buildErrorCode(expectedCode, expectedMsg)

	if err.Code != expectedCode {
		t.Errorf("Expected error code %d, but got %d", expectedCode, err.Code)
	}

	if err.Msg != expectedMsg {
		t.Errorf("Expected error message '%s', but got '%s'", expectedMsg, err.Msg)
	}
}

//test 函数使用错误代码和错误消息的预期值调用函数，然后使用语句检查返回的结构是否具有预期值。
//如果任何语句失败，测试函数将使用该函数报告错误。
