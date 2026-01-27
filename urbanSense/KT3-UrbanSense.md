# Go-Specifične Bezbednosne Ranjivosti u UrbanSense Sistemu

## Uvod

UrbanSense sistem koristi Go backend servis kao centralnu komponentu za prijem, obradu i skladištenje podataka prikupljenih sa pametnih gradskih senzora. Go je izabran zbog svoje performanse, ugrađene podrške za konkurentnost i bogate standardne biblioteke. Ipak, kao i svaki programski jezik, Go poseduje specifične bezbednosne izazove koji proizilaze iz njegovog runtime okruženja, modela izvršavanja i ekosistema zavisnosti.

Cilj ovog dokumenta je identifikacija i analiza bezbednosnih ranjivosti koje su **usko povezane sa Go jezikom**, kao i razmatranje potencijalnih scenarija napada i načina njihovog ublažavanja u kontekstu UrbanSense sistema.

---

## 1. Ranjivosti u Go standardnoj biblioteci i runtime okruženju

### Opis ranjivosti  
Go standardna biblioteka je opsežna i često se koristi kao primarni oslonac za mrežnu komunikaciju, parsiranje podataka i kriptografske operacije. Ranjivosti u samoj biblioteci mogu direktno ugroziti aplikaciju, čak i kada je aplikacioni kod ispravno implementiran.

### Specifičnost u Go  
Za razliku od jezika koji se više oslanjaju na spoljne biblioteke, Go aplikacije u velikoj meri zavise od standardne biblioteke. Greške u runtime-u ili paketima kao što su `net/http`, `archive` ili `crypto` mogu izazvati uskraćivanje usluge ili neočekivane prekide rada.

### Mogući scenario napada  
Napadač može poslati posebno oblikovan ulaz (npr. HTTP zaglavlja ili arhivske podatke) koji aktivira poznatu ranjivost u standardnoj biblioteci, dovodeći do preopterećenja memorije ili pada Go servisa.

### Moguće mitigacije  
- Redovno ažuriranje Go verzije  
- Praćenje zvanične Go baze ranjivosti  
- Izbegavanje zastarelih verzija standardne biblioteke  

### Reference  
- https://go.dev/doc/security/vuln  
- https://go.dev/doc/security/vuln/database  

---

## 2. Zavisnosti i Go module ekosistem (Supply Chain rizici)

### Opis ranjivosti  
Go koristi `go.mod` i `go.sum` za upravljanje zavisnostima, ali ne sprečava automatski korišćenje ranjivih modula. Aplikacija može biti kompromitovana kroz zavisnost, iako sopstveni kod ne sadrži greške.

### Specifičnost u Go  
Go module sistem omogućava lako uključivanje velikog broja paketa, ali bez obavezne bezbednosne validacije tokom build procesa. Bez dodatnih alata, ranjive verzije mogu ostati neprimećene.

### Mogući scenario napada  
Ranjiv MQTT klijent ili NoSQL drajver može omogućiti napadaču neovlašćen pristup ili manipulaciju podacima koji se obrađuju u UrbanSense sistemu.

### Moguće mitigacije  
- Korišćenje `govulncheck` alata  
- Ograničavanje broja eksternih zavisnosti  
- Redovno ažuriranje modula  

### Reference  
- https://go.dev/doc/security/vuln/database  
- https://www.reddit.com/r/golang/comments/19dk7t4/vulnerabilities_and_best_practices/  

---

## 3. Konkurentnost i iscrpljivanje resursa (gorutine i kanali)

### Opis ranjivosti  
Go omogućava jednostavno kreiranje gorutina, ali ne ograničava njihov broj. Neadekvatno upravljanje konkurentnim izvršavanjem može dovesti do curenja gorutina i iscrpljivanja sistemskih resursa.

### Specifičnost u Go  
Go model konkurentnosti je osnovna karakteristika jezika. Problemi vezani za gorutine i kanale ne pojavljuju se u ovom obliku u tradicionalnim jezicima sa drugačijim modelima izvršavanja.

### Mogući scenario napada  
Napadač može generisati veliki broj zahteva ili poruka sa senzora, pri čemu svaki zahteva pokretanje nove gorutine, što može dovesti do uskraćivanja usluge.

### Moguće mitigacije  
- Ograničavanje broja gorutina  
- Korišćenje worker pool obrazaca  
- Praćenje i kontrola životnog ciklusa gorutina  

### Reference  
- https://my.f5.com/manage/s/article/K000152727  
- https://go.dev/doc/security/vuln/  

---

## 4. Ignorisanje grešaka i upotreba panic mehanizma

### Opis ranjivosti  
Go zahteva eksplicitnu obradu grešaka, ali ne primorava programera da ih zaista obradi. Ignorisane greške ili neadekvatna upotreba `panic` mehanizma mogu dovesti do nepredvidivog ponašanja sistema.

### Specifičnost u Go  
Go nema izuzetke u klasičnom smislu, već se oslanja na povratne vrednosti grešaka. Ovo zahteva disciplinu u pisanju koda, jer se greške lako mogu zanemariti.

### Mogući scenario napada  
Neobrađena greška tokom obrade senzorskih podataka može izazvati `panic` i oboriti ceo backend servis, čime se prekida prikupljanje podataka.

### Moguće mitigacije  
- Dosledna obrada svih grešaka  
- Ograničena i kontrolisana upotreba `panic` i `recover`  
- Centralizovano logovanje grešaka  

### Reference  
- https://www.reddit.com/r/golang/comments/19dk7t4/vulnerabilities_and_best_practices/  

---

## 5. Korišćenje `unsafe` i `cgo` paketa

### Opis ranjivosti  
Paketi `unsafe` i `cgo` omogućavaju izlazak iz Go memory safety modela. Njihova upotreba može dovesti do grešaka koje inače nisu moguće u čistom Go kodu.

### Specifičnost u Go  
Go je po prirodi memory-safe jezik. Korišćenjem ovih mehanizama gube se osnovne bezbednosne garancije koje Go pruža.

### Mogući scenario napada  
Greška u radu sa pointerima ili nativnim kodom može omogućiti curenje memorije ili nepredvidivo ponašanje aplikacije.

### Moguće mitigacije  
- Izbegavanje `unsafe` i `cgo` gde god je moguće  
- Stroga revizija koda koji ih koristi  
- Izolacija niskonivojskog koda  

### Reference  
- https://my.f5.com/manage/s/article/K000152727  

---

## STRIDE analiza Go backend servisa

| STRIDE kategorija | Opis pretnje | Zašto je relevantno za Go |
|------------------|-------------|---------------------------|
| Spoofing | Lažno predstavljanje izvora podataka | Go servisi često direktno komuniciraju sa mrežnim izvorima |
| Tampering | Manipulacija podacima u obradi | Greške u konkurentnosti mogu dovesti do nekonzistentnih stanja |
| Repudiation | Nemogućnost dokazivanja aktivnosti | Neadekvatno logovanje grešaka i događaja |
| Information Disclosure | Curenje podataka | Panic i stack trace mogu otkriti interne detalje |
| Denial of Service | Iscrpljivanje resursa | Neograničene gorutine i ranjivosti u runtime-u |
| Elevation of Privilege | Neovlašćeni pristup | Korišćenje unsafe i ranjivih zavisnosti |
