package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	template := "template-us.dcm"
	outfile := fmt.Sprintf("US_%d.dcm", time.Now().UnixNano())

	// 1. Copia template → file nuovo
	if err := copyFile(template, outfile); err != nil {
		log.Fatalf("copy failed: %v", err)
	}

	// 2. Aggiorna alcuni tag (DCMTK: dcmodify)
	update := [][]string{
		{"-i", fmt.Sprintf("(0010,0010)=DOE^JOHN")},
		{"-i", fmt.Sprintf("(0010,0020)=PATID_%d", time.Now().Unix())},
		{"-i", fmt.Sprintf("(0008,0020)=%s", time.Now().Format("20060102"))},
		{"-i", fmt.Sprintf("(0008,0030)=%s", time.Now().Format("150405"))},
	}

	for _, u := range update {
		args := append(u, outfile)
		cmd := exec.Command("dcmodify", args...)
		if err := cmd.Run(); err != nil {
			log.Fatalf("dcmodify: %v", err)
		}
	}

	// 3. Invia al PACS
	cmd := exec.Command("storescu",
		"-aec", "MYPACS",
		"127.0.0.1", "3000",
		outfile,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("Invio %s ...", outfile)
	if err := cmd.Run(); err != nil {
		log.Fatalf("storescu error: %v", err)
	}

	log.Println("Inviato!")
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
