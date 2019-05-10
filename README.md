Usage: formatSSHkey [sshkeyfilename] [clientname]

formatSSHkey will take an ssh keyfile and determine what format it's in (RFC4716/SSH2 or Openssh).
It will then convert the ssh keyfile to the other format and write it to either "ssh2.txt" or "openssh.txt".
The MD5 hash fingerprint of the openssh format key is printed.
Finally, the client name and key are formatted for copy-pasting into a JSON key-value pair format.
formatSSHkey takes two positional arguments:

[sshkeyfilename]                An ssh Keyfile name (e.g. sshkey.txt or CLIENT_SSH2.pub)
[clientname]                    A client name in all caps (e.g. FONCIA or NORDEA)

EXAMPLE
formatSSHkey ./my_key_file.pub clientname