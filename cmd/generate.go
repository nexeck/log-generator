package cmd

import (
    "math/rand"
    "time"

    "github.com/Pallinder/go-randomdata"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
    Use:   "generate",
    Short: "A brief description of your command",
    Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

    Run: func(cmd *cobra.Command, args []string) {
        var logLevel int
        for true {
            logLevel = rand.Intn(4)
            getLogEvent(logLevel).
                Str("Noun", randomdata.Noun()).
                Int( "Number", randomdata.Number(0,10000)).
                Str("IPv4", randomdata.IpV4Address()).
                Str("IPv6", randomdata.IpV6Address()).
                Bool("Bool", randomdata.Boolean()).
                Msg(randomdata.SillyName())
            time.Sleep(time.Duration(logInterval) * time.Millisecond)
        }
    },
}

func init() {
    rootCmd.AddCommand(generateCmd)
}

func getLogEvent(logLevel int) *zerolog.Event {
    switch logLevel {
    case 0:
        return log.Debug()
    case 1:
        return log.Info()
    case 2:
        return log.Warn()
    case 3:
        return log.Error()
    default:
        return log.Panic()
    }
}
