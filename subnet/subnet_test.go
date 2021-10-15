package subnet

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

var testProviders map[string]*schema.Provider
var testProvider *schema.Provider

func init() {
	testProvider = Provider()
	testProviders = map[string]*schema.Provider{
		"subnet": testProvider,
	}
}

func Test_checkCIDR(t *testing.T) {
	cidr := "10.69.32.0/20"
	ip := "10.69.36.88"
	res, err := checkCIDR(cidr, ip)
	assert.NoError(t, err)
	assert.True(t, res)
}

func Test_MinMax(t *testing.T) {
	testList := []int{0, 0, 1}
	min, max, idxMin, idxMax := MinMax(testList)
	assert.Equal(t, 0, min)
	assert.Equal(t, 1, max)
	assert.Equal(t, 0, idxMin)
	assert.Equal(t, 2, idxMax)

	testList = []int{10, 1}
	min, max, idxMin, idxMax = MinMax(testList)
	assert.Equal(t, 1, min)
	assert.Equal(t, 10, max)
	assert.Equal(t, 1, idxMin)
	assert.Equal(t, 0, idxMax)
}

func Test_subnet_single(t *testing.T) {
	name := "data.subnet_single.test"
	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: `
data "subnet_single" "test" {
	cidr = "10.69.32.0/20"
	ip = "10.69.36.88"
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "cidr", "10.69.32.0/20"),
					resource.TestCheckResourceAttr(name, "ip", "10.69.36.88"),
					resource.TestCheckResourceAttr(name, "included", "true"),
				),
			},
		},
	})
}

func Test_subnet_list(t *testing.T) {
	name := "data.subnet_list.test"
	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: `
data "subnet_list" "test" {
	cidr_list = ["10.69.32.0/20","10.75.32.0/20"]
	ip = "10.69.36.88"
}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "cidr_list.0", "10.69.32.0/20"),
					resource.TestCheckResourceAttr(name, "cidr_list.1", "10.75.32.0/20"),
					resource.TestCheckResourceAttr(name, "cidr_list.#", "2"),
					resource.TestCheckResourceAttr(name, "ip", "10.69.36.88"),
					resource.TestCheckResourceAttr(name, "included", "true"),
					resource.TestCheckResourceAttr(name, "included_subnet_cidr", "10.69.32.0/20"),
					resource.TestCheckResourceAttr(name, "included_subnet_index", "0"),
				),
			},
		},
	})
}
