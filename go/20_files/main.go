package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 📘 filesOp → get file info (name, size, modified time)
func filesOp(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("❌ Error while opening file:", err)
		return
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println("❌ Error reading file info:", err)
		return
	}

	fmt.Printf("📁 File Name: %s\n", fileInfo.Name())
	fmt.Printf("📏 Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("🕒 Modified: %v\n", fileInfo.ModTime())

	// Detect if file or directory
	if fileInfo.IsDir() {
		fmt.Println("📂 Type: Directory")
	} else {
		fmt.Println("📄 Type: Regular File")
	}
}

// 📖 fileRead → read entire file manually using os.Open + .Read()
func fileRead(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("❌ Error while opening file for reading:", err)
		return
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println("❌ Error reading file info:", err)
		return
	}

	buf := make([]byte, fileInfo.Size())
	_, err = f.Read(buf)
	if err != nil {
		fmt.Println("❌ Error reading file:", err)
		return
	}

	fmt.Println("📄 File Extension:", filepath.Ext(fileInfo.Name()))
	fmt.Println("📖 File Content:")
	fmt.Println(string(buf))
}

// 🧠 ReadFileOS → simplified file reading using os.ReadFile
func ReadFileOS(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("❌ Error reading file:", err)
		return
	}
	fmt.Println("📚 File read using os.ReadFile():")
	fmt.Println(string(data))
}

// 📂 copyFileToAnotherFile → copy content from one file to another
func copyFileToAnotherFile(srcPath string) {
	destPath := "./copied_example.txt"

	data, err := os.ReadFile(srcPath)
	if err != nil {
		fmt.Println("❌ Error reading source file:", err)
		return
	}

	err = os.WriteFile(destPath, data, 0644)
	if err != nil {
		fmt.Println("❌ Error writing destination file:", err)
		return
	}

	fmt.Println("✅ File copied successfully to:", destPath)
}

func main() {
	filePath := "./example.txt"

	// ✍️ Create or truncate file
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("❌ Error while creating the file:", err)
		return
	}
	defer file.Close()

	fmt.Println("✅ File created successfully:", file.Name())

	// Write to file
	_, err = file.WriteString("Hello GOLANG file write via code\n")
	if err != nil {
		fmt.Println("❌ Unable to write text into file:", err)
		return
	}
	fmt.Println("✍️ File written successfully.")

	fmt.Println()
	filesOp(filePath) // get info

	fmt.Println()
	fileRead(filePath) // manual read

	fmt.Println()
	ReadFileOS(filePath) // simplified read

	fmt.Println()
	copyFileToAnotherFile(filePath) // copy file
}
