package back

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/RDLrpl/fenrir/utility"
	"github.com/schollz/progressbar/v3"
)

func AutoDownloadCAAU() error {
	temp, base, lk := CAAUOSCHECK()
	if temp == "" {
		return fmt.Errorf("%s", lk)
	}

	resp, err := http.Get(lk)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Read Guide! CAAU lib site is broken. You can install CAAU manually%s", resp.Status)
	}

	err = os.MkdirAll(filepath.Dir(temp), os.ModePerm)
	if err != nil {
		return err
	}

	out, err := os.Create(temp)
	if err != nil {
		return err
	}
	defer out.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"chromium downloading",
	)
	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		return err
	}
	out.Close()
	fmt.Println("\nUnpacking CAAU...")
	if err := unzip(temp, base); err != nil {
		return fmt.Errorf("unpack error: %v", err)
	}

	err = os.RemoveAll("./.tmp")
	if err != nil {
		return fmt.Errorf("delete error: %v", err)
	}
	return nil
}

func CAAUOSCHECK() (string, string, string) {

	if runtime.GOOS == "linux" && runtime.GOARCH == "amd64" {
		dwnllink := utility.FenrirCAAUpack["linux-amd"]

		base := "./fpkg/CAAU/linuxamd"
		temp := "./.tmp/linuxchromium.zip"

		return temp, base, dwnllink
	}
	if runtime.GOOS == "windows" && runtime.GOARCH == "amd64" {
		dwnllink := utility.FenrirCAAUpack["windows-amd"]

		base := "./fpkg/CAAU/windowsamd"
		temp := "./.tmp/windowschromium.zip"

		return temp, base, dwnllink
	} else {
		rt := "Unfortunately, your OS doesn't support the fenrirCAAU package.\n" +
			"You need to use Windows x64 or Linux x64.\n" +
			"Alternatively, you can modify the code to make the library work.\n" +
			"It will not work on Android or iOS, but should work on ARM devices like macOS."

		return "", "", rt
	}
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
