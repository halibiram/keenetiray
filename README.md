# V2Ray Keenetic Control Panel

*English below / Türkçe versiyonu aşağıdadır.*

---

## EN

This project is a V2Ray proxy application designed specifically for Keenetic routers. It provides a web-based management interface to control the V2Ray service on the router.

### Features (Planned)

*   **Web UI**: A responsive web interface to manage V2Ray server configurations.
*   **V2Ray Core Integration**: Manages the V2Ray-core process for traffic routing.
*   **System Integration**: Integrates with Keenetic RouterOS for startup and network management.
*   **Multiple Protocol Support**: Supports VMess, VLESS, Shadowsocks, etc.

### Installation on Keenetic (Planned)

**CAUTION:** This project is still under development and is not ready for direct installation on Keenetic devices. The following steps outline the target installation process for a future release.

1.  **Download Package**: Download the `.opkg` or `.tar.gz` package appropriate for your Keenetic device's architecture (ARM/MIPS) from the project's `Releases` page.
2.  **Upload to Device**: Upload the downloaded package to the `/tmp` or `/opt` directory on the device using `scp` or the file manager in the Keenetic web interface.
3.  **Install via SSH**: Connect to your Keenetic device via SSH and run the following commands:
    ```bash
    # Example for an opkg package:
    opkg install /tmp/v2ray-keenetic-*.opkg
    ```
4.  **Configure via Web UI**: After installation is complete, you can find the V2Ray Control Panel under the "Applications" menu or a similar section in the Keenetic web interface (usually at `192.168.1.1`). You can configure your V2Ray server settings from there.

### Project Structure

*   `/cmd/v2ray-keenetic`: The Go backend application (API server).
*   `/web`: The Vue.js frontend application (Web UI).
*   `/internal`: Internal Go packages for the backend.
*   `/scripts`: Build and deployment scripts.
*   `/configs`: Example configuration files.
*   `/docs`: Project documentation.

### Development

This project is in the early stages of development.

#### Backend
The backend is written in Go using the Gin framework. To run the backend:
```bash
# From the project root
# Note: Building the binary is currently not supported in this environment due to size limits.
# Use 'go run' for development.
go run ./cmd/v2ray-keenetic
```

#### Frontend
The frontend is a Vue.js 3 application built with Vite. To run the frontend development server:
```bash
# From the project root
cd web
npm install
npm run dev
```
The frontend will be available at `http://localhost:5173` and will proxy API requests to the backend running on `http://localhost:8080`.

---

## TR

Bu proje, Keenetic router'lar için özel olarak tasarlanmış bir V2Ray proxy uygulamasıdır. Router üzerindeki V2Ray servisini kontrol etmek için web tabanlı bir yönetim arayüzü sağlar.

### Özellikler (Planlanan)

*   **Web Arayüzü**: V2Ray sunucu yapılandırmalarını yönetmek için duyarlı bir web arayüzü.
*   **V2Ray Çekirdek Entegrasyonu**: Trafik yönlendirme için V2Ray-core sürecini yönetir.
*   **Sistem Entegrasyonu**: Başlangıç ve ağ yönetimi için Keenetic RouterOS ile entegre olur.
*   **Çoklu Protokol Desteği**: VMess, VLESS, Shadowsocks vb. destekler.

### Keenetic Cihazına Kurulum (Planlanan)

**DİKKAT:** Bu proje henüz geliştirme aşamasındadır ve Keenetic cihazlarına doğrudan kuruluma hazır değildir. Aşağıdaki adımlar, projenin gelecekteki bir sürümü için hedeflenen kurulum sürecini özetlemektedir.

1.  **Paketin İndirilmesi**: Projenin `Releases` sayfasından Keenetic cihazınızın mimarisine (ARM/MIPS) uygun `.opkg` veya `.tar.gz` paketini indirin.
2.  **Cihaza Yükleme**: İndirilen paketi `scp` veya Keenetic web arayüzündeki dosya yöneticisi aracılığıyla cihazın `/tmp` veya `/opt` dizinine yükleyin.
3.  **SSH ile Kurulum**: SSH üzerinden Keenetic cihazınıza bağlanın ve aşağıdaki komutları çalıştırın:
    ```bash
    # Örnek olarak opkg paketi için:
    opkg install /tmp/v2ray-keenetic-*.opkg
    ```
4.  **Web Arayüzünden Yapılandırma**: Kurulum tamamlandıktan sonra, Keenetic web arayüzünde (genellikle `192.168.1.1`) "Uygulamalar" veya benzeri bir menü altında V2Ray Kontrol Paneli'ni bulabilirsiniz. Buradan V2Ray sunucu ayarlarınızı yapabilirsiniz.

### Proje Yapısı

*   `/cmd/v2ray-keenetic`: Go ile yazılmış arka uç uygulaması (API sunucusu).
*   `/web`: Vue.js ile yazılmış ön uç uygulaması (Web arayüzü).
*   `/internal`: Arka uç için dahili Go paketleri.
*   `/scripts`: Derleme ve dağıtım betikleri.
*   `/configs`: Örnek yapılandırma dosyaları.
*   `/docs`: Proje dokümantasyonu.

### Geliştirme

Bu proje geliştirmenin erken aşamalarındadır.

#### Arka Uç (Backend)
Arka uç, Gin framework'ü kullanılarak Go dilinde yazılmıştır. Arka ucu çalıştırmak için:
```bash
# Proje kök dizininden
# Not: Bu ortamda dosya boyutu limitleri nedeniyle binary oluşturma desteklenmemektedir.
# Geliştirme için 'go run' kullanın.
go run ./cmd/v2ray-keenetic
```

#### Ön Uç (Frontend)
Ön uç, Vite ile oluşturulmuş bir Vue.js 3 uygulamasıdır. Geliştirme sunucusunu çalıştırmak için:
```bash
# Proje kök dizininden
cd web
npm install
npm run dev
```
Ön uç `http://localhost:5173` adresinde erişilebilir olacak ve API isteklerini `http://localhost:8080` adresinde çalışan arka uca yönlendirecektir.
