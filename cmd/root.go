package cmd

import (
    "fmt"
    "os"
    "time"

    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "log-generator",
    Short: "A brief description of your application",
    Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
    // Uncomment the following line if your bare application
    // has an action associated with it:
    //  Run: func(cmd *cobra.Command, args []string) { },
}

var logLevel string
var logOutput string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize(initLogger)

    rootCmd.PersistentFlags().StringVar(&logLevel, "level", "debug", "Enable debug output")
    rootCmd.PersistentFlags().StringVar(&logOutput, "output", "text", "asd")
}

// initConfig reads in config file and ENV variables if set.
func initLogger() {
    zerolog.TimeFieldFormat = time.RFC3339Nano

    switch logLevel {
    case "debug":
        zerolog.SetGlobalLevel(zerolog.DebugLevel)
    case "info":
        zerolog.SetGlobalLevel(zerolog.InfoLevel)
    case "warn":
        zerolog.SetGlobalLevel(zerolog.WarnLevel)
    case "error":
        zerolog.SetGlobalLevel(zerolog.ErrorLevel)
    case "fatal":
        zerolog.SetGlobalLevel(zerolog.FatalLevel)
    case "panic":
        zerolog.SetGlobalLevel(zerolog.PanicLevel)
    }

    switch logOutput {
    case "text":
        log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
    case "json":
    }
}
