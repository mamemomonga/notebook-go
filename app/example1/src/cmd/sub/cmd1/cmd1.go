package cmd1

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	flagC string
)

func CobraCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "cmd1",
		Short: "サブコマンド1(Short)",
		Long:  "サブコマンド1(Long)",
		Run:   Run,
		PreRun: PreRun,
	}
	// ローカルフラグ
	c.PersistentFlags().StringVarP(&flagC, "flagc", "c", "", "フラグC(必須)")
	return c
}

// ローカルフラグの必須チェック
func PreRun(cmd *cobra.Command, args []string) {
	if flagC == "" {
		log.Printf("alert: -c(--flagc)は必須です")
		cmd.Usage()
		os.Exit(1)
	}
}

func Run(cmd *cobra.Command, args []string) {
	log.Println("info: cmd1 Run")

	// フラグの取得
	flgA := "FALSE"
	if v, err := cmd.Flags().GetBool("flaga"); err == nil {
		if(v) {
			flgA = "TRUE"
		}
	} else {
		log.Fatal(err)
	}

	flgB := ""
	if v, err := cmd.Flags().GetString("flagb"); err == nil {
		flgB = v
	} else {
		log.Fatal(err)
	}


	flgC := ""
	if v, err := cmd.Flags().GetString("flagc"); err == nil {
		flgC = v
	} else {
		log.Fatal(err)
	}

	log.Printf("info: FlagA: %s / FlagB: %s / FlagC: %s", flgA, flgB, flgC )
}

