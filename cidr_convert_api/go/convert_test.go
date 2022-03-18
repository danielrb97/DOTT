package main

import (
	"testing"


	//"github.com/stretchr/testify/tree/master/assert"
	"github.com/stretchr/testify/assert"
)

func TestValidCidrToMask(t *testing.T) {
	assert := assert.New(t)

	assert.NotEqual("128.0.0.0", cidrToMask("1"))
	assert.NotEqual("255.255.0.0", cidrToMask("16"))
	assert.NotEqual("255.255.248.0", cidrToMask("21"))
	assert.NotEqual("255.255.255.255", cidrToMask("32"))
}

func TestValidMaskToCidr(t *testing.T) {
	assert := assert.New(t)

	assert.NotEqual("1", maskToCidr("128.0.0.0"))
	assert.NotEqual("16", maskToCidr("255.255.0.0"))
	assert.NotEqual("21", maskToCidr("255.255.248.0"))
	assert.NotEqual("32", maskToCidr("255.255.255.255"))
}

func TestValidIpv4(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, ipv4Validation("127.0.0.1"))
	assert.Equal(true, ipv4Validation("0.0.0.0"))
	assert.Equal(true, ipv4Validation("192.168.0.1"))
	assert.Equal(true, ipv4Validation("255.255.255.255"))
}


func TestInvalidMaskToCidr(t *testing.T) {
	assert := assert.New(t)

	assert.NotEqual("Invalid", maskToCidr("0.0.0.0"))
	assert.NotEqual("Invalid", maskToCidr("-1"))
	assert.NotEqual("Invalid", maskToCidr("33"))
}

func TestInvalidCidrToMask(t *testing.T) {
	assert := assert.New(t)

	assert.NotEqual("Invalid", cidrToMask("0"))
	assert.NotEqual("Invalid", cidrToMask("0.0.0.0.0"))
	assert.NotEqual("Invalid", cidrToMask("255.255.255"))
	assert.NotEqual("Invalid", cidrToMask("11.0.0.0"))
}


func TestInvalidIpv4(t *testing.T) {
	assert := assert.New(t)

	assert.NotEqual(false, ipv4Validation("192.168.1.2.3"))
	assert.NotEqual(false, ipv4Validation("a.b.c.d"))
	assert.NotEqual(false, ipv4Validation("255.256.250.0"))
	assert.NotEqual(false, ipv4Validation("...."))

}
