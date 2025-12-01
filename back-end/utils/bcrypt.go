/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 11:06:06
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 17:31:05
 * @FilePath: \ginManager\utils\bcrypt.go
 * @Description: bcrypt 密码加密工具
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */

package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

const (
	lower   = "abcdefghijklmnopqrstuvwxyz"
	upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
	special = "!@#$^&*()-_=+[]{}<>?"
	all     = lower + upper + digits + special
)

// 生成哈希（存库）
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 校验密码（登录时）
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)
	return err == nil
}

// GeneratePassword 生成指定长度的强密码
func GeneratePassword(length int) (string, error) {
	if length < 4 {
		return "", fmt.Errorf("length must be at least 4 to include all character types")
	}

	// 确保每种类型至少出现一次
	password := make([]byte, length)
	charsets := []string{lower, upper, digits, special}

	for i, chars := range charsets {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		password[i] = chars[n.Int64()]
	}

	// 剩余位随机填充
	for i := len(charsets); i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(all))))
		if err != nil {
			return "", err
		}
		password[i] = all[n.Int64()]
	}

	// 打乱顺序
	return shuffle(password), nil
}

// shuffle 打乱字节切片
func shuffle(src []byte) string {
	dst := make([]byte, len(src))
	copy(dst, src)

	for i := len(dst) - 1; i > 0; i-- {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		dst[i], dst[j.Int64()] = dst[j.Int64()], dst[i]
	}
	return string(dst)
}
