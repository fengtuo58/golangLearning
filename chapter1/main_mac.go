/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2019 WireGuard LLC. All Rights Reserved.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/ipc"
	"golang.zx2c4.com/wireguard/tun"
)

const (
	ExitSetupSuccess = 0
	ExitSetupFailed  = 1
)

const (
	ENV_WG_TUN_FD             = "WG_TUN_FD"
	ENV_WG_UAPI_FD            = "WG_UAPI_FD"
	ENV_WG_PROCESS_FOREGROUND = "WG_PROCESS_FOREGROUND"
)

const testInput = `private_key=882ecf8a353056fe35ed674f0b7aaf204f2f278b04f3c72ac8ec9afb73f9494d
replace_peers=true
public_key=dd0a6cf10ef61f5a5e998de6a04ad75644035263d9d6a2e85a91472db8097051
endpoint=42.159.91.157:51820
persistent_keepalive_interval=10
replace_allowed_ips=true
allowed_ip=0.0.0.0/0`

var dict map[string]string //保存原dns

func execCmd(cmdStr string) (res string, err error) {
	args := strings.Split(cmdStr, " ")
	cmd := exec.Command(args[0], args[1:]...)
	ret, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func isIp(ip string) (b bool) {
	if m, _ := regexp.MatchString("^(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)$", ip); !m {
		return false
	}
	return true
}

func getGateway() (gateway string) {
	strCmd := "netstat -nr -f inet"
	response, err := execCmd(strCmd)
	if err != nil {
		return
	}
	lines := strings.Split(response, "\n")
	str := ""
	for _, line := range lines {
		pound := strings.Index(line, "default")
		if pound < 0 {
			continue
		}
		str = strings.TrimSpace(line)
		break
	}
	if str != "" {
		lines := strings.Split(str, " ")
		for _, line := range lines {
			if line == "" {
				continue
			}

			if isIp(line) {
				return line
			}
		}
	}
	return ""
}

func getOldDns() error {
	dict = make(map[string]string)
	strCmd := "networksetup -listallnetworkservices"
	response, err := execCmd(strCmd)
	if err != nil {
		return err
	}
	lines := strings.Split(response, "\n")
	for _, line := range lines {
		pound := strings.Index(line, "is disabled")
		if pound > 0 || line == "" {
			continue
		}
		line = strings.TrimSpace(line)
		strCmd := "networksetup"

		cmd := exec.Command(strCmd, "-getdnsservers", line)
		ret, err := cmd.Output()
		response = string(ret)
		if err != nil {
			dict[line] = "empty"
		} else {
			pound := strings.Index(response, line)
			if pound < 0 {
				dict[line] = strings.TrimRight(response, "\n")
			} else {
				dict[line] = "empty"
			}
		}
	}
	return nil
}

func setDns(dnsArray ...string) error {
	var args = []string{"networksetup", "-setdnsservers", ""}
	args = append(args, dnsArray...)
	for service := range dict {
		args[2] = service
		cmd := exec.Command(args[0], args[1:]...)
		_, err := cmd.Output()
		if err != nil {
			return err
		}
	}
	return nil
}

func recoderDns() error {
	setDns("empty")
	var args = []string{"networksetup", "-setdnsservers", ""}
	for service := range dict {
		args[2] = service
		args = append(args, strings.Split(dict[service], "\n")...)
		cmd := exec.Command(args[0], args[1:]...)
		_, err := cmd.Output()
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	interfaceName := "utun"

	logLevel := func() int {
		switch os.Getenv("LOG_LEVEL") {
		case "debug":
			return device.LogLevelDebug
		case "info":
			return device.LogLevelInfo
		case "error":
			return device.LogLevelError
		case "silent":
			return device.LogLevelSilent
		}
		return device.LogLevelInfo
	}()

	device.RoamingDisabled = false //mac:false  ios:true

	tun, err := tun.CreateTUN(interfaceName, device.DefaultMTU)
	if err == nil {
		realInterfaceName, err2 := tun.Name()
		if err2 == nil {
			interfaceName = realInterfaceName
		}
	}

	logger := device.NewLogger(
		logLevel,
		fmt.Sprintf("(%s) ", interfaceName),
	)

	logger.Info.Println("Starting wireguard-go version", device.WireGuardGoVersion)

	logger.Debug.Println("Debug log enabled")

	if err != nil {
		logger.Error.Println("Failed to create TUN device:", err)
		os.Exit(ExitSetupFailed)
	}

	fileUAPI, err := func() (*os.File, error) {
		uapiFdStr := os.Getenv(ENV_WG_UAPI_FD)
		if uapiFdStr == "" {
			return ipc.UAPIOpen(interfaceName)
		}

		fd, err := strconv.ParseUint(uapiFdStr, 10, 32)
		if err != nil {
			return nil, err
		}

		return os.NewFile(uintptr(fd), ""), nil
	}()

	if err != nil {
		logger.Error.Println("UAPI listen error:", err)
		os.Exit(ExitSetupFailed)
		return
	}

	device := device.NewDevice(tun, logger)
	logger.Info.Println("Device started")
	logger.Info.Println(interfaceName)
	uapi, err := ipc.UAPIListen(interfaceName, fileUAPI)
	if err != nil {
		logger.Error.Println("Failed to listen on uapi socket:", err)
		os.Exit(ExitSetupFailed)
	}

	setError := device.IpcSetOperation(bufio.NewReader(strings.NewReader(testInput)))
	if setError != nil {
		logger.Error.Println(setError)
		return
	}

	device.Up()

	errs := make(chan error)
	term := make(chan os.Signal, 1)

	go func() {
		for {
			conn, err := uapi.Accept()
			if err != nil {
				errs <- err
				return
			}
			go device.IpcHandle(conn)
		}
	}()

	logger.Info.Println("UAPI listener started")
	var strCmd string
	addressIp := "10.253.0.4"
	strCmd = fmt.Sprintf("ifconfig %s inet %s/32 %s alias",
		interfaceName, addressIp, addressIp)
	execCmd(strCmd)

	strCmd = fmt.Sprintf("%s up", interfaceName)
	execCmd(strCmd)

	strCmd = fmt.Sprintf("route -q -n add -inet 0.0.0.0/1 -interface %s",
		interfaceName)
	execCmd(strCmd)

	strCmd = fmt.Sprintf("route -q -n add -inet 128.0.0.0/1 -interface %s",
		interfaceName)
	execCmd(strCmd)
	serverIp := "42.159.91.157"
	gateway := getGateway()
	logger.Info.Println(gateway)
	strCmd = fmt.Sprintf("route -q -n add -inet %s -gateway %s",
		serverIp, gateway)
	execCmd(strCmd)

	//设置DNS
	err = getOldDns()
	if err != nil {
		return
	}

	err = setDns("1.1.1.1", "8.8.8.8")
	if err != nil {
		return
	}

	// wait for program to terminate

	signal.Notify(term, os.Interrupt)
	signal.Notify(term, os.Kill)
	signal.Notify(term, syscall.SIGTERM)

	select {
	case <-term:
	case <-errs:
	case <-device.Wait():
	}

	// clean up

	uapi.Close()
	device.Close()

	//恢复dns
	recoderDns()

	logger.Info.Println("Shutting down")
}
