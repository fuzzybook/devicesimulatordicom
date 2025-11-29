✅ Alcuni software / simulatori già pronti
SIMUS

SIMUS è un simulatore open-source (parte di una toolbox per MATLAB) per imaging ecografico computazionale — genera campi acustici e segnali RF, e può simulare immagini ecografiche per sonde lineari o convex. 
sciencedirect.com

Ottimo se vuoi fare ricerca, sviluppo o “simulazione a basso livello” (propagazione onde, beamforming, prove di algoritmi).

Limite: non è “plug-and-play” per generare DICOM come da macchinario reale — è più per ricerca/sviluppo.

ImFusion 3D Ultrasound Suite (o simili strumenti di simulazione basati su CT/MR)

Ci sono segnalazioni di utenti che usano ImFusion per “simulare sweep ecografici” a partire da dati DICOM (CT/MR). 
ImFusion Support Forum

Potenzialmente utile se vuoi trasformare dataset radiologici (CT/MR) in “simulazioni ecografiche”.

Svantaggi: spesso la versione “demo” può essere limitata (export, numero di frame, licenze). 
ImFusion Support Forum

Simulatori didattici / formativi “full-stack”

Se invece cerchi simulazione clinica o formativa (come “provare ad usare un ecografo su paziente virtuale”), ci sono piattaforme commerciali e semi-commerciali:

Soluzione	Note
SonoSim (con il suo “SonoSimulator®”)	Simulator per apprendimento ecografia: offre casi reali (normalità + patologie), è pensato per training, non per generazione DICOM da zero. 
sonosim.com
+1

Intelligent Ultrasound / HeartWorks, ScanTrainer, BodyWorks Eve, ecc.	Simulatori “high-fidelity” per ecografia clinica/education, a volte con manichini o interfacce dedicate. 
healthysimulation.com
+2
medicalexpo.com
+2

Soluzioni VR / 3D immersive come 3D Organon Ultrasound Simulator	Permettono “scansioni” virtuali su modelli 3D, senza necessità di manichino fisico. 
3D ORGANON
Emulatore DICOM “finto”

Infine, se il tuo scopo è testare un sistema PACS / flusso DICOM (non la generazione realistica dell’immagine ecografica), esiste anche un approccio più semplice:

Esistono “DICOM simulator / generator node” usati per test di reti DICOM / PACS, che possono generare studi “dummy”. 
pukka-j.com
+2
UK Digital Imaging Forum
+2

Questo può servire per verificare che il tuo server DIMSE + PACS smistino immagini correttamente, senza dover avere un vero ecografo o un simulatore avanzato.

⚠️ A cosa prestare attenzione / cosa aspettarsi

Realismo vs praticità: simulazioni open-source come SIMUS o ImFusion sono più “teoriche/ricerca”, meno “user friendly”. I sistemi commerciali sono più completi ma costano e di solito sono pensati per training, non per generazione DICOM in serie.

Licenza / usabilità: verifica sempre licenze (open-source vs. commerciali), restrizioni su export, simulazioni, “uso in produzione”.

Output DICOM: generare un’immagine ecografica “realistica” e confezionarla in uno studio DICOM come se fosse da un vero ecografo è complesso, e pochi simulatori lo fanno “chiavi in mano”.

Obiettivo: è importante sapere cosa vuoi ottenere — test PACS/DIMSE, sviluppo algoritmi, training operatore, simulazione paziente, ecc. Ogni soluzione ha il suo dominio.

🎯 Il mio consiglio in base a cosa vuoi fare

Dipende dal tuo obiettivo:

Vuoi testare e sviluppare un flusso DICOM / PACS → usa un “DICOM simulator / generator node” (più semplice, leggero).

Vuoi ricerca / sviluppo algoritmi di imaging / beamforming → guarda a SIMUS o simili simulatori open-source.

Vuoi simulare un ecografo “uso clinico/training” → orientati a SonoSim, Intelligent Ultrasound/HeartWorks/ScanTrainer, oppure a soluzioni VR come 3D Organon.

Vuoi trasformare CT/MR in ecografie simulate → un tool tipo ImFusion può essere utile (ma con limitazioni).

Se vuoi — posso provare a trovare 5–10 soluzioni open-source attive oggi (free / MIT / GPL / academic) che puoi scaricare subito e sperimentare per simulare un ecografo (sia per immagini, sia per DICOM). Vuoi che le cerchi?

Hai detto:
devo testare e sviluppare un flusso DICOM / PACS
ChatGPT ha detto:

Perfetto — per testare e sviluppare un flusso DICOM/PACS non ti serve un vero ecografo, e nemmeno un simulatore fisico complesso.
Quello che serve davvero è un DICOM modality simulator (un “ecografo software”) che parla DIMSE (C-STORE, C-FIND, C-MOVE…) e invia/stampa/esegue query come una modality reale.

Ti elenco le soluzioni migliori — open-source, gratuite e commerciali — che usano davvero DIMSE/DCMTK e possono simulare un ecografo in modo fedele al protocollo.

✅ 1. DCMTK + storescu / findscu / movescu (soluzione più semplice & affidabile)

È lo standard de facto per testare PACS senza una modality vera.

Con DCMTK puoi:

simulare invio immagini: storescu

simulare query worklist o PACS: findscu

simulare retrieve: movescu

generare studi di test con dcmj2pnm, dcmodify, img2dcm

Per simulare un ecografo:

Prepari un set di file DICOM (puoi usare template US).

Invochi:

storescu -aec PACS_AE PACS_IP PACS_PORT *.dcm


Il tuo PACS / servizio Go riceve immagini esattamente come da ecografo reale.

👉 Questo è ciò che fanno quasi tutte le software house quando integrano PACS.

Vuoi? Ti preparo:

un simulatore ecosgrafico completo in Go + DCMTK (wrapper storescu)

generazione template DICOM US realistica (study/series/instance UIDs, modality=US, SOP, timestamps, ecc.)

✅ 2. DVTk — Modality Emulator (open-source, molto completo)

DVTk
 è uno dei tool più potenti per testare e “emulare” dispositivi DICOM.
Ha una GUI e può simulare una modality ecografica con:

C-ECHO

C-STORE (invio DICOM)

C-FIND (worklist query)

C-MOVE

SCP e SCU per tutti i servizi principali

Supporta anche:

validazione DICOM

log DIMSE dettagliati

editing dataset

Pro:

gratuito, molto maturo

ottimo per validazione formale DICOM

multi-modalità

Contro:

solo Windows

non integrabile facilmente in pipeline automatizzate (è più GUI che headless)

✅ 3. Orthanc + Orthanc Tools (open-source)

Orthanc è un mini-PACS open-source, ma puoi usarlo anche come modality simulator:

orthanc --simulate modality.json

o via REST API: upload DICOM → Orthanc → forward automatico verso il tuo PACS

Puoi configurarlo per:

comportarsi come un ecografo che invia studi (store-scu)

generare worklist (MWL)

simulare un PACS che risponde ai FIND/MOVE

Ottimo per testare workflow HL7 + DICOM insieme.

✅ 4. dcm4che / dcm4chee tools

Open-source Java toolkit usato da molti vendor.

Comandi:

dcm2xml / xml2dcm per modifiche rapide

dcmqr per C-FIND/C-MOVE

storescu equivalente a quello DCMTK

Potente per pipeline automatiche.

✅ 5. Soluzioni commerciali già pronte (per validazione industriale)

Se ti serve simulazione di un ecografo certificato DICOM (per validazione FDA/CE), esistono simulatori professionali:

PukkaJ DICOM Simulator Node

generatore massivo di casi di test DICOM

coda, flussi, simulazione di modality reali (US, CT, MR)

usato in ambiente clinico e in validazione industriale.

Laurel Bridge (Compass, Navigator)

strumenti enterprise per simulare molte modality contemporaneamente.

costosi, ma ottimi per QA di prodotti PACS.

✅ Quale usare per il tuo caso?

Hai detto che stai già sviluppando un server DIMSE in Go → allora la soluzione migliore è:

👉 DCMTK + un tuo simulatore di ecografo in Go (SCU)

Con questo ottieni:

controllo totale del workflow

automatizzazione perfetta

test ripetibili

nessuna GUI, tutto scriptabile

Il simulatore invia:

1 study → n series → m images
nel formato e con i metadata che vuoi (come un vero ecografo).

📦 Esempio: simulatore di ecografo in Go (SCU) usando DCMTK storescu

Ti preparo un simulatore minimal (ma realistico):

Step 1 — genera un template DICOM US

Template (es. template-us.dcm) con campi principali:

Modality = US

SOP Class = Ultrasound Image Storage

PatientID, StudyInstanceUID, SeriesInstanceUID, ecc.

Te lo posso generare su misura se vuoi.

Step 2 — codice Go

Esegue:

duplicazione di un template US

modifica dei tag (UID, timestamp, patient, ecc.)

invio via storescu

simulator.go (inviatore)

Avvio:

go run simulator.go


Questo ti simula un ecografo che invia un’immagine DICOM.

Vuoi generare 200 immagini? Basta un loop.

Vuoi simulare più serie? Set di Study/Series Instance UID diversi.

Vuoi simulare un flusso continuo? Timer + generatore.

🎯 CONCLUSIONE

Per testare un flusso DICOM/PACS:

✔ DCMTK è lo standard per simulare modality → ti basta

✔ DVTk è eccellente se vuoi GUI e validazione

✔ Orthanc è potente per simulare sia client che server

✔ Soluzioni commerciali → solo se fai QA certificato

Posso prepararti un simulatore avanzato in Go che:

genera studi completi (US)

invia immagini in burst o real-time

simula worklist

supporta C-FIND, C-STORE, C-MOVE

produce log DIMSE come un vero ecografo

Vuoi che te lo scriva?