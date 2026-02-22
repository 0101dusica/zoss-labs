# Stablo napada

## Glavna pretnja

N0 predstavlja agregiranu pretnju nad UrbanSense sistemom: neautorizovani akteri mogu ugroziti pouzdanost rada sistema i tačnost podataka kroz zloupotrebu komunikacionih i servisnih komponenti. Posledice uključuju unošenje lažnih senzorskih podataka (integritet) i degradaciju ili prekid rada servisa (dostupnost).

- N0: Kompromitacija pouzdanosti UrbanSense sistema kroz neautorizovan pristup i iscrpljivanje resursa
  - N1: MQTT neautorizovan publish (lažni senzor) [Praktični]
  - N2: DoS nad Go API (neograničen request body / resursna potrošnja) [Praktični]
  - N3: HTTP request smuggling u Go `net/http` (CVE-2025-22871) [Teorijski]
  - N4: DoS nad Mosquitto brokerom kroz crafted CONNECT paket (CVE-2017-7651) [Teorijski]

---

## Opisi čvorova

### **N0: Kompromitacija pouzdanosti UrbanSense sistema kroz neautorizovan pristup i iscrpljivanje resursa**
- **Opis:** Glavna pretnja obuhvata scenarije u kojima napadač može narušiti integritet senzorskih podataka ili dostupnost centralnih komponenti (MQTT broker, Go servis), čime se kompromituje obrada, agregacija i potencijalno alarmiranje.
- **Pogođene komponente:** MQTT sloj, Go backend servis, mrežna komunikacija.
- **Bezbednosni cilj:** Integritet i dostupnost (primarno); poverljivost (sekundarno).
- **Posledice:** Lažni podaci u sistemu, degradacija performansi, prekid prijema/obrade merenja.
- **Mitigacija (sažeto):** Autentifikacija i autorizacija (MQTT), limitiranje resursa (API), ažuriranje ranjivih verzija komponenti.

---

### **N1: MQTT neautorizovan publish (lažni senzor) [Praktični]**
- **Opis:** U ranjivoj konfiguraciji brokera omogućeno je anonimno povezivanje i publish bez ACL pravila. Napadač publikuje poruke na topic koji backend prihvata, simulirajući legitimni senzor.
- **Pogođene komponente:** MQTT broker (Mosquitto) i subscriber/ingest deo backend-a.
- **Bezbednosni cilj:** Integritet (primarno); poverljivost (sekundarno ako je moguć subscribe).
- **Preduslovi:** Broker dostupan; `allow_anonymous true`; nema ACL ograničenja.
- **Očekivani ishod:** Backend/observer vidi lažnu poruku na `sensors/+/data`.
- **Mitigacija (sažeto):** `allow_anonymous false`, password auth + ACL (least privilege), opcionalno TLS/mTLS.

---

### **N2: DoS nad Go API (neograničen request body / resursna potrošnja) [Praktični]**
- **Opis:** Endpoint prihvata velike payload-e bez limita i timeouts. Napadač šalje prevelike request body-je ili veći broj paralelnih zahteva, izazivajući povećanje CPU/memorije i degradaciju dostupnosti.
- **Pogođene komponente:** Go API (ingest endpoint), sistemski resursi hosta.
- **Bezbednosni cilj:** Dostupnost (primarno).
- **Preduslovi:** Endpoint javno dostupan; nema limita veličine tela; nema timeouts/rate limiting.
- **Očekivani ishod:** Povećana latencija i/ili opterećenje resursa; potencijalno timeouts.
- **Mitigacija (sažeto):** `http.MaxBytesReader`, server timeout-i (`ReadHeaderTimeout`, `ReadTimeout`, `WriteTimeout`), rate limiting.

---

### **N3: HTTP request smuggling u Go `net/http` (CVE-2025-22871) [Teorijski]**
- **Opis:** U određenim verzijama Go `net/http` parser može drugačije interpretirati neispravan chunked format u odnosu na proxy sloj. Neusklađeno parsiranje može omogućiti “smuggled” request u istoj konekciji.
- **Pogođene komponente:** Reverse proxy/gateway + Go backend (`net/http`).
- **Bezbednosni cilj:** Integritet (primarno); poverljivost i autorizacija (sekundarno).
- **Preduslovi:** Ranjiva Go verzija; proxy okruženje sa parser mismatch scenarijem.
- **Očekivani ishod:** Backend obradi dodatni skriveni request mimo očekivanog toka.
- **Mitigacija (sažeto):** Ažuriranje Go verzije (fix), striktno parsiranje, hardening proxy sloja, korelaciono logovanje.

---

### **N4: DoS nad Mosquitto brokerom kroz crafted CONNECT paket (CVE-2017-7651) [Teorijski]**
- **Opis:** U ranjivim verzijama Mosquitto brokera, maliciozni CONNECT paket može izazvati visoku potrošnju memorije. Ponavljanjem konekcija može doći do OOM i prekida rada brokera.
- **Pogođene komponente:** MQTT broker (Mosquitto).
- **Bezbednosni cilj:** Dostupnost (primarno).
- **Preduslovi:** Ranjiva verzija brokera; broker dostupan napadaču; nema network limit-a.
- **Očekivani ishod:** Rast memorije, degradacija, crash brokera.
- **Mitigacija (sažeto):** Ažuriranje brokera, network rate limiting/connection limiting, izolacija brokera, dodatna auth/ACL zaštita.