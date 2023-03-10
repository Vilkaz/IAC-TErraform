package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	)

func TestTerraformAwsExample(t *testing.T) {
	t.Parallel()

	// Generate a random name to avoid naming conflicts
	uniqueID := random.UniqueId()

	// Give the example Terraform code an AWS region to test in
	awsRegion := "eu-central-1"

	// Give the example Terraform code a name prefix for all resources
	expectedNamePrefix := "example-" + uniqueID

	// Construct the Terraform options with the test folder path, VPC name, and AWS region
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"name_prefix": expectedNamePrefix,
			"aws_region":  awsRegion,
		},
	}

	// At the end of the test, destroy the Terraform resources using the terraformOptions
	defer terraform.Destroy(t, terraformOptions)

	// Deploy the example Terraform code
	terraform.InitAndApply(t, terraformOptions)

	// Check that the EC2 instance exists
	instanceID := terraform.Output(t, terraformOptions, "instance_id")
	instance := aws.GetEc2Instance(t, awsRegion, instanceID)
	assert.Equal(t, instanceID, *instance.InstanceId)

	// Check that the instance is running
	assert.Equal(t, "running", *instance.State.Name)

	// Check that the instance has the expected name tag
	nameTag := "Name"
	expectedNameTagValue := expectedNamePrefix + "-web"
	actualNameTagValue := aws.GetEc2InstanceTag(t, awsRegion, instanceID, nameTag)
	assert.Equal(t, expectedNameTagValue, actualNameTagValue)

	// Check that the instance has the expected public IP address
	expectedPublicIP := terraform.Output(t, terraformOptions, "public_ip")
	actualPublicIP := *instance.PublicIpAddress
	assert.Equal(t, expectedPublicIP, actualPublicIP)

	// Check that the instance has the expected NGINX server installed and running
	publicURL := "http://" + actualPublicIP
	assert.Eventually(t, func() bool {
		response, err := aws.GetUrlE(t, publicURL)
		if err != nil {
			return false
		}
		return response.StatusCode == 200
	}, 5*time.Minute, 10*time.Second)

	// Check that the instance has the expected security group rules
	expectedIngress := []aws.IpPermission{
		{
			FromPort:   aws.Int(80),
			ToPort:     aws.Int(80),
			IpProtocol: aws.String("tcp"),
			IpRanges: []aws.IpRange{
				{
					CidrIp: aws.String("0.0.0.0/0"),
				},
			},
		},
	}
	actualIngress := aws.GetSecurityGroupIngress(t, awsRegion, expectedNamePrefix+"-sg")
	assert.ElementsMatch(t, expectedIngress, actualIngress)
}
