# Kontrolna Tačka 3 – Bezbednosna analiza (UrbanSense)

Ovaj repozitorijum sadrži rezultate KT3 bezbednosne analize za sistem **UrbanSense** (Smart City Environmental Monitoring System).
Fokus KT3 je identifikacija realnih pretnji i napada nad arhitekturom sistema, izrada stabla napada, kao i implementacija ranjive i mitigovane verzije za praktične napade.

## Sadržaj

KT3 obuhvata:
- Opis glavne pretnje (u kontekstu sistema UrbanSense)
- Opise četiri napada (2 praktična + 2 teorijska)
- Stablo napada (sa ID čvorovima i objašnjenjima)
- Mitigacije za svaki napad (teorijske za sve, praktične + demonstracija za 2 praktična)
- Implementacije ranjive i mitigovane verzije (za praktične napade)

---

## Glavna pretnja (high-level)

Glavna pretnja u okviru UrbanSense sistema odnosi se na kompromitaciju pouzdanosti rada sistema kroz kombinaciju neadekvatno zaštićene komunikacije i nedostatka kontrola potrošnje resursa na servisnim komponentama. U praksi, to omogućava neautorizovanim akterima da (1) ubace lažne senzorske podatke u sistem i naruše integritet podataka koji se koriste kao izvor istine, ili da (2) degradiraju dostupnost ključnih servisa slanjem prevelikih ili previše učestalih zahteva, čime se otežava ili onemogućava obrada legitimnih merenja.

Pretnja je mapirana na sledeće osetljive resurse i bezbednosne ciljeve:
- **Podaci sa senzora** i njihova obrada: primarno je ugrožen **integritet**, jer lažne poruke mogu dovesti do pogrešnih zaključaka i odluka.
- **Sistem za alarmiranje i obaveštavanje**: integritet merenja direktno utiče na tačnost alarmiranja, dok degradacija servisa ugrožava blagovremenu reakciju.
- **Dostupnost pipeline-a (MQTT broker i Go API)**: primarno je ugrožena **dostupnost**, jer DoS scenariji mogu prekinuti prijem i obradu podataka ili značajno povećati latenciju.

U okviru KT3, ova glavna pretnja je razložena kroz četiri napada: dva praktična (MQTT neautorizovan publish i DoS nad Go API), kao i dva teorijska napada zasnovana na realnim CVE zapisima (request smuggling u Go `net/http` i DoS nad Mosquitto brokerom kroz crafted CONNECT paket). Detaljna razrada, mitigacije i retest nalaze se u odgovarajućim `exploitX.md` fajlovima i u stablu napada (`attack_tree.md`).

---

## Napadi obuhvaćeni analizom

U okviru KT3 analizirana su 4 napada:

### Praktični napad 1 - MQTT neautorizovan publish (lažni senzor)
- **Cilj:** demonstracija narušavanja integriteta podataka slanjem poruka na MQTT topic bez adekvatne autentifikacije/ACL kontrole.
- **Uticaj:** lažni/izmenjeni senzorski podaci u obradi i skladištu, potencijalno pogrešno alarmiranje.
- **Status:** Implementirano (ranjivo + mitigovano).

Dokument: `exploits/exploit1.md`

### Praktični napad 2 - DoS nad Go API servisom (unbounded request body / resource exhaustion)
- **Cilj:** demonstracija narušavanja dostupnosti kroz zadržavanje konekcija / iscrpljivanje resursa usled neadekvatnih timeout podešavanja i ograničenja.
- **Uticaj:** degradacija performansi ili pad API servisa, prekid obrade i prijema podataka.
- **Status:** Implementirano (ranjivo + mitigovano).

Dokument: `exploits/exploit2.md`

### Teorijski napad 3 - Request smuggling u Go `net/http` (CVE)
- **Cilj:** analiza realne ranjivosti iz sveta (CVE) i njenog potencijalnog uticaja u scenariju reverse proxy / više-hop okruženja.
- **Status:** Teorijska analiza (bez implementacije).

Dokument: `exploits/exploit3.md`

### Teorijski napad 4 - Mosquitto DoS kroz crafted CONNECT packet (CVE)
- **Cilj:** analiza realne ranjivosti iz sveta (vendor advisory/CVE) i njenog uticaja na MQTT sloj.
- **Status:** Teorijska analiza (bez implementacije).

Dokument: `exploits/exploit4.md`

---

## Stablo napada

Stablo napada se nalazi u fajlu:
- `attack_tree.md`

U stablu su svi čvorovi označeni jedinstvenim ID-jevima (npr. `N0`, `N1`, `N1.1`...), a ispod stabla su dati opisi:
- uslov(i) za uspeh napada,
- pogođene komponente,
- bezbednosni cilj koji je narušen,
- veze sa mitigacijama.

---

## Struktura repozitorijuma

- `README.md` – ovaj fajl
- `exploits/` – opisi svih napada
  - `exploit1.md` – MQTT neautorizovan publish
  - `exploit2.md` – DoS nad Go API servisom
  - `exploit3.md` – Request smuggling u Go `net/http`
  - `exploit4.md` – Mosquitto DoS kroz crafted CONNECT packet
- `attack_tree.md` – Stablo napada sa opisima čvorova
- `implementation/vulnerable/` – Implementacija ranjivih verzija (za praktične napade)
- `implementation/mitigated/` – Implementacija mitigovanih verzija (za praktične napade)

---

## Kako koristiti ovaj repozitorijum (pregled)

1. Pročitati **glavnu pretnju** i kontekst u ovom README-u.
2. Pogledati **attack_tree.md** (stablo + opisi čvorova).
3. Otvoriti `exploits/exploit1.md` i `exploits/exploit2.md`:
   - koraci napada,
   - dokaz uspeha,
   - mitigacija,
   - retest nakon mitigacije.
4. Pregledati implementacije:
   - `implementation/vulnerable/` (ranjivo)
   - `implementation/mitigated/` (mitigovano)

> Napomena: Za praktične napade dostupni su video snimci napada i mitigacije.

---

## Reference (osnovne)

> Svaki `exploitX.md` sadrži posebnu listu referenci relevantnih za konkretan napad.
---