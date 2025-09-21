# V2Ray Keenetic Control Panel

*English below / Türkçe versiyonu aşağıdadır.*

---

## EN

This project is a V2Ray proxy application designed specifically for Keenetic routers. It provides a web-based management interface to control the V2Ray service on the router.

## ✨ Features

- 🌐 **Web Management Interface**: Intuitive Vue.js-based control panel
- 🔄 **Multiple Protocols**: VMess, VLESS, Shadowsocks, Trojan support
- 📱 **QR Code Import**: Easy server configuration via QR codes
- 📊 **Real-time Monitoring**: Connection status and traffic statistics
- 🔒 **Secure**: HTTPS support and authentication
- 🌍 **Multi-language**: English and Turkish support
- 📱 **Mobile Friendly**: Responsive design for all devices
- 🚀 **Auto Updates**: Automatic configuration updates
- ⚖️ **Load Balancing**: Multiple server support with failover

## 📋 Prerequisites

- Keenetic router with shell access (SSH/Telnet)
- v2ray-core binary installed and accessible in PATH
- At least 32MB free storage space
- Internet connection for initial setup

## 🔧 Supported Router Models

- Keenetic Giga (KN-1010/1011)
- Keenetic Ultra (KN-1810)
- Keenetic Titan (KN-1820)
- Keenetic Hero 4G (KN-2310)
- And other NDMS-based models

## 🚦 Quick Start

### 1. Check Prerequisites
```bash
# Verify v2ray-core installation
v2ray --version

# Check available space
df -h
```

### 2. Download and Install
```bash
# Download for your architecture (replace with actual version)
wget https://github.com/halibiram/keenetiray/releases/latest/download/keenetiray_linux_arm.tar.gz -O /tmp/keenetiray.tar.gz

# Create installation directory
mkdir -p /opt/keenetiray
cd /opt/keenetiray

# Extract and set permissions
tar -xzf /tmp/keenetiray.tar.gz
chmod +x keenetiray

# Run the application
./keenetiray
```

### 3. Access Web Interface
Open your browser and navigate to: `http://[ROUTER_IP]:8888`

Default credentials:
- Username: `admin`
- Password: `admin123` (change immediately!)

### Project Structure
*   `/cmd/v2ray-keenetic`: The Go backend application (API server).
*   `/web`: The Vue.js frontend application (Web UI).
*   ... and other development folders.

### Development
This project is in the early stages of development. To run it locally for development, follow these steps.

#### Backend
```bash
# From the project root
go run ./cmd/v2ray-keenetic
```

#### Frontend
```bash
# From the project root
cd web
npm install
npm run dev
```

---

## TR

Bu proje, Keenetic router'lar için özel olarak tasarlanmış bir V2Ray proxy uygulamasıdır. Router üzerindeki V2Ray servisini kontrol etmek için web tabanlı bir yönetim arayüzü sağlar.

## ✨ Özellikler

- 🌐 **Web Yönetim Arayüzü**: Sezgisel Vue.js tabanlı kontrol paneli
- 🔄 **Çoklu Protokol Desteği**: VMess, VLESS, Shadowsocks, Trojan desteği
- 📱 **QR Kod ile İçe Aktarma**: QR kodları aracılığıyla kolay sunucu yapılandırması
- 📊 **Gerçek Zamanlı İzleme**: Bağlantı durumu ve trafik istatistikleri
- 🔒 **Güvenli**: HTTPS desteği ve kimlik doğrulama
- 🌍 **Çoklu Dil Desteği**: İngilizce ve Türkçe desteği
- 📱 **Mobil Uyumlu**: Tüm cihazlar için duyarlı tasarım
- 🚀 **Otomatik Güncellemeler**: Otomatik yapılandırma güncellemeleri
- ⚖️ **Yük Dengeleme**: Yedekleme ile çoklu sunucu desteği

## 📋 Ön Gereksinimler

- Kabuk erişimi (SSH/Telnet) olan Keenetic router
- PATH içinde erişilebilir ve kurulu v2ray-core binary'si
- En az 32MB boş depolama alanı
- İlk kurulum için internet bağlantısı

## 🔧 Desteklenen Router Modelleri

- Keenetic Giga (KN-1010/1011)
- Keenetic Ultra (KN-1810)
- Keenetic Titan (KN-1820)
- Keenetic Hero 4G (KN-2310)
- Ve diğer NDMS tabanlı modeller

## 🚦 Hızlı Başlangıç

### 1. Ön Gereksinimleri Kontrol Edin
```bash
# v2ray-core kurulumunu doğrulayın
v2ray --version

# Kullanılabilir alanı kontrol edin
df -h
```

### 2. İndirin ve Kurun
```bash
# Mimarınıza uygun sürümü indirin (gerçek sürümle değiştirin)
wget https://github.com/halibiram/keenetiray/releases/latest/download/keenetiray_linux_arm.tar.gz -O /tmp/keenetiray.tar.gz

# Kurulum dizinini oluşturun
mkdir -p /opt/keenetiray
cd /opt/keenetiray

# Arşivi çıkartın ve izinleri ayarlayın
tar -xzf /tmp/keenetiray.tar.gz
chmod +x keenetiray

# Uygulamayı çalıştırın
./keenetiray
```

### 3. Web Arayüzüne Erişin
Tarayıcınızı açın ve şu adrese gidin: `http://[ROUTER_IP]:8888`

Varsayılan kimlik bilgileri:
- Kullanıcı adı: `admin`
- Şifre: `admin123` (hemen değiştirin!)

### Proje Yapısı
*   `/cmd/v2ray-keenetic`: Go ile yazılmış arka uç uygulaması (API sunucusu).
*   `/web`: Vue.js ile yazılmış ön uç uygulaması (Web arayüzü).
*   ... ve diğer geliştirme klasörleri.

### Geliştirme
Bu proje geliştirmenin erken aşamalarındadır. Geliştirme amacıyla yerel olarak çalıştırmak için bu adımları izleyin.

#### Arka Uç (Backend)
```bash
# Proje kök dizininden
go run ./cmd/v2ray-keenetic
```

#### Ön Uç (Frontend)
```bash
# Proje kök dizininden
cd web
npm install
npm run dev
```
