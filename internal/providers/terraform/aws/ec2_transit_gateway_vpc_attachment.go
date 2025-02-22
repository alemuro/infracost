package aws

import (
	"github.com/infracost/infracost/internal/schema"

	"github.com/shopspring/decimal"
)

func GetEC2TransitGatewayVpcAttachmentRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_ec2_transit_gateway_vpc_attachment",
		RFunc: NewEC2TransitGatewayVpcAttachment,
		ReferenceAttributes: []string{
			"transit_gateway_id",
		},
	}
}

func NewEC2TransitGatewayVpcAttachment(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	region := d.Get("region").String()
	transitGatewayRefs := d.References("transit_gateway_id")
	if len(transitGatewayRefs) > 0 {
		region = transitGatewayRefs[0].Get("region").String()
	}

	var gbDataProcessed *decimal.Decimal

	if u != nil && u.Get("monthly_gb_data_processed").Exists() {
		gbDataProcessed = decimalPtr(decimal.NewFromFloat(u.Get("monthly_gb_data_processed").Float()))
	}

	return &schema.Resource{
		Name: d.Address,
		CostComponents: []*schema.CostComponent{
			transitGatewayAttachmentCostComponent(region, "TransitGatewayVPC"),
			transitGatewayDataProcessingCostComponent(region, "TransitGatewayVPC", gbDataProcessed),
		},
	}
}

func transitGatewayAttachmentCostComponent(region string, operation string) *schema.CostComponent {
	return &schema.CostComponent{
		Name:           "Transit gateway attachment",
		Unit:           "hours",
		UnitMultiplier: 1,
		HourlyQuantity: decimalPtr(decimal.NewFromInt(1)),
		ProductFilter: &schema.ProductFilter{
			VendorName: strPtr("aws"),
			Region:     strPtr(region),
			Service:    strPtr("AmazonVPC"),
			AttributeFilters: []*schema.AttributeFilter{
				{Key: "usagetype", ValueRegex: strPtr("/TransitGateway-Hours/")},
				{Key: "operation", Value: strPtr(operation)},
			},
		},
	}
}

func transitGatewayDataProcessingCostComponent(region string, operation string, gbDataProcessed *decimal.Decimal) *schema.CostComponent {
	return &schema.CostComponent{
		Name:           "Data processed",
		Unit:           "GB",
		UnitMultiplier: 1,
		HourlyQuantity: gbDataProcessed,
		ProductFilter: &schema.ProductFilter{
			VendorName: strPtr("aws"),
			Region:     strPtr(region),
			Service:    strPtr("AmazonVPC"),
			AttributeFilters: []*schema.AttributeFilter{
				{Key: "usagetype", ValueRegex: strPtr("/TransitGateway-Bytes/")},
				{Key: "operation", Value: strPtr(operation)},
			},
		},
	}
}
