package main

import (
	"fmt"
	"github.com/imdario/identita"
	"github.com/spf13/cobra"
)

var IdentitaCmd = &cobra.Command{
	Use:   "identita",
	Short: "Identita is a cryptographic system to provide unique and secure personal identification",
	Long: `A cryptographic system to provied unique and secure personal identification based in government-issued proof of identities.

Documentation at https://dario.im/identita/`,
}

var genkeyCmd = &cobra.Command{
	Use:   "genkey",
	Short: "Generate Identita key",
	Long:  `Generate Identita key, based on Curve25519. One key is for one Identita Authority (IA) and it is used to sign payloads.`,
	Run:   genkey,
}

var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign a payload",
	Run:   sign,
}

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a signature",
	Run:   verify,
}

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Create Identita Personas from JSON files following the specification",
	Run:   issue,
}

var (
	out, key, password     string
	outFlagDefinition      = [...]string{"out", "o", "", "output file  - default stdout"}
	keyFlagDefinition      = [...]string{"key", "k", "", "key to use"}
	passwordFlagDefinition = [...]string{"password", "p", "", "Identita Persona's password"}
)

func genkey(cmd *cobra.Command, args []string) {
	if err := identita.GenerateKey(out); err != nil {
		panic(err)
	}
}

func sign(cmd *cobra.Command, args []string) {
	for _, in := range args {
		fmt.Printf("%s: ", in)
		if err := identita.SignFile(key, in); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("done")
		}
	}
}

func verify(cmd *cobra.Command, args []string) {
	for _, in := range args {
		fmt.Printf("%s: ", in)
		if err := identita.VerifyFile(key, in); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ok")
		}
	}
}

func issue(cmd *cobra.Command, args []string) {
	for _, in := range args {
		fmt.Printf("%s: ", in)
		id, err := identita.IssueFromFile(key, in, password)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(id))
		}
	}
}

func main() {
	IdentitaCmd.AddCommand(genkeyCmd, signCmd, verifyCmd, issueCmd)
	genkeyCmd.Flags().StringVarP(&out, outFlagDefinition[0], outFlagDefinition[1], outFlagDefinition[2], outFlagDefinition[3])
	signCmd.Flags().StringVarP(&key, keyFlagDefinition[0], keyFlagDefinition[1], keyFlagDefinition[2], keyFlagDefinition[3])
	signCmd.Flags().StringVarP(&out, outFlagDefinition[0], outFlagDefinition[1], outFlagDefinition[2], outFlagDefinition[3])
	verifyCmd.Flags().StringVarP(&key, keyFlagDefinition[0], keyFlagDefinition[1], keyFlagDefinition[2], keyFlagDefinition[3])
	issueCmd.Flags().StringVarP(&key, keyFlagDefinition[0], keyFlagDefinition[1], keyFlagDefinition[2], keyFlagDefinition[3])
	issueCmd.Flags().StringVarP(&password, passwordFlagDefinition[0], passwordFlagDefinition[1], passwordFlagDefinition[2], passwordFlagDefinition[3])
	IdentitaCmd.Execute()
}
