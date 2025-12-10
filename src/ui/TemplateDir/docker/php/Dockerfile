# Allow selecting the PHP base image at build time
ARG PHP_BASE
FROM $PHP_BASE

# Update packages and install dependencies + SSH server
RUN apt-get update && apt-get install -y \
    openssh-server \
    git \
    curl \
    zip \
    vim \
    unzip \
    libicu-dev \
    bash \
    sudo \
    libzip-dev \
    libpng-dev \
    libjpeg-dev \
    libfreetype6-dev \
    libonig-dev \
    libxml2-dev \
    libcurl4-openssl-dev \
    libssl-dev

# Install Composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

# Install PHP extensions
RUN docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install -j$(nproc) \
    pdo \
    pdo_mysql \
    mysqli \
    intl \
    opcache \
    zip \
    gd \
    bcmath \
    soap \
    mbstring \
    exif \
    pcntl \
    sockets \
    xml

# Install Xdebug
RUN pecl install xdebug
COPY ./xdebug.ini "${PHP_INI_DIR}/conf.d/xdebug.ini"

# تنظیم SSH (قبل از ساخت user)
RUN mkdir -p /var/run/sshd /run/sshd \
    && sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/' /etc/ssh/sshd_config \
    && sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin no/' /etc/ssh/sshd_config \
    && sed -i 's/UsePAM yes/UsePAM no/' /etc/ssh/sshd_config

# Create a non-root user
ARG UID
ARG GID
RUN groupadd -g $GID developer && \
    useradd -u $UID -g developer -d /home/developer -s /bin/bash -m developer && \
    chown -R developer:developer /home/developer && \
    usermod -aG sudo developer


RUN echo "developer ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers


RUN mkdir -p /var/log/xdebug && \
    chown -R developer:developer /var/log/xdebug






User developer

# Set working directory
WORKDIR /var/www/html
