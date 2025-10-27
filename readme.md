<!DOCTYPE html>
<html lang="fa" dir="rtl">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>PHPVM – Dockerized PHP Environment</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Tahoma", sans-serif;
            line-height: 1.6;
            background-color: #f9f9f9;
            color: #333;
            max-width: 900px;
            margin: 20px auto;
            padding: 25px;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
        }
        h1, h2, h3 {
            border-bottom: 2px solid #eee;
            padding-bottom: 10px;
            color: #222;
        }
        h1 {
            font-size: 2.5em;
        }
        h2 {
            font-size: 2em;
            margin-top: 40px;
        }
        h3 {
            font-size: 1.5em;
            border-bottom: 1px solid #eee;
        }
        ul {
            padding-right: 20px;
        }
        li {
            margin-bottom: 10px;
        }
        code {
            font-family: "Courier New", Courier, monospace;
            background-color: #eee;
            padding: 2px 6px;
            border-radius: 4px;
            direction: ltr;
            display: inline-block;
        }
        pre {
            background-color: #2d2d2d;
            color: #f1f1f1;
            padding: 15px;
            border-radius: 5px;
            overflow-x: auto;
            direction: ltr;
        }
        pre code {
            background-color: transparent;
            padding: 0;
            color: inherit;
        }
        blockquote {
            background-color: #f0f7ff;
            border-right: 5px solid #007bff;
            margin-right: 0;
            padding: 15px 20px;
            font-style: italic;
        }
        video {
            display: block;
            margin: 20px 0;
            border-radius: 5px;
            background-color: #000;
        }
        hr {
            border: 0;
            border-top: 2px solid #eee;
            margin: 40px 0;
        }
    </style>
</head>
<body>

    <h1>PHPVM – Dockerized PHP Environment</h1>

    <blockquote>
        <p>این برنامه تماما با ChatGPT تولید شده و هدفش استفاده راحت از PHP نسخه ۸ و بالاتر با Docker است.</p>
    </blockquote>

    <hr>

    <h2>🚀 ویژگی‌ها</h2>
    <ul>
        <li>استفاده از <strong>Docker</strong> برای راه‌اندازی سریع PHP ۸+ بدون نصب مستقیم روی سیستم.</li>
        <li>دستورات محلی <strong>php</strong>, <strong>composer</strong>, <strong>composerexec</strong> مستقیماً در پوشه پروژه فعال هستند.</li>
        <li><strong>کانفیگ انعطاف‌پذیر</strong>: می‌توانید فایل‌های پروژه را تغییر دهید و دوباره build کنید؛ برنامه انتزاع جدید نمی‌سازد.</li>
        <li>پورت‌های <strong>xdebug و FPM</strong> روی ماشین لوکال قابل دسترسی هستند.</li>
        <li>اتصال به سرویس‌های داخلی لوکال با استفاده از <code>host.docker.internal</code>.</li>
        <li>پورت <strong>8000</strong> به <strong>8007</strong> مپ شده؛ قابل تغییر در <code>docker-compose.yml</code>.</li>
        <li>دستور داده شده بعد از پایان نصب رو روی شلتون اعمال کنید این یه بار نیازه.</li>
    </ul>

    <hr>

    <h2>⚙️ نصب و اجرا</h2>
    <h3>Linux (تست شده)</h3>

<pre><code class="language-bash"># Build برنامه
go build -o phpvm

# اجرا در مسیر پروژه
./phpvm
</code></pre>

    <hr>

    <h2>🎥 مشاهده ویدیو</h2>
    <p>برای مشاهده توضیحات بیشتر در مورد نحوه استفاده و اجرای PHPVM، می‌توانید ویدیوی آموزشی را تماشا کنید:</p>

    <video width="100%" style="max-width: 800px;" controls>
      <source src="phpvm.mp4" type="video/mp4">
      مرورگر شما از پخش این ویدیو پشتیبانی نمی‌کند.
    </video>

    <p>
      <strong>📌 ویدیو مربوط به راه‌اندازی و استفاده از PHPVM است.</strong>
    </p>

    <hr>

    <h2>🔹 نکات مهم</h2>
    <ul>
        <li>تمام فایل‌ها و کانتینرها تحت مالکیت کاربر فعلی ساخته می‌شوند.</li>
        <li>فایل‌های موجود <strong>overwrite نمی‌شوند</strong> مگر شما بخواهید.</li>
        <li>کانفیگ‌ها و <code>.env</code> کاملاً در اختیار شما هستند و می‌توانید تغییر دهید و دوباره build بگیرید.</li>
        <li>Linux تست شده و پایدار است؛ Windows هنوز تست نشده است.</li>
    </ul>

    <hr>

    <h2>🎉 موفق باشید!</h2>

</body>
</html>
