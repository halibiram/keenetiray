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

### Installation on Keenetic

**Prerequisites:**
*   You must have shell access to your Keenetic router (e.g., via SSH).
*   You need the `v2ray-core` binary installed on your router and available in the system's `PATH`. This application is a *controller* for `v2ray-core`; it does not include it.

**Steps:**
1.  **Download on Router**: Connect to your router via SSH. Find the package URL for your router's architecture from the project's [GitHub Releases](https://github.com/your-username/v2ray-keenetic/releases) page. Use `wget` to download it to the `/tmp` directory.
    ```bash
    # Example for ARM architecture. Replace URL with the correct one from the releases page.
    wget https://github.com/your-username/v2ray-keenetic/releases/download/v1.0.0/v2ray-keenetic_linux_arm.tar.gz -O /tmp/v2ray-keenetic.tar.gz
    ```

2.  **Install on Router**: Perform the following steps:
    *   Navigate to a persistent directory. The user data partition is a good choice (`/opt` is common).
        ```bash
        mkdir -p /opt/v2ray-keenetic
        cd /opt/v2ray-keenetic
        ```
    *   Extract the archive:
        ```bash
        tar -xzf /tmp/v2ray-keenetic_linux_arm.tar.gz -C .
        ```
    *   Make the binary executable:
        ```bash
        chmod +x ./v2ray-keenetic
        ```

4.  **Run the Application**: You can now run the application:
    ```bash
    ./v2ray-keenetic
    ```
    The application will start, and the web interface will be available at `http://<your-router-ip>:8080`.

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

### Keenetic Cihazına Kurulum

**Ön Gereksinimler:**
*   Keenetic router'ınıza shell erişiminiz (örneğin, SSH ile) olmalıdır.
*   Router'ınızda `v2ray-core` binary'sinin kurulu ve sistem `PATH`'inde erişilebilir olması gerekmektedir. Bu uygulama, `v2ray-core` için bir *kontrolcüdür*; onu içermez.

**Adımlar:**
1.  **Paketi Oluşturma**: Geliştirme makinenizden derleme betiğini çalıştırarak router'ınızın mimarisine (örneğin, `arm` veya `mips`) uygun paketi oluşturun:
    ```bash
    ./scripts/build.sh arm
    ```
    Bu komut, `dist/v2ray-keenetic_linux_arm.tar.gz` adında bir dosya oluşturacaktır.

2.  **Router'a Yükleme**: Oluşturulan `.tar.gz` arşivini router'ınızın geçici bir dizinine (örneğin, `/tmp`) kopyalayın. Bunun için `scp` kullanabilirsiniz:
    ```bash
    scp dist/v2ray-keenetic_linux_arm.tar.gz root@192.168.1.1:/tmp/
    ```

2.  **Router'da Kurulum**: Aşağıdaki adımları izleyin:
    *   Kalıcı bir dizine gidin. Kullanıcı veri bölümü (`/opt` yaygındır) iyi bir seçimdir.
        ```bash
        mkdir -p /opt/v2ray-keenetic
        cd /opt/v2ray-keenetic
        ```
    *   Arşivi çıkartın:
        ```bash
        tar -xzf /tmp/v2ray-keenetic.tar.gz -C .
        ```
    *   Binary'yi çalıştırılabilir yapın:
        ```bash
        chmod +x ./v2ray-keenetic
        ```

4.  **Uygulamayı Çalıştırma**: Artık uygulamayı çalıştırabilirsiniz:
    ```bash
    ./v2ray-keenetic
    ```
    Uygulama başlayacak ve web arayüzü `http://<router-ip-adresiniz>:8080` adresinde erişilebilir olacaktır.

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
