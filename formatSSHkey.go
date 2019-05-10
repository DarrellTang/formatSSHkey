package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println("sshFormatKey: " + e.Error())
		os.Exit(0)
	}
}

func findKeyFormat(sshkey string) string {
	if strings.Contains(sshkey, "---- BEGIN SSH2 PUBLIC KEY ----") {
		return "ssh2"
	} else if strings.Contains(sshkey, "ssh-rsa") {
		return "openssh"
	} else {
		return "wrongformat"
	}

}

func replaceNewlines(sshkey string) string {
	if strings.Contains(sshkey, "\r\n") {
		return strings.ReplaceAll(sshkey, "\r\n", "\\n")
	} else {
		return strings.ReplaceAll(sshkey, "\n", "\\n")
	}
}

func replaceDoubleQuotes(sshkey string) string {
	return strings.ReplaceAll(sshkey, "\"", "\\\"")
}

func replaceComments(sshkey, lasteight string) string {
	indexOfComment := strings.Index(sshkey, "Comment: ")
	indexOfLastDoubleQuote := strings.LastIndex(sshkey, "\"")
	oldcomment := sshkey[indexOfComment:indexOfLastDoubleQuote]
	return strings.Replace(sshkey, oldcomment, "Comment: \\\"Key-ID "+lasteight+"\\", 1)
}

func convertToSSH2(sshkey, keyfilepath string) string {
	var (
		cmdOutput []byte
		err       error
	)
	cmdName := "ssh-keygen"
	cmdArgs := []string{"-e", "-f", keyfilepath}
	if cmdOutput, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running ssh-keygen command: ", err)
		os.Exit(1)
	}
	err = ioutil.WriteFile("ssh2.txt", cmdOutput, 0644)
	check(err)
	fmt.Println("\nWrote RFC 4716/SSH2 key to ssh2.txt")
	return string(cmdOutput)
}

func convertToOpenssh(sshkey, keyfilepath string) string {
	var (
		cmdOutput []byte
		err       error
	)
	cmdName := "ssh-keygen"
	cmdArgs := []string{"-i", "-f", keyfilepath}
	if cmdOutput, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running ssh-keygen command: ", err)
		os.Exit(1)
	}
	err = ioutil.WriteFile("openssh.txt", cmdOutput, 0644)
	check(err)
	fmt.Println("\nWrote Openssh key to openssh.txt")
	return string(cmdOutput)
}

func getFingerprint(keyfile string) (string, string) {
	var (
		cmdOutput []byte
		err       error
	)
	cmdName := "ssh-keygen"
	cmdArgs := []string{"-l", "-E", "MD5", "-f", keyfile}
	if cmdOutput, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running ssh-keygen command: ", err)
		os.Exit(1)
	}
	lasteight := string(cmdOutput)[45:56]
	lasteight = strings.Replace(lasteight, ":", "", -1)
	lasteight = strings.ToUpper(lasteight)
	return string(cmdOutput), lasteight
}

func printHelp() {
	fmt.Println("Usage: formatSSHkey [sshkeyfilename] [clientname]")
	fmt.Println()
	fmt.Println("formatSSHkey will take an ssh keyfile and determine what format it's in (RFC4716/SSH2 or Openssh).")
	fmt.Println("It will then convert the ssh keyfile to the other format and write it to either \"ssh2.txt\" or \"openssh.txt\".")
	fmt.Println("The MD5 hash fingerprint of the openssh format key is printed.")
	fmt.Println("Finally, the client name and key are formatted for copy-pasting into a JSON key-value pair format.")
	fmt.Println("formatSSHkey takes two positional arguments:")
	fmt.Println()
	fmt.Println("[sshkeyfilename]		An ssh Keyfile name (e.g. sshkey.txt or CLIENT_SSH2.pub)")
	fmt.Println("[clientname]			A client name in all caps (e.g. FONCIA or NORDEA)")
	fmt.Println()
	fmt.Println("EXAMPLE")
	fmt.Println("formatSSHkey ./my_key_file.pub clientname")
	fmt.Println()
	os.Exit(0)
}

func main() {
	var (
		keyfilepath string
		clientname  string
		sshkey      string
		keyformat   string
		fingerprint string
		lasteight   string
	)

	if len(os.Args) == 3 {
		keyfilepath = os.Args[1]
		clientname = strings.ToUpper(os.Args[2])
	} else {
		printHelp()
	}

	keyfile, err := ioutil.ReadFile(keyfilepath)
	check(err)

	sshkey = string(keyfile)
	fmt.Println("The current keyfile is:")
	fmt.Println("==================================================")
	fmt.Print(sshkey)
	fmt.Println("==================================================")

	keyformat = findKeyFormat(sshkey)
	switch keyformat {
	default:
		fmt.Println("This is not a properly formatted SSH Key. formatSSHkey will now exit.")
		os.Exit(0)
	case "ssh2":
		fmt.Println("\nThis is an RFC 4716/SSH2 formatted SSH Key.")
		fmt.Println("Converting to Openssh format.")
		fmt.Println("==================================================")
		fmt.Print(convertToOpenssh(sshkey, keyfilepath))
		fmt.Println("==================================================")
		fmt.Println("\nThe fingerprint of this key is:")
		fingerprint, lasteight = getFingerprint("openssh.txt")
		fmt.Println(fingerprint)
		fmt.Println("==================================================")
	case "openssh":
		fmt.Println("\nThis is an Openssh formatted SSH Key.")
		fmt.Println("\nConverting to RFC 4716/SSH2 format.")
		sshkey = convertToSSH2(sshkey, keyfilepath)
		fmt.Println("==================================================")
		fmt.Print(sshkey)
		fmt.Println("==================================================")
		fmt.Println("\nThe fingerprint of this key is:")
		fingerprint, lasteight = getFingerprint(keyfilepath)
		fmt.Println(fingerprint)
		fmt.Println("==================================================")
	}

	sshkey = replaceNewlines(sshkey)
	sshkey = replaceDoubleQuotes(sshkey)
	sshkey = replaceComments(sshkey, lasteight)
	fmt.Println("Copy and paste the below text into the platform databag json file")
	fmt.Println("==================================================")
	fmt.Println("\""+clientname+"\" :", "\""+sshkey+"\"")
	fmt.Println("==================================================")

}
