package main

import (
	"testing"
)

func TestFindKeyFormat(t *testing.T) {
	tables := []struct {
		key    string
		format string
	}{
		{`---- BEGIN SSH2 PUBLIC KEY ----
		Comment: "2048-bit RSA, converted by dtang@fr@US-WS-314 from OpenSSH"
		AAAAB3NzaC1yc2EAAAADAQABAAABAQCeyOzR4zGPxDdVySvVkX6dBWrzTNRb8menTXtPmG
		OD+EQxw5Lna4/Dg2uH4Gi9ZNnVhOqwYXvGPr8DMQwVykdwVhc7Z1Ez+2XMa2py9i/MnK5O
		veobE7MWqBJ3RgwBK+JI1MbMQeDuiTSPJv+tt7tAJ6sJFigbgfgnineUo4AUs2A8KshPnV
		sJmBgSWggN0Kk+xSp5ScmkS+vh7Dm2szM6dUw384nJ464U66ej7NPl+olnbgj1IVqIegiI
		zofV0XFkSGCyinlwq6UqQ3eNAg9fOv9rG9TULjrcCH3/GAN7x1E4WCKrEWoHuEU+BZRFiK
		lecz7BaJmbTY0oMaO+BXwx
		---- END SSH2 PUBLIC KEY ----`, "ssh2"},
		{"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCeyOzR4zGPxDdVySvVkX6dBWrzTNRb8menTXtPmGOD+EQxw5Lna4/Dg2uH4Gi9ZNnVhOqwYXvGPr8DMQwVykdwVhc7Z1Ez+2XMa2py9i/MnK5OveobE7MWqBJ3RgwBK+JI1MbMQeDuiTSPJv+tt7tAJ6sJFigbgfgnineUo4AUs2A8KshPnVsJmBgSWggN0Kk+xSp5ScmkS+vh7Dm2szM6dUw384nJ464U66ej7NPl+olnbgj1IVqIegiIzofV0XFkSGCyinlwq6UqQ3eNAg9fOv9rG9TULjrcCH3/GAN7x1E4WCKrEWoHuEU+BZRFiKlecz7BaJmbTY0oMaO+BXwx dtang@fr@US-WS-314", "openssh"},
		{"i've got a lovely bunch of coconuts", "wrongformat"},
	}

	for _, table := range tables {
		keyformat := findKeyFormat(table.key)
		if keyformat != table.format {
			t.Errorf("Key format of %s was wrong, got %s, want %s", table.key, keyformat, table.format)
		}
	}
}

func TestReplaceNewlines(t *testing.T) {
	tables := []struct {
		str    string
		result string
	}{
		{`---- BEGIN SSH2 PUBLIC KEY ----
Comment: "2048-bit RSA, converted by dtang@fr@US-WS-314 from OpenSSH"
AAAAB3NzaC1yc2EAAAADAQABAAABAQCeyOzR4zGPxDdVySvVkX6dBWrzTNRb8menTXtPmG
OD+EQxw5Lna4/Dg2uH4Gi9ZNnVhOqwYXvGPr8DMQwVykdwVhc7Z1Ez+2XMa2py9i/MnK5O
veobE7MWqBJ3RgwBK+JI1MbMQeDuiTSPJv+tt7tAJ6sJFigbgfgnineUo4AUs2A8KshPnV
sJmBgSWggN0Kk+xSp5ScmkS+vh7Dm2szM6dUw384nJ464U66ej7NPl+olnbgj1IVqIegiI
zofV0XFkSGCyinlwq6UqQ3eNAg9fOv9rG9TULjrcCH3/GAN7x1E4WCKrEWoHuEU+BZRFiK
lecz7BaJmbTY0oMaO+BXwx
---- END SSH2 PUBLIC KEY ----
`, "---- BEGIN SSH2 PUBLIC KEY ----\\nComment: \"2048-bit RSA, converted by dtang@fr@US-WS-314 from OpenSSH\"\\nAAAAB3NzaC1yc2EAAAADAQABAAABAQCeyOzR4zGPxDdVySvVkX6dBWrzTNRb8menTXtPmG\\nOD+EQxw5Lna4/Dg2uH4Gi9ZNnVhOqwYXvGPr8DMQwVykdwVhc7Z1Ez+2XMa2py9i/MnK5O\\nveobE7MWqBJ3RgwBK+JI1MbMQeDuiTSPJv+tt7tAJ6sJFigbgfgnineUo4AUs2A8KshPnV\\nsJmBgSWggN0Kk+xSp5ScmkS+vh7Dm2szM6dUw384nJ464U66ej7NPl+olnbgj1IVqIegiI\\nzofV0XFkSGCyinlwq6UqQ3eNAg9fOv9rG9TULjrcCH3/GAN7x1E4WCKrEWoHuEU+BZRFiK\\nlecz7BaJmbTY0oMaO+BXwx\\n---- END SSH2 PUBLIC KEY ----\\n"},

		{`



`, "\\n\\n\\n\\n"},
	}

	for _, table := range tables {
		result := replaceNewlines(table.str)
		if result != table.result {
			t.Errorf("Replacing new line characters of %s was wrong, got %s, want %s", table.str, result, table.result)
		}
	}
}

func TestReplaceDoubleQuotes(t *testing.T) 
	
}

// func TestReplaceComments(t *testing.T)  {

// }

// func TestConvertToSSH2(t *testing.T)  {

// }

// func TestConvertToOpenssh(t *testing.T)  {

// }

// func TestGetFingerprint(t *testing.T)  {

// }
