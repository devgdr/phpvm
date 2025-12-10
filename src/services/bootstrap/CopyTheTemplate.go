package bootstrap

import (
	"io/fs"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"

	"phpvm/src/ui/TemplateDir"
)

// CopyTheTemplate copies all embedded template files into the project folder
// ensuring the current user owns the files (on Linux/macOS).
func CopyTheTemplate(projectPathDirectory string) {
	// گرفتن uid و gid کاربر فعلی (Linux/macOS)
	var uid, gid int
	if runtime.GOOS != "windows" {
		usr, err := user.Current()
		if err != nil {
			log.Println("⚠ Warning: could not get current user info:", err)
		} else {
			uid, _ = strconv.Atoi(usr.Uid)
			gid, _ = strconv.Atoi(usr.Gid)
		}
	}

	err := fs.WalkDir(TemplateDir.TemplateFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		targetPath := filepath.Join(projectPathDirectory, path)

		if d.IsDir() {
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return err
			}
			// تنظیم مالکیت فقط در Linux/macOS
			if runtime.GOOS != "windows" && uid != 0 && gid != 0 {
				_ = os.Chown(targetPath, uid, gid)
			}
		} else {
			// اگر فایل موجود است، آن را overwrite نکن
			if _, err := os.Stat(targetPath); err == nil {
				log.Println("⚠ Skipping existing file:", targetPath)
				return nil
			}

			data, err := TemplateDir.TemplateFS.ReadFile(path)
			if err != nil {
				return err
			}

			if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
				return err
			}

			if err := os.WriteFile(targetPath, data, 0644); err != nil {
				return err
			}

			// تنظیم مالکیت فقط در Linux/macOS
			if runtime.GOOS != "windows" && uid != 0 && gid != 0 {
				_ = os.Chown(targetPath, uid, gid)
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error copying template: %v", err)
	}

	// ایجاد پوشه App اگر وجود ندارد
	appPath := filepath.Join(projectPathDirectory, "App")
	if err := os.MkdirAll(appPath, 0755); err != nil {
		log.Fatalf("Error creating App folder: %v", err)
	}
	if runtime.GOOS != "windows" && uid != 0 && gid != 0 {
		_ = os.Chown(appPath, uid, gid)
	}

	log.Println("✅ Template copied successfully into:", projectPathDirectory)
}
