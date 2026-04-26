// Package config handles YAML configuration file operations.
package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// BackupManager handles creation of backup files.
type BackupManager struct {
	backupDir string
}

// NewBackupManager creates a new backup manager.
// If backupDir is empty, backups are created in the same directory as the original file.
func NewBackupManager(backupDir string) *BackupManager {
	return &BackupManager{
		backupDir: backupDir,
	}
}

// CreateBackup creates a backup of the specified file with a timestamp.
// Returns the path to the backup file.
func (b *BackupManager) CreateBackup(filePath string) (string, error) {
	// Check if source file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("source file does not exist: %s", filePath)
	}

	// Generate backup filename
	backupPath := b.generateBackupPath(filePath)

	// Create backup directory if needed
	backupDir := filepath.Dir(backupPath)
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create backup directory: %w", err)
	}

	// Copy file
	if err := copyFile(filePath, backupPath); err != nil {
		return "", fmt.Errorf("failed to create backup: %w", err)
	}

	return backupPath, nil
}

// generateBackupPath generates a backup file path with timestamp.
func (b *BackupManager) generateBackupPath(originalPath string) string {
	dir := filepath.Dir(originalPath)
	base := filepath.Base(originalPath)
	ext := filepath.Ext(base)
	name := base[:len(base)-len(ext)]

	timestamp := time.Now().Format("20060102_150405")

	if b.backupDir != "" {
		dir = b.backupDir
	}

	return filepath.Join(dir, fmt.Sprintf("%s.backup.%s%s", name, timestamp, ext))
}

// copyFile copies a file from src to dst.
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return err
	}

	// Copy file permissions
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(dst, info.Mode())
}

// CleanupOldBackups removes backup files older than the specified duration.
func (b *BackupManager) CleanupOldBackups(pattern string, maxAge time.Duration) error {
	// Implementation for cleanup - finds files matching pattern and removes old ones
	// This is a placeholder for future enhancement
	return nil
}

// DefaultBackupManager is a default instance for simple backup operations.
var DefaultBackupManager = NewBackupManager("")

// CreateDefaultBackup creates a backup using the default manager.
func CreateDefaultBackup(filePath string) (string, error) {
	return DefaultBackupManager.CreateBackup(filePath)
}
