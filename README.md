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
    # This URL is an EXAMPLE for the ARM architecture.
    # Replace it with the actual URL from the GitHub Releases page.
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
        tar -xzf /tmp/v2ray-keenetic.tar.gz -C .
        ```
    *   Make the binary executable:
        ```bash
        chmod +x ./v2ray-keenetic
        ```

3.  **Run the Application**: You can now run the application:
    ```bash
    ./v2ray-keenetic
    ```
    The application will start, and the web interface will be available at `http://<your-router-ip>:8080`.

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
1.  **Router'a İndirme**: Router'ınıza SSH ile bağlanın. Projenin [GitHub Releases](https://github.com/your-username/v2ray-keenetic/releases) sayfasından router'ınızın mimarisine uygun paket URL'sini bulun. `wget` kullanarak paketi `/tmp` dizinine indirin.
    ```bash
    # Bu URL, ARM mimarisi için bir ÖRNEKTİR.
    # Onu GitHub Releases sayfasındaki gerçek URL ile değiştirin.
    wget https://github.com/your-username/v2ray-keenetic/releases/download/v1.0.0/v2ray-keenetic_linux_arm.tar.gz -O /tmp/v2ray-keenetic.tar.gz
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

3.  **Uygulamayı Çalıştırma**: Artık uygulamayı çalıştırabilirsiniz:
    ```bash
    ./v2ray-keenetic
    ```
    Uygulama başlayacak ve web arayüzü `http://<router-ip-adresiniz>:8080` adresinde erişilebilir olacaktır.

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
