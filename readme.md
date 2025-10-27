# PHPVM – Dockerized PHP Environment

> این برنامه تماما با ChatGPT تولید شده و هدفش استفاده راحت از PHP نسخه ۸ و بالاتر با Docker است.

---

## 🚀 ویژگی‌ها

- استفاده از **Docker** برای راه‌اندازی سریع PHP ۸+ بدون نصب مستقیم روی سیستم.
- دستورات محلی **php**, **composer**, **composerexec** مستقیماً در پوشه پروژه فعال هستند.
- **کانفیگ انعطاف‌پذیر**: می‌توانید فایل‌های پروژه را تغییر دهید و دوباره build کنید؛ برنامه انتزاع جدید نمی‌سازد.
- پورت‌های **xdebug و FPM** روی ماشین لوکال قابل دسترسی هستند.
- اتصال به سرویس‌های داخلی لوکال با استفاده از `host.docker.internal`.
- پورت **8000** به **8007** مپ شده؛ قابل تغییر در `docker-compose.yml`.
-دستور داده شده بعد از پایان نصب رو روی شلتون اعمال کنید این یه بار نیازه.

---

## ⚙️ نصب و اجرا

### Linux (تست شده)

```bash
# Build برنامه
go build -o phpvm

# اجرا در مسیر پروژه
./phpvm

## 🎥 مشاهده ویدیو


<p align="center">
  <strong>برای مشاهده توضیحات، روی تصویر زیر کلیک کنید:</strong>
  <br>
  <a href="phpvm.mp4" title="Click to watch video">
    <img src="thumbnail.png" alt="PHPVM Video Thumbnail" style="max-width:800px; border-radius:10px; box-shadow: 0 4px 10px rgba(0,0,0,0.2);">
  </a>
</p>

> 📌 ویدیو مربوط به راه‌اندازی و استفاده از PHPVM است.
