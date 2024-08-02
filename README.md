# Yerel Ağ Üzerinde WebSocket Uygulaması

Bu proje, Go programlama dili kullanılarak geliştirilmiş basit bir WebSocket uygulamasıdır.

Bu uygulama, yerel ağ üzerinde bulunan CLIENT'ların (istemcilerin) mesajlarını yerel ağ üzerinde 
oluşturulmuş bir SERVER'a (sunucuya) göndermesini ve sunucunun bu mesajları yerel ağ üzerinde bulunan 
diğer istemcilere yayımlanmasını sağlamaktadır.

# Yerel Ağ Üzerinde WebSocket Uygulaması için Gereksinimler
- Go Programlama Dili Kaynak Dosyaları (https://go.dev/dl/go1.22.5.windows-amd64.msi)
- Gorilla WebSocket Paketi (https://github.com/gorilla/websocket)

NOT: Aktif olarak kullanmakta olduğunuz Shell ekranını (CMD, PowerShell, GitBash vb.) içerisinde
Yerel Ağ Üzerinde WebSocket Uygulamasına ait proje dizinine giderek WebSocket uygulamaları
için geliştirilmiş olan bu Gorilla paketinin indirme işlemini aşağıdaki shell komutu ile
gerçekleştirebilirsiniz.

	- go get github.com/gorilla/websocket

# Yerel Ağ Üzerinde WebSocket Uygulaması için gerekli olan kurulumların yapılmasının ardından yine
kullanmakta olduğunuz Shell ortamı içerisinde;
	
	ipconfig komutunu çalıştırarak Server olarak tanımlayacağınız yerel bilgisayarınızın
	IP bilgilerini kopyalayın veya bir yere not edin.

# Daha sonrasında kaynak kodları içerisinde yer alan public/script.js dosyasını herhangi bir derleyici
üzerinden açmanız gerekmektedir.
	
Script dosyası içerisinde birinci satırda parantez içinde yer alan ---> 'ws://**.**.**.**:8080/ws'
şeklinde hali hazırda yazılmış olan komut bünyesinde var olan yıldızlı alanlar yerine Server olarak
tanımlamak istediğiniz yerel bilgisayarınızın IP bilgilerini girerek dosyayı kaydediniz!

# Bu işlemin ardından Yerel Ağ Üzerinde WebSocket Uygulaması kaynak dosyalarının yer aldığı dizin
içerisinde var olan "main.exe" uygulamasını çalıştırarak WINDOWS tarafından istenilen izinlere 
izin veriniz.

# TEBRİKLER! Yerel ağ sunucunuzu oluşturmuş bulunmaktasınız. Artık sizler ile aynı ağ merkezini paylaşan
CLIENT'lar (istemciler), herhangi bir WEB tarayıcı üzerinden "http://**.**.**.**:8080" adlı URL'yi
ziyaret ederek bu basit Yerel Ağ WebSocket uygulamasına bağlanabilir ve aynı ağ merkezini paylaştığı
diğer CLIENT'lar (istemciler) ile haberleşebilmektedir.

NOT: Unutmayın, sizin isteğiniz doğrultusunda "main.exe" dosyası çalıştırılmamış veya kapatılmış olması
halinde uygulama çalışmayacaktır. Bu basit WebSocket uygulamasını kullanmak istediğiniz durumlarda
"main.exe" dosyasını çalıştırmalısınız.