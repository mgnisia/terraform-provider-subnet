package subnet

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func checkCIDR(cidr, ip string) (b bool, err error) {
	_, ipv4Net, err := net.ParseCIDR(cidr)
	if err != nil {
		return false, err
	}
	return ipv4Net.Contains(net.ParseIP(ip)), err
}

func getValue(attribute string, d *schema.ResourceData, diag diag.Diagnostics) (string, error) {
	var v interface{}
	if v = d.Get(attribute); v != nil {
		return v.(string), nil
	} else {
		return "", fmt.Errorf("value %s couldn't be read from d", attribute)
	}
}

func getCount(attribute string, d *schema.ResourceData, diag diag.Diagnostics) (int, error) {
	var v interface{}
	if v = d.Get(attribute); v != nil {
		return v.(int), nil
	} else {
		return 0, fmt.Errorf("(int) value %s couldn't be read from d", attribute)
	}
}

func dataSubnetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidrValue, err := getValue("cidr", d, diags)
	if err != nil {
		diag.FromErr(err)
	}
	ipValue, err := getValue("ip", d, diags)
	if err != nil {
		diag.FromErr(err)
	}
	ipCheck, err := checkCIDR(cidrValue, ipValue)
	if err != nil {
		diag.FromErr(err)
	}
	if err := d.Set("included", ipCheck); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func dataSubnetListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidrCount, err := getCount("cidr_list.#", d, diags)
	if err != nil {
		diag.FromErr(err)
	}
	ipValue, err := getValue("ip", d, diags)
	if err != nil {
		diag.FromErr(err)
	}
checkList:
	for i := 0; i < cidrCount; i++ {
		cidrValue, err := getValue(fmt.Sprintf("cidr_list.%d", i), d, diags)
		if err != nil {
			diag.FromErr(err)
		}
		ipCheck, err := checkCIDR(cidrValue, ipValue)
		if err != nil {
			diag.FromErr(err)
		}
		if err := d.Set("included", ipCheck); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("included_subnet_cidr", cidrValue); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("included_subnet_index", i); err != nil {
			return diag.FromErr(err)
		}
		if ipCheck {
			break checkList
		}
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func MinMax(cidrlist []int) (min, max, idxMin, idxMax int) {
	max = cidrlist[0]
	min = cidrlist[0]
	for idx, value := range cidrlist {
		if max < value {
			max = value
			idxMax = idx
		} else if min > value {
			idxMin = idx
			min = value
		}
	}
	return min, max, idxMin, idxMax
}

func dataSubnetCompareRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	cidrCount, err := getCount("cidr_list.#", d, diags)
	if err != nil {
		diag.FromErr(err)
	}
	// Get largest cidr
	cidrSize := make([]int, 0)
	cidrs := make([]string, 0)
	for i := 0; i < cidrCount; i++ {
		cidrValueCurrent, err := getValue(fmt.Sprintf("cidr_list.%d", i), d, diags)
		if err != nil {
			return diag.FromErr(err)
		}
		_, net, err := net.ParseCIDR(cidrValueCurrent)
		if err != nil {
			return diag.FromErr(err)
		}
		size, _ := net.Mask.Size()
		cidrs = append(cidrs, cidrValueCurrent)
		cidrSize = append(cidrSize, size)
	}
	// To obtain the largest and lowest cidr in the list
	// we assume that comparing the subnet size e.g /20 vs /24 is sufficient
	_, _, idxMin, idxMax := MinMax(cidrSize)
	// the lowest means in the networking context the larget subnet
	if err := d.Set("cidr_largest", cidrs[idxMin]); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cidr_largest_index", idxMin); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cidr_lowest", cidrs[idxMax]); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cidr_lowest_index", idxMax); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
