package main

import (
	"testing"
)

func TestThreatType(t *testing.T) {
	app := NewApplication()

	info, err := app.ThreatType(`example.com/hello%2Fworld%3Ffoo%3Dbar`)

	if err != nil {
		t.Fatal(err)
		return
	}

	if info.Threat != ttNone {
		t.Fatal("ThreatType should be NONE")
		return
	}
}

func TestThreatTypeInsertOne(t *testing.T) {
	app := NewApplication()

	url := "dieutribenhkhop.com/parking/pay/rd.php?id=10"

	app.Database.Insert(HashURL(url))

	info, err := app.ThreatType(url)

	if err != nil {
		t.Fatal(err)
		return
	}

	if info.Threat != ttMalware {
		t.Fatalf("ThreatType for `%s` should be MALWARE instead of %s", url, info.Threat)
		return
	}
}

func TestThreatTypeInsertMany(t *testing.T) {
	app := NewApplication()

	// NOTES(yorman): sample URLs obtained from the "Malware Domain List"
	// available at -> https://www.malwaredomainlist.com/mdl.php
	//
	// All domains on this list should be considered dangerous. If you do not
	// know what you are doing here, it is recommended you leave right away.
	// This website is a resource for security professionals and enthusiasts.
	malicious := []string{
		"textspeier.de",
		"photoscape.ch/Setup.exe",
		"sarahdaniella.com/swift/SWIFT%20$.pdf.ace",
		"amazon-sicherheit.kunden-ueberpruefung.xyz",
		"alegroup.info/ntnrrhst",
		"fourthgate.org/Yryzvt",
		"dieutribenhkhop.com/parking/",
		"dieutribenhkhop.com/parking/pay/rd.php?id=10",
		"ssl-6582datamanager.de/",
		"privatkunden.datapipe9271.com/",
		"www.hjaoopoa.top/admin.php?f=1.gif",
		"up.mykings.pw:8888/update.txt",
		"down.mykings.pw:8888/ver.txt",
		"down.mykings.pw:8888/ups.rar",
		"fo5.a1-downloader.org/g2v9s1.php?id=yourname@yourdomain.com",
		"falconsafe.com.sg/api/get.php?id=aW5mb0BzYXBjdXBncmFkZXMuY29t",
		"www.lifelabs.vn/api/get.php?id=aW5mb0BzYXBjdXBncmFkZXMuY29t",
		"61kx.uk-insolvencydirect.com/sending_data/in_cgi/bbwp/cases/Inquiry.php",
		"daralasnan.com/wp-content/plugins/mkazaqbya/vmywyvz4.php",
		"www.studiolegaleabbruzzese.com/wp-content/plugins/urxwhbnw3ez/flight_4832.pdf",
		"raneevahijab.id/adnin/box/workspace/",
		"kingskillz.ru/~kingskil/Prince/Man/lucy/mine/shit.exe",
		"www.family-partners.fr/data.dpg",
		"elmissouri.fr/data.dpg",
		"art-archiv.ru/images/animated-number/docum-arhiv.exe",
		"catjogger.win/ganel/gate.php",
		"tscl.com.bd/m/RI%20XIN%20QUOTATION%20LIST.zip",
		"ad.getfond.info",
		"jessisjewels.com/disk/update/postmaster/en/?ar=yourname@yourdomain.com",
		"structured.blackswanstore.com/plc/header.js",
		"ross.starvingmillionaire.org/unveiled/dropdown.js?ver=496e05e1aea0a9c4655800e8a7b9ea28",
		"ad.9tv.co.il/serv4/www/delivery/ajs.php?zoneid=37&cb=54350405237&charset=utf-8",
		"giants.yourzip.co/static/quotes.js?ver=d58072be2820e8682c0a27c0518e805e",
		"evans.babajilab.in/specimen/1479491/tire-something-detect-five-what-knot-unknown-entertain-stiff",
		"tahit.wastech2016.in/xcqrsw3.html",
		"pogruz.wanyizhao.net/ceqxwu3.html",
		"livre.wasastation.fi/ceqxwu3.html",
		"sanya.vipc2f.com/ceqxwu3.html",
		"tanner.alicerosenmanmemorial.com/hggfgl3.html",
		"wuvac.agwebdigital.com/dsgajo3.html",
		"rufex.ajfingenieros.cl/dsgajo3.html",
		"cqji.artidentalkurs.com/vdgqb3.html",
		"unlink.altitude.lv/vdgqb3.html",
		"vitaly.agricolacolhue.cl/rncbu3.html",
		"geil.alon3.tk/rncbu3.html",
		"honor.agitaattori.fi/rncbu3.html",
		"gojnox.boxtomarket.com/yxmvr3.html",
		"womsy.bobbutcher.net/rtuee3.html",
		"bonjo.bmbsklep.pl/jvoxyj3.html",
		"pybul.bestfrozenporn.nl/jvoxyj3.html",
		"soxorok.ddospower.ro/lwwxx3.html",
		"funkucck.bluerobot.cl/lwwxx3.html",
		"wixx.caliptopis.cl/lwwxx3.html",
		"mepra.blautechnology.cl/pwigd3.html",
		"wopper.bioblitzgaming.ca/pwigd3.html",
		"losas.cabanaslanina.com.ar/wkicrz3.html",
		"losos.caliane.com.br/wkicrz3.html",
		"pumpkin.brisik.net/rvgkm3.html",
		"decorator.crabgrab.cl/rjavgx3.html",
		"scanty.colormark.cl/rjavgx3.html",
		"coffeol.com/fend/raw_server.exe",
		"www.pgathailand.com/which.exe",
		"euro-vertrieb.com/hosteurope/KIS-Login.htm",
		"www.jcmarcadolib.com/hbc/a.php",
		"lexu.goggendorf.at/nukgfr2.html",
		"victor.connectcloud.ch/nukgfr2.html",
		"molla.gato1000.cl/edmiu2.html",
		"hmora.fred-build.tk/odbsx2.html",
		"peeg.fronterarq.cl/odbsx2.html",
		"adv.riza.it/www/delivery/ajs.php?zoneid=51&cb=96020978060",
		"borat.elticket.com.ar/pkge2.html",
		"lay.elticket.com.ar/tslwo2.html",
		"plank.duplicolor.cl/zbtqvc2.html",
		"pave.elisecries.com/zbtqvc2.html",
		"spread.diadanoivabh.com.br/alvbh2.html",
		"smilll.depozit.hr/bgaldb2.html",
		"soros.departamentosejecutivos.cl/venak2.html",
		"stock.daydreamfuze.com/rxdjna2.html",
		"vdula.czystykod.pl/rxdjna2.html",
		"produla.czatgg.pl/rxdjna2.html",
		"aircraft.evote.cl/ybluq2.html",
		"absurdity.flarelight.com/xdnkn2.html",
		"pacan.gofreedom.info/omrjy2.html",
		"pacman.gkgar.com/omrjy2.html",
		"terem.eltransbt.ro/ysfmgl2.html",
		"likes.gisnetwork.net/ysfmgl2.html",
		"personal.editura-amsibiu.ro/rdxzmt2.html",
		"above.e-rezerwacje24.pl/uzjuz2.html",
		"headless.ebkfwd.com/uzjuz2.html",
		"higher.dwebsi.tk/uzjuz2.html",
		"crops.dunight.eu/uzjuz2.html",
		"invention.festinolente.cl/ajuijm2.html",
		"erupt.fernetmoretti.com.ar/ajuijm2.html",
		"stork.escortfinder.cl/ajuijm2.html",
		"vomit.facilitandosonhos.com.br/ajuijm2.html",
		"trifle.ernstenco.be/ajuijm2.html",
		"www.ywvcomputerprocess.info/errorreport/ty5ug6h4ndma4/",
		"cosmos.furnipict.com/gsvot2.html",
		"milf.gabriola.cl/gsvot2.html",
		"cosmos.felago.es/gsvot2.html",
	}

	for _, url := range malicious {
		if !app.Database.Insert(HashURL(url)) {
			t.Fatalf("cannot insert `%s` into the database\n", url)
			return
		}
	}

	var err error
	var info ThreatInfo

	for _, url := range malicious {
		if info, err = app.ThreatType(url); err != nil {
			t.Fatal(err)
			return
		}

		if info.Threat != ttMalware {
			t.Fatalf("ThreatType for `%s` should be MALWARE instead of %s", url, info.Threat)
			return
		}
	}
}
