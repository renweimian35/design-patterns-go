// template 模板模式
package template

import "testing"

func TestTemplate(t *testing.T) {
	commonDriver := NewDriver(&CommonTemplate{})
	commonDriver.Template()
}
