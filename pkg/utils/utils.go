/************************************************************
 * Author:        jackey
 * Date:        2022/11/24
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// GetNowDateTime 获取当前时间字符串
// 格式：2006-01-02 15:04:05
func GetNowDateTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	result := tm.Format("2006-01-02 15:04:05")
	return result
}

// DeleteSliceString 删除字符串切片指定元素。
func DeleteSliceString(a []string, elem string) []string {
	j := 0
	for _, v := range a {
		if v != elem {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

// CheckInStringSlice 查看字符串是否在切片内
func CheckInStringSlice(a []string, elem string) bool {
	for _, val := range a {
		if val == elem {
			return true
		}
	}
	return false
}

// GetRange 获取随机数
func GetRange(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}

// MD5 生成32位MD5
func MD5(text string) string {
	if text == "" {
		return ""
	}
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// WriteFile 写入文本文件，追加模式，如果不存在自动创建文件
func WriteFile(path string, data string) error {
	dir := filepath.Dir(path)
	_ = CreateMutiDir(dir)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes := []byte(data + "\n")
	_, err = file.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

// CreateMutiDir 调用os.MkdirAll递归创建文件夹
func CreateMutiDir(filePath string) error {
	_, err := os.Stat(filePath) //os.Stat获取文件信息
	if err != nil && !os.IsExist(err) {
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			fmt.Println("CreateMutiDir,error info:", err)
			return err
		}
	}
	return nil
}
