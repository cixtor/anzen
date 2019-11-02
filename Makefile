all:
	docker build -t babywaf_ariados:latest -f ariados/Dockerfile .
	docker build -t babywaf_dugtrio:latest -f dugtrio/Dockerfile .
	docker-compose up -d

clean:
	docker-compose down
	docker image rm babywaf_ariados
	docker image rm babywaf_dugtrio

test-insert:
	echo "Insert a couple of malicious URLs:"
	curl -XPOST "http://localhost:8080/insert/MALWARE/sarahdaniella.com/swift/SWIFT%2520%24.pdf.ace"
	curl -XPOST "http://localhost:8080/insert/POTENTIALLY_HARMFUL/amazon-sicherheit.kunden-ueberpruefung.xyz"
	curl -XPOST "http://localhost:8080/insert/SOCIAL_ENGINEERING/dieutribenhkhop.com/parking/pay/rd.php%3Fid%3D10"
	curl -XPOST "http://localhost:8080/insert/SOCIAL_ENGINEERING/www.hjaoopoa.top/admin.php%3Ff%3D1.gif"
	curl -XPOST "http://localhost:8080/insert/UNWANTED_SOFTWARE/down.mykings.pw:8888/ups.rar"
	curl -XPOST "http://localhost:8080/insert/SOCIAL_ENGINEERING/fo5.a1-downloader.org/g2v9s1.php%3Fid%3Dyourname%40yourdomain.com"
	curl -XPOST "http://localhost:8080/insert/UNWANTED_SOFTWARE/falconsafe.com.sg/api/get.php%3Fid%3DaW5mb0BzYXBjdXBncmFkZXMuY29t"
	curl -XPOST "http://localhost:8080/insert/POTENTIALLY_HARMFUL/www.lifelabs.vn/api/get.php%3Fid%3DaW5mb0BzYXBjdXBncmFkZXMuY29t"
	curl -XPOST "http://localhost:8080/insert/MALWARE/61kx.uk-insolvencydirect.com/sending_data/in_cgi/bbwp/cases/Inquiry.php"
	curl -XPOST "http://localhost:8080/insert/SOCIAL_ENGINEERING/daralasnan.com/wp-content/plugins/mkazaqbya/vmywyvz4.php"
	curl -XPOST "http://localhost:8080/insert/UNWANTED_SOFTWARE/www.studiolegaleabbruzzese.com/wp-content/plugins/urxwhbnw3ez/flight_4832.pdf"
	curl -XPOST "http://localhost:8080/insert/SOCIAL_ENGINEERING/kingskillz.ru/%7Ekingskil/Prince/Man/lucy/mine/shit.exe"
	curl -XPOST "http://localhost:8080/insert/MALWARE/art-archiv.ru/images/animated-number/docum-arhiv.exe"
	curl -XPOST "http://localhost:8080/insert/POTENTIALLY_HARMFUL/tscl.com.bd/m/RI%2520XIN%2520QUOTATION%2520LIST.zip"
	curl -XPOST "http://localhost:8080/insert/UNWANTED_SOFTWARE/jessisjewels.com/disk/update/postmaster/en/%3Far%3Dyourname%40yourdomain.com"
	curl -XPOST "http://localhost:8080/insert/MALWARE/structured.blackswanstore.com/plc/header.js"
	curl -XPOST "http://localhost:8080/insert/SOCIAL_ENGINEERING/ross.starvingmillionaire.org/unveiled/dropdown.js%3Fver%3D496e05e1aea0a9c4655800e8a7b9ea28"
	curl -XPOST "http://localhost:8080/insert/SOCIAL_ENGINEERING/ad.9tv.co.il/serv4/www/delivery/ajs.php%3Fzoneid%3D37%26cb%3D54350405237%26charset%3Dutf-8"
	curl -XPOST "http://localhost:8080/insert/UNWANTED_SOFTWARE/giants.yourzip.co/static/quotes.js%3Fver%3Dd58072be2820e8682c0a27c0518e805e"
	curl -XPOST "http://localhost:8080/insert/MALWARE/evans.babajilab.in/specimen/1479491/tire-something-detect-five-what-knot-unknown-entertain-stiff"
	curl -XPOST "http://localhost:8080/insert/MALWARE/adv.riza.it/www/delivery/ajs.php%3Fzoneid%3D51%26cb%3D96020978060"
	curl -XPOST "http://localhost:8080/insert/SOCIAL_ENGINEERING/www.ywvcomputerprocess.info/errorreport/ty5ug6h4ndma4/"

test-retrieve:
	echo "Retrieve a couple of malicious URLs:"
	curl -XGET "http://localhost:8080/urlinfo/1/sarahdaniella.com/swift/SWIFT%2520%24.pdf.ace"
	curl -XGET "http://localhost:8080/urlinfo/1/amazon-sicherheit.kunden-ueberpruefung.xyz"
	curl -XGET "http://localhost:8080/urlinfo/1/dieutribenhkhop.com/parking/pay/rd.php%3Fid%3D10"
	curl -XGET "http://localhost:8080/urlinfo/1/www.hjaoopoa.top/admin.php%3Ff%3D1.gif"
	curl -XGET "http://localhost:8080/urlinfo/1/down.mykings.pw:8888/ups.rar"
	curl -XGET "http://localhost:8080/urlinfo/1/fo5.a1-downloader.org/g2v9s1.php%3Fid%3Dyourname%40yourdomain.com"
	curl -XGET "http://localhost:8080/urlinfo/1/falconsafe.com.sg/api/get.php%3Fid%3DaW5mb0BzYXBjdXBncmFkZXMuY29t"
	curl -XGET "http://localhost:8080/urlinfo/1/www.lifelabs.vn/api/get.php%3Fid%3DaW5mb0BzYXBjdXBncmFkZXMuY29t"
	curl -XGET "http://localhost:8080/urlinfo/1/61kx.uk-insolvencydirect.com/sending_data/in_cgi/bbwp/cases/Inquiry.php"
	curl -XGET "http://localhost:8080/urlinfo/1/daralasnan.com/wp-content/plugins/mkazaqbya/vmywyvz4.php"
	curl -XGET "http://localhost:8080/urlinfo/1/www.studiolegaleabbruzzese.com/wp-content/plugins/urxwhbnw3ez/flight_4832.pdf"
	curl -XGET "http://localhost:8080/urlinfo/1/kingskillz.ru/%7Ekingskil/Prince/Man/lucy/mine/shit.exe"
	curl -XGET "http://localhost:8080/urlinfo/1/art-archiv.ru/images/animated-number/docum-arhiv.exe"
	curl -XGET "http://localhost:8080/urlinfo/1/tscl.com.bd/m/RI%2520XIN%2520QUOTATION%2520LIST.zip"
	curl -XGET "http://localhost:8080/urlinfo/1/jessisjewels.com/disk/update/postmaster/en/%3Far%3Dyourname%40yourdomain.com"
	curl -XGET "http://localhost:8080/urlinfo/1/structured.blackswanstore.com/plc/header.js"
	curl -XGET "http://localhost:8080/urlinfo/1/ross.starvingmillionaire.org/unveiled/dropdown.js%3Fver%3D496e05e1aea0a9c4655800e8a7b9ea28"
	curl -XGET "http://localhost:8080/urlinfo/1/ad.9tv.co.il/serv4/www/delivery/ajs.php%3Fzoneid%3D37%26cb%3D54350405237%26charset%3Dutf-8"
	curl -XGET "http://localhost:8080/urlinfo/1/giants.yourzip.co/static/quotes.js%3Fver%3Dd58072be2820e8682c0a27c0518e805e"
	curl -XGET "http://localhost:8080/urlinfo/1/evans.babajilab.in/specimen/1479491/tire-something-detect-five-what-knot-unknown-entertain-stiff"
	curl -XGET "http://localhost:8080/urlinfo/1/adv.riza.it/www/delivery/ajs.php%3Fzoneid%3D51%26cb%3D96020978060"
	curl -XGET "http://localhost:8080/urlinfo/1/www.ywvcomputerprocess.info/errorreport/ty5ug6h4ndma4/"

benchmark-safe:
	wrk -t4 -c100 -d60s --latency "http://localhost:8080/urlinfo/1/cixtor.com/hello%2fworld"

benchmark-unsafe:
	curl -XPOST "http://localhost:8080/insert/MALWARE/twitter-scanner.com/hello%2fworld"
	wrk -t4 -c100 -d60s --latency "http://localhost:8080/urlinfo/1/twitter-scanner.com/hello%2fworld"
