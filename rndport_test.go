package rndport

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAddress(t *testing.T) {
	port, err := GetAddress()
	require.NoError(t, err)
	require.IsType(t, "", port)
	_, err = strconv.Atoi(port)
	require.NoError(t, err)

	port2, err := GetAddress()
	require.NoError(t, err)
	require.NotEqual(t, port, port2)

	ip := "172.1.0.1"
	formatedAddress, err := GetAddress(fmt.Sprintf("%s:%%d", ip))
	require.NoError(t, err)
	addressPort := strings.Split(formatedAddress, ":")
	require.Equal(t, ip, addressPort[0])
	_, err = strconv.Atoi(addressPort[1])
	require.NoError(t, err)
}
