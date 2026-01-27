# KT1 - Konceptualni i arhitekturni opis - Smart City Environmental Monitoring System

<p align="center">
  <b>Smart City Environmental Monitoring System</b>
</p>
<p align="center">
  <img src="images/logo.png" alt="UrbanSense Logo" width=auto height="200px"/>
</p>

# 1. Domen problema i poslovni kontekst

## 1.1. Kratak opis domena

**UrbanSense** predstavlja sistem za pametne gradske senzore namenjen praćenju i analizi **kvaliteta životne sredine** u urbanim sredinama. Sistem omogućava prikupljanje podataka u realnom vremenu sa distribuiranih senzora postavljenih širom grada, kao što su senzori za kvalitet vazduha, nivo buke, temperaturu i vlažnost vazduha.

Prikupljeni podaci se centralizovano obrađuju, skladište i analiziraju kako bi se omogućio uvid u trenutno stanje životne sredine, detekcija potencijalno opasnih vrednosti i donošenje informisanih odluka od strane gradskih službi i drugih relevantnih aktera. Sistem je namenjen kako institucionalnim korisnicima (gradskim upravama), tako i široj javnosti kroz transparentan pristup agregiranim informacijama.

## 1.2. Učesnici (akteri) i njihove uloge
- Citizens (Građani) - Krajnji korisnici sistema koji imaju uvid u javno dostupne podatke o kvalitetu vazduha i drugim ekološkim parametrima u različitim delovima grada. Mogu koristiti sistem za informisanje, planiranje aktivnosti i praćenje dugoročnih trendova.
- City Authorities (Gradske uprave i institucije) - Institucionalni korisnici odgovorni za praćenje stanja životne sredine, donošenje odluka, reagovanje na prekoračenja dozvoljenih vrednosti i planiranje ekoloških mera. Koriste detaljne izveštaje, alarme i istorijske analize.
- Environmental Agencies (Ekološke agencije i inspekcije) - Koriste podatke sistema za nadzor, analizu usklađenosti sa propisima i izradu stručnih izveštaja o stanju životne sredine.
- System Administrators (Administratori sistema) - Zaduženi za upravljanje infrastrukturom sistema, konfiguraciju senzora, upravljanje korisničkim ulogama, nadzor nad radom sistema i pregled audit logova.
- IoT Sensors / Edge Devices (Senzori i edge uređaji) - Fizički uređaji postavljeni na terenu koji kontinuirano mere ekološke parametre i šalju podatke ka centralnom sistemu.
- External Data Consumers (Eksterni sistemi i servisi) - Treće strane (npr. istraživačke institucije ili gradski open-data portali) koje koriste agregirane ili anonimizovane podatke putem API-ja.

## 1.3. Poslovni procesi koje softver podržava
- Prikupljanje podataka sa senzora - Kontinuirano očitavanje ekoloških parametara sa distribuiranih IoT senzora i siguran prenos podataka ka centralnom sistemu.
- Obrada i validacija podataka - Provera integriteta i validnosti podataka, filtriranje anomalija i priprema za dalju obradu.
- Skladištenje i agregacija podataka - Čuvanje sirovih i agregiranih podataka u cilju podrške istorijskoj analizi i dugoročnom praćenju trendova.
- Analitika i detekcija kritičnih stanja - Analiza podataka radi identifikacije prekoračenja dozvoljenih vrednosti, detekcije zagađenja i drugih potencijalno rizičnih situacija.
- Vizualizacija i izveštavanje - Prikaz podataka kroz dashboarde, mape i izveštaje prilagođene različitim grupama korisnika.
- Alarmiranje i obaveštavanje - Generisanje upozorenja i notifikacija za gradske službe u slučaju detekcije kritičnih ili opasnih vrednosti.
- Administracija i nadzor sistema - Upravljanje senzorima, korisnicima, pristupima, konfiguracijom sistema i audit evidencijama.

# 2. Arhitektura zamišljenog softvera

## 2.1. Arhitekturalne karakteristike

UrbanSense sistem je projektovan kao modularna, event-driven arhitektura prilagođena radu sa **distribuiranim IoT senzorima** i obradom podataka u realnom vremenu. Arhitektura je zamišljena tako da omogući jasno razdvajanje odgovornosti, skalabilnost po funkcionalnim celinama i jednostavno proširenje sistema u budućnosti.

Osnovne arhitekturalne karakteristike sistema uključuju:
- IoT/Edge sloj za prikupljanje podataka sa terena (senzori).
- Centralni backend servis odgovoran za prijem, obradu i analizu podataka.
- Asinhronu obradu događaja korišćenjem message broker-a za pouzdan i skalabilan protok podataka.
- Poliglotnu perzistenciju, gde se različiti tipovi podataka skladište u skladištima prilagođenim njihovoj prirodi (npr. vremenske serije).
- Prezentacioni sloj za vizualizaciju podataka i interakciju sa korisnicima.

Event-driven pristup omogućava da sistem efikasno obradi veliki broj merenja bez direktne zavisnosti između senzora i potrošača podataka, čime se smanjuje međuzavisnost komponenti i povećava otpornost sistema na greške.

## 2.2. Predložene tehnologije i njihove uloge

U skladu sa zahtevima KT1 za tim od jednog člana, sistem koristi minimum tri tehnologije, pri čemu nisu sve bazirane na klasičnim web aplikacijama i relacionim bazama podataka.

### 1 - IoT senzori i edge uređaji

**Environmental IoT Sensors**
- Fizički uređaji postavljeni u urbanom okruženju koji mere parametre kao što su kvalitet vazduha, nivo buke, temperatura i vlažnost.
Njihova uloga je generisanje mernih podataka u realnom vremenu i slanje podataka ka centralnom sistemu. Ovi uređaji predstavljaju početnu tačku sistema i ključni izvor podataka.

### 2 - Backend servis za obradu podataka

**Data Ingestion & Processing Service (Go)**

Centralni backend servis odgovoran za:

- prijem podataka sa senzora,
- validaciju i osnovnu obradu merenja,
- generisanje događaja za dalju analizu i skladištenje.

Go je odabran zbog:

- dobre podrške za konkurentnost,
- efikasne obrade velikog broja događaja,
- pogodnosti za rad u IoT i event-driven sistemima.

### 3 - Message broker (event-driven komunikacija)

**MQTT ili Apache Kafka**

Message broker služi kao centralni mehanizam za asinhronu razmenu podataka između senzora, backend servisa i drugih komponenti sistema.

Ovaj sloj omogućava:

- pouzdan prenos podataka,
- obradu velikog broja događaja,
- decoupling senzora od sistema za analitiku i skladištenje.

Ovakav pristup je posebno pogodan za sisteme sa velikim brojem IoT uređaja i nepredvidivim obrascima saobraćaja.

### 4 - Skladište podataka

**Time-series / NoSQL Database (npr. InfluxDB ili MongoDB)**

Služi za skladištenje merenja i agregiranih podataka o stanju životne sredine.

Prednosti ovog tipa skladišta:
- optimizacija za vremenske serije,
- fleksibilna šema podataka,
- efikasna istorijska analiza i agregacija.

### 5 - Prezentacioni sloj

**Web Dashboard (React ili sličan frontend okvir)**

Omogućava vizualizaciju podataka kroz grafike, mape i izveštaje.
Namenjen je građanima i gradskim institucijama za pregled trenutnog i istorijskog stanja životne sredine.

## 2.3. Integracije i spoljašnji sistemi

Sistem je projektovan tako da omogući integraciju sa spoljašnjim sistemima putem API-ja, uključujući:

- gradske open-data portale,
- istraživačke institucije,
- ekološke informacione sisteme.

Integracije su zamišljene kao **read-only** u osnovnoj verziji sistema, čime se smanjuje kompleksnost i bezbednosni rizici u ranim fazama projekta.

# 3. Grupe slučajeva korišćenja (Use-case groups)

## 3.1. Citizens (Građani) – Public Web Dashboard

**Pregled trenutnog stanja životne sredine**
- Pregled aktuelnih vrednosti kvaliteta vazduha, nivoa buke i drugih ekoloških parametara po gradskim zonama.

**Vizualizacija i mape**
- Prikaz merenja na interaktivnim mapama i grafikonima radi lakšeg razumevanja prostorne raspodele i promena u vremenu.

**Istorijski podaci i trendovi**
- Pregled istorijskih merenja i uočavanje dugoročnih trendova u kvalitetu životne sredine.

## 3.2. City Authorities & Environmental Agencies - Institutional Dashboard

**Detaljna analitika i izveštaji**
- Napredna analiza podataka, agregacija po vremenskim periodima i generisanje izveštaja za potrebe planiranja i donošenja odluka.

**Praćenje prekoračenja dozvoljenih vrednosti**
- Detekcija situacija u kojima merenja prelaze definisane pragove i zahtevaju reakciju nadležnih službi.

**Alarmiranje i obaveštavanje**
- Prijem upozorenja i notifikacija u realnom vremenu u slučaju kritičnih ekoloških stanja.

## 3.3. System Administrators - Administration Interface

**Upravljanje senzorima i uređajima**
- Registracija, konfiguracija, aktivacija i deaktivacija IoT senzora u sistemu.

**Upravljanje korisnicima i ulogama**
- Definisanje pristupa institucionalnih korisnika i kontrola njihovih privilegija.

**Nadzor rada sistema**
- Praćenje dostupnosti komponenti sistema, osnovnih performansi i stanja komunikacije sa senzorima.

**Audit i evidencija aktivnosti**
- Pregled sistemskih i administrativnih logova radi praćenja promena i aktivnosti u sistemu.

## 3.4. External Systems - Data Access API

**Pristup agregiranim podacima**
- Omogućavanje eksternim sistemima da pristupe anonimizovanim i agregiranim podacima putem API-ja.

**Integracija sa open-data platformama**
- Objavljivanje odabranih skupova podataka za javnu i istraživačku upotrebu.

# 4. Osetljivi resursi (Sensitive assets) i bezbednosni ciljevi

## 4.1. Podaci sa senzora (Environmental Sensor Data)

**Opis:**

Sirovi i obrađeni podaci koje generišu gradski senzori (kvalitet vazduha, nivo buke, temperatura, vlažnost).

**Bezbednosni ciljevi:**
- **Integritet** - sprečiti neovlašćenu izmenu merenja koja bi mogla dovesti do pogrešnih zaključaka ili odluka.
- **Dostupnost** - podaci moraju biti dostupni za analizu i izveštavanje u realnom vremenu.
- **Poverljivost (sekundarni cilj)** - u slučajevima kada se podaci mogu indirektno povezati sa konkretnim lokacijama ili obrascima kretanja.

**Prioritet:** Visok

## 4.2. Lokacije senzora i topologija mreže

**Opis:**

Podaci o tačnoj lokaciji senzora, njihovoj raspodeli po gradu i međusobnoj povezanosti.

**Bezbednosni ciljevi:**
- **Integritet** - osigurati da informacije o lokacijama budu tačne i ažurne.
- **Poverljivost** - sprečiti zloupotrebu informacija o lokaciji senzora (npr. fizičko onesposobljavanje ili sabotaža).

**Prioritet:** Srednji

## 4.3. Sistem za alarmiranje i obaveštavanje

**Opis:**

Mehanizam za generisanje upozorenja i notifikacija u slučaju detekcije prekoračenja dozvoljenih vrednosti ili anomalija u merenjima.

**Bezbednosni ciljevi:**
- **Integritet** - sprečiti lažne ili manipulativne alarme.
- **Dostupnost** - sistem mora biti pouzdan i dostupan u kritičnim situacijama.

**Prioritet:** Visok

## 4.4. Administrativni pristup i privilegije

**Opis:**

Administratorski nalozi, uloge i privilegije koje omogućavaju upravljanje senzorima, konfiguracijom sistema i korisničkim pristupima.

**Bezbednosni ciljevi:**
- **Poverljivost** - zaštita kredencijala i pristupnih podataka.
- **Integritet** - sprečiti neovlašćene izmene konfiguracije i privilegija.
- **Auditabilnost** - obezbediti trag aktivnosti administrativnih korisnika.

**Prioritet:** Visok

## 4.5. Istorijski i agregirani podaci o kvalitetu vazduha

**Opis:**

Dugoročno skladišteni i agregirani podaci koji se koriste za analizu trendova, izveštavanje i planiranje.

**Bezbednosni ciljevi:**
- **Dostupnost** - kontinuirana dostupnost za analizu i javni uvid.
- **Integritet** - očuvanje tačnosti istorijskih podataka.

**Prioritet:** Srednji

## 4.6. Korisnički podaci institucionalnih korisnika

**Opis:**

Podaci o gradskim službama, ekološkim agencijama i njihovim korisničkim nalozima (identiteti, uloge, pristupi).

**Bezbednosni ciljevi:**
- **Poverljivost** - zaštita identiteta i pristupnih podataka.
- **Integritet** - sprečiti neovlašćenu eskalaciju privilegija.

**Prioritet:** Srednji

## 4.7. Regulativni i zakonski aspekti

**Opis:**

Iako UrbanSense sistem primarno obrađuje ekološke podatke, pojedini resursi (npr. podaci povezani sa lokacijama senzora ili institucionalnim korisnicima) mogu potpasti pod **regulative o zaštiti podataka**, kao što je **GDPR**, u slučajevima kada se podaci mogu indirektno povezati sa identifikovanim ili identifikabilnim subjektima.

Zbog toga je neophodno obezbediti:
- minimizaciju prikupljenih podataka,
- kontrolu pristupa,
- auditabilnost i transparentnost obrade podataka.